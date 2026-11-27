package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

// fakeTransport intercepts the client's requests and answers with canned
// responses, so the tests never touch the network. The page payloads are raw
// JSON so the rows are passed through exactly as Baserow would send them.
type fakeTransport struct {
	pagesJSON map[int]string // page -> raw JSON array of rows
	statuses  map[int]int    // page -> HTTP status, default 200
	count     int            // value for the envelope's "count" field
	nextPage  int            // last page number that advertises a next link

	requests []*http.Request
}

func (f *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	f.requests = append(f.requests, req)

	q := req.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}

	resp := &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Request: req}
	if status, ok := f.statuses[page]; ok {
		resp.StatusCode = status
		return resp, nil
	}

	next := ""
	if page < f.nextPage {
		nextQuery := url.Values{}
		for k, v := range req.URL.Query() {
			nextQuery[k] = v
		}
		nextQuery.Set("page", strconv.Itoa(page+1))
		next = req.URL.Scheme + "://" + req.URL.Host + req.URL.Path + "?" + nextQuery.Encode()
	}
	nextJSON, _ := json.Marshal(next)
	envelope := fmt.Sprintf(`{"count":%d,"next":%s,"previous":null,"results":%s}`,
		f.count, nextJSON, f.pagesJSON[page])
	resp.Body = io.NopCloser(bytes.NewReader([]byte(envelope)))
	resp.Header.Set("Content-Type", "application/json")
	return resp, nil
}

// patchTransport swaps http.DefaultTransport (used by run's default client)
// for the fake and returns a restore function.
func patchTransport(fake http.RoundTripper) func() {
	original := http.DefaultTransport
	http.DefaultTransport = fake
	return func() { http.DefaultTransport = original }
}

// fixturePagesJSON provides page payloads with fields in a deliberate,
// non-alphabetical order and with JSON values of mixed types (number, string,
// array), exactly as Baserow would send them.
var fixturePagesJSON = map[int]string{
	1: `[
		{"Model": "Apertus 8B", "Params": 8000000000, "Notes": "line one\nline two", "Flags": ["open weights", "EU | trained"]},
		{"Model": "Apertus 70B"}
	]`,
	2: `[{"Model": "Apertus Mini"}]`,
}

func TestFetchRowsAndRenderJSON(t *testing.T) {
	transport := &fakeTransport{
		pagesJSON: fixturePagesJSON,
		count:     3,
		nextPage:  2,
	}
	client := &http.Client{Transport: transport}

	cfg := config{
		baseURL: "https://baserow.test", apiKey: "test-key",
		tableID: 42, viewID: 7, pageSize: 200,
		outPath: filepath.Join(t.TempDir(), "out.json"),
	}
	snap, err := fetchRows(cfg, client)
	if err != nil {
		t.Fatalf("fetchRows: %v", err)
	}

	if snap.count != 3 {
		t.Errorf("count = %d, want 3", snap.count)
	}
	if len(snap.rows) != 3 {
		t.Errorf("fetched %d rows, want 3", len(snap.rows))
	}
	if len(transport.requests) != 2 {
		t.Fatalf("made %d requests, want 2 (one per page)", len(transport.requests))
	}

	first := transport.requests[0].URL
	if first.Path != "/api/database/rows/table/42/" {
		t.Errorf("endpoint path = %q", first.Path)
	}
	if got := first.Query(); got.Get("user_field_names") != "true" ||
		got.Get("view_id") != "7" || got.Get("size") != "200" || got.Get("page") != "1" {
		t.Errorf("first request query = %v", got)
	}
	if got := transport.requests[1].URL.Query().Get("page"); got != "2" {
		t.Errorf("second request page = %q, want 2", got)
	}
	if got := transport.requests[0].Header.Get("Authorization"); got != "Token test-key" {
		t.Errorf("Authorization header = %q", got)
	}

	// Rows are kept as raw JSON, byte-for-byte as the API sent them.
	if got := string(snap.rows[0]); !strings.Contains(got, `"Model": "Apertus 8B"`) {
		t.Errorf("first row = %s, want the untouched API row", got)
	}

	generated := time.Date(2026, 11, 24, 14, 32, 47, 0, time.UTC)
	data, err := renderJSON(snap, generated)
	if err != nil {
		t.Fatalf("renderJSON: %v", err)
	}

	var export exportFile
	if err := json.Unmarshal(data, &export); err != nil {
		t.Fatalf("export is not valid JSON: %v\n%s", err, data)
	}
	if export.Meta.Table != 42 || export.Meta.View != 7 || export.Meta.Count != 3 {
		t.Errorf("meta = %+v, want table 42, view 7, count 3", export.Meta)
	}
	if !export.Meta.Generated.Equal(generated) {
		t.Errorf("meta.generated = %v, want %v", export.Meta.Generated, generated)
	}
	if !strings.Contains(export.Meta.Source, "/api/database/rows/table/42/") {
		t.Errorf("meta.source = %q", export.Meta.Source)
	}
	if len(export.Rows) != 3 {
		t.Errorf("export has %d rows, want 3", len(export.Rows))
	}

	// Values keep their JSON types and content: no Markdown escaping.
	var row0 map[string]any
	if err := json.Unmarshal(export.Rows[0], &row0); err != nil {
		t.Fatalf("row is not valid JSON: %v", err)
	}
	if params, ok := row0["Params"].(float64); !ok || params != 8000000000 {
		t.Errorf("Params = %#v, want the number 8000000000", row0["Params"])
	}
	if notes, _ := row0["Notes"].(string); notes != "line one\nline two" {
		t.Errorf("Notes = %q, want a plain string with the original newline", notes)
	}
	flags, ok := row0["Flags"].([]any)
	if !ok || len(flags) != 2 || flags[1] != "EU | trained" {
		t.Errorf("Flags = %#v, want [\"open weights\", \"EU | trained\"] with the pipe intact", row0["Flags"])
	}
}

func TestRenderJSONEmpty(t *testing.T) {
	snap := &snapshot{tableID: 1, count: 0}
	data, err := renderJSON(snap, time.Now().UTC())
	if err != nil {
		t.Fatalf("renderJSON: %v", err)
	}
	var export exportFile
	if err := json.Unmarshal(data, &export); err != nil {
		t.Fatalf("export is not valid JSON: %v\n%s", err, data)
	}
	if export.Rows == nil || len(export.Rows) != 0 {
		t.Errorf("empty export should carry an empty rows array, got %s", data)
	}
	if strings.Contains(string(data), `"view"`) {
		t.Errorf("view 0 should be omitted from the metadata, got %s", data)
	}
}

func TestGetJSONErrorHints(t *testing.T) {
	testcases := []struct {
		status int
		hint   string
	}{
		{http.StatusUnauthorized, "valid database token"},
		{http.StatusForbidden, "no read access"},
		{http.StatusNotFound, "check the -table and -view IDs"},
	}
	for _, tc := range testcases {
		transport := &fakeTransport{statuses: map[int]int{1: tc.status}}
		client := &http.Client{Transport: transport}
		_, err := getJSON(client, "key", "https://baserow.test/api/database/rows/table/42/")
		if err == nil {
			t.Fatalf("status %d: expected an error", tc.status)
		}
		if !strings.Contains(err.Error(), tc.hint) {
			t.Errorf("status %d: error %q misses hint %q", tc.status, err, tc.hint)
		}
	}
}

func TestRunEndToEnd(t *testing.T) {
	restore := patchTransport(&fakeTransport{
		pagesJSON: map[int]string{1: fixturePagesJSON[1]},
		count:     2,
		nextPage:  1,
	})
	defer restore()

	out := filepath.Join(t.TempDir(), "data", "baserow", "out.json")
	env := func(key string) string {
		switch key {
		case "BASEROW_URL":
			return "https://baserow.test"
		case "BASEROW_API_KEY":
			return "test-key"
		}
		return ""
	}
	var stdout, stderr bytes.Buffer
	code := run([]string{"-table", "42", "-view", "7", "-out", out}, env, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run exit code = %d, stderr: %s", code, stderr.String())
	}
	content, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("reading output: %v", err)
	}
	for _, want := range []string{"\"meta\"", "\"rows\"", "\"Model\": \"Apertus 8B\""} {
		if !strings.Contains(string(content), want) {
			t.Errorf("output file misses %q:\n%s", want, content)
		}
	}
	if !strings.Contains(stdout.String(), "wrote "+out+" (2 rows)") {
		t.Errorf("unexpected summary output: %s", stdout.String())
	}
}

func TestRunRequiresTableAndKey(t *testing.T) {
	env := func(key string) string {
		if key == "BASEROW_API_KEY" {
			return "test-key"
		}
		return ""
	}
	var stdout, stderr bytes.Buffer
	if code := run([]string{}, env, &stdout, &stderr); code == 0 {
		t.Error("run without -table should fail")
	}

	stderr.Reset()
	if code := run([]string{"-table", "42"}, func(string) string { return "" }, &stdout, &stderr); code == 0 {
		t.Error("run without BASEROW_API_KEY should fail")
	}
	if !strings.Contains(stderr.String(), "BASEROW_API_KEY") {
		t.Errorf("error should mention BASEROW_API_KEY, got: %s", stderr.String())
	}
}

func TestDefaultOutPath(t *testing.T) {
	if got := defaultOutPath(42, 0); got != filepath.Join("data", "baserow", "table-42.json") {
		t.Errorf("defaultOutPath(42, 0) = %q", got)
	}
	if got := defaultOutPath(42, 7); got != filepath.Join("data", "baserow", "table-42-view-7.json") {
		t.Errorf("defaultOutPath(42, 7) = %q", got)
	}
}

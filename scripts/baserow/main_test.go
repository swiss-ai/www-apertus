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
// JSON so the field order is preserved exactly as Baserow would send it.
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
// non-alphabetical order, exactly as Baserow would send them.
var fixturePagesJSON = map[int]string{
	1: `[
		{"Model": "Apertus 8B", "Params": 8000000000, "Notes": "line one\nline two", "Flags": ["open weights", "EU | trained"]},
		{"Model": "Apertus 70B"}
	]`,
	2: `[{"Model": "Apertus Mini"}]`,
}

func TestFetchRowsAndRender(t *testing.T) {
	transport := &fakeTransport{
		pagesJSON: fixturePagesJSON,
		count:     3,
		nextPage:  2,
	}
	client := &http.Client{Transport: transport}

	cfg := config{
		baseURL: "https://baserow.test", apiKey: "test-key",
		tableID: 42, viewID: 7, pageSize: 200,
		outPath: filepath.Join(t.TempDir(), "out.md"),
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

	// Column order follows the API response, not Go's random map order.
	if got := strings.Join(columnOrder(snap.rows), ","); got != "Model,Params,Notes,Flags" {
		t.Errorf("column order = %q, want Model,Params,Notes,Flags", got)
	}

	md := renderMarkdown(snap, time.Now().UTC())
	for _, want := range []string{
		"Table:     42",
		"View:      7",
		"| Model | Params | Notes | Flags |",
		"| --- | --- | --- | --- |",
		"| Apertus 8B | 8000000000 | line one<br>line two | open weights, EU \\| trained |",
		"| Apertus 70B |  |  |  |",
		"| Apertus Mini |  |  |  |",
	} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown misses %q\ngot:\n%s", want, md)
		}
	}
	if strings.Contains(md, " EU | trained ") {
		t.Error("unescaped pipe found in rendered markdown")
	}
}

func TestRenderMarkdownEmpty(t *testing.T) {
	snap := &snapshot{tableID: 1, count: 0}
	md := renderMarkdown(snap, time.Now().UTC())
	if !strings.HasPrefix(md, "<!--") {
		t.Errorf("empty export should keep the metadata comment, got %q", md)
	}
	if strings.Contains(md, "|") {
		t.Errorf("empty export should contain no table, got %q", md)
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

	out := filepath.Join(t.TempDir(), "data", "baserow", "out.md")
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
	if !strings.Contains(string(content), "| Model | Params | Notes | Flags |") {
		t.Errorf("output file misses the markdown table:\n%s", content)
	}
	if !strings.Contains(stdout.String(), "2 rows, 4 columns") {
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
	if got := defaultOutPath(42, 0); got != filepath.Join("data", "baserow", "table-42.md") {
		t.Errorf("defaultOutPath(42, 0) = %q", got)
	}
	if got := defaultOutPath(42, 7); got != filepath.Join("data", "baserow", "table-42-view-7.md") {
		t.Errorf("defaultOutPath(42, 7) = %q", got)
	}
}

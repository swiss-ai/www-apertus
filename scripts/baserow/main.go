// Command baserow exports the rows of a Baserow table view as a Markdown
// table, ready to be committed to the Hugo site's data/ folder and rendered
// with the baserow-table shortcode (layouts/shortcodes/baserow-table.html).
//
// Configuration is read from the environment:
//
//	BASEROW_URL      API base URL, e.g. https://api.baserow.io (default)
//	BASEROW_API_KEY  database token created under Baserow "Settings > API tokens"
//
// The API key must never appear on the command line; use flags only for the
// non-secret parameters:
//
//	go run ./scripts/baserow -table 123 -view 45
//	go run ./scripts/baserow -table 123 -view 45 -out data/baserow/roadmap.md
//
// Rows are fetched from
//
//	GET {BASEROW_URL}/api/database/rows/table/{table}/?user_field_names=true&view_id={view}
//
// so the filters, sortings and column visibility of the view apply, and field
// values come back in their human readable form (see
// https://api.baserow.io/api/redoc/#operation/list_database_table_rows).
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	defaultBaseURL   = "https://api.baserow.io"
	defaultPageSize  = 200
	maxPages         = 10000
	maxResponseBytes = 64 << 20
	userAgent        = "www-apertus-baserow-fetch"
)

// row is a Baserow row whose fields are kept in the order they appear in the
// API response, so the generated Markdown columns match the table layout.
type row struct {
	keys   []string
	values map[string]any
}

// snapshot is the complete, paginated result of one export run.
type snapshot struct {
	tableID   int
	viewID    int
	count     int
	sourceURL string
	rows      []row
}

type config struct {
	baseURL  string
	apiKey   string
	tableID  int
	viewID   int
	outPath  string
	pageSize int
}

func main() {
	os.Exit(run(os.Args[1:], os.Getenv, os.Stdout, os.Stderr))
}

func run(args []string, getenv func(string) string, stdout, stderr io.Writer) int {
	cfg, err := parseFlags(args, getenv, stderr)
	if err != nil {
		return 2 // the message has already been printed
	}

	snap, err := fetchRows(cfg, &http.Client{Timeout: 30 * time.Second})
	if err != nil {
		fmt.Fprintf(stderr, "baserow: %v\n", err)
		return 1
	}

	markdown := renderMarkdown(snap, time.Now().UTC())
	if err := os.MkdirAll(filepath.Dir(cfg.outPath), 0o755); err != nil {
		fmt.Fprintf(stderr, "baserow: creating output directory: %v\n", err)
		return 1
	}
	if err := os.WriteFile(cfg.outPath, []byte(markdown), 0o644); err != nil {
		fmt.Fprintf(stderr, "baserow: writing %s: %v\n", cfg.outPath, err)
		return 1
	}

	fmt.Fprintf(stdout, "wrote %s (%d rows, %d columns)\n",
		cfg.outPath, len(snap.rows), len(columnOrder(snap.rows)))
	return 0
}

func parseFlags(args []string, getenv func(string) string, stderr io.Writer) (config, error) {
	fs := flag.NewFlagSet("baserow", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		fmt.Fprint(stderr, usage)
	}
	fail := func(err error) (config, error) {
		fmt.Fprintf(stderr, "baserow: %v\n\n", err)
		fmt.Fprint(stderr, usage)
		return config{}, err
	}

	var (
		tableFlag    = fs.Int("table", 0, "Baserow table ID (required)")
		databaseFlag = fs.Int("database", 0, "alias for -table")
		viewFlag     = fs.Int("view", 0, "Baserow view ID (applies the view's filters, sortings and hidden columns)")
		outFlag      = fs.String("out", "", "output Markdown path (default data/baserow/table-<id>[-view-<id>].md)")
		pageSizeFlag = fs.Int("page-size", defaultPageSize, "rows per request (Baserow allows at most 200)")
		urlFlag      = fs.String("url", "", "Baserow API base URL (overrides $BASEROW_URL)")
	)
	if err := fs.Parse(args); err != nil {
		return config{}, err
	}
	if fs.NArg() > 0 {
		return fail(fmt.Errorf("unexpected argument %q", fs.Arg(0)))
	}

	tableID := *tableFlag
	if tableID == 0 {
		tableID = *databaseFlag
	}
	if tableID <= 0 {
		return fail(fmt.Errorf("the -table flag is required"))
	}
	if *pageSizeFlag <= 0 {
		return fail(fmt.Errorf("invalid -page-size %d", *pageSizeFlag))
	}

	apiKey := getenv("BASEROW_API_KEY")
	if apiKey == "" {
		apiKey = getenv("BASEROW_API_TOKEN")
	}
	if apiKey == "" {
		return fail(fmt.Errorf("BASEROW_API_KEY is not set; create a database token in Baserow under Settings > API tokens"))
	}

	baseURL := *urlFlag
	if baseURL == "" {
		baseURL = getenv("BASEROW_URL")
	}
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	outPath := *outFlag
	if outPath == "" {
		outPath = defaultOutPath(tableID, *viewFlag)
	}

	return config{
		baseURL:  strings.TrimRight(baseURL, "/"),
		apiKey:   apiKey,
		tableID:  tableID,
		viewID:   *viewFlag,
		outPath:  outPath,
		pageSize: *pageSizeFlag,
	}, nil
}

func defaultOutPath(tableID, viewID int) string {
	name := fmt.Sprintf("table-%d", tableID)
	if viewID != 0 {
		name += fmt.Sprintf("-view-%d", viewID)
	}
	return filepath.Join("data", "baserow", name+".md")
}

func fetchRows(cfg config, client *http.Client) (*snapshot, error) {
	query := url.Values{}
	query.Set("user_field_names", "true")
	query.Set("size", strconv.Itoa(cfg.pageSize))
	if cfg.viewID != 0 {
		query.Set("view_id", strconv.Itoa(cfg.viewID))
	}
	endpoint := fmt.Sprintf("%s/api/database/rows/table/%d/", cfg.baseURL, cfg.tableID)

	snap := &snapshot{tableID: cfg.tableID, viewID: cfg.viewID}
	for page := 1; page <= maxPages; page++ {
		query.Set("page", strconv.Itoa(page))
		reqURL := endpoint + "?" + query.Encode()
		if page == 1 {
			snap.sourceURL = reqURL
		}

		body, err := getJSON(client, cfg.apiKey, reqURL)
		if err != nil {
			return nil, err
		}

		var envelope struct {
			Count   int               `json:"count"`
			Next    *string           `json:"next"`
			Results []json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(body, &envelope); err != nil {
			return nil, fmt.Errorf("%s returned invalid JSON: %w", reqURL, err)
		}
		snap.count = envelope.Count

		for _, raw := range envelope.Results {
			r, err := decodeRow(raw)
			if err != nil {
				return nil, fmt.Errorf("%s returned a malformed row: %w", reqURL, err)
			}
			snap.rows = append(snap.rows, r)
		}

		if envelope.Next == nil || *envelope.Next == "" || len(envelope.Results) == 0 {
			return snap, nil
		}
	}
	return nil, fmt.Errorf("gave up after %d pages of table %d", maxPages, cfg.tableID)
}

func getJSON(client *http.Client, apiKey, reqURL string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		hint := ""
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			hint = " (check that BASEROW_API_KEY is a valid database token)"
		case http.StatusForbidden:
			hint = " (the database token has no read access to this table)"
		case http.StatusNotFound:
			hint = " (check the -table and -view IDs)"
		}
		return nil, fmt.Errorf("%s: %s%s: %s",
			truncate(reqURL, 120), resp.Status, hint, truncate(strings.TrimSpace(string(body)), 300))
	}
	return body, nil
}

// decodeRow unmarshals one row while preserving the field order of the JSON
// object, which Go's map decoding would discard.
func decodeRow(raw json.RawMessage) (row, error) {
	r := row{values: map[string]any{}}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()

	tok, err := dec.Token()
	if err != nil {
		return r, err
	}
	if delim, ok := tok.(json.Delim); !ok || delim != '{' {
		return r, fmt.Errorf("expected a JSON object, got %v", tok)
	}
	for dec.More() {
		keyTok, err := dec.Token()
		if err != nil {
			return r, err
		}
		key, ok := keyTok.(string)
		if !ok {
			return r, fmt.Errorf("expected a field name, got %v", keyTok)
		}
		var val any
		if err := dec.Decode(&val); err != nil {
			return r, err
		}
		if _, seen := r.values[key]; !seen {
			r.keys = append(r.keys, key)
		}
		r.values[key] = val
	}
	if _, err := dec.Token(); err != nil { // closing brace
		return r, err
	}
	return r, nil
}

// columnOrder returns the union of all field names in first-seen order.
func columnOrder(rows []row) []string {
	var cols []string
	seen := make(map[string]bool)
	for _, r := range rows {
		for _, k := range r.keys {
			if !seen[k] {
				seen[k] = true
				cols = append(cols, k)
			}
		}
	}
	return cols
}

// renderMarkdown produces the data-file contents: an HTML comment with the
// export metadata followed by the rows as a GitHub Flavored Markdown table.
// With no rows only the comment is written.
func renderMarkdown(snap *snapshot, generated time.Time) string {
	var b strings.Builder
	view := "-"
	if snap.viewID != 0 {
		view = strconv.Itoa(snap.viewID)
	}
	fmt.Fprintf(&b, `<!--
  Baserow table snapshot, generated by scripts/baserow - do not edit.
  Table:     %d
  View:      %s
  Rows:      %d
  Source:    %s
  Generated: %s
-->
`, snap.tableID, view, snap.count, snap.sourceURL, generated.Format(time.RFC3339))

	cols := columnOrder(snap.rows)
	if len(snap.rows) == 0 || len(cols) == 0 {
		return b.String()
	}

	b.WriteString("\n")
	headers := make([]string, len(cols))
	for i, c := range cols {
		headers[i] = escapeCell(c)
	}
	b.WriteString("| " + strings.Join(headers, " | ") + " |\n")

	b.WriteString("|")
	for range cols {
		b.WriteString(" --- |")
	}
	b.WriteString("\n")

	for _, r := range snap.rows {
		cells := make([]string, len(cols))
		for i, c := range cols {
			cells[i] = formatCell(r.values[c])
		}
		b.WriteString("| " + strings.Join(cells, " | ") + " |\n")
	}
	return b.String()
}

// formatCell turns any Baserow field value into Markdown table cell text.
func formatCell(v any) string {
	return escapeCell(rawCell(v))
}

// rawCell renders a value without Markdown escaping; arrays are joined so the
// final cell is escaped exactly once.
func rawCell(v any) string {
	switch x := v.(type) {
	case nil:
		return ""
	case string:
		return x
	case json.Number:
		return x.String()
	case bool:
		return strconv.FormatBool(x)
	case []any:
		parts := make([]string, 0, len(x))
		for _, e := range x {
			parts = append(parts, rawCell(e))
		}
		return strings.Join(parts, ", ")
	case map[string]any:
		// Link-to-table, lookup and similar fields return nested objects;
		// prefer a display key, otherwise fall back to compact JSON.
		for _, key := range []string{"value", "name", "title", "visible_name"} {
			if display, ok := x[key]; ok {
				return rawCell(display)
			}
		}
		encoded, err := json.Marshal(x)
		if err != nil {
			return fmt.Sprint(x)
		}
		return string(encoded)
	default:
		encoded, err := json.Marshal(x)
		if err != nil {
			return fmt.Sprint(x)
		}
		return string(encoded)
	}
}

// escapeCell makes a value safe inside a GFM table cell: pipes cannot split
// the cell and line breaks become <br>.
func escapeCell(s string) string {
	s = strings.ReplaceAll(s, "|", `\|`)
	s = strings.ReplaceAll(s, "\r\n", "<br>")
	s = strings.ReplaceAll(s, "\n", "<br>")
	s = strings.ReplaceAll(s, "\r", "<br>")
	return s
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

const usage = `Export a Baserow table view as a Markdown table for Hugo's data/ folder.

Usage:
  go run ./scripts/baserow -table <id> [-view <id>] [flags]

Flags:
  -table int      Baserow table ID (required; -database is an alias)
  -view int       Baserow view ID; its filters, sortings and hidden columns apply
  -out string     output Markdown path (default data/baserow/table-<id>[-view-<id>].md)
  -page-size int  rows per request, max 200 (default 200)
  -url string     Baserow API base URL (overrides $BASEROW_URL)

Environment:
  BASEROW_URL      API base URL, default https://api.baserow.io
  BASEROW_API_KEY  database token from Baserow "Settings > API tokens" (required)
`

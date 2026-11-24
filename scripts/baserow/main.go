// Command baserow exports the rows of a Baserow table view as a JSON
// snapshot, ready to be committed to the Hugo site's data/ folder and
// rendered at build time with the showcase-cards shortcode
// (layouts/shortcodes/showcase-cards.html).
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
//	go run ./scripts/baserow -table 123 -view 45 -out data/baserow.json
//
// Rows are fetched from
//
//	GET {BASEROW_URL}/api/database/rows/table/{table}/?user_field_names=true&view_id={view}
//
// so the filters, sortings and column visibility of the view apply, and field
// values come back in their human readable form (see
// https://api.baserow.io/api/redoc/#operation/list_database_table_rows).
//
// The snapshot file has the shape
//
//	{"meta": {"table": …, "view": …, "count": …, "source": …, "generated": …},
//	 "rows": [ …the rows, exactly as the API returned them… ]}
//
// Hugo templates load it with os.ReadFile + transform.Unmarshal, which keeps
// the original field names and value types (no Markdown escaping round-trip).
package main

import (
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

// snapshot is the complete, paginated result of one export run. Rows are kept
// as raw JSON, so they are written back byte-for-byte as Baserow sent them.
type snapshot struct {
	tableID   int
	viewID    int
	count     int
	sourceURL string
	rows      []json.RawMessage
}

// exportFile is the on-disk shape of a snapshot.
type exportFile struct {
	Meta meta              `json:"meta"`
	Rows []json.RawMessage `json:"rows"`
}

// meta documents where a snapshot came from and when it was taken.
type meta struct {
	Table     int       `json:"table"`
	View      int       `json:"view,omitempty"`
	Count     int       `json:"count"`
	Source    string    `json:"source"`
	Generated time.Time `json:"generated"`
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

	data, err := renderJSON(snap, time.Now().UTC())
	if err != nil {
		fmt.Fprintf(stderr, "baserow: encoding JSON: %v\n", err)
		return 1
	}
	if err := os.MkdirAll(filepath.Dir(cfg.outPath), 0o755); err != nil {
		fmt.Fprintf(stderr, "baserow: creating output directory: %v\n", err)
		return 1
	}
	if err := os.WriteFile(cfg.outPath, data, 0o644); err != nil {
		fmt.Fprintf(stderr, "baserow: writing %s: %v\n", cfg.outPath, err)
		return 1
	}

	fmt.Fprintf(stdout, "wrote %s (%d rows)\n", cfg.outPath, len(snap.rows))
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
		outFlag      = fs.String("out", "", "output JSON path (default data/baserow/table-<id>[-view-<id>].json)")
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
	return filepath.Join("data", "baserow.json")
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
		snap.rows = append(snap.rows, envelope.Results...)

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

// renderJSON produces the data-file contents: the export metadata followed by
// the rows exactly as the API returned them. With no rows the rows array is
// empty rather than null, so downstream templates can range over it directly.
func renderJSON(snap *snapshot, generated time.Time) ([]byte, error) {
	rows := snap.rows
	if rows == nil {
		rows = []json.RawMessage{}
	}
	data, err := json.MarshalIndent(exportFile{
		Meta: meta{
			Table:     snap.tableID,
			View:      snap.viewID,
			Count:     snap.count,
			Source:    snap.sourceURL,
			Generated: generated.UTC().Round(time.Second),
		},
		Rows: rows,
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

const usage = `Export a Baserow table view as a JSON snapshot for Hugo's data/ folder.

Usage:
  go run ./scripts/baserow -table <id> [-view <id>] [flags]

Flags:
  -table int      Baserow table ID (required; -database is an alias)
  -view int       Baserow view ID; its filters, sortings and hidden columns apply
  -out string     output JSON path (default data/baserow.json)
  -page-size int  rows per request, max 200 (default 200)
  -url string     Baserow API base URL (overrides $BASEROW_URL)

Environment:
  BASEROW_URL      API base URL, default https://api.baserow.io
  BASEROW_API_KEY  database token from Baserow "Settings > API tokens" (required)
`

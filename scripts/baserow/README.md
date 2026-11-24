# baserow-fetch — export Baserow views as JSON for Hugo

`scripts/baserow` is a small Go utility (stdlib only) that pulls the rows of a
[Baserow](https://baserow.io) table **view** through the REST API and writes
them as a JSON snapshot into the site's `data/` folder. Hugo consumes those
snapshots at build time with the `showcase-cards` and `showcase-updated`
shortcodes (`layouts/shortcodes/`), which turn the rows into HTML cards —
no client-side JavaScript and no Markdown escaping round-trip.

It is a self-contained Go module, so it does not interfere with Hugo's module
graph (`go.mod` / `hugo.toml` at the repository root).

## Prerequisites

- Go 1.21+
- A Baserow **database token**: in Baserow go to *Settings → API tokens*,
  create a token and give it read access to the database containing your
  table. (The docs page in `apertus-program/Baserow/developer/Backend API.html`
  links the full API reference.)

## Configuration

| Environment variable | Purpose                                              | Default                  |
| -------------------- | ---------------------------------------------------- | ------------------------ |
| `BASEROW_API_KEY`    | Database token (required; `BASEROW_API_TOKEN` also accepted) | –                 |
| `BASEROW_URL`        | API base URL                                         | `https://api.baserow.io` |

Keep the token out of the shell history and the repository: put it in your
environment or a `.env` file (already git-ignored), e.g.

```sh
set -a; source .env; set +a
```

## Usage

Table and view IDs are visible in the Baserow web UI URL when you open a view:
`…/database/<database-id>/table/<table-id>/<view-id>/…`.

```sh
# Writes data/baserow.json
go run ./scripts/baserow -table 123 -view 45

# Custom output path
go run ./scripts/baserow -table 123 -view 45 -out data/baserow.json

# Self-hosted instance
BASEROW_URL=https://baserow.example.org go run ./scripts/baserow -table 123 -view 45
```

Flags:

| Flag           | Purpose                                                          |
| -------------- | ---------------------------------------------------------------- |
| `-table`       | Table ID (required; `-database` is an alias)                     |
| `-view`        | View ID; its filters, sortings and hidden columns are applied    |
| `-out`         | Output path (default `data/baserow.json`) |
| `-page-size`   | Rows per request, at most 200 (default 200)                      |
| `-url`         | Base URL, overrides `$BASEROW_URL`                               |

The tool paginates through the whole view
(`GET /api/database/rows/table/<table>/?user_field_names=true&view_id=<view>`)
and overwrites the output file with a JSON document:

```json
{
  "meta": {
    "table": 874809,
    "view": 2051696,
    "count": 72,
    "source": "https://api.baserow.io/api/database/rows/table/…",
    "generated": "2026-11-24T14:32:47Z"
  },
  "rows": [ … ]
}
```

The rows are passed through **byte-for-byte as the API returned them** (field
names in their human readable form, numbers as numbers, booleans as booleans,
multi-selects as arrays). Commit the generated file, so the site builds
without API credentials.

## Rendering in Hugo

Snapshots are not part of Hugo's `.Site.Data` tree; templates load them with
`os.ReadFile` + `transform.Unmarshal`, which preserves the original field
names and types. Two shortcodes do this for you in any content page:

```markdown
{{</* showcase-cards types="Fine-tune,Hosting" */>}}
{{</* showcase-updated href="https://baserow.io/database/…" */>}}
```

| Shortcode          | Purpose                                                            |
| ------------------ | ------------------------------------------------------------------ |
| `showcase-cards`   | Rows as `<a class="card">` boxes with a type badge                 |
| `showcase-updated` | "Snapshot of the project database, generated …" provenance line    |

`showcase-cards` parameters:

| Parameter  | Purpose                                                            |
| ---------- | ------------------------------------------------------------------ |
| `src`      | Snapshot path relative to `data/` (default `baserow.json`) |
| `types`    | Comma-separated `Type` values to include; omit for all types       |
| `badges`   | `"false"` hides the type badge on each card                        |
| `active`   | `"false"` also renders rows with `Active = false`                  |
| `truncate` | Maximum description length in characters (default 240)             |

Templates can call the underlying partial directly:

```go-html-template
{{ partial "showcase-cards.html" (dict "types" "Fine-tune") }}
```

The shortcodes render nothing and log a Hugo warning when the snapshot is
missing or no row matches, so a forgotten export never breaks a build.

## Tests

```sh
go test ./scripts/baserow
```

The tests cover pagination, authentication headers and JSON passthrough
(field order, value types, special characters) using a fake HTTP transport
(no network access needed).

## Optional: refreshing in CI

Snapshots are committed, so the regular deploy workflow needs no secret. To
refresh them automatically, add a workflow (or a step in the existing deploy
workflow) before the build:

```yaml
- name: Refresh Baserow snapshots
  env:
    BASEROW_URL: ${{ secrets.BASEROW_URL }}       # if self-hosted
    BASEROW_API_KEY: ${{ secrets.BASEROW_API_KEY }}
  run: |
    go run ./scripts/baserow -table 123 -view 45
    # commit & push changes, e.g. with EndBug/add-and-commit
```

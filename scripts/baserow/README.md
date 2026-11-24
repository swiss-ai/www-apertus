# baserow-fetch — export Baserow views as Markdown for Hugo

`scripts/baserow` is a small Go utility (stdlib only) that pulls the rows of a
[Baserow](https://baserow.io) table **view** through the REST API and writes
them as a GitHub Flavored Markdown table into the site's `data/` folder. Hugo
renders those snapshots with the `baserow-table` shortcode
(`layouts/shortcodes/baserow-table.html`).

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
# Writes data/baserow/table-123-view-45.md
go run ./scripts/baserow -table 123 -view 45

# Custom output path
go run ./scripts/baserow -table 123 -view 45 -out data/baserow/roadmap.md

# Self-hosted instance
BASEROW_URL=https://baserow.example.org go run ./scripts/baserow -table 123 -view 45
```

Flags:

| Flag           | Purpose                                                        |
| -------------- | -------------------------------------------------------------- |
| `-table`       | Table ID (required; `-database` is an alias)                   |
| `-view`        | View ID; its filters, sortings and hidden columns are applied  |
| `-out`         | Output path (default `data/baserow/table-<id>[-view-<id>].md`) |
| `-page-size`   | Rows per request, at most 200 (default 200)                    |
| `-url`         | Base URL, overrides `$BASEROW_URL`                             |

The tool paginates through the whole view
(`GET /api/database/rows/table/<table>/?user_field_names=true&view_id=<view>`)
and overwrites the output file with:

- an HTML comment carrying the export metadata (table/view/row count, source
  URL, timestamp), and
- the rows as a Markdown table. Column order follows the Baserow layout, cell
  values are rendered human-readably (`|` escaped, line breaks become `<br>`,
  single/multi-selects and dates appear as in the UI).

Commit the generated file, so the site builds without API credentials.

## Rendering in Hugo

Embed a snapshot in any content page:

```markdown
{{</* baserow-table src="baserow/table-123-view-45.md" caption="Source: [Baserow](https://baserow.io)" */>}}
```

| Parameter | Purpose                                             |
| --------- | --------------------------------------------------- |
| `src`     | Path of the snapshot relative to `data/` (required) |
| `caption` | Optional caption below the table (Markdown allowed) |
| `class`   | Extra CSS classes for the wrapping `<figure>`       |

Templates can call the underlying partial directly:

```go-html-template
{{ partial "baserow-table.html" (dict "src" "baserow/table-123-view-45.md") }}
```

The shortcode renders nothing and logs a Hugo warning when the snapshot is
missing or empty, so a forgotten export never breaks a build.

## Tests

```sh
go test ./scripts/baserow
```

The tests cover pagination, authentication headers, column ordering and
Markdown escaping using a fake HTTP transport (no network access needed).

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

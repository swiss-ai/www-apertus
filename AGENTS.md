# AGENTS.md

This file provides guidance to agents, such as Claude Code (claude.ai/code) when working with code in this repository.

## Getting Started

### Prerequisites
- **Go** (v1.21 or later) - Required for Hugo modules
- **Hugo Extended** (v0.140 or later) - Required for Sass processing

### Common Commands

```bash
# Start development server
hugo server

# Start development server with drafts enabled
hugo server -D

# Start server accessible from other devices
hugo server --bind 0.0.0.0

# Build for production
hugo --gc --minify

# Build with production environment
HUGO_ENV=production hugo --gc --minify

# Update Hugo modules
hugo mod get -u

# Create new content page
hugo new <section>/<filename>.md
```

## Project Architecture

This is a Hugo static website using the Lotus Docs theme. The structure follows standard Hugo conventions:

### Content Structure
- `content/` - Markdown content files
  - `_index.md` - Homepage content
  - `about.md`, `charter.md`, `documentation.md`, `faq.md`, `get-started.md`, `research.md` - Main pages

### Templates
- `layouts/` - Custom Hugo templates
  - `_default/` - Base templates (overrides theme)
  - `partials/` - Reusable template snippets
  - `shortcodes/` - Custom shortcodes for content formatting

### Assets
- `assets/` - Source assets processed by Hugo
- `static/` - Static files served as-is

### Configuration
- `hugo.toml` - Hugo configuration file
- `go.mod` - Go module dependencies (for Hugo modules)

## Custom Shortcodes

The site includes several custom shortcodes for content formatting:

### Columns
```markdown
{{< columns cols="3" >}}
### Column 1
Content here.
<--->
### Column 2
Content here.
<--->
### Column 3
Content here.
{{< /columns >}}
```

### Cards
```markdown
{{< cards cols="2" >}}
{{< card title="Card Title" href="/link" >}}
Card description.
{{< /card >}}
{{< /cards >}}
```

### Tabs
```markdown
{{< tabs name="example" >}}
{{% tab name="tab1" title="First Tab" %}}
Tab content here.
{{% /tab %}}
{{% tab name="tab2" title="Second Tab" %}}
More content.
{{% /tab %}}
{{< /tabs >}}
```

## Theme

The site uses the [Lotus Docs](https://lotusdocs.dev/) theme imported as a Hugo module. The base template in `layouts/_default/baseof.html` overrides the theme's default template to customize the site's appearance.

## Deployment

The site is automatically deployed to GitHub Pages via GitHub Actions when changes are pushed to the `main` branch. The build output is generated in the `public/` directory.

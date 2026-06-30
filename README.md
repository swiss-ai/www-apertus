# APERTVS

Official website for [Apertus](https://apertus-ai.org), a fully open foundation model for sovereign AI developed by the Swiss AI Initiative. This is a static website featuring a minimal, research-focused design. It serves as the public face of the Apertus project, providing information about the models, research, documentation, and governance.

## Tech Stack

- **[Hugo](https://gohugo.io/)** (v0.140+, extended) - Static site generator
- **[Bootstrap 5](https://getbootstrap.com/)** - CSS framework (loaded via CDN)
- **[Lotus Docs](https://lotusdocs.dev/)** - Hugo theme (customized)
- **GitHub Actions** - CI/CD deployment to GitHub Pages

## Prerequisites

Before running locally, ensure you have:

- **Go** (v1.21 or later, required by Hugo) - [Install the latest version](https://go.dev/learn/)
- **Hugo** (v0.140 or later) - [Installation guide](https://gohugo.io/installation/)

**Ubuntu/Debian:**
```bash
sudo apt install golang-go hugo
```

**macOS:**
```bash
brew install golang hugo
```

**Windows:**
```bash
choco install golang hugo-extended
```

## Getting Started

```bash
cd www-apertus
hugo server
```

Then open [http://localhost:1313](http://localhost:1313)

### Full Setup (fresh machine)

1. **Install Go** (required for Hugo modules):
   ```bash
   # Ubuntu - via snap (recommended)
   sudo snap install go --classic

   # Verify
   go version
   ```

2. **Install Hugo Extended** (must be extended version for Sass):
   ```bash
   # Ubuntu - download .deb
   wget https://github.com/gohugoio/hugo/releases/download/v0.142.0/hugo_extended_0.142.0_linux-amd64.deb
   sudo dpkg -i hugo_extended_0.142.0_linux-amd64.deb

   # Verify (should show "extended")
   hugo version
   ```

3. **Clone and run:**
   ```bash
   git clone https://github.com/swiss-ai/www-apertus.git
   cd www-apertus
   hugo server
   ```

4. **Open your browser:**
   Navigate to [http://localhost:1313](http://localhost:1313)

The site will automatically reload when you make changes.

### Troubleshooting

**"binary with name 'go' not found"**
- Install Go: `sudo snap install go --classic`

**"template for shortcode not found"**
- Ensure `layouts/shortcodes/` directory exists (not `_shortcodes`)
- Ensure `layouts/partials/` directory exists (not `_partials`)

**Theme header/SVG "Layer_1" appearing**
- Ensure `layouts/_default/baseof.html` exists to override the Lotus Docs theme's base template

**Hugo modules not downloading**
- Run `hugo mod get -u` to update modules
- Ensure Go is installed and in PATH

### Useful Commands

```bash
# Development server with drafts
hugo server -D

# Development server accessible from other devices
hugo server --bind 0.0.0.0

# Build for production
hugo --gc --minify

# Build with environment
HUGO_ENV=production hugo --gc --minify
```

## Project Structure

```
www-apertus/
├── archetypes/          # Templates for new content
├── assets/              # Source assets (processed by Hugo)
├── content/             # Markdown content
│   ├── _index.md        # Homepage content
│   ├── about.md         # About page
│   ├── charter.md       # Governance charter
│   ├── documentation.md # Documentation links
│   ├── faq.md           # FAQ page
│   ├── get-started.md   # Getting started guide
│   └── research.md      # Research publications
├── layouts/             # Custom Hugo templates
│   ├── _default/
│   │   ├── baseof.html  # Base template (overrides theme)
│   │   └── single.html  # Single page template
│   ├── partials/
│   │   ├── navbar.html  # Navigation bar
│   │   ├── hero.html    # Hero section
│   │   └── head.html    # HTML head
│   ├── shortcodes/      # Custom shortcodes
│   │   ├── cards.html   # Card grid container
│   │   ├── card.html    # Individual card
│   │   ├── columns.html # Multi-column layout
│   │   ├── tabs.html    # Tab container
│   │   ├── tab.html     # Individual tab
│   │   ├── button.html  # Styled button
│   │   └── hero.html    # Hero shortcode
│   └── index.html       # Homepage layout
├── static/              # Static files (served as-is)
│   ├── favicon.ico
│   └── images/
│       ├── apertvs.svg  # Logo mark
├── hugo.toml            # Hugo configuration
├── go.mod               # Go module dependencies
└── README.md
```

## Customization

### Colors

The site uses a simple color palette defined inline in the templates:

| Purpose | Color | Hex |
|---------|-------|-----|
| Primary (buttons, links) | Light Blue | `#49BED8` |
| Accent (details) | Bright Red | `#E02D24` |
| Text | Dark Gray | `#1a1a1a` |
| Muted Text | Gray | `#6b7280` |
| Background | White/Light Blue | `#ffffff` / `#e8f7fa` |

### Navigation

Edit `layouts/partials/navbar.html` to modify navigation links.

### Homepage Content

Edit `content/_index.md` to update homepage content. The frontmatter controls the hero section:

```yaml
---
title: "Apertus"
mission: "A Fully Open Foundation Model for Sovereign AI"
description: "Your description here..."
hero_primary_label: "Explore Models"
hero_primary_url: "https://huggingface.co/..."
hero_secondary_label: "Read the Paper"
hero_secondary_url: "https://arxiv.org/..."
---
```

### Shortcodes

The site includes several custom shortcodes for content formatting:

**Columns:**
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

**Cards:**
```markdown
{{< cards cols="2" >}}
{{< card title="Card Title" href="/link" >}}
Card description.
{{< /card >}}
{{< /cards >}}
```

**Tabs:**
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

## Deployment

The site is automatically deployed to GitHub Pages via GitHub Actions when changes are pushed to the `main` branch.

### Manual Deployment

To build the site manually:

```bash
hugo --gc --minify
```

The generated site will be in the `public/` directory.

### GitHub Pages Setup

The deployment workflow is configured in `.github/workflows/hugo.yaml`. It:

1. Installs Hugo and Go
2. Builds the site with `hugo --gc --minify`
3. Deploys to GitHub Pages

## Adding New Pages

1. Create a new markdown file in `content/`:
   ```bash
   hugo new docs/getting-started.md
   ```

2. Edit the frontmatter and content:
   ```yaml
   ---
   title: "Getting Started"
   description: "How to get started with Apertus"
   ---

   Your content here...
   ```

3. The page will be available at `/docs/getting-started/`

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Make your changes
4. Test locally with `hugo server`
5. Commit your changes (`git commit -am 'Add improvement'`)
6. Push to the branch (`git push origin feature/improvement`)
7. Open a Pull Request

## License

This website is licensed under [Apache 2.0](LICENSE).

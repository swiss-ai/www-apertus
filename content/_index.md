---
title: "Apertus"
description: "Switzerland's first large-scale open, multilingual language model"
# Hero
hero:
  enable: true
  weight: 10
  template: hero

  backgroundImage:
    path: "images/templates/hero"
    filename:
      desktop: "gradient-desktop.webp"
      mobile: "gradient-mobile.webp"

  badge:
    text: v0.1.0
    color: primary # primary, secondary, success, danger, warning, info, light, dark
    pill: false # boolean
    soft: true # boolean

  title: "Lotus Docs"
  subtitle: A lightweight, **modern documentation** theme for Hugo. Easily customised for building **fast**, **secure**, and **SEO-friendly** documentation sites.

  image:
    path: "images" # path to image under configured assets directory. default 'images'
    filename: "lotus_docs_screenshot.png" # filename of your hero image (including file extension)
    alt: "Lotus Docs Screenshot" # Optional but recommended
    boxShadow: true # default 'false' (excludes .svg images)
    rounded: true # round the image corners? default 'false' (excludes .svg images)

  ctaButton:
    icon: rocket_launch
    btnText: "Get Started"
    url: "/docs/quickstart/#create-a-new-lotus-docs-site"
  cta2Button:
    icon: construction
    btnText: "In Development"
    url: "https://github.com/colinwilson/lotusdocs"

  info: "**Open Source** MIT Licensed."
---

{{< hero 
    title="Apertus" 
    subtitle="Switzerland's First Open Multilingual AI" 
    description="A collaborative effort between EPFL, ETH Zurich, and CSCS – trained on 15 trillion tokens across 1,000+ languages, including Swiss German, Romansh, and many underrepresented languages."
>}}

{{% button href="https://publicai.co" icon="rocket" %}}Try the Demo{{% /button %}}
{{% button href="/docs" icon="book" style="secondary" %}}Read the Docs{{% /button %}}

{{< /hero >}}

---

## Why Apertus?

Apertus – Latin for "open" – reflects our commitment to transparency in AI development. All architecture, weights, and training data are openly documented and available.

{{< columns >}}

### Multilingual
Trained on a diverse dataset with **40% non-English content**, supporting over 1,000 languages including underrepresented ones.

<--->

### Transparent
Fully open-source with documented architecture, weights, and training data. Built on Swiss values of trust and accountability.

<--->

### Accessible
Available in **8B** and **70B** parameter sizes under Apache 2.0 license – suitable for education and commercial use.

{{< /columns >}}

---

## Get Started

{{< cards >}}

{{< card title="Download Models" icon="download" link="https://huggingface.co/ApertusLLM" >}}
Access both 8B and 70B models on Hugging Face with free download options.
{{< /card >}}

{{< card title="Try the Demo" icon="play" link="https://publicai.co" >}}
Experience Apertus through our interactive chat and API interface at PublicAI.co.
{{< /card >}}

{{< card title="Read the Docs" icon="book" link="/docs" >}}
Comprehensive documentation for deployment, fine-tuning, and integration.
{{< /card >}}

{{< /cards >}}

---

## News & Publications


### Technical Report Released
Our comprehensive technical report detailing Apertus's architecture, training methodology, and benchmark results is now available.

[Read the Tech Report →](/publications/tech-report)

---

### Swiss AI Initiative Launch
Apertus is the flagship project of the Swiss AI Initiative, a joint effort between EPFL, ETH Zurich, and CSCS.

[Learn More →](/about/swiss-ai-initiative)


---

## For Every User

Whether you're a developer, researcher, or just curious about AI, there's a path for you.

{{< tabs >}}

{{% tab title="Users" %}}

### Start Using Apertus Today

- **[Try the Demo](https://publicai.co)** – No setup required, start chatting immediately
- **[FAQ](/docs/faq)** – Common questions and quick answers
- **[Use Cases](/docs/use-cases)** – Explore chatbots, translation, educational tools, and more

{{% /tab %}}

{{% tab title="Developers" %}}

### Build with Apertus

- **[Hugging Face Models](https://huggingface.co/ApertusLLM)** – Download and integrate into your projects
- **[API Documentation](/docs/api)** – RESTful API reference and examples
- **[Deployment Guides](/docs/deployment)** – Deploy on Transformers, vLLM, SGLang, llama.cpp, or MLX

{{% /tab %}}

{{% tab title="Researchers" %}}

### Advance the Science

- **[Research Papers](/publications)** – Academic publications and citations
- **[Training Data](/docs/data)** – Documentation on our multilingual dataset
- **[Benchmarks](/docs/benchmarks)** – Performance metrics and comparisons

{{% /tab %}}

{{% tab title="Contributors" %}}

### Join the Community

- **[GitHub Repository](https://github.com/ApertusLLM)** – Contribute code, report issues, suggest features
- **[Community Guidelines](/community/guidelines)** – How to participate effectively
- **[Roadmap](/community/roadmap)** – See what's coming and help shape the future

{{% /tab %}}

{{< /tabs >}}

---

## Transparency & Compliance

Apertus is built on a foundation of ethical AI development and legal compliance.

| Principle | Implementation |
|-----------|----------------|
| **Data Integrity** | Built only from publicly available data, respecting opt-outs and removing personal data |
| **Swiss Regulations** | Compliant with Swiss data protection and copyright laws |
| **EU AI Act** | Meets transparency requirements for foundation models |
| **Reproducibility** | Full open-source code, weights, and training data documentation |

---

## Part of the Swiss AI Initiative

Apertus exemplifies Swiss collaboration – bringing together interdisciplinary expertise from research, engineering, and infrastructure across Switzerland's leading institutions.

- EPFL
- ETH Zurich  
- CSCS (Swiss National Supercomputing Centre)

---

## Get Involved

We welcome feedback and contributions to drive Apertus's continued development.

### Ready to explore?

{{% button href="https://publicai.co" icon="rocket" %}}Try the Demo{{% /button %}}
{{% button href="https://huggingface.co/ApertusLLM" icon="download" style="secondary" %}}Download Models{{% /button %}}
{{% button href="https://github.com/ApertusLLM" icon="github" style="outline" %}}View on GitHub{{% /button %}}

---

### Apertus

Switzerland's open multilingual AI model

#### Quick Links
- [Home](/)
- [Documentation](/docs)
- [Publications](/publications)
- [About](/about)

#### Resources
- [Hugging Face](https://huggingface.co/ApertusLLM)
- [GitHub](https://github.com/ApertusLLM)
- [Demo](https://publicai.co)
- [API Docs](/docs/api)

#### Community
- [Contributing](/community/contributing)
- [Code of Conduct](/community/code-of-conduct)
- [FAQ](/docs/faq)
- [Contact](/contact)

#### Legal
- [Privacy Policy](/legal/privacy)
- [Terms of Use](/legal/terms)
- [License (Apache 2.0)](/legal/license)

---

© 2026 Swiss AI Initiative. A collaboration between EPFL, ETH Zurich, and CSCS.

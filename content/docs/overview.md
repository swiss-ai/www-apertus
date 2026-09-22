---
weight: 100
title: "Overview"
description: "Apertus is an open source LLM developed in Switzerland"
icon: "circle"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Introduction"]
categories: [""]
author: "Apertus Project"
---

Apertus is a fully open source Large Language Model (LLM) developed in Switzerland.
This documentation shows you how to get started with the LLM, whether as user, researcher, or advanced contributor: we are maintaining this knowledge base for you, and could [✉️ use your feedback](/contact/).
This website is running on the [Hugo](https://gohugo.io/) static site engine and [Lotus docs](https://lotusdocs.dev/) theme. The source code is available on [GitHub](https://github.com/swiss-ai/www-apertus).

- **For casual users:** There are a number of apps in our [Community Showcase](/pages/get-showcase), as well as a range of providers providing high quality inference services for Apertus listed on our [Getting Started](/pages/get-started) page.
- **For developers:** The providers listed on the [Getting Started](/pages/get-started) page offer services in Apertus LLM hosting online for you to use via API. We test and verify each listing, but make no guarantees about their uptime or performance.
- **For self-hosting:** You can download one of the Apertus model distributions for free from Hugging Face and deploy it locally into a desktop tool, such as the ones listed in our [Deployment Guides](guides/). Check if you have enough system resources (e.g. 16 GB for BF16 versions, ~6 GB for quantized builds), or look for community-supported quantized releases.
- **For Apertus contributors:** If you have questions or constructive feedback about your experience with Apertus, it is easy to engage with the developer and user community - just look for the official Swiss AI projects on Hugging Face, GitHub, and social media - see our [Contact page](/contact) for links on how to get involved.

⬅️ [Back to the home page](https://apertus-ai.org)

About the project
-----------------

The Apertus model is part of the [Swiss AI Initiative](https://www.swiss-ai.org/), which started in late 2023. This is a platform for many [data science projects](https://www.swiss-ai.org/swiss-ai-projects), including the development of several foundation models.

The Apertus project began in earnest with the launch of **Apertus 1.0 in September 2025**, when EPFL, ETH Zurich, and the Swiss National Supercomputing Centre (CSCS) released Switzerland's first large-scale, fully open, multilingual large language model (LLM). 
Key highlights of the LLM project, as [pre-announced in July](https://ethz.ch/en/news-and-events/eth-news/news/2025/07/a-language-model-built-for-the-public-good.html) and described in the [official September announcement](/articles/2025-09-apertus-1-0/), include:

-   **Multilingualism**: Trained on more than 15 trillion tokens across 1,500+ languages, 40% non-English - equal usage cost across languages - see [public summary](https://github.com/swiss-ai/apertus-legal/blob/main/apertus_1/Apertus_EU_Public_Summary.pdf)
-   **Performance**: This is a large model (8 billion and 70 billion parameters), and it will be continue to be actively optimized.
-   **Open & Transparent**: Published under [Apache-2.0 license](https://huggingface.co/swiss-ai/Apertus-70B-Instruct-2509/blob/main/LICENSE.txt) - including source code, weights, and open training data.
-   **Data Privacy**: Developed with attention to the GDPR, the EU AI Act, and Swiss data protection law - see [Fan et al 2025](https://arxiv.org/abs/2504.06219)
-   **Infrastructure**: Developed on the new [Alps supercomputer at CSCS](https://www.cscs.ch/computers/alps) with over 10,000 NVIDIA GH200 Grace-Hopper chips
-   **Global Reach**: Research and borderless applications in mind, for sovereign and international public-interest AI.

The model was distinguished by its complete openness—architecture, weights, training data, and recipes were all publicly documented—alongside compliance with Swiss data protection and copyright law and EU AI Act transparency requirements. From the start, access was provided through strategic partner Swisscom's sovereign AI platform, Hugging Face downloads, and the Public AI Inference Utility, positioning Apertus as a "blueprint" for trustworthy, sovereign AI built for the public good.

Through early and mid-2026, progress shifted from release to ecosystem building and external validation. The second [Open-Source LLM Builders Summit](/articles/2026-02-builder-summit/) gathered the research community to figure out how global collaboration on open LLMs could actually work, addressing knowledge sharing, regulatory fragmentation, and infrastructure dependencies. The [Swiss AI SME Circle](/articles/2026-03-sme-circle/) connected businesses with the project's engineers, repeatedly surfacing demand for data sovereignty, practical integration paths, and hyperscaler independence. 

Real-world adoption followed: the Canton of Ticino deployed an [in-house translation system](/articles/2026-03-ticino/) built on a model fine-tuned by Artificialy (scoring 94% on their benchmarks after fine-tuning), an [independent assessment](/articles/2026-02-gpai/) by Trinity College Dublin's AI Accountability Lab ranked Apertus's training-data transparency summary among the best in the industry, and [the family expanded](/articles/2026-06-apertus-mini/) with the **Apertus 1.1 Mini** (June 2026) collection — distilled and quantized models for resource-constrained and offline use. Scientific recognition came in July with the Apertus technical report's acceptance at the ACL 2026 Main Conference in San Diego.

By late 2026 the project had matured from research artifact to operational, widely deployed infrastructure. **Apertus 1.5** (July 2026) [extended the original](/articles/2026-07-apertus-1-5/) models via continued pretraining with additional trillions of tokens, adding native image understanding (and experimental audio), an optional "thinking mode" for reasoning, a four-fold larger 262k-token context window, and improved instruction-following and tool use—all while preserving full openness. 

The September 2026 [general-availability announcement](/articles/2026-09-apertus-1-5-ga/) showcases a broad ecosystem of third-party providers offering the model as cloud APIs or on-premise deployments. With growing community engagement (startup award winners, the interpretability platform Apertus Claritas, thousands of newsletter subscribers, media coverage a year in), the project looks ahead to **Apertus 2.0** in 2027: a magnitude-larger model developed to the same standard of transparency.

Sources
-------

For further information:

- [Official Apertus Website](https://www.apertus-ai.org/) (apertus-ai.org)
- [Swiss AI Initiative](https://www.swiss-ai.org/apertus) (swiss-ai.org)
- [CSCS Announcement](https://www.cscs.ch/science/computer-science-hpc/2025/apertus-a-fully-open-transparent-multilingual-language-model) (cscs.ch)
- [EPFL AI Center](https://ai.epfl.ch/) (ai.epfl.ch)
- [ETH AI Center](https://ai.ethz.ch/) (ai.ethz.ch)

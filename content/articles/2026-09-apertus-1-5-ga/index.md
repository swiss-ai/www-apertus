---
title: "Inference Ecosystem"
date: 2026-09-17T00:00:00+01:00
draft: false
author: "ETH EPFL CSCS"
tags: ["Apertus", "Release", "Announcement"]
featured_image: "inference-ecosystem-2026-09.jpeg"
summary: "Getting started with Apertus is easier than ever."
comments: false
---


#### No longer limited to the research lab or the local workstation, Apertus is being deployed across a vibrant ecosystem from on‑premise deployments to high‑throughput APIs. Find out more in this update.

Apertus 1.5 adds image understanding competitive with leading open-weight models at scale, experimental audio understanding, stronger tool use, and thinking mode from our second-generation post-training pipeline. You can read more about this in our [July announcement](https://apertus-ai.org/articles/2026-07-apertus-1-5/).

Some of these capabilities need to be enabled through configuration changes and upgrades on the inference side. Pioneering organizations have gone above-and-beyond to support our work this summer, in some cases contributing their engineer's expertise and code commits to key open source components to benefit the entire community.

<small>§ &nbsp; <i>All logos shown above are trademarks of their respective owners.</i></small>


Over the summer, we became thanks to our enthusiastic community the highest downloaded model from Europe on Hugging Face, with over 4 million downloads in total across official Apertus releases [@swiss-ai](https://huggingface.co/swiss-ai)

![A graph of Apertus Model downloads over time](/images/pub/huggingface-stats-0926.png)

| Model               | Release | Downloads/month |
|---------------------|---------|-----------------|
| **Apertus 8B Instruct** | 2025    | 472k            |
| Mistral Small 3.1 24B | 2025    | 462k            |
| Ministral 3 14B Instruct | 2025    | 417k            |
| Devstral Small 2 24B   | 2025    | 275k            |
| **Apertus v1.5 8B**       | 2026    | 201k            |
| Mistral Medium 3.5 128B| 2026    | 89k             |
| Mistral Small 4        | 2026    | 54k             |
| **Apertus 70B Instruct**  | 2025    | 22k             |
| **Apertus v1.5 70B**      | 2026    | 17k             |

<small>Monthly top model downloads, main official repositories, as of 20.9.2026 from <a href="https://huggingface.co/docs/hub/models-download-stats">Hugging Face</a></small>

# Community Evaluation

An outreach and community function connects external contributors and deployment engineers with the development process and collects their findings to improve model compatibility. [1] The public website [^1] includes the charter, usage guides, benchmark resources, model documentation, links to downstream projects; newsletters, social media, and multilingual media monitoring complement it. Direct involvement of the engineering team has included presentations to science and industry, round-tables with companies, support of hackathons, and hands-on workshops.

One month before the Apertus 1.5 release, 15 organizations evaluated the 8B and 70B model weights, represented by roughly 50 engineers and experts. This pre-release exercise covered multiple areas of deployment readiness: quantization, alignment, multimodal input, tokenizer distribution, and general compatibility. It encouraged the production of quantizations and community builds for several inference frameworks.

We are grateful to evaluators and contributors at [Artificialy](https://huggingface.co/artificialy), [Begasoft](https://www.begasoft.ch/themenbeitraege/zukunft-der-ki), [Exoscale](https://huggingface.co/exoscale), [Federal Court](https://bger.ch/) (BGER), [Infomaniak](https://huggingface.co/Infomaniak-AI), [Liip](https://www.liip.ch/en/blog/apertus-1-5-first-impressions-from-using-switzerland-s-updated-ai-model), [OnPrem.ai](https://www.onprem.ai/en/knowhow/run-apertus-v15-70b-single-nvidia-rtx-6000/), [Phoeniqs](https://phoeniqs.com/en), [Public AI](https://publicai.co/stories/apertus-1-5), [Puzzle AG](https://www.puzzle.ch/service-products/ai), [stepping stone](https://www.stepping-stone.ch/en/products/artificial-intelligence/ai-on-demand-powered-by-swiss-ai-initiative), [Swisscom](https://docs.cloud.swisscom.ch/guide/cloud-services/aip/models/apertus-1_5_70B), [Switch](https://www.switch.ch/en) and [VSHN](https://www.vshn.ch/en/) who participated in the pre-release feedback round.

A limited-scale performance evaluation was run with several providers of the 70B model. To give an idea of the difference in the real-world speed of the various APIs on offer, the table below shows anonymized scores on simulated workload performance (Latency, Throughput) using the open source tool [GuideLLM](https://vllm-project.github.io/guidellm/main/), in order of ascending latency:

| Provider | Latency (s) | Input (tokens/s) | Output (tokens/s) | TTFT (ms) |
|----------|-------------|----------------------------|-----------------------------|-----------|
| CH       | 1.89762 | 179.696881 | 67.938525 | 103.629446 |
| P3       | 2.086661 | 157.557343 | 61.516734 | 87.940463 |
| P5       | 2.181321 | 150.90462 | 58.917091 | 171.180878 |
| P1       | 2.307436 | 143.597144 | 56.354363 | 469.813888 |
| P4       | 3.634175 | 92.761177 | 35.291866 | 81.59709 |
| P2       | 4.296362 | 79.530041 | 29.930671 | 223.610248 |

<small>CH is our own research data center. TTFT denotes time to first token. Lower latency and TTFT, and higher throughput, are better.</small>

---

# General Availability

The fully open model is [available for download](https://huggingface.co/collections/swiss-ai/apertus-v15) in two configurations - an 8B version that runs on many laptops and workstations, and the full-scale 70B parameter model for server-class hardware. 

We are glad to announce general availability of third-party services that provide access and support of Apertus 1.5: 


| Provider         | Models | Location¹ | Per-Token² | Docs |
|------------------|--------|----------|------------------|------|
| [Swisscom](#swisscom)         | 70B | 🇨🇭 |  | [`cloud.swisscom.ch`](https://docs.cloud.swisscom.ch/guide/cloud-services/aip/models/apertus-1_5_70B) |
| [Infomaniak](#infomaniak)       | 70B | 🇨🇭 | ✅ | [`infomaniak.com`](https://www.infomaniak.com/en/hosting/ai-services/open-source-models) |
| [PHOENIQS](#phoeniqs)         | 70B | 🇨🇭 | ✅ | [`kvant.cloud`](https://documentation.kvant.cloud/maas/active-models/) |
| [Safe Swiss Cloud](#safe-swiss-cloud) | 70B | 🇨🇭 | ✅ | [`safeswisscloud.com`](https://safeswisscloud.com/en/private-ai/) |
| [OnPrem.ai](#onprem-ai)        | 70B | 🇨🇭 |  | [`onprem.ai`](https://www.onprem.ai/en/ai-llm-models/) |
| [stepping stone](#stepping-stone)    | 8B | 🇨🇭 |  | [`stepping-stone.ch`](https://www.stepping-stone.ch/en/products/artificial-intelligence/ai-on-demand-powered-by-swiss-ai-initiative) |
| [Public AI](#public-ai)        | 8B, 70B | 🌍 | ✅ | [`publicai.co`](https://platform.publicai.co) |
| [Featherless](/docs/deploy/featherless/)        | 70B | 🌍 | ✅ | [`featherless.ai`](https://featherless.ai/models?query=apertus) |
| [AWS Sagemaker](/docs/deploy/sagemaker/) | 8B, 70B | 🌍 |  | [`aws.amazon.com`](https://aws.amazon.com/blogs/alps/switzerlands-open-source-apertus-llms-now-available-on-amazon-sagemaker-ai/) |
| [Microsoft Azure](/docs/deploy/azure/) | 8B, 70B | 🌍 |  | [`github.com`](https://github.com/Azure-Samples/swiss-llm-quickstart/blob/main/azure-virtual-machine/README.md) |
| [Google Vertex AI](/docs/deploy/vertex/) | 8B, 70B | 🌍 |  | [`cloud.google.com`](https://docs.cloud.google.com/gemini-enterprise-agent-platform/models/open-models/deploy-custom-vllm) |

<small>
Information is current as of September 2026. Notes:
1. Physical location of the datacenter where user data is processed.
2. Per-token rates are available for metered Apertus 1.5 model usage.
</small>

---

# Providers

You can find the links to current general-access providers who we are working with on our [Get Started page](https://apertus-ai.org/pages/get-started/). More technical instructions for deploying our models to several additional cloud providers are also available in the technical [deployment guides](/docs/deploy/). Please contact us if there are others that you are using.

### Swisscom

> "Switzerland shouldn't only consume AI, it should be able to build it. Apertus 1.5 shows the Swiss research community can deliver a genuinely open, genuinely capable model — and Swisscom's job is to make it usable: hosted in Switzerland, hardened for regulated industries, available from day one." — _Sarah Levy, Head of [Swiss AI Platform](https://www.swisscom.ch/en/business/enterprise/offer/platforms-applications/data-driven-business/swiss-ai-platform.html)_

Swisscom is offering 1000 keys for Apertus during the [Swiss {ai} Weeks](https://ai-weeks.ch/events?tag=Hackathon) hackathons

70B  ✅ High availability  ✅ Security certification  ✅ Strategic Partner

Documentation: [cloud.swisscom.ch](https://docs.cloud.swisscom.ch/guide/cloud-services/aip/models/apertus-1_5_70B)

### Infomaniak

> "Robust and sovereign European cloud offering LLM access with full data protection, 100 % renewable energy, and a clear no‑logs policy---letting you integrate any open‑source model into your applications while keeping your data securely hosted in Swiss data centers."

The [Euria app](https://euria.infomaniak.com/) supports Apertus 1.5 on mobile phones and in kSuite

70B  ✅ Swiss data center  ✅ Per-token rates ✅ Code contributor

To get started: [Discover Euria](https://www.infomaniak.com/en/support/faq/2924/discover-euria-and-the-apertus-model), connect to [AI services](https://www.infomaniak.com/en/hosting/ai-services/open-source-models)

### PHOENIQS

> "With Apertus 1.5's support for agents, tools and out-of-the-box EU AI Act compliance, thanks to Switzerland, the Continent now joins the race hereby dominated by the Americans and Chinese."

The PHOENIQS AI platform brings Apertus to regulated enterprise and public sector workloads.

70B  ✅ Swiss data center ✅ Enterprise scalability ✅ Consulting support

See [model documentation](https://documentation.kvant.cloud/maas/active-models/)


### OnPrem AI

> "Replace any cloud AI with local enterprise AI servers. It is literally plug&play, thanks to compatible APIs and the latest LLM models, managed through a user-friendly interface."

Offers quantized, full-featured [Apertus builds](https://huggingface.co/onprem-ai/) at roughly half the rated GPU power of the reference deployment, approximately 35% less energy per generated token than the FP8 alternative, and 99% of FP8 MMLU quality in a 48 GiB checkpoint.

70B  ✅ On premise deployment ✅ Consulting support ✅ Code contributor

Details: [Blog post](https://www.onprem.ai/en/knowhow/run-apertus-v15-70b-single-nvidia-rtx-6000/), [Models](https://www.onprem.ai/en/ai-llm-models/)


### stepping stone

> "Use AI models directly on Swiss infrastructure --- without sending data abroad. stepping stone offers leading open-source models as a managed service: ready to use, with data sovereignty and personalised support."

The model deployment is [openly documented](https://wiki.stoney-cloud.com/wiki/AI_on_demand:_apertus-ai/Apertus-v1.5-8B) and operated on Swiss infrastructure.

8B  ✅ Swiss data center ✅ ISO certification ✅ Consulting support

Consult the [solution offer](https://www.stepping-stone.ch/en/products/artificial-intelligence/ai-on-demand-powered-by-swiss-ai-initiative)

### Safe Swiss Cloud

> "Choose from a rich catalog of sovereign LLMs -- all with the same strict privacy and compliance guarantees. Safe Swiss Cloud's Private AI (PAI) services combine a broad selection of open-source LLMs with a consistent security, privacy and compliance foundation. You keep full control over data, infrastructure and model choice, while we provide the sovereign hosting and operational excellence"

Apertus 1.5 is available, optimized for multilingual dialogue use cases.

70B  ✅ Swiss data center ✅ ISO certification ✅ Per-token rates

Get started with fully [private AI](https://safeswisscloud.com/en/private-ai/)


### Public AI

> "The Public AI Inference Utility is a nonprofit, open-source project. Our team builds products and organizes advocacy to support the work of public AI model builders like the Swiss National AI Initiative, AI Singapore, AI Sweden, and the Barcelona Supercomputing Center."

An Apertus demo is freely accessible on Hugging Face and on [publicai.co](https://publicai.co), while a not-for-profit [cooperative](https://publicai.ch) brings local focus to the vision.

8B & 70B  ✅ Free demo ✅ Large community ✅ Social enterprise

Join the active &amp; [growing community](https://publicai.co/stories/apertus-1-5)


---

# What This Means for You

From academic research groups that demand strict data sovereignty to startups that require seamless API access, the breadth of providers ensures that Apertus can be plugged into your workflow in the way that best matches your operational and ethical requirements.

- On‑Premise deployments leverage the model securely behind your own firewalls, ideal for handling sensitive datasets or regulated environments.  
- Cloud and SaaS partners deliver ready‑to‑use endpoints, automatic scaling, and simplified billing, letting you focus on building rather than infrastructure.  
- Community initiatives like Public AI keep a free or low-cost gateway running on shared compute, preserving inclusivity with an easy way to demo new services.

Please check back on our [Get Started](https://apertus-ai.org/pages/get-started/) website for links, contact us if you are aware of other providers, or have a question about getting Apertus running on another platform.


### Looking Ahead

Just as Apertus 1.5 expanded its knowledge base and instruction‑following prowess, this growing partnership network expands its reach without diluting the open‑source principle that powers it. 

We thank every provider and contributor who has turned the vision of a truly universal, multilingual LLM into a practical reality, fully in the sense of the Open project that is one of our primary research goals. Here is what's next on the [Apertus Roadmap](https://apertus-ai.org/pages/about/#roadmap):

**Apertus 2.0 - 2027** - a magnitude larger model, developed to the same standard of transparency and trustworthiness: deeper understanding, better agentic capabilities, a more globally fair vocabulary, reliability in long-running processes, and more.

Further development of Apertus 2.5 will be based on feedback from our users. Stay tuned for the latest updates via our Inside [Apertus newsletter](https://apertus-ai.org/subscribe/). Join us at upcoming events to learn what innovative solutions emerge as developers worldwide integrate these models into new research, products, and services.

If you build with or use Apertus, reach out for a [Showcase](https://apertus-ai.org/showcase/) listing or through our [Contact page](https://apertus-ai.org/contact/). Don't hesitate to send us a link or screenshot of a cool prompt or use case! We will promote them in presentations, give coverage in the [News area](http://apertus-ai.org/news), as well as make them visible to our network at the [ETH](https://ai.ethz.ch/news-and-events/ai-center-news.html) and [EPFL](https://ai.epfl.ch/news/) AI Centers and [CSCS](https://www.cscs.ch/publications/news).

---

🌍 Tag **#Apertus** and connect with us on social media: [Mastodon](https://fosstodon.org/@apertus/116975431194025200) / [LinkedIn](https://www.linkedin.com/posts/apertus_apertus-opensourcellm-swissaiinitiative-activity-7486428926039203840-JBvM?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAByCdcBz1u0JhDLT9NpOkRcCuok5UzgWXQ) / [X](https://x.com/apertusllm/status/2080661538474451147) / [Bluesky\
](https://bsky.app/profile/did:plc:ytwqtr3ykzq6nzr7kg2465ca/post/3mrfkp6q4zc2z)

⭕️ Join the discussion in the community forums on[  Hugging](https://huggingface.co/swiss-ai/Apertus-v1.5-8B/discussions)[  Face](https://huggingface.co/swiss-ai/Apertus-v1.5-70B/discussions) / [Discord](https://discord.gg/t9TY8FsJd)

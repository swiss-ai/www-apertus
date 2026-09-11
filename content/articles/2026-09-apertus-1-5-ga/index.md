---
title: "Apertus 1.5"
date: 2026-07-24T00:00:00+01:00
draft: false
author: "ETH EPFL CSCS"
tags: ["Apertus", "Release", "Announcement"]
featured_image: "apertus-keyvisual-1-5.jpg"
summary: "Multimodal input, reasoning, longer context, and better instruction following"
comments: false
---


## Apertus 1.5 has demonstrated the ability to handle complex instruction-following, coding, and multimodal reasoning through several trillion additional training tokens in a fully open model. No longer limited to the research lab or the local workstation, Apertus is being deployed across a vibrant ecosystem from on‑premise deployments to high‑throughput APIs. Find out more in this update.

Being among the best fully open models at this scale for following instructions, Apertus 1.5 adds image understanding competitive with leading open-weight models at this scale, audio understanding, stronger tool use, and reasoning capability from our second-generation post-training pipeline. You can read more about this in our [July announcement](https://apertus-ai.org/articles/2026-07-apertus-1-5/).

Some of these capabilities need to be enabled through configuration changes and upgrades on the inference side. We would like to express our gratitude to a group of pioneering organizations going above-and-beyond to support our work over the summer, in some cases contributing their engineer's expertise and code commits to key open source components to benefit the entire community.

The fully open model is available for download on Hugging Face since July in two configurations - an 8B version that runs on many laptops and workstations, and the full-scale 70B parameter model for server-class hardware. For most people who do not wish to bother with installing and configuring an LLM, we are glad to announce general availability of cloud services that provide it. Here is an overview of the third-party platforms currently available for accessing Apertus 1.5. You can also find the links to them on our [dedicated web page](https://apertus-ai.org/pages/get-started/). Please contact us if there are others that you are aware of, and we will add them.

### Swisscom

"The Enterprise Service Cloud (ESC) is the Swiss private cloud for companies, offering IT infrastructure, security, and managed services from a single source -- fully integrated, automated, locally hosted, and flexibly scalable."

Swisscom is offering 1000 API keys during the Swiss {ai} Weeks

70B  ✅ High availability  ✅ Security certification  ✅ Strategic Partner

<https://docs.cloud.swisscom.ch/guide/cloud-services/aip/models/apertus-1_5_70B> 

### Infomaniak

"Robust and sovereign European cloud offering LLM access with full data protection, 100 % renewable energy, and a clear no‑logs policy---letting you integrate any open‑source model into your applications while keeping your data securely hosted in Swiss data centers."

The Euria app now supports Apertus natively

70B  ✅ Swiss data center  ✅ Sustainable design ✅ Code contributor

<https://www.infomaniak.com/en/hosting/ai-services/open-source-models> 

### OnPrem.ai

"Replace any cloud AI with local enterprise AI servers. It is literally plug&play, thanks to compatible APIs and the latest LLM models, managed through a user-friendly interface."

The best open Swiss model -- sovereign, air-gap capable, with vision, tool calling and a 192k context profile (229k measured maximum), at roughly half the rated GPU power of the reference deployment, approximately 35% less energy per generated token than the FP8 alternative, and 99% of FP8 MMLU quality in a 48 GiB checkpoint.

70B  ✅ On premise deployment  ✅ Consulting support  ✅ Code contributor

<https://www.onprem.ai/en/knowhow/run-apertus-v15-70b-single-nvidia-rtx-6000/> 

### Stepping Stone

"Use AI models directly on Swiss infrastructure --- without sending data abroad. stepping stone offers leading open-source models as a managed service: ready to use, with data sovereignty and personalised support."

Apertus --- the Swiss AI Initiative's fully open language model: developed by ETH Zurich, EPFL and CSCS, with 8 billion parameters and support for over 1,000 languages. The model weights, training data and methods are fully documented --- operated by stepping stone on Swiss infrastructure.

8B  ✅ Swiss data center  ✅ ISO certification  ✅ Consulting support

<https://www.stepping-stone.ch/en/products/artificial-intelligence/ai-on-demand-powered-by-swiss-ai-initiative> 

### PHOENIQS

"A Swiss sovereign cloud and AI provider that gives organizations full control over their digital infrastructure and data, enabling secure innovation, strategic independence, and scalable growth."

Apertus 1.5 is available on the PHOENIQS AI platform, bringing Switzerland's sovereign, open-source multimodal model to regulated enterprise and public sector workloads.

70B  ✅ Swiss data center  ✅ Enterprise scalability  ✅ Consulting support

<https://documentation.kvant.cloud/maas/active-models/> 

[LinkedIn post](https://www.linkedin.com/posts/phoenix-technologies-ag_apertus-15-available-on-phoeniqs-ai-platform-activity-7486417413090050048-pH8K)

### Public AI 

"The Public AI Inference Utility is a nonprofit, open-source project. Our team builds products and organizes advocacy to support the work of public AI model builders like the Swiss National AI Initiative, AI Singapore, AI Sweden, and the Barcelona Supercomputing Center."

Apertus 1.5 makes Switzerland's open language models more useful in practice: multimodal input, reasoning, longer context, and better instruction following.

8B & 70B  ✅ Free demo  ✅ Large community

<https://publicai.co/stories/apertus-1-5> 

### Safe Swiss Cloud

"Choose from a rich catalog of sovereign LLMs -- all with the same strict privacy and compliance guarantees. Safe Swiss Cloud's Private AI (PAI) services combine a broad selection of open-source LLMs with a consistent security, privacy and compliance foundation. You keep full control over data, infrastructure and model choice, while we provide the sovereign hosting and operational excellence"

Apertus 1.5 is available, optimized for multilingual dialogue use cases.

70B  ✅ Swiss data center  ✅ ISO certification  ✅ Integration support

<https://safeswisscloud.com/en/private-ai/> 

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

🌍 Tag **#Apertus** and connect with us on social media: [Mastodon](https://fosstodon.org/@apertus/116975431194025200) / [LinkedIn](https://www.linkedin.com/posts/apertus_apertus-opensourcellm-swissaiinitiative-activity-7486428926039203840-JBvM?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAByCdcBz1u0JhDLT9NpOkRcCuok5UzgWXQ) / [X](https://x.com/apertusllm/status/2080661538474451147) / [Bluesky\
](https://bsky.app/profile/did:plc:ytwqtr3ykzq6nzr7kg2465ca/post/3mrfkp6q4zc2z)

⭕️ Join the discussion in the community forums on[  Hugging](https://huggingface.co/swiss-ai/Apertus-v1.5-8B/discussions)[  Face](https://huggingface.co/swiss-ai/Apertus-v1.5-70B/discussions) / [Discord](https://discord.gg/t9TY8FsJd)

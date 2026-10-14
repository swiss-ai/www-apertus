---
title: "Apertus 1.5"
date: 2026-10-14T00:00:00+01:00
draft: false
author: "ETH EPFL CSCS"
tags: ["Apertus", "Release", "Announcement"]
featured_image: "technical-report.jpg"
summary: "Multimodal input, reasoning, longer context, and better instruction following"
comments: false
---

#### Apertus 1.5, our latest model release, extends the 8B and 70B Apertus 1.0 models with image understanding, an optional thinking mode, a four times longer context window, improved instruction-following, and better tool use. 

Apertus 1.5 is a continued pretraining of Apertus 1.0: we added 4 trillion tokens of text and multimodal training data to the 8B model and 2 trillion tokens to the 70B model. As with every Apertus release, the models are fully open: open [weights](https://huggingface.co/collections/swiss-ai/apertus-v15), open [data](https://github.com/swiss-ai/apertus-data-transparency/), open [values](https://apertus-ai.org/pages/charter/), and full training [details](https://apertus-ai.org/pages/research/). 

The updated [8B and 70B models](https://huggingface.co/collections/swiss-ai/apertus-v15) were released to Hugging Face on July 24, 2026 under an Apache 2.0 license. Full details on the training process are now available in the project Technical Report [arXiv:2610.15542](https://arxiv.org/abs/2610.15542) [cs.CL]

## Key Enhancements

The following capabilities are highlighted in the new release, with references to sections of the report:

- *Native Image Understanding:* [§2] Apertus 1.5 accepts images alongside text, bringing documents, diagrams, and photos into the conversation. The models can also process spoken language, though this remains experimental.
- *Thinking Mode:* [§4.2] The models can be switched into a thinking mode to reason about the input before answering, which improves performance on reasoning problems.
- *Long Context:* [§3.3] The context window grows to 262,144 tokens, a four-fold increase over Apertus 1.0.
- *Improved Instruction-Following:* [§4.2] Better instruction adherence than Apertus 1.0 gives more predictable and accurate responses.
- *Improved Tool Use:* [§5.8] The models are trained for better tool integration, allowing more effective use of external tools and APIs.
- *Open Values:* [§4.6] Stronger adherence to the [Apertus Charter](/pages/charter/), which brings transparency to the model’s values and principles.

Screenshot from developer testing:

![Screenshot of Apertus](Apertus-1.5-Wasserkraftwerks_Niederried-Radelfingen,_Schnittzeichnung.jpg)

_Multimodal processing: Apertus 1.5 70B interprets the contents of a technical drawing of a [Water power plant](https://commons.wikimedia.org/wiki/File:Wasserkraftwerks_Niederried-Radelfingen,_Schnittzeichnung.jpg) (CC BY 4.0) based on an image and short prompt, replying with a description in the form of a text response._

## Evaluation Results

Alongside the extended model capabilities, we also significantly expanded the scope of evaluations. Beyond the existing text-based evaluations used for Apertus 1.0, we added multi-modal evaluations (text & audio), and thinking-mode evaluations. 

We also added a series of text-based benchmarks, as well as extended benchmarking for translation and multilingual capabilities, safety, and tool use. These can be found in our post-training [evaluation repository](https://github.com/swiss-ai/evals-post-train). Our results for text-based evaluations are in the [model evaluation card](https://evaleval-general-eval-card.hf.space/models/swiss-ai/Apertus-v1.5-8B).

---

<iframe src="https://evaleval-general-eval-card.hf.space/embed/model/reported-metrics/swiss-ai/Apertus-v1.5-8B" width="100%" height="680" frameborder="0" style="border:1px solid #e5e5e5;border-radius:4px;" loading="lazy" title="Reported metrics"></iframe>

---

Apertus collaborates with the Every Eval coalition to ensure fully-transparent and reproducible benchmarking. [Apertus 1.5 8B](https://evaleval-general-eval-card.hf.space/models/swiss-ai/Apertus-v1.5-8B) and [Apertus 1.5 70B](https://evaleval-general-eval-card.hf.space/models/swiss-ai/Apertus-v1.5-70B) perform well against other fully-open peer models of comparable size in standard knowledge-based benchmarks, showcasing the effectiveness of the continued pre-training. 

While still trailing frontier open-weight models, the model's new release has also significantly improved in coding and math capabilities. Thinking further improves performance on challenging reasoning-heavy benchmarks like AIME. The multimodal performance of the model is closer to current leading open-weight models like Gemma 4 12B and others. 

Seven intermediate repositories have been added to the **[Apertus 1.5 collection](https://huggingface.co/collections/swiss-ai/apertus-v15)**, including the base weights, SFT, RLVR and DPO. The sharing of intermediate training checkpoints allow researchers to investigate how each stage changes model behavior.

## Model Availability

[Proton Lumo](https://proton.me/support/apertus-lumo) and [Infomaniak Euria](https://www.infomaniak.com/en/support/faq/2924/discover-euria-and-the-apertus-model) are two widely used apps that have recently integrated Apertus 1.5. A [discussion thread](https://github.com/swiss-ai/www-apertus/discussions/66) has previews of these deployments, which are linked in the **[Community Showcase](https://apertus-ai.org/pages/get-showcase/)**. 

Access to Apertus 1.5 in the form of an API for developers has been announced by a range of providers, including [Swisscom](https://docs.cloud.swisscom.ch/guide/cloud-services/aip/models/apertus-1_5_70B), [Infomaniak](https://www.infomaniak.com/en/hosting/ai-services/open-source-models), [Phoeniqs](https://documentation.kvant.cloud/maas/active-models/), and other companies listed in our **[General Availability guide](https://apertus-ai.org/articles/2026-09-apertus-1-5-ga/)**. The Technical Report [§7] describes how we work with external evaluators to test and prepare the models.

Technical instructions for running the models are in the model cards of the [Apertus 1.5 8B](https://huggingface.co/swiss-ai/Apertus-v1.5-8B) and [Apertus 1.5 70B](https://huggingface.co/swiss-ai/Apertus-v1.5-70B) releases on Hugging Face and in our **[Quickstart](https://apertus-ai.org/docs/quickstart/)**.

## How to Get Involved

First of all, thank you for reading this blog post and being part of a great community! We just celebrated five million downloads in our Hugging Face analytics 🤗 and we are not slowing down. At the recent [AI+X Summit](https://plusx.ai), our team more widely announced the roadmap for **Apertus 2.0**: sparse/MoE architecture, 500–700B total parameters, extended training schedule, with agents, mathematics and coding as priorities.

In summary, we invite you to read our reports, download the models, explore the intermediate checkpoints, finetune the weights, and reproduce our evaluations. Questions that are not already covered in [our publications](https://apertus-ai.org/pages/research/) or [website FAQ](https://apertus-ai.org/docs/faq/) are always welcome, and we would also love to hear which topics would be most useful to the community for future deep dives.

If you build with or use Apertus, fill out the [Showcase form](https://apertus-ai.org/showcase/), and reach out through the options on our [Contact page](https://apertus-ai.org/contact/). Looking forward to highlight more work, learn what you are building, and hear how Apertus can be improved!

### Citation

> Apertus Project, et al. “Apertus 1.5: From Open Text Models to Open Multimodal Reasoners.” arXiv:2610.15542. Preprint, arXiv, October 11, 2026. https://doi.org/10.48550/arXiv.2610.15542.

<a name="coverage"></a>

### Online Coverage

- [Apertus 1.5: Building the next generation of open AI infrastructure](https://actu.epfl.ch/news/apertus-15-building-the-next-generation-of-open--2/) - EPFL - 24.7.2026
- [Apertus 1.5: Building the next generation of open AI infrastructure](https://ai.ethz.ch/news-and-events/ai-center-news/2026/07/apertus-15-building-the-next-generation-of-open-ai-infrastructure.html) - ETH Zurich - 24.7.2026
- [Apertus 1.5: Building the next generation of open AI infrastructure](https://www.cscs.ch/science/computer-science-hpc/2026/apertus-15-building-the-next-generation-of-open-ai-infrastructure) - CSCS - 24.7.2026
- [With 1.5x the love, from Switzerland](https://publicai.co/stories/apertus-1-5) - Public AI - 24.7.2026
- [IA: l’EPFL et l’ETH Zurich lancent le modèle Apertus 1.5](https://www.xavierstuder.com/2026/07/apertus-1-5-llm-ia-intelligence-artificielle-epfl-eth-zurich/) - Xavier Studer - 25.7.2026
- [Schweizer KI Apertus bekommt Update](https://www.inside-it.ch/schweizer-ki-apertus-bekommt-update-20260727) - inside-it.ch - 27.7.2026
- [LLM Benchmark Evaluation - Apertus 1.5-8B](https://blog.nlp-lab.ai/2026/07/29/Apertus15Bench.html) - DS-NLP Lab - 29.7.2026
- [Apertus 1.5 : l’IA souveraine suisse lance sa nouvelle version](https://dcod.ch/2026/07/31/apertus-1-5-ia-souveraine-suisse/) - Marc Barbezat - 31.7.2026
- [Run Apertus v1.5 70B on a single NVIDIA RTX 6000](https://www.onprem.ai/en/knowhow/run-apertus-v15-70b-single-nvidia-rtx-6000/) - onprem.ai - 3.8.2026
- [The State of the Art in Open Source AI for Swiss Legal Tasks](https://github.com/JoelNiklaus/SwissLegalEvals/blob/main/blog/swiss-legal-evals-2026.md#the-state-of-the-art-in-open-source-ai-for-swiss-legal-tasks) - Joel Niklaus - 4.8.2026
- [Helvetra moves to Apertus 1.5: what changes, and what we're watching](https://www.helvetra.ch/news/apertus-1-5-what-it-means-for-helvetra) - 4.8.2026
- [Shrinking Apertus 1.5 8B for an 8 GB Laptop GPU](https://digitalpathlines.ch/2026/08/12/shrinking-apertus-1-5-8b-for-an-8-gb-laptop-gpu/) - Emmanuel Belo - 12.8.2026
- [Apertus 1.5 - first impressions from using Switzerland’s updated AI model](https://www.liip.ch/en/blog/apertus-1-5-first-impressions-from-using-switzerland-s-updated-ai-model) - Josef Kruckenberg - 26.8.2026
- [A year after its launch, it is time to take stock of Switzerland’s AI model Apertus](https://www.swissinfo.ch/eng/swiss-ai/one-year-on-has-swiss-ai-model-apertus-lived-up-to-the-hype/91984276) - Sara Ibrahim - 2.9.2026

#### \* * * * *

Please contact us if you would like to see your review listed here.
For updates and further information, subscribe to our [Inside Apertus newsletter](/subscribe).

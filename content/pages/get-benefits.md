---
title: "A Practical Guide to Open Source AI for Impact"
---

**For organizations seeking open, multilingual, audit-friendly AI without vendor lock‑in, Apertus is a ready-to-build-on model under Apache 2.0, EU AI‑Act‑friendly, already multilingual, and simple to run locally or via vetted open endpoints. Fine-tune it on your own mission-specific data to maximise relevance and impact while staying transparent and affordable.**

This guide synthesizes insights to explain how organizations of all types can safely, affordably, and effectively use Apertus to advance their missions—from accessible education and business support, to local language preservation and community empowerment.

See also: [further links](#resources) below.

### 🔑 Key benefits of Apertus

1. **Multilingual & Affordable**  
   - Apertus supports ~1,800 languages—ideal for organizations who want help-desk chat in local dialects without licensing overhead.

2. **Transparency & Ethics‑by‑Design**  
   - **EU AI Act–ready** safeguards mean you can meet reporting and audit requirements from the start as your build out AI safety layers.

3. **Open License**  
   - **Apache 2.0** lets you develop bespoke tools, redistribute as part of your service stack, or embed in existing software stacks without fear of infringement.

4. **Extensibility**  
   - Specialise Apertus on your domain: legal briefs for legal‑aid NGOs, health terminology for clinics, education content for NGOs running literacy programs, etc.

5. **No Vendor Lock‑in**  
   - You’re never tied to a SaaS provider; move hosting if you outgrow one vendor or switch providers for legal or budget reasons.

6. **Strong Swiss Academic & Compute Backbone**  
   - CSCS & EPFL provide fast inference access for validated Swiss‑based projects.

### 🧾 Fictional “Value Story” to use in Proposals

> *“Our Nonprofit *Lingua‑Aid* serves displaced populations across French, Arabic, and Swahili-speaking regions. By adopting Apertus (open, multilingual, EU‑compliant), we reduce per-user translation costs by ~40% and maintain privacy​—no vendor-locked SaaS, no per-token fees beyond our own compute. Fine-tuning on our aid-guidance library boosted first-call accuracy to 91% in appropriate local dialects, allowing our volunteers to focus on complex cases rather than babysitting chatbots.”*

For actual case studies, please visit our [Community Showcase](/pages/get-showcase/) and [News Archive](/news).

### 📌 Practical Checklist before you go Live

| Item | Action | Why |
|------|--------|-----|
| **Data Source Review** | Use only **your own data** or widely licenced datasets; avoid mixed licensing that could trigger copyright alerts. | Keeps you compliant and reinforces donor trust. |
| **PII Mitigation** | Enable **output filtering** and **assistant profile** settings to block personal info leaks. | Satisfies GDPR/AI‑Act transparency expectations. |
| **Rate Limits & Budgeting** | Start with *Free tiers*; monitor usage; select managed endpoints or on‑prem as traffic grows; budget compute per 1k tokens (often < $0.001). | Avoids surprise bills but still gives you elasticity. |
| **Backup & Versioning** | Keep record of exact `model-id` or git commit before fine-tuning; store fine‑tuning configs in repo (Apache-licensed). | Efficient rollback if new checkpoint breaks something. |
| **Funding Leverage** | Tie use of Apertus to **Swiss academic partnerships**, **EU open‑source grants**, or corporate sponsors interested in **Swiss‑neutral AI**. | Demonstrates sustainability and good governance for funders. |
| **Documentation & Tutorials** | Document your fine-tuning steps and evaluation results; link back to Apertus docs/examples. | Builds community knowledge, creating a virtuous cycle. |

# Resources

Materials to encourage further exploration & self-assisted studies:

Apertus Resources
-----------------

**Research topics** & publications\
<https://apertus-ai.org/pages/research/>

**Technical documentation**\
<https://apertus-ai.org/docs/>

**Community of practice** (Showcase)\
<https://apertus-ai.org/pages/get-showcase/>

**NGO Workshop** materials (10.11.2026)\
[Slide presentation](https://docs.google.com/presentation/d/e/2PACX-1vSKRWA8CyqLoMgv345NA6YRmZbuORWwq4aTFGRF45hSDNWcUxoBXLjAz9U_ucaW-jojgvkG4XCmKZE8/pub?start=false&loop=false&delayms=3000)

Research Platforms
------------------

**Swiss AI Projects** - an interactive map\
<https://www.swiss-ai.org/swiss-ai-projects>

**Swiss AI Research Platform** (CSCS)\
<https://serving.swissai.svc.cscs.ch/faq>

**RCP - AI Inference Platform** (EPFL)\
<https://www.epfl.ch/research/facilities/rcp/ai-inference-as-a-service/>

**Renku** - open research platform (SDSC)\
<https://docs.renkulab.io/en/stable/>

**Continuing education** programs (EPFL)\
[Extension School](https://www.epfl.ch/education/continuing-education/), [AI and Machine Learning Essentials](https://www.formation-continue-unil-epfl.ch/en/formation/ai-ml-essentials/)

Tutorials / Demos
-----------------

**Apertus RAG Demo in Python**\
Ask questions about the science and training of Apertus.\
<https://ask.apertus-ai.org>

**Apertus Local demo**\
Scripts to bootstrap Apertus 1.5 on your local machine.\
<https://huggingface.co/spaces/faruk-zahiragic/apertus-local>

**Chat with MeditronFO**\
Apertus-70B fine-tuned on the open MeditronFO medical data.\
<https://epflight.github.io/Meditron/#chat>

**Apertus Mini in your Pocket, or in your Browser**\
Distillations and Quantizations of our models into more compact formats.\
<https://apertus-ai.org/articles/2026-06-apertus-mini/>

**Awesome AI Switzerland**\
A curated list of Apertus links from the Swiss AI ecosystem.\
<https://github.com/rnckp/awesome-ai-switzerland>

**Inside the Machines** (Andrea Fanfani, 2026)\
A plain-language guide to how modern AI models actually work --- no maths required.\
<https://www.fanfani.net/tutorials/inside-the-machines/>

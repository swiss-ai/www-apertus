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

Apertus is an open source Large Languag Model (LLM) developed in Switzerland.
This documentation shows you how to get started with the LLM, whether as user, researcher, or advanced contributor: we are maintaining this knowledge base for you, and could [✉️ use your feedback](mailto:llm-requests@swiss-ai.org).

The model development team is part of the [Swiss AI Initiative](https://www.swiss-ai.org/), which started in late 2023. This is a platform for over 80 data science projects including the LLM development. Key highlights of the LLM project, as [announced in July](https://ethz.ch/en/news-and-events/eth-news/news/2025/07/a-language-model-built-for-the-public-good.html), include:

-   **Multilingualism**: Trained on more than 15 trillion tokens across 1,500+ languages, 40% non-English - equal usage cost across languages - see [@epfml](http://huggingface.co/datasets/epfml)
-   **Performance**: This is a large model (8 billion and 70 billion parameters), trained on a lot of tokens, and it will be continue to be actively optimized.
-   **Open & Transparent**: Published under Apache-2.0 license - including source code, weights, and open training data.
-   **Data Privacy**: Compliant with GDPR, EU AI Act, and Swiss data protection laws - see [Fan et al 2025](https://arxiv.org/abs/2504.06219)
-   **Infrastructure**: Developed on the new Alps supercomputer at CSCS with over 10,000 NVIDIA GH200 Grace-Hopper chips
-   **Global Reach**: Research and borderless applications in mind, for sovereign and international public-interest AI.

Tech specs
----------

The Swiss LLM is trained on the [Alps supercomputer](https://en.wikipedia.org/wiki/Alps_(supercomputer)), operational at CSCS since [September 2024](https://www.netzwoche.ch/news/2024-09-17/neue-forschungsinfrastruktur-supercomputer-alps-eingeweiht):

-   10,752 NVIDIA GH200 Grace-Hopper chips
-   Computing power: 270-435 PFLOPS
-   [Ranked 6th on the TOP500 list](https://top500.org/lists/top500/2024/06/) (June 2024)

The Swiss LLM was trained on approximately 15 trillion tokens. Particularly noteworthy is the high proportion of non-English data (40%) and coverage of over 1,500 languages, including rare ones like Romansh or Zulu. The data was ethically sourced - without illegal scraping, respecting `robots.txt` and copyright requirements. While this limits access to certain specialized information, CSCS emphasizes: «For general tasks, this doesn't lead to measurable performance losses.»

For more technical references, see the [Sources](#sources) further below.

### Initial benchmarks

See the [Evaluation section](https://huggingface.co/swiss-ai/Apertus-8B-2509#evaluation) of the Apertus Model Card, and Section 5 of the [Tech Report](https://github.com/swiss-ai/apertus-tech-report/blob/main/Apertus_Tech_Report.pdf) for more data. This is an initial independent evaluation, and we expect more to come soon:

| Model | MMLU (Knowledge) | Global-MMLU (Multilingual) | GSM8K (Math) | HumanEval (Code) | RULER @32k (Long Context) |
| --- | --- | --- | --- | --- | --- |
| Claude 3.5 Sonnet | 88.7% | --- | 96.4% | 92.0% | --- |
| Llama 3.1 70B | 83.6% | --- | 95.1% | 80.5% | --- |
| Apertus-70B | 69.6% | 62.7% | 77.6% | 73.0% | 80.6% |
| Apertus-8B | 60.9% | 55.7% | 62.9% | 67.0% | 69.5% |

> "Notes on Comparability: The prompt setups differ between models (shot numbers and chain-of-thought configurations). Global-MMLU and RULER values are not available in the official documentation for the comparison models. The 70B variant convinces in general knowledge and multilingual tasks, but remains behind the top models in mathematics and programming."

Source: [effektiv.ch](https://www.effektiv.ch/en/blog/apertus-release), [lifearchitect.ai](https://lifearchitect.ai/models-table/)

Performance comparison
----------------------

| Model | Parameters | Openness | Language Coverage | Training Hardware | Strengths |
| --- | --- | --- | --- | --- | --- |
| Swiss LLM | 8B / 70B | Open Source, Weights, Data | >1,500 | Alps: 10,752 GH200 GPUs | Linguistic diversity, data privacy, transparency |
| GPT-4.5 | ~2T (estimated) | Proprietary | ~80 - 120 | Azure: ~25,000 A100 GPUs | Creativity, natural conversation, agentic planning |
| Claude 4 | Not published | Proprietary | ? | Anthropic: Internal clusters | Adaptive reasoning, coding |
| Llama 4 | 109B / 400B | Open Weight | 12, with 200+ in training | Meta: ~20,000 H100 GPUs | Multimodality, large community, agentic tasks |
| Grok 4 | ~1.8T MoE | Proprietary | ? | Colossus: 200,000 H100 GPUs | Reasoning, real-time data, humor... |

Source: [effektiv.ch](https://www.effektiv.ch/en/blog/swiss-llm)

Sources
-------

For further information:

-   [Swiss AI Initiative](https://www.swiss-ai.org/) (swiss-ai.org)
-   [July Announcement](https://ethz.ch/en/news-and-events/eth-news/news/2025/07/a-language-model-built-for-the-public-good.html) (ethz.ch)
-   [ETH Zurich AI Center](https://ai.ethz.ch/) (ai.ethz.ch)
-   [EPFL Machine Learning Lab](https://www.epfl.ch/labs/mlo/) (epfl.ch)
-   [Apertus Tech Session](https://drive.google.com/file/d/1NYiK648rdX0ypRSfcIcSdnL36suGh9ce/preview) (swiss-ai-weeks.ch)
-   [Can the Swiss LLM Compete?](https://www.effektiv.ch/en/blog/swiss-llm) (effektiv.ch)
-   [First experiences and tips for users of Ollama, LM Studio, vLLM](https://log.alets.ch/110/#using-ollama) (alets.ch)
-   [AlgorithmWatch statement](https://algorithmwatch.ch/de/apartus-eth-epfl/) & [position paper](https://algorithmwatch.ch/en/ai-regulation/)
-   [Alps Supercomputer ranking](https://top500.org/lists/top500/2024/06/) (TOP500.org)
-   [Modern Transformer Design Guide](https://rohitbandaru.github.io/blog/Transformer-Design-Guide-Pt2/) (Rohit Bandaru)
-   [GLU Variants Improve Transformer](https://arxiv.org/abs/2002.05202) (Shazeer 2020)

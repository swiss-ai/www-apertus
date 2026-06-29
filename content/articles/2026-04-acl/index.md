---
title: "Apertus paper at ACL 2026"
date: 2026-04-25T10:00:00+01:00
draft: false
author: "ETH EPFL CSCS"
tags: ["Apertus", "Announcement", "Research"]
featured_image: "apertus-paper.jpg"
summary: "Our technical report has been accepted at a leading conference for AI & NLP."
comments: false
---

##### The Swiss National AI Initiative and the Apertus team are pleased to announce that the technical report underlying the Apertus v1 large language model (LLM) has been accepted for presentation at the ACL 2026 Main Conference, one of the most prestigious venues in the field of AI & Natural Language Processing (NLP). Passing peer review highlights the quality of the scientific work of Apertus, and underscores our commitment to advancing open and compliant AI research.

The accepted paper, titled [Apertus: Democratizing Open and Compliant LLMs for Global Language Environments](https://arxiv.org/abs/2509.14233) (arXiv:2509.14233), presents a step towards improved open source model ecosystems, addressing several fundamental challenges, including:

1. Data Compliance: Many existing models rely on data that is not openly available or does not respect opt-out exclusions, leading to potential legal and ethical issues.
2. Multilingual Representation: Most models focus primarily on English, neglecting the diverse linguistic needs of global users.

To address these limitations, the Apertus team developed a novel approach that:

- Pretrains models exclusively on openly available data, respecting robots.txt exclusions and filtering out toxic and personally identifiable content.
- Adopts the "Goldfish objective" during pretraining to suppress verbatim memorization of data, while maintaining model performance.
- Trains on 15 trillion tokens of text from over 1,800 languages, with approximately 40% of pretraining data allocated to non-English content.

The report introduces two models, both available on the **@swiss-ai** [Hugging Face page](https://huggingface.co/collections/swiss-ai/apertus-llm), each in the Base and Instruct (further trained for chat support) variants, at the 8B- and 70B-scale. All of them are trained on the 15T token datamix with similar training parameters.

The report also includes safety considerations, warning that LLM models such as Apertus may hallucinate and generate unsafe or toxic outputs. Deployment requires additional testing and alignment for specific use cases.

The ACL 2026 invitation is a significant recognition of the Apertus project's technological prowess in global NLP academia. The conference, which accepts [around 20%](https://www.aclweb.org/aclwiki/Conference_acceptance_rates) from thousands of submitted papers, is considered a world-class venue for AI research.

Our accepted paper will be presented at the [ACL 2026 Main Conference](https://2026.aclweb.org/) in San Diego on July 2-7, 2026 - this is where the Apertus team will share their findings and discuss potential applications of the model. The research was already [available as a pre-print](https://arxiv.org/abs/2509.14233) since the model release date, listed along with other relevant publications on our [Research page](/pages/research/).

### About Apertus

Apertus is a fully open suite of large language models designed to address systemic shortcomings in today’s AI model ecosystem. The project is part of the [Swiss National AI Initiative](https://swiss-ai.org), committed to advancing open and responsible AI research. The Apertus team is composed of researchers from leading Swiss institutions, including [ETH Zurich](https://ai.ethz.ch), [EPFL](https://ai.epfl.ch), and Swiss universities in collaboration with engineers from the [Swiss National Supercomputing Centre](https://cscs.ch) (CSCS). For more information and contact details, please [visit our website](/).

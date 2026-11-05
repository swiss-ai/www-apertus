---
title: "Evaluation"
description: "How Apertus models are evaluated, and how to reproduce the results"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-11-05T10:00:00+01:00"
toc: true
tags: ["Technical"]
categories: [""]
author: "Apertus Project"
---

The Swiss AI Initiative is running experiments and training an LLM on the [Alps supercomputer](https://en.wikipedia.org/wiki/Alps_(supercomputer)), operational at CSCS since [September 2024](https://www.netzwoche.ch/news/2024-09-17/neue-forschungsinfrastruktur-supercomputer-alps-eingeweiht). About 50% of the capacity was used in the Apertus 1.0 release. Here are some technical details of the total capacity of Alps:

-   10,752 NVIDIA GH200 Grace-Hopper chips
-   Computing power: 270-435 PFLOPS
-   [Ranked 6th on the TOP500 list](https://top500.org/lists/top500/2024/06/) (June 2024)

For more technical references, see the [Sources](#sources) further below.

### How Apertus models are evaluated

Continuous evaluations are run from our post-training [evaluation repository](https://github.com/swiss-ai/evals-post-train), which orchestrate various benchmarking suites. Four evaluation frameworks run side by side, each with its own folder, entry point and environment:

| Framework | Used for |
|---|---|
| Our fork of [lm-evaluation-harness](https://github.com/swiss-ai/lm-evaluation-harness) | Standard benchmarks (MMLU, GSM8K, ...), task suites, LLM judges |
| [Inspect AI](https://inspect.aisi.org.uk/) | Benchmarks with their own scorer, judge or simulated user (tau2-bench, HLE, ...) |
| [Harbor](https://github.com/harbor-framework/harbor) | Agentic coding and terminal benchmarks (SWE-bench, Terminal-Bench, ...) |
| [EvalScope](https://github.com/modelscope/evalscope) | SWE-bench Verified and Pro |

Models are either loaded in the evaluation job — [vLLM](https://docs.vllm.ai/en/latest/) by default, with SGLang, Hugging Face and Megatron-LM as alternative backends — or evaluated against an already-running OpenAI-compatible endpoint, such as the [CSCS serving API](https://api.swissai.svc.cscs.ch/v1). A dedicated [evaluation service](https://github.com/swiss-ai/evals-svc) orchestrates each run end to end, from requests to the model through scoring with the correct tokenizer and chat template. 

The results can be visualized directly in our dashboard, pushed to [Weights & Biases](https://wandb.ai) for further review, or exported to the [Every Eval Ever](https://github.com/evaleval/every_eval_ever) datastore and Hugging Face (see the Evaluation Cards section below).

Beyond text, the multimodal evaluations (image & audio in Apertus 1.5) were run through the [MLLM-eval-suite](https://github.com/swiss-ai/MLLM-eval-suite) pipeline, which merges two major multimodal evaluation frameworks, [lmms-eval](https://github.com/swiss-ai/lmms-eval) and [VLMEvalKit](https://github.com/swiss-ai/VLMEvalKit). This release also added thinking-mode evaluations, and extended benchmarking for translation, safety, and tool use - this is covered in our [Technical Report](/articles/2026-07-apertus-1-5/).

### Evaluation Cards

Apertus collaborates with the [EvalEval Coalition](https://evalevalai.com) to ensure fully transparent and reproducible benchmarking. We publish the Apertus 1.5 results in the [Every Eval Ever (EEE)](https://github.com/evaleval/every_eval_ever) schema, a shared format that records not only scores, but also who ran the evaluation, how the model was accessed, the generation settings, and how each metric should be interpreted: details that are routinely missing from leaderboards and that often explain why similar-looking scores differ.

The results are shown in [Eval Cards](https://huggingface.co/blog/eee-community-evals), which combine the run data with benchmark and model metadata into one interpretable record, and are also available through Hugging Face Community Evals. This is the same infrastructure the [UK AI Security Institute](https://huggingface.co/blog/evaleval-aisi) uses to publish its evaluations openly. Since our text evaluations run on our public fork of lm-evaluation-harness, which EEE supports with a dedicated converter, anyone can rerun them and compare their numbers against ours under identical settings.

Structured evaluation reporting also matters beyond research. EvalEval has [mapped Evaluation Cards](https://evalevalai.com/infrastructure/2026/09/28/mapping-evaluation-cards-to-ai-governance-requirements/) to AI governance requirements, including Article 55 of the EU AI Act and the GPAI Code of Practice, and found that the EEE schema can represent 56 of the 60 evaluation-reporting elements these frameworks call for.

### Reproducibility

All recipes, ablations, and data mixtures are documented in the [Technical Reports](/pages/research/). We aim to popularize open model development by making the v1.5 training process public and accessible. Intermediate checkpoints are hosted in the [Apertus 1.5 collection](https://huggingface.co/collections/swiss-ai/apertus-v15), including the base weights, SFT, RLVR and DPO, as well as curated training data mixes and custom datasets in the [Apertus 1.5 datasets collection](https://huggingface.co/collections/swiss-ai/apertus-v15-datasets). Sharing them allows researchers to investigate how each stage changes the model behavior, and to reproduce our evaluations.

### Initial benchmarks

See the [Evaluation section](https://huggingface.co/swiss-ai/Apertus-8B-2509#evaluation) of the Apertus Model Card, and Section 5 of the original [Tech Report](https://arxiv.org/abs/2509.14233) for more data. For Apertus 1.5, the [release article](/articles/2026-07-apertus-1-5/) and the Eval Cards described above are the primary references.

Independent benchmarks of Apertus models include:

- [Swiss Legal Evals](https://github.com/JoelNiklaus/SwissLegalEvals) ([results](https://github.com/JoelNiklaus/SwissLegalEvals/blob/main/blog/swiss-legal-evals-2026.md)) — Swiss legal tasks, where Apertus 1.5 is on par with leading open-weight models
- [Last Translation benchmark](https://last-translation-benchmark.vilda.net/leaderboard-results) ([paper](https://huggingface.co/papers/2609.04173)) — machine translation

And others, which we share through our [Zotero group library](https://www.zotero.org/groups/6385576/apertus) (Apertus citations). Please contact us if you are the author of an independent benchmark.

### Sources

- [Apertus 1.5 release article](/articles/2026-07-apertus-1-5/)
- [EvalEval Coalition](https://evalevalai.com)
- [Every Eval Ever](https://github.com/evaleval/every_eval_ever)
- [GPAI Training Transparency](/articles/2026-02-gpai/)
- [Flower Lizzy benchmark](https://flower.ai/blog/2026-04-15-releasing-lizzy7b-model/)
- [Llama-GENBA-10B paper](https://arxiv.org/abs/2509.05668)
- [Apertus Tech Session](https://drive.google.com/file/d/1NYiK648rdX0ypRSfcIcSdnL36suGh9ce/preview) (swiss-ai-weeks.ch)
- [Can the Swiss LLM Compete?](https://www.effektiv.ch/en/blog/swiss-llm) (effektiv.ch)
- [AlgorithmWatch statement](https://algorithmwatch.ch/de/apartus-eth-epfl/) & [position paper](https://algorithmwatch.ch/en/ai-regulation/)
- [Alps Supercomputer ranking](https://top500.org/lists/top500/2024/06/) (TOP500.org)

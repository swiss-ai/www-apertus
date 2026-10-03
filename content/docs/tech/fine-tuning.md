---
title: "Fine-tuning"
description: "Notes on tinkering with Apertus models"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Technical"]
categories: [""]
author: "Apertus Project"
---


Fine-tuning is essential to adapt Apertus to your domain or task, whether you are working on specialized knowledge, improving performance on a certain dataset, or creating a custom application. This is a general class of techniques which add or modify the weights of the LLM, tinkering with the capabilities, and they are used in the development and maintenance of the models themselves. 

Apertus [fine-tuning recipes](https://github.com/swiss-ai/apertus-finetuning-recipes) are available: these are sample configurations for training on a range of hardware levels. In the future we hope to also provide more references from our community.

## Learning Path for Fine-tuning LLMs

Here is a structured learning path from beginner to advanced for finetuning, developed with help from Apertus 1.5:

### 1. Foundations 

_Survey the landscape, obtain resources, and set up your environment._

- Familiarize yourself with the open-source releases of Apertus.
- Verify hardware requirements (suggested VRAM: 8GB+ for 4B, 12GB+ for 8B, 32GB+ for 70B).
- Access the official weights / checkpoints from Hugging Face.
- Clone a template for your fine-tuning experiments (e.g., [apertus-finetuning-recipes](https://github.com/swiss-ai/apertus-finetuning-recipes))
- Download additional datasets from a repository (e.g., [swiss-ai/datasets](https://huggingface.co/swiss-ai/datasets) for the SFT mixture, annotated completion pool, preference pairs, or other datasets curated by the Apertus team)

### 2. Evaluations

_I cannot learn that which I do not understand. I cannot understand what is not measured._

- We use many repeated evaluations to observe the baseline performance gains.
- Read the evaluations section of the Apertus Technical Report.
- Explore benchmark leaderboards covering performance across LLM models.
- Try to decide which tests or evaluation suites are relevant to your use case.
- Install the EleutherAI-based [swiss-ai/lm-evaluation-harness](https://github.com/swiss-ai/lm-evaluation-harness) maintained by the Apertus team.
- Check out [huggingface/lighteval](https://github.com/huggingface/lighteval), among other awesome evaluation tools.

### 3. Techniques

_Develop a post-training pipeline to bring your model forward._

**Supervised Fine-Tuning** (SFT) is part of the second generation of the Apertus post-training pipeline, delivering features like thinking mode, more reliable instruction following, and more effective tool use. Implement SFT using a mixture similar to the one in Apertus 1.5 ([swiss-ai/Apertus-v1.5-SFT-mix](https://huggingface.co/datasets/swiss-ai/Apertus-v1.5-SFT-mix), which focuses on instruction following tasks). Our [finetuning-recipes](https://github.com/swiss-ai/apertus-finetuning-recipes) cover the following SFT modes:

1. LoRA Fine-tuning (for small model on 1 GPU)
2. Full-parameter fine-tuning (for large model on 4+ GPUs)
3. Multi-node training (3 nodes x 4 GPUs)

Further techniques to explore, for which we will soon release more code and instructions:

- **Reinforcement Learning with Verifiable Rewards** (RLVR) is a core post-training technique used by Apertus for improving tool use and reliability. It is often most effective in combination with SFT. A specific method of this is GRPO ([Shao et al., 2024](https://arxiv.org/abs/2402.03300))
- **Direct Preference Optimization** (DPO) preceding online preference optimization was decisive for the post-training of Apertus 1.5.
- **Self-Distillation Policy Optimization** (SDPO) is used for [Apertus Charter](https://apertus-ai.org/pages/charter/) alignment and enables the "thinking mode."
- **Long-Context Extension** enables Apertus to process more data in each prompt.
- **Continued Pre-training (CPT)** adds self-supervised data without restarting training.

### 4. Optimization

_Push performance further on specific benchmarks or integrate capabilities._

- Leverage the released training code to replicate or modify stages of the Apertus development.
- Try to address remaining gaps mentioned in the Conclusion of the tech report, or ideas from other LLM teams.
- Experiment with data filtering and mixed-stage training schedules.
- Use the training code to optimize hyperparameters and long-context handling.
- Use the evaluation harnesses to benchmark progress against released and third-party models (e.g., compare your DPO results against the Overall Average of the current 8B model).
- Publish your results to arXiv, alphaXiv, or even just your own blog. 
- Tag your posts with #Apertus on social media, or [contact us](https://apertus-ai.org/contact) about your results, feedback or questions.

### 5. Safety Considerations

- Avoid training on legally unclear, harmful or sensitive data.
- Do not try to circumvent guardrails like the [Apertus Charter](/pages/charter/).
- Use the [official guidelines](/pages/documentation) for responsible AI development.
- See [Rethinking Safety in LLM Fine-tuning](https://arxiv.org/abs/2508.12531) ... and please share your insights!

## References

For more insights:

- **[Apertus Technical Report](https://www.apertus-ai.org/pages/research/)**
- **Our interpretability journal, [Apertus Claritas](https://apertus-claritas.org)**
- [Introducing Torchforge](https://pytorch.org/blog/introducing-torchforge/) (PyTorch)
- [TRL toolkit for SFT, PEFT/QLoRA ++](https://github.com/huggingface/trl) (Hugging Face)
- [Fine Tune OlmoEarth](https://allenai.github.io/rslearn/examples/FinetuneOlmoEarth/) (AllenAI)
- [Fine-tuning now available for GPT‑4o](https://openai.com/index/gpt-4o-fine-tuning/) (OpenAI)
- [The Ultimate Guide to Fine-Tuning LLMs](https://arxiv.org/html/2408.13296v2) (CeADAR)

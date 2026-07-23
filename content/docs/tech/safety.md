---
title: "Safety"
description: "An overview of LLM safety features, and how they are applied in Apertus"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Technical"]
categories: [""]
author: "Apertus Project"
---

In the development and deployment of LLMs, guardrails are mechanisms used to shape the output of an AI system and ensure appropriate behavior. These controls help manage responses across key domains including topic restrictions, predefined interaction flows, style preferences, structured data extraction, and other content guidelines

In the Apertus LLM, various safety features are being designed to prevent the model from accessing or generating sensitive information, such as personal data, copyrighted content, or toxic material. Guardrails in the training process are based on data filtering, model features like the Goldfish objective, parameter unlearning, and data-influence unlearning to prevent models from accessing or generating sensitive information. 

We strive to create models that are safe and reliable, while also respecting data privacy and ownership rights. This is an active area of research, that in practice very much depends on guardrails being implemented in the system within which the LLM is being deployed. For more information, visit these additional resources:

- [Apertus Research papers](/pages/research/)
- [NVidia NeMo Guardrails](https://docs.nvidia.com/nemo/guardrails/latest/index.html)
- [Meta LlamaFirewall](https://meta-llama.github.io/PurpleLlama/LlamaFirewall/)
- [EvalGuard & alternatives](https://evalguard.ai/alternatives)

Here is an overview of some of these general concepts:

### Data Filtering

During pretraining, Apertus models are trained on openly available data that has been filtered to remove sensitive information. This filtering process respects `robots.txt` exclusions and removes non-permissive, toxic, and personally identifiable content.

### Memorization Prevention

The Goldfish objective is a key component of the Apertus approach. It is designed to suppress verbatim recall of data during pretraining. This means that the model is trained to forget specific pieces of information it has learned, making it less likely to memorize sensitive data.

### Parameter Unlearning

During post-training, Apertus models can be fine-tuned with guardrails to remove unwanted information from their parameters. This process involves identifying and adjusting the specific parameters where the unwanted data is stored, effectively removing it from the model.

### Data-Influence Unlearning

In addition to parameter unlearning, Apertus models can also use data-influence unlearning to degrade the model's perception of the data to be forgotten and reinforce the importance of the remaining data. This approach helps ensure that the model forgets the unwanted information while maintaining its overall functionality and accuracy.

### Model Evaluation

Apertus models are evaluated on their ability to generate safe outputs, striving to ensure that they do not access or produce sensitive information. This evaluation process is an important part of the guardrail mechanism.

### Safety Advisory

The safety advisory in the Apertus Tech Report and Model Card warns that the models may hallucinate and generate unsafe or toxic outputs. This highlights the importance of careful deployment and testing of the models to ensure they are used safely and responsibly.

### Open Source Transparency

The Apertus project is committed to open source and transparency. The guardrail mechanisms are designed to be reproducible and transparent, allowing researchers and developers to understand how they work and apply them in their own projects.

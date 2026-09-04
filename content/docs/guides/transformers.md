---
aliases: ["/docs/guides/apertus-1-0/transformers/"]
title: "Transformers" 
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-09-03T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: ["guides"]
author: "Apertus Project"
---

This document outlines how to use the [Transformers library](https://huggingface.co/docs/transformers/index) from Hugging Face with Apertus.

> 🚧 **Information for Apertus 1.5** - we are currently working on integrating changes from the [latest release](https://huggingface.co/swiss-ai/Apertus-v1.5-8B#how-to-use). Apertus 1.5 accepts interleaved text, image, and audio inputs and generates text. The model does not generate audio or images. The integration is not yet part of a released Transformers version (upstreaming is in progress). Until then, install transformers from our branch:

```bash
pip install "transformers[torch,vision,audio] @ git+https://github.com/swiss-ai/transformers.git@3797303dda74844e3d1f8977ff5518bb91f818b4"
```

## Quickstart

Run a command like this first to install the library, as well as PyTorch, using a package manager:

```bash
pip install -U transformers torch

```

With this sample Python code, you can load and prompt Apertus:

```python
import torch
from transformers import AutoModelForMultimodalLM, AutoProcessor

MODEL_ID = "swiss-ai/Apertus-v1.5-8B"  # or "swiss-ai/Apertus-v1.5-70B"

processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID, dtype="auto", device_map="auto"
).eval()


def generate(messages, max_new_tokens=256, **template_kwargs):
    inputs = processor.apply_chat_template(
        messages,
        add_generation_prompt=True,
        tokenize=True,
        return_dict=True,
        return_tensors="pt",
        **template_kwargs,
    ).to(model.device)
    with torch.inference_mode():
        output_ids = model.generate(**inputs, max_new_tokens=max_new_tokens)
    return processor.decode(
        output_ids[0, inputs["input_ids"].shape[-1]:], skip_special_tokens=True
    )

# prepare the model input
prompt = "Give me a brief explanation of gravity in simple terms."
messages = [
    {"role": "system", "content": "You are a concise and helpful assistant."},
    {"role": "user", "content": prompt}
]
print(generate(messages))
```
```
```

**Tip**:

- Check the model card for instructions on using Image or Audio inputs.
- We recommend setting `temperature=0.8` and `top_p=0.9` in the sampling parameters.

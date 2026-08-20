---
aliases: ["/docs/guides/transformers/"]
title: "Transformers" 
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: [""]
author: "Apertus Project"
---

This document outlines how to use the [Transformers library](https://huggingface.co/docs/transformers/index) with Apertus.

_We are currently working on integrating changes for the Apertus 1.5 release. Please stay tuned for updated instructions here._

Run a command like this first to install the library using a package manager:

```bash
pip install -U transformers
```

With this sample Python code, you can load and prompt Apertus in the library from Hugging Face:

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

model_name = "swiss-ai/Apertus-8B-Instruct-2509"
device = "cuda"  # for GPU usage or "cpu" for CPU usage

# load the tokenizer and the model
tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForCausalLM.from_pretrained(
    model_name,
).to(device)

# prepare the model input
prompt = "Give me a brief explanation of gravity in simple terms."
messages_think = [
    {"role": "user", "content": prompt}
]

text = tokenizer.apply_chat_template(
    messages_think,
    tokenize=False,
    add_generation_prompt=True,
)
model_inputs = tokenizer([text], return_tensors="pt", add_special_tokens=False).to(model.device)

# Generate the output
generated_ids = model.generate(**model_inputs, max_new_tokens=32768)

# Get and decode the output
output_ids = generated_ids[0][len(model_inputs.input_ids[0]) :]
print(tokenizer.decode(output_ids, skip_special_tokens=True))
```

**Tip**:

> We recommend setting `temperature=0.8` and `top_p=0.9` in the sampling parameters.

---
title: "Featherless.ai" 
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-09-11T11:11:00+01:00"
lastmod: "2026-09-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: ["guides"]
author: "Apertus Project"
---

Featherless.ai launched the [Open world](https://featherless.ai/openworld) campaign in 2026, hosting various versions of [Apertus models](https://featherless.ai/models?v=c&query=apertus) on their platform. 

They have generously provided starter keys for recent hackathons at the EPFL and ETHZ - let us know if you need extra access for your event.


The following instructions are based on the [Quickstart guide](https://featherless.ai/docs/quickstart-guide) in the official Featherless.ai documentation. See also the [Zero to AI blog post](https://featherless.ai/blog/zero-to-ai-deploying-language-models-without-the-infrastructure-headache).

> Featherless AI is a serverless AI inference platform. Our goal is to make all AI models available for serverless inference and we’ve started with large language models (e.g. Qwen, Llama, Mistral, DeepSeek, RWKV). We provide inference via API to a continually expanding library of open-weight models, including the most popular models for role-playing, creative writing, coding assistance, and more.

![Screenshot from Featherless AI](/images/guides/featherless-model.png)

The API interface is OpenAI compatible, meaning any client program that works with OpenAI as an AI/inference provider can be reconfigured with little effort:

1.  Sign up for an account at [Featherless](https://featherless.ai/register)
2.  Get your API key from the [dashboard](https://featherless.ai/account/api-keys)
3.  Search for the Apertus model you want to use on the [model page](https://featherless.ai/models?v=c&query=apertus)
4.  Make your first call using the OpenAI or Featherless API as in the snippets below.

## OpenAI SDK

```py
from openai import OpenAI

client = OpenAI(
  base_url="https://api.featherless.ai/v1",
  api_key="FEATHERLESS_API_KEY",
)

response = client.chat.completions.create(
  model='swiss-ai/Apertus-v1.5-70B',
  messages=[
    {"role": "system", "content": "You are Apertus, a helpful assistant."},
    {"role": "user", "content": "Hello!"}
  ],
)
print(response.choices[0].message.content)
```

## Featherless API (direct)

You can also make direct requests to our API endpoints (most important of which being [/completions and /chat/completions](https://featherless.ai/docs/completions)) to integrate Featherless into any software application.

```py
import requests

response = requests.post(
    url="https://api.featherless.ai/v1/chat/completions",
    headers={
        "Content-Type": "application/json",
        "Authorization": "Bearer FEATHERLESS_API_KEY"
    },
    json={
        "model": "swiss-ai/Apertus-v1.5-70B",
        "messages": [
            {"role": "system", "content": "You are Apertus, a helpful assistant."},
            {"role": "user", "content": "Hello! How are you?"}
        ]
    }
)
print(response.json()["choices"][0]["message"]["content"])
```

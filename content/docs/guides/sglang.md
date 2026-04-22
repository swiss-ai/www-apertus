---
title: "SGlang"
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: [""]
author: "Apertus Project"
---

**SGLang** is a high-performance [open source](https://github.com/sgl-project/sglang) serving framework for large language models and multimodal models. It is designed to deliver low-latency and high-throughput inference across a wide range of setups, from a single GPU to large distributed clusters. 

It is recommended by the engineering team of Apertus, and our [Scratchpad repo](https://github.com/swiss-ai/Scratchpad
) has an example configuration. You can also find usage guidelines from [Hugging Face](https://huggingface.co/docs/inference-endpoints/main/en/engines/sglang).

Please visit the [official website](https://sglang.io/) and [documentation](https://docs.sglang.io/) for further deployment instructions. For practical examples and tutorials, open the [SGLang quickstart](https://docs.sglang.io/quickstart). 

![Architecture diagram](https://wilsonwu.me/en/blog/2025/getting-started-with-sglang/1-sglang-architecture.png)
_Diagram from [Getting started with SGLang](https://wilsonwu.me/en/blog/2025/getting-started-with-sglang/) by Wilson Wu_


### Overview

Key features of SGLang include:

- **Efficient Backend Runtime**: Uses RadixAttention, prefix caching, and other optimizations to speed up inference. This reduces latency and increases throughput, making it ideal for production environments.
- **Flexible Frontend Language**: Offers intuitive APIs for generating text, controlling model behavior, and managing complex prompts. This allows developers to craft detailed, context-dependent responses without diving into low-level details.
- **Model Compatibility**: Supports a wide range of models (e.g., Llama, Gemma, Qwen), making it easy to experiment with different set-ups.
- **Open-Source and Community-Driven**: Backed by a growing community, ensuring active development and support. This means frequent updates, new features, and access to community examples and troubleshooting resources.

For LLM developers, SGLang streamlines the deployment process with:

1. **Quick Deployment**: With minimal configuration, you can set up a server to serve a model, leveraging your GPU resources efficiently.
2. **Customizable**: Provides control over model parameters (e.g., max context length, temperature, batch size) to fine-tune performance and behavior.
3. **Streaming Support**: Ideal for real-time applications like chatbots or live interactions, allowing continuous generation without blocking.
4. **Production-Ready**: Supports scaling, concurrency, and integration with cloud platforms (e.g., via Kubernetes), making it suitable for large-scale use cases.

For production use, consider deploying SGLang with Kubernetes, e.g., using the [OME operator](https://github.com/sgl-project/ome) or [SkyPilot](https://docs.skypilot.co/en/latest/examples/serving/sglang.html) for cloud-scale setups.


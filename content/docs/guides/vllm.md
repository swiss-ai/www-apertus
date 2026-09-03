---
aliases: ["/docs/guides/apertus-1-0/vllm/"]
title: "vLLM"
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: ["guides"]
author: "Apertus Project"
---

vLLM is a community-driven, efficient library for deploying large language models (LLMs) for inference and serving. Initially developed by the Sky Computing Lab at UC Berkeley, vLLM has evolved through contributions from academia and industry. While the library itself is user-friendly, deployment can be challenging due to its reliance on NVIDIA libraries and CUDA tools. However, for IT teams using DevOps tools like Kubernetes, vLLM’s multiplatform support and integration capabilities are significant advantages.

For more information visit the [vLLM website](https://vllm.ai), where you will find the official documentation.

Support work for vLLM compatibility is taking place in https://github.com/swiss-ai/vllm

> **Information for Apertus 1.5** - we are currently working on integrating changes from the latest release. Please stay tuned for updated instructions here. See the [vLLM Omni Patch](#omni-patch) section below for early adopter details.

# Quickstart

We provide you with a quick script to set up Apertus in your environment with vLLM, which assumes you have:

- Docker and Docker Compose installed
- NVIDIA drivers and CUDA toolkit - make adjustments for AMD or other accelerators

This example assumes you are using the [Apertus 1.5 8B model](https://huggingface.co/swiss-ai/Apertus-v1.5-8B). We provide a simple `docker-compose.yml` for a basic setup. To configure your deployment, you will probably want to set the following environment variables:

- **`HF_TOKEN`**: A Hugging Face API token with "Read access to contents of all repos you can access". Generate this in your [Hugging Face settings](https://huggingface.co/settings/token). You can also put this into an environment (`.env`) file.
- **`HF_MODEL`**: Set to `swiss-ai/Apertus-v1.5-8B` or your preferred version of the Apertus model.

Example `docker-compose.yml` for the Nightly (not for production!) build of vLLM:

```yaml
version: '3'
services:
  vllm:
    image: ghcr.io/vllm/vllm:nightly
    environment:
      HF_TOKEN: $HF_TOKEN
      HF_MODEL: $HF_MODEL
      # Optional: Adjust max-model-len based on your GPU VRAM
      # Set to 8K for better results on smaller GPUs (default: 64K)
      MAX_MODEL_LEN: 4096  # 4K tokens (adjust based on your system)
    ports:
      - "8000:8000"  # Expose port 8000 for the default API
    volumes:
      - ./models:/app/models  # Optional: mount local model directory
```

The `MAX_MODEL_LEN` option controls the maximum input length of the model. For a 16GB GPU, 4K tokens might be more practical. The default 64K requires more VRAM.

Start your service as follows:

```bash
docker-compose up -d
```

Navigate to `http://localhost:8000` in your browser (or the server’s public IP if deployed elsewhere). 

You will then need to configure an API key to interact with the model from software like OpenWebUI. See the vLLM documentation for details. You can also find an official guide to using vLLM with Docker in the [vLLM documentation](https://docs.vllm.ai/en/stable/deployment/docker/).

#### Tips

- The `swiss-ai/Apertus-v1.5-8B` model is optimized for smaller GPUs, you may wish to choose one of the other model alternatives depending on your hardware and needs.
- **Software Licensing**: If using CUDA acceleration, ensure you comply with NVIDIA’s licensing terms for using their proprietary libraries.
- **Stored Secrets**: Ensure the `HF_TOKEN` is not exposed in your `.env` file or Docker environment variables. Use a secrets manager if possible.
- **Performance Tuning**: Experiment with different `MAX_MODEL_LEN` values for your hardware.
- **Kubernetes**: Explore vLLM’s Kubernetes support for scaling and managing multiple instances.
- **Custom Environments**: Use `docker-compose.yml` to mount local models or data directories for custom use cases.
- **Contributing**: Share feedback or PRs on the [Apertus model-specific forum](https://huggingface.co/swiss-ai).

## Troubleshooting

- **CUDA Version**: Check if your version is compatible with vLLM (CUDA 12+ recommended).
- **VRAM Issues**: If you encounter memory errors, reduce `MAX_MODEL_LEN` or upgrade your GPU.
- **Docker Errors**: Verify your Docker setup and ensure permissions are correct.
- **Something else**: Search through [GitHub issues](https://github.com/vllm/vllm/issues) and start a new one if you find a new bug.
- **Get Help**: For production use, evaluate your specific needs and adjust configurations according to enterprise guidelines, and potentially seek professional support of the vLLM team.

## vLLM Omni Patch

The Apertus 1.5 release has not been completely merged in the Transformers and vLLM libraries. In particular to use multimodal features and tool calling capabilities, a custom configuration (i.e. Dockerfile) is recommended.

The Apertus team has provided vLLM plugins until the Apertus logic is folded into the core vLLM release. You can find these in the [apertus-omni-tokenizer](https://github.com/swiss-ai/apertus-omni-tokenizer/blob/main/parsers/vllm/README.md) repository, which has documentation for vLLM.

Here is an overview of the compatibility by version:

| vLLM version range | What’s in vLLM already | Should you **use the plugin** from here? |
| --- | --- | --- |
| **≤ v0.10.2** (pre-refactor / older layouts) | Only special-token reasoning delimiters, no `apertus` parser | ❌ — plugin doesn’t fit |
| **v0.19 ‑ v0.21** | Layout supported by plugin but **Parser not upstreamed yet** | ✅ Only way |
| **v0.22 ‑ v0.24** | Upstream finally ships identical parser **without the plugin registration line** | ✅ Only if you need to <br>• freeze parser version <br>• iterate without freq vLLM upgrades <br>• hot-fix a production parser bug <br>• use reasoning-tracking independent of engine patch |
| **v0.24+ w/ same layout** | Same as v0.22+ — upstream parser exists | ✅ only if you need any of the bullets above (otherwise just `--tool-call-parser apertus --reasoning-parser apertus`) |

Start vllm from the commandline with the `apertus` parsers enabled like this:

```bash
# Built-in on v0.22+
vllm serve ... --tool-call-parser apertus --reasoning-parser apertus

# Plugin version on any compatible range
vllm serve ... --tool-call-parser apertus --tool-parser-plugin ./plugins/apertus_tool_parser.py \
               --reasoning-parser apertus --reasoning-parser-plugin ./plugins/apertus_reasoning_parser.py
```

If you are experiencing issues with tools (MCP), we recommend explicitly disabling thinking mode with this parameter:

`--default-chat-template-kwargs.enable_thinking true`

Please visit our [README](https://github.com/swiss-ai/apertus-omni-tokenizer/blob/main/parsers/vllm/README.md) for more deteails, and [contact us](/contact) if you have questions.

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

vLLM is a community-driven, efficient library for deploying large language models (LLMs) for inference and serving. Initially developed by the Sky Computing Lab at UC Berkeley, vLLM has evolved through contributions from academia and industry. While the library itself is user-friendly, deployment can be challenging due to its reliance on NVIDIA libraries and CUDA tools. For IT teams using DevOps tools like Kubernetes, vLLM’s multiplatform support and integration capabilities are significant advantages.

For more information visit the [vLLM website](https://vllm.ai), where you will find the official documentation.
Support work for vLLM compatibility with Apertus is taking place in https://github.com/swiss-ai/vllm

> **Information for Apertus 1.5** - We are currently working on adding support for our models to upstream vLLM and transformers releases. In the meantime, you can use our modified version of [vLLM](https://github.com/swiss-ai/vllm/tree/a601a9d998ddeb488f0c17e8512874b116aa7658) and [transformers](https://github.com/swiss-ai/transformers/tree/3797303dda74844e3d1f8977ff5518bb91f818b4) to run the models. We have pre-installed the dependencies in a [Docker image](https://github.com/swiss-ai/model-launch/pkgs/container/vllm_apertus_1.5_release) which is available in the GitHub Container Registry. See the [Thinking Mode](#thinking-mode) and [Omni Patch](#omni-patch) sections below for additional early adopter details.

# Quickstart

We provide you with a quick example script to set up Apertus in your environment with Docker. Information for users of [Kubernetes](https://docs.vllm.ai/en/latest/deployment/k8s/) and other platforms is available in the [vLLM documentation](https://docs.vllm.ai/en/latest/).

You may also need NVIDIA drivers and CUDA toolkit - make adjustments for Intel, Mac, or other accelerators using the appropriate [Pre-built images](https://docs.vllm.ai/en/latest/deployment/docker/#pre-built-images). 

The source Dockerfile used to build the image is available in [this repository](https://github.com/swiss-ai/model-launch/blob/main/images/vllm_apertus_1.5_release/Dockerfile).
You can pull the image with the following command:

- amd64 architecture:

```bash
docker pull ghcr.io/swiss-ai/vllm_apertus_1.5_release:latest-amd64
```

- arm64 architecture:

```bash
docker pull ghcr.io/swiss-ai/vllm_apertus_1.5_release:latest-arm64
```

You can use the following commands to run the models with vLLM:

- 8B model:

```bash
vllm serve swiss-ai/Apertus-v1.5-8B \
--chat-template-content-format string \
--gpu-memory-utilization 0.6 \
--max-model-len 262144 \
--enable-auto-tool-choice \
--tool-call-parser apertus
```

- 70B model:

```bash
vllm serve swiss-ai/Apertus-v1.5-70B \
--chat-template-content-format string \
--tensor-parallel-size 4 \
--gpu-memory-utilization 0.8 \
--max-model-len 262144 \
--enable-auto-tool-choice \
--tool-call-parser apertus
```

Depending on your hardware, you may need to adjust `--tensor-parallel-size`, `--gpu-memory-utilization`, and `--max-model-len` (e.g. lower `--max-model-len` if you run out of memory). On some hardware configurations, CUDA Graph capture may fail with `--tensor-parallel-size > 1` due to the fused all-reduce RMS optimization. If this occurs, launch vLLM with `--compilation-config.pass_config.fuse_allreduce_rms false`.

### How to use

This server can be queried in the same format as OpenAI API. For example, to list the models:

```bash
curl http://localhost:8000/v1/models
```

You can pass in the argument `--api-key` or environment variable `VLLM_API_KEY` to enable the server to check for API key in the header. See the vLLM documentation for details. You can also find an official guide to using vLLM with Docker in the [vLLM documentation](https://docs.vllm.ai/en/stable/deployment/docker/), as well as many more [deployment examples](https://docs.vllm.ai/en/latest/examples/).

## Tips

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

### Thinking Mode

To enable thinking mode, set `--reasoning-parser` and `--default-chat-template-kwargs.enable_thinking` as shown below. The tool-call flags are intentionally omitted: tool calling is unsupported in thinking mode, so we don't recommend combining the two.

- 8B model:

```bash
vllm serve swiss-ai/Apertus-v1.5-8B \
--served-model-name swiss-ai/Apertus-v1.5-8B-thinking \
--chat-template-content-format string \
--gpu-memory-utilization 0.6 \
--max-model-len 262144 \
--reasoning-parser apertus \
--default-chat-template-kwargs.enable_thinking true
```

- 70B model:

```bash
vllm serve swiss-ai/Apertus-v1.5-70B \
  --served-model-name swiss-ai/Apertus-v1.5-70B-thinking \
  --chat-template-content-format string \
  --tensor-parallel-size 4 \
  --gpu-memory-utilization 0.8 \
  --max-model-len 262144 \
  --reasoning-parser apertus \
  --default-chat-template-kwargs.enable_thinking true
```

### Omni Patch

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

You can enable thinking mode (please note: currently not recommended in combination with tools) with this parameter:

`--default-chat-template-kwargs.enable_thinking true`

Please visit our [README](https://github.com/swiss-ai/apertus-omni-tokenizer/blob/main/parsers/vllm/README.md) for more details on the above, and [contact us](/contact) if you have questions.

---
aliases: ["/docs/guides/vllm/"]
title: "vLLM"
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: [""]
author: "Apertus Project"
---

vLLM is a community-driven, efficient library for deploying large language models (LLMs) for inference and serving. Initially developed by the Sky Computing Lab at UC Berkeley, vLLM has evolved through contributions from academia and industry. While the library itself is user-friendly, deployment can be challenging due to its reliance on NVIDIA libraries and CUDA tools. However, for IT teams using DevOps tools like Kubernetes, vLLM’s multiplatform support and integration capabilities are significant advantages.

For more information visit the [vLLM website](https://vllm.ai), where you will find the official documentation.

# Quickstart

We provide you with a quick script to set up Apertus in your environment with vLLM, which assumes you have:

- Docker and Docker Compose installed
- NVIDIA drivers and CUDA toolkit (CUDA 11.8 or later recommended) - make adjustments for AMD or other accelerators

This example assumes you are using the [Apertus 8B Instruct model](https://huggingface.co/swiss-ai/Apertus-8B-Instruct-2509). We provide a simple `docker-compose.yml` and `nginx.conf` for a basic setup. 

To configure your deployment, you must set the following environment variables:

- **`HF_TOKEN`**: A Hugging Face API token with "Read access to contents of all repos you can access". Generate this in your [Hugging Face settings](https://huggingface.co/settings/token).
- **`HF_MODEL`**: Set to `swiss-ai/Apertus-8B-Instruct-2509` or your preferred model.

Example `docker-compose.yml`:

```yaml
version: '3'
services:
  vllm:
    image: ghcr.io/vllm/vllm:nightly
    environment:
      HF_TOKEN: YOUR_HF_TOKEN_HERE
      HF_MODEL: swiss-ai/Apertus-8B-Instruct-2509
      # Optional: Adjust max-model-len based on your GPU VRAM
      # Set to 8K for better results on smaller GPUs (default: 64K)
      MAX_MODEL_LEN: 4096  # 4K tokens (adjust based on your system)
    ports:
      - "8000:8000"  # Expose port 8000 for Open WebUI
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro  # Mount NGINX config
      - ./models:/app/models  # Optional: mount local model directory
    depends_on:
      - nginx
  nginx:
    image: nginx:latest
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
    ports:
      - "80:80"
```

**Note**: The `MAX_MODEL_LEN` option controls the maximum input length of the model. For a 16GB GPU, 4K tokens might be more practical. The default 64K requires more VRAM.

The provided `nginx.conf` is a basic configuration to route traffic to vLLM. Ensure it’s placed in the same directory as `docker-compose.yml`.

Start your service as follows:

```bash
docker-compose up -d
```

Navigate to `http://localhost:8000` in your browser (or the server’s public IP if deployed elsewhere). 

#### Tips

You will need to configure an API key to interact with the model from software like OpenWebUI.

The `swiss-ai/Apertus-8B-Instruct-2509` model is optimized for smaller GPUs, you may wish to choose one of the other model alternatives depending on your hardware and needs.

If using CUDA acceleration, ensure you comply with NVIDIA’s licensing terms for using their proprietary libraries.

Ensure the `HF_TOKEN` is not exposed in your `.env` file or Docker environment variables. Use a secrets manager if possible.

#### Optimization

- **Performance Tuning**: Experiment with different `MAX_MODEL_LEN` values for your hardware.
- **Kubernetes**: Explore vLLM’s Kubernetes support for scaling and managing multiple instances.
- **Custom Environments**: Use `docker-compose.yml` to mount local models or data directories for custom use cases.
- **Contributing**: Share feedback or PRs on the [Apertus model-specific forum](https://huggingface.co/swiss-ai).

### Troubleshooting

- **CUDA Version**: Check if your version is compatible with vLLM (CUDA 11.8+ recommended).
- **VRAM Issues**: If you encounter memory errors, reduce `MAX_MODEL_LEN` or upgrade your GPU.
- **Docker Errors**: Verify your Docker setup and ensure permissions are correct.
- **Something else**: Search through [GitHub issues](https://github.com/vllm/vllm/issues) and start a new one if you find a new bug.

By following these steps, you can quickly deploy a vLLM instance with Open WebUI using Apertus on a GPU. For production use, evaluate your specific needs and adjust configurations according to enterprise guidelines, and potentially seek professional support of the vLLM team.

---
title: "Google Vertex AI" 
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-08-11T11:11:00+01:00"
lastmod: "2026-08-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: ["guides"]
author: "Apertus Project"
draft: true
---

_THIS DOCUMENT IS A WORK IN PROGRESS_

Deploying Apertus on Google Cloud [Vertex AI Model Garden](https://cloud.google.com/vertex-ai/model-garden) requires using a customized vLLM container image. Since Apertus 1.5 introduces a novel multimodal architecture for native image and audio processing, standard pre-built Model Garden containers may not yet fully support its capabilities.

## 1. Hugging Face Access

Generate a Hugging Face Access Token with **Read** permissions to download the weights from your container.

Note: some versions of the Apertus models have gated access. You may need to accept the usage terms on the model card before the container can download the weights. 

## 2. Build a Custom vLLM Container

Vertex AI requires specific health and prediction routes. Custom containers can handle requests using the recommended structure. You must wrap the latest multimodal-capable vLLM release in a FastAPI server that routes `/health` for health checks and `/predict` for inferences.

- Create a repository in Google Cloud Artifact Registry:

```bash
gcloud artifacts repositories create vllm-repo\
  --repository-format=docker\
  --location=us-central1
```

- Write a Dockerfile that installs `vllm` and copies in your FastAPI server script (TODO: share an example)
- Next, build and push the image using Cloud Build:

```bash
gcloud builds submit --tag us-central1-docker.pkg.dev/YOUR_PROJECT/vllm-repo/vllm-apertus:latest
```

## 3. Register the Model in Vertex AI

Use the `google-cloud-aiplatform` Python SDK to upload your model configuration to the Model Registry. You will need to pass your Hugging Face token securely as an environment variable so the container can download the model at startup.

```python
from google.cloud import aiplatform

aiplatform.init(project="YOUR_PROJECT", location="us-central1")

model = aiplatform.Model.upload(
    display_name="apertus-1-5-8b",
    serving_container_image_uri="us-central1-docker.pkg.dev/YOUR_PROJECT/vllm-repo/vllm-apertus:latest",
    serving_container_environment_variables={
        "HF_TOKEN": "YOUR_HUGGINGFACE_TOKEN"
    },
    serving_container_command=[
        "python3", "your_fastapi_server.py",
        "--model", "swiss-ai/Apertus-v1.5-8B",
        "--trust-remote-code"
    ]
)
```

## 4. Create an Endpoint and Deploy

Create an online prediction endpoint and deploy the model with an appropriate resource allocation. The 8B model fits on a single NVIDIA A100 GPU, but if you are deploying the 70B version or leveraging Apertus's full 262,144 token context window, you will need multiple GPUs with at least 140 GB RAM - significantly less for quantized versions, such as NPV4 (TODO: link to example builds).

```python
endpoint = aiplatform.Endpoint.create(display_name="apertus-endpoint")

model.deploy(
    endpoint=endpoint,
    machine_type="a2-highgpu-1g",
    accelerator_type="NVIDIA_TESLA_A100",
    accelerator_count=1,
    deploy_request_timeout=1800
)
```

## 5. Access your Endpoint via API

Use the Google Vertex AI interface, or external tools to query your new edpoint.

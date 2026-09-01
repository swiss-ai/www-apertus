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
---



Deploying Apertus 1.5 on Google Cloud Vertex AI Model Garden requires using a customized vLLM container image. Because Apertus 1.5 introduces a novel multimodal architecture for native image and audio processing, standard pre-built Model Garden containers may not yet fully support its capabilities.

**1.Acquire Hugging Face Access:**Apertus models are gated.

You must accept the model's usage license before the container can download the weights.

1.  Navigate to the `swiss-ai/Apertus-v1.5-8B` or `swiss-ai/Apertus-v1.5-70B` repository on Hugging Face.

2.  Accept the model terms.

3.  Generate a Hugging Face Access Token with **Read** permissions.

**2.Build and Push a Custom vLLM Container:**Vertex AI requires specific health and prediction routes.

Vertex AI expects custom containers to handle requests using a specific structure. You must wrap the latest multimodal-capable vLLM release in a FastAPI server that routes `/health` for health checks and `/predict` for inferences.

1.  Create a repository in Google Cloud Artifact Registry:

Bash

```
gcloud artifacts repositories create vllm-repo\
  --repository-format=docker\
  --location=us-central1

```

1.  Write a Dockerfile that installs `vllm` and copies in your FastAPI server script.

2.  Build and push the image using Cloud Build:

Bash

```
gcloud builds submit --tag us-central1-docker.pkg.dev/YOUR_PROJECT/vllm-repo/vllm-apertus:latest

```

**3.Register the Model in Vertex AI:**Link the custom image to the Hugging Face weights.

Use the `google-cloud-aiplatform` Python SDK to upload your model configuration to the Model Registry. You will need to pass your Hugging Face token securely as an environment variable so the container can download the model at startup.

Python

```
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

**4.Create an Endpoint and Deploy:**Allocate appropriate GPU hardware.

Create an online prediction endpoint and deploy the model. The 8B model fits on a single NVIDIA A100 GPU, but if you are deploying the 70B version or leveraging Apertus's full 262,144 token context window, you will need multiple GPUs.

Python

```
endpoint = aiplatform.Endpoint.create(display_name="apertus-endpoint")

model.deploy(
    endpoint=endpoint,
    machine_type="a2-highgpu-1g",
    accelerator_type="NVIDIA_TESLA_A100",
    accelerator_count=1,
    deploy_request_timeout=1800
)

```

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

For LLM developers on the Apertus team, SGLang streamlines the deployment process with quick server setup, leveraging GPU resources efficiently. Detailed control over model parameters (e.g., max context length, temperature, batch size) allows us to fine-tune performance and behavior.

For production use, consider deploying SGLang with Kubernetes, e.g., using the [OME operator](https://github.com/sgl-project/ome) or [SkyPilot](https://docs.skypilot.co/en/latest/examples/serving/sglang.html) for cloud-scale setups. You can use the Python SDK or API tools like cURL to interact with SGLang, maintaining compatibility with existing workflows.

To get started with SGLang:

1. **Install SGLang**:  
   ```  
   pip install sglang huggingface-hub  
   ```  

2. **Start the Service**:  
   - **Via CLI**:  
     ```  
     python -m sglang.launcher --model swiss-ai/Apertus-8B-Instruct-2509 --host 0.0.0.0 --port 30000 --tp 1 --max-model-len 8192  
     ```  
   - **Or with `sglang` command (if supported by your version):**  
     ```  
     sglang serve -m swiss-ai/Apertus-8B-Instruct-2509 --host 0.0.0.0 --port 30000  
     ```  

3. **Test Inference**:  
   - **Example with cURL**:  
     ```bash  
     curl http://127.0.0.1:30000/v1/chat/completions\  
       -H "Content-Type: application/json"\  
       -H "Authorization: Bearer EMPTY"\  
       -d '{"model": "swiss-ai/Apertus-8B-Instruct-2509",  
            "messages": [  
              {"role": "system", "content": "You are a helpful assistant."},  
              {"role": "user", "content": "Explain SGLang in 3 sentences."}  
            ],  
            "max_tokens": 128,  
            "temperature": 0.2  
          }'  
     ```  

For more advanced use cases, explore the [official documentation](https://sglang.io/docs) and community resources.

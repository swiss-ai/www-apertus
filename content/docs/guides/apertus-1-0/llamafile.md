---
aliases: ["/docs/guides/llamafile/"]
title: "Llamafile" 
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-06-11T11:11:00+01:00"
lastmod: "2026-06-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: [""]
author: "Apertus Project"
---


[Llamafiles](https://docs.mozilla.ai/llamafile) are open-source AI models packaged as a single file, that runs on your laptop, offline.

Built on top of [llama.cpp](https://github.com/ggerganov/llama.cpp), this enables **offline, cross-platform inference** with minimal setup: no Python, CUDA, or complex dependencies required. A Llamafile bundles:
- The model weights (quantized for efficiency).
- The inference engine (llama.cpp).
- A built-in server for serving the model via a REST API.

### Using Apertus Models with Llamafile

Pre-built Llamafiles for the Apertus 8B model, and soon for smaller versions, are available here:
- [Pre-built llamafiles](https://docs.mozilla.ai/llamafile/getting-started/pre-built-llamafiles)

Example (Linux/macOS):
```bash
# Download a Llamafile
curl -L https://huggingface.co/mozilla-ai/llamafile_0.10/resolve/main/Apertus-8B-Instruct-2509.llamafile -o apertus.llamafile
chmod +x apertus.llamafile
```

Execute the Llamafile directly:
```bash
./apertus.llamafile

```
This starts a local server (default: `http://localhost:8080`). Interact via:
- **CLI**: `curl http://localhost:8080/completions -d '{"prompt": "Your text here"}'`
- **API**: Send POST requests to the `/completions` endpoint (OpenAI-compatible format).

### Custom Builds (Advanced)

To create a Llamafile from an Apertus model in GGUF format:

```bash
# Install llamafile (requires Docker or Podman)
curl -L https://github.com/Mozilla-Ocho/llamafile/releases/download/0.8.13/llamafile -o llamafile
chmod +x llamafile

# Convert a GGUF model to Llamafile
./llamafile --model /path/to/apertus-8b.gguf --out apertus.llamafile
```

See our [ollama page](/docs/guides/apertus-1-0/ollama/) for some recommended GGUF builds.

### **Resources**

- [Llamafile GitHub](https://github.com/Mozilla-Ocho/llamafile)
- [llama.cpp](https://github.com/ggerganov/llama.cpp) (underlying engine)
- [Simplon-Off](https://simplon-off.ch/), a friendly guide to Llamafiles made at a recent hackathon in Switzerland.


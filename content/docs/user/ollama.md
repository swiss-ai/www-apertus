---
aliases: ["/docs/guides/apertus-1-0/ollama/"]
title: "Ollama" 
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-09-02T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: ["guides"]
author: "Apertus Project"
---

Ollama is a tool that lets you run and interact with open source large language models (LLMs) on your local machine.

Models downloaded from the [Ollama library](https://ollama.com/library) can be configured and managed most easily with an elegant chat interface. Community support is available from contributors in our network:

- [Michel Rosselli](https://ollama.com/MichelRosselli/apertus-v1.1) (v1.1)
- [Michel Rosselli](https://ollama.com/MichelRosselli/apertus) (v1.0 8B, 70B)

GGUF (a [binary format](https://github.com/ggml-org/ggml/blob/master/docs/gguf.md) for models) support for Apertus is available through community remixes. We have verified Apertus releases by:

- [Andreas Martin](https://huggingface.co/andreasmartin/apertus-v1.5-8b-text-Q8_0-GGUF) (v1.5 8B)
- [AI Singapore](https://huggingface.co/aisingapore/Apertus-SEA-LION-v4-8B-IT-GGUF) (SEA-LION 8B)
- [Bartowski](https://huggingface.co/bartowski/swiss-ai_Apertus-8B-Instruct-2509-GGUF) (v1.0 8B, 70B)
- [redponike](https://huggingface.co/redponike/Apertus-8B-Instruct-2509-GGUF) (v1.0 8B, 70B)
- [Unsloth AI](https://huggingface.co/unsloth/Apertus-8B-Instruct-2509-GGUF) (v1.0 8B, 70B)

To get started, just run a command like this in your console where you have Ollama installed:

`ollama pull MichelRosselli/apertus-v1.1`

Please note that we cannot support community builds directly. Feel free to share your experiences in the discussion areas on Hugging Face.

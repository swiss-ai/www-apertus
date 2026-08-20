---
aliases: ["/docs/guides/ollama/"]
title: "Ollama" 
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: [""]
author: "Apertus Project"
---

Ollama is a tool that lets you run and interact with open source large language models (LLMs) on your local machine. 

Models downloaded from the [Ollama library](https://ollama.com/library) can be configured and managed most easily with an elegant chat interface. We are working to make Apertus available here in the future.

As of the 0.12.6 release of Ollama, GGUF (a [binary format](https://github.com/ggml-org/ggml/blob/master/docs/gguf.md) for models) support for Apertus is available through community remixes. We have tested the following 8B releases by:

- [bartowski](https://huggingface.co/bartowski/swiss-ai_Apertus-8B-Instruct-2509-GGUF)
- [redponike](https://huggingface.co/redponike/Apertus-8B-Instruct-2509-GGUF) 
- [unsloth](https://huggingface.co/unsloth/Apertus-8B-Instruct-2509-GGUF)

To get started, just run a command like this in your console where you have Ollama installed:

`ollama pull hf.co/bartowski/swiss-ai_Apertus-8B-Instruct-2509-GGUF:Q4_K_M`

Feel free to share your experiences with the community in the [discussion area](https://huggingface.co/swiss-ai/Apertus-8B-Instruct-2509/discussions/3).

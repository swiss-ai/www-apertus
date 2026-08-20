---
aliases: ["/docs/guides/openwebui/"]
title: "Open WebUI"
description: "Instructions for users of Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Users"]
categories: [""]
author: "Apertus Project"
---

[**Open WebUI**](https://github.com/open-webui/open-webui) is a self-hosted, ChatGPT-style interface. As long as a serving API is OpenAI-compatible, you can point a local Open WebUI instance at a cloud endpoint and chat with models like Apertus. This is a friendly UI that lets you chat with your models without needing to know the details of the API. It is also a popular choice for institutions and small to medium-sized teams.

Please visit the [official guide](https://docs.openwebui.com/getting-started/) for further details. On this page we provide some advice for configuring this product to work well with our models.

### 1\. Get your API key

Your Apertus provider (see our [Getting Started](/pages/get-started/) page) will have a method of generating an API key to connect with the model. They will also provide you with a URL through which models can be accessed. This is usually an OpenAI-compatible service with a `v1` at the end. Set this up as environment variables with these commands on the terminal or launch script:

```
export APERTUS_SERVING_KEY="sk-rc-..."           # your key from an API key page
export APERTUS_SERVING_URL="https://api.../v1"   # link to your API endpoint
```

### 2\. Run Open WebUI with Docker

The quickest path is to bake the connection details straight into the `docker run` command:

```
docker run -d\
  -p 3000:8080\
  -e OPENAI_API_BASE_URL=$APERTUS_SERVING_URL\
  -e OPENAI_API_KEY=$APERTUS_SERVING_KEY\
  -v open-webui:/app/backend/data\
  --name open-webui\
  --restart always\
  ghcr.io/open-webui/open-webui:main
```

-   `-p 3000:8080` exposes the UI on <http://localhost:3000>.
-   `OPENAI_API_BASE_URL` points Open WebUI at our OpenAI-compatible endpoint.
-   `OPENAI_API_KEY` authenticates every request with your key.
-   `-v open-webui:/app/backend/data` keeps your accounts, chats, and settings across restarts.

Then open <http://localhost:3000>, create a local admin account (this stays on your machine), and start chatting.

### 3\. (Optional) Configure the connection from the UI instead

If you'd rather not put the key in the `docker run` command, start the container without the two `-e` flags and add the connection from the UI:

1.  Go to **Settings → Admin Settings → Connections**.
2.  Under **OpenAI API**, set:
    -   **API Base URL**: `https://..../v1`
    -   **API Key**: your secret key
3.  Save. Our models will now appear in the model picker.

### 4\. Pick a model

Open WebUI auto-discovers everything served at `GET /v1/models`. Select one from the dropdown, for example:

-   `swiss-ai/Apertus-70B-Instruct-2509`
-   `swiss-ai/Apertus-8B-Instruct-2509`

The available list changes over time --- the model picker always reflects what's live. You can also see the full list any time with:

```
curl $APERTUS_SERVING_URL/models\
  -H "Authorization: Bearer $APERTUS_SERVING_KEY"
```

### Verify the endpoint directly

If something isn't connecting, confirm the API works on its own before debugging Open WebUI:

```
curl -X POST $APERTUS_SERVING_URL/chat/completions\
  -H "Content-Type: application/json"\
  -H "Authorization: Bearer $APERTUS_SERVING_KEY"\
  -d '{
    "model": "swiss-ai/Apertus-8B-Instruct-2509",
    "messages": [{"role": "user", "content": "Who is Pablo Picasso?"}],
    "stream": false
  }'
```

That's it --- a fully local UI, our models, your key. 🎉

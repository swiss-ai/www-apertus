---
title: "Apertus 1.5"
date: 2026-07-24T00:00:00+01:00
draft: false
author: "ETH EPFL CSCS"
tags: ["Apertus", "Release", "Announcement"]
featured_image: "apertus-keyvisual-1-5.jpg"
summary: "Our fully open models become more useful in practice: multimodal input, reasoning, longer context, and better instruction following"
comments: false
---

#### Apertus 1.5, our latest model release, extends the 8B and 70B Apertus 1.0 models with image understanding, an optional thinking mode, a four times longer context window, improved instruction-following, and better tool use.

Apertus 1.5 is a continued pretraining of Apertus 1.0: we added 4 trillion tokens of text and multimodal training data to the 8B model and 2 trillion tokens to the 70B model. As with every Apertus release, the models are fully open: open weights, open data, open values, and full training details.

The updated 8B and 70B models are released on July 24, 2026. We are working with inference providers to make them available on multiple platforms.

### Key Enhancements

- Native Image Understanding: Apertus 1.5 accepts images alongside text, bringing documents, diagrams, and photos into the conversation. The models can also process spoken language, though this remains experimental.
- Thinking Mode: The models can be switched into a thinking mode to reason about the input before answering, which improves performance on reasoning problems.
- Long Context: The context window grows to 262,144 tokens, a four-fold increase over Apertus 1.0.
- Improved Instruction-Following: Better instruction adherence than Apertus 1.0 gives more predictable and accurate responses.
- Improved Tool Use: The models are trained for better tool integration, allowing more effective use of external tools and APIs.
- Open Values: Stronger adherence to the [Apertus Charter](/pages/charter/), which brings transparency to the model's values and principles.

![Screenshot of Apertus](Apertus-1.5-Wasserkraftwerks_Niederried-Radelfingen,_Schnittzeichnung.jpg)

_Multimodal processing: Apertus 1.5 70B interprets the contents of a technical drawing of a [Water power plant](https://commons.wikimedia.org/wiki/File:Wasserkraftwerks_Niederried-Radelfingen,_Schnittzeichnung.jpg) (CC BY 4.0) based on an image and short prompt, replying with a description in the form of a text response._

A technical report with further details, along with detailed benchmark results, training pipelines, and intermediate checkpoints, will be published in the coming weeks. Technical instructions for running the models can be found in the model cards of the [Apertus 1.5 8B](https://huggingface.co/swiss-ai/Apertus-v1.5-8B) and [Apertus 1.5 70B](https://huggingface.co/swiss-ai/Apertus-v1.5-70B) releases on Hugging Face.

### How to Get Involved

If you build with or use Apertus, please reach out through our [Contact page](/contact/). We would love to highlight your work, learn what you are building, and hear how Apertus can be improved.

Join the discussion in the community forums on Hugging Face and GitHub. At our regular [Swiss AI SME Circle](/articles/2026-06-sme-circle/) events, small and medium enterprises meet our engineers in person to exchange ideas and experiences.

Want to help build fully open, responsibly trained, and truly global open source AI? The Apertus team is growing, with engineer positions open at ETH Zurich and EPFL \[[1](https://www.jobs.ethz.ch/job/view/JOPG_ethz_R5qUjBUy23vspsktAL), [2](https://www.jobs.ethz.ch/job/view/JOPG_ethz_HDwKLCU4qbu6ray4fL), [3](https://www.jobs.ethz.ch/job/view/JOPG_ethz_kWhKpJAIZB2OmXNQEH), [4](https://www.jobs.ethz.ch/job/view/JOPG_ethz_Z6YaLwO1IE0WjNz0Le), [5](https://careers.epfl.ch/job/Lausanne-Post-Training-Engineer-Apertus/1164610655/), [6](https://dlab.epfl.ch/2026-07-06-safety-pretraining-engineer-positions/)\].

#### \* * * * *

For updates and further information, subscribe to our [Inside Apertus newsletter](/subscribe).

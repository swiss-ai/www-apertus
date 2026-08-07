---
title: "About"
---

## About Apertus

**Apertus is Switzerland's first large-scale, fully open, multilingual language model. Developed by researchers at ETH Zurich, EPFL, and the Swiss National Supercomputing Centre (CSCS), it represents a new approach to foundation model development: built by public institutions, designed for the public good.**

The name comes from the Latin word for "open", reflecting the model's defining characteristic. Unlike commercial models developed behind closed doors, Apertus makes its architecture, weights, training data, and methods fully accessible.

## Roadmap

<style>
  .roadmap { margin: 2rem 0 1rem; }
  .roadmap-item { position: relative; padding: 0 0 1.75rem 2.5rem; }
  .roadmap-item::before {
    content: ""; position: absolute; left: 7px; top: 20px; bottom: -4px;
    width: 2px; background: #e5e7eb;
  }
  .roadmap-item:last-child::before { display: none; }
  .roadmap-item::after {
    content: ""; position: absolute; left: 0; top: 4px;
    width: 16px; height: 16px; border-radius: 50%;
    background: #49BED8; border: 3px solid #49BED8;
  }
  .roadmap-item.planned::after { background: #fff; border-color: #c0c4cc; }
  .roadmap-item.latest .roadmap-content {
    background: #f0f9fc; border: 1px solid #49BED8; border-radius: 12px;
    padding: 1rem 1.25rem;
  }
  .roadmap-header { display: flex; flex-wrap: wrap; align-items: baseline; gap: 0.5rem; margin-bottom: 0.5rem; }
  .roadmap-title { font-size: 1.15rem; font-weight: 600; color: #1a1a1a; }
  .roadmap-weights { font-size: 0.85rem; white-space: nowrap; }
  .roadmap .badge-released { background: #d1fae5; color: #065f46; }
  .roadmap .badge-latest { background: #dbeafe; color: #1e40af; }
  .roadmap .badge-planned { background: #f3f4f6; color: #6b7280; }
  .roadmap-item p { margin-bottom: 0; }
  .roadmap-item ul { margin: 0; padding-left: 1.1rem; }
  .roadmap-item li { margin-bottom: 0.35rem; }
  .roadmap-item li:last-child { margin-bottom: 0; }
  .roadmap-note { color: #9ca3af; font-size: 0.9rem; margin-bottom: 2.5rem; }
</style>

<div class="roadmap">
  <div class="roadmap-item" id="apertus-1-0">
    <div class="roadmap-content">
      <div class="roadmap-header">
        <a class="roadmap-title" href="/articles/2025-09-apertus-1-0/">Apertus 1.0</a>
        <span class="badge badge-released">Released · September 2025</span>
        <a class="roadmap-weights" href="https://huggingface.co/collections/swiss-ai/apertus-v1-68b699e65415c231ace3b059" target="_blank" rel="noopener">Weights ↗</a>
      </div>
      <ul>
        <li>Fully open: architecture, weights, training data, and methods all public, released under the Apache 2.0 licence, which allows commercial use.</li>
        <li>Dense 8B and 70B models, among the largest fully open models available, with the 8B among the best dense models at the 8-10B scale.</li>
        <li>Pretrained on over 1000 languages, using only data from public, versioned sources under open licences, with opt-outs respected for web data and a published, peer-reviewed mechanism that effectively prevents memorisation of training data.</li>
        <li>Instruction following, aligned to the open values of the <a href="/pages/charter/">Apertus Charter</a>.</li>
      </ul>
    </div>
  </div>
  <div class="roadmap-item" id="apertus-1-1">
    <div class="roadmap-content">
      <div class="roadmap-header">
        <a class="roadmap-title" href="/articles/2026-06-apertus-mini/">Apertus 1.1</a>
        <span class="badge badge-released">Released · May 2026</span>
        <a class="roadmap-weights" href="https://huggingface.co/collections/swiss-ai/apertus-mini-6a1ed4b1b7983e243a7e6e5a" target="_blank" rel="noopener">Weights ↗</a>
      </div>
      <p>Distillation of Apertus into 0.5B, 1.5B, and 4B models for mobile and browser deployment: competitive performance at a fraction of the training cost.</p>
    </div>
  </div>
  <div class="roadmap-item latest" id="apertus-1-5">
    <div class="roadmap-content">
      <div class="roadmap-header">
        <a class="roadmap-title" href="/articles/2026-07-apertus-1-5/">Apertus 1.5</a>
        <span class="badge badge-latest">Latest release · July 2026</span>
        <a class="roadmap-weights" href="https://huggingface.co/collections/swiss-ai/apertus-v15-6a63014ea6c937b7ecc6048a" target="_blank" rel="noopener">Weights ↗</a>
      </div>
      <p>A continuation of Apertus 1.0, and among the best fully open models at this scale for following instructions. Adds image understanding competitive with leading open-weight models at this scale, audio understanding, stronger tool use, and reasoning capability from our second-generation post-training pipeline. Stronger adherence to the <a href="/pages/charter/">Apertus Charter</a>.</p>
    </div>
  </div>
  <div class="roadmap-item planned" id="apertus-2-0">
    <div class="roadmap-content">
      <div class="roadmap-header">
        <span class="roadmap-title">Apertus 2.0</span>
        <span class="badge badge-planned">Planned · Q1 2027</span>
      </div>
      <p>A magnitude larger model, developed to the same standard of transparency and trustworthiness: multimodal understanding, a more globally fair vocabulary, and more.</p>
    </div>
  </div>
  <div class="roadmap-item planned" id="apertus-2-5">
    <div class="roadmap-content">
      <div class="roadmap-header">
        <span class="roadmap-title">Apertus 2.5</span>
        <span class="badge badge-planned">Planned · Q3 2027</span>
      </div>
      <p>Further development based on feedback from our partners and developers.</p>
    </div>
  </div>
</div>

<p class="roadmap-note"><em>Planned dates are targets and may shift.</em></p>

Please visit the <a href="/docs/faq">Frequently Asked Questions</a> and our [Documentation](/pages/documentation) section for details. 
Use the [Contact page](/contact) to share ideas and questions with the Apertus team. 

---

## Swiss AI Initiative

**The [Swiss AI Initiative](https://swiss-ai.org) is a national research effort to build capabilities for advanced, trustworthy AI systems for the benefit of society. Launched in December 2023, it is co-led by ETH Zurich and EPFL. Apertus is one of several projects within the Initiative, which spans research in foundation models, AI safety, applications, and education.**

The Initiative represents one of the world's largest open science efforts dedicated to AI foundation models. It pursues an open science approach to increase transparency and access for startups, SMEs, large companies, and the public sector in Switzerland and beyond.
Please visit [swiss-ai.org](https://www.swiss-ai.org/) to learn more.

_We are grateful to the support of all our strategic partners, members, and community._

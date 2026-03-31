---
title: "Documentation"
---

## Technical Information

<p class="section-intro">
  The <a href="/docs/faq">Frequently Asked Questions</a> cover common issues.
  See our <a href="/pages/research">Research collection</a> for an in-depth look at the architecture, training, data mix, and evaluation results.
  Deployment &amp; compliance information follows:</p>

<div class="card-grid">
  <a href="https://huggingface.co/swiss-ai/Apertus-70B-Instruct-2509/blob/main/USAGE_POLICY.md" class="card" style="text-decoration: none;">
    <h4>Apertus Usage Policy</h4>
    <p>Terms and conditions for model use</p>
  </a>
  <a href="/pages/charter" class="card" style="text-decoration: none;">
    <h4>Swiss AI Charter</h4>
    <p>The internal constitution of Apertus</p>
  </a>
  <a href="https://huggingface.co/swiss-ai/Apertus-70B-2509/blob/main/Apertus_EU_Public_Summary.pdf" class="card" style="text-decoration: none;">
    <h4>EU Public Summary</h4>
    <p>Public summary for EU AI Act compliance</p>
  </a>
  <a href="https://huggingface.co/swiss-ai/Apertus-70B-2509/blob/main/Apertus_EU_Code_of_Practice.pdf" class="card" style="text-decoration: none;">
    <h4>EU Code of Practice</h4>
    <p>Code of practice documentation</p>
  </a>
</div>

<p class="section-intro">
  Download the <b>official 
  <a href="https://huggingface.co/swiss-ai">swiss-ai</a> releases</b> on Hugging Face.
</p>

<div class="card-grid">
  <div class="card">
    <img src="https://huggingface.co/front/assets/huggingface_logo-noborder.svg" alt="Hugging Face" style="height:48px;margin-bottom:2rem;float:left">
    <h3>Apertus 8B</h3>
    <p>8 billion parameters · Apache 2.0 license</p>
    Download →
    <a href="https://huggingface.co/swiss-ai/Apertus-8B-2509">Base</a> or
    <a href="https://huggingface.co/swiss-ai/Apertus-8B-Instruct-2509">Instruct</a>
  </div>
  <div class="card">
    <img src="https://huggingface.co/front/assets/huggingface_logo-noborder.svg" alt="Hugging Face" style="height:48px;margin-bottom:2rem;float:left">
    <h3>Apertus 70B</h3>
    <p>70 billion parameters · Apache 2.0 license</p>
    Download →
    <a href="https://huggingface.co/swiss-ai/Apertus-70B-2509">Base</a> or
    <a href="https://huggingface.co/swiss-ai/Apertus-70B-Instruct-2509">Instruct</a>
  </div>
</div>
<p class="section-intro">
  'Instruct' versions are further trained for chat support.
</p>

<div class="card-grid">
  <div class="card">
    <a href="/docs/guides/lmstudio/" style="text-decoration: none;">
      <h4>LM Studio</h4>
      <p>Quick inference on Mac, Linux and Windows.</p>
    </a>
    <p>Over 50 others at 
      <a href="https://alternativeto.net/software/lm-studio/?license=opensource" target="_blank">AlternativeTo</a> ...</p>
    <span class="badge badge-supported">Desktop</span>
  </div>
  <div class="card">
    <a href="/docs/guides/vllm/" style="text-decoration: none;">
      <h4>vLLM</h4>
      <p>Recommended for self-hosted inference servers.</p>
    </a>
    <p>Use
      <a href="/docs/guides/sglang/">SGlang</a>
      for advanced scaling needs.
    </p>
  <span class="badge badge-supported">Server</span>
  </div>
</div>

<p class="section-intro">
  <a href="https://huggingface.co/models?sort=trending&search=apertus" target="_blank">
    Community builds</a> 
  exist in formats like
  <a href="https://huggingface.co/models?library=gguf&sort=trending&search=apertus" target="_blank">MLX</a>
  and
  <a href="https://huggingface.co/models?library=mlx&sort=trending&search=apertus" target="_blank">GGUF</a>.
  Please be aware if you are using a downstream version of Apertus, and seek technical support from the publisher in case of issues.
  <br><a href="/contact">Contact us</a> in case of security, legal, or other concerns.
</p>

---

## Codebase

<div class="card-grid">
  <a href="https://github.com/orgs/swiss-ai/repositories?q=apertus" class="card" style="text-decoration: none;">
    <img src="https://github.githubassets.com/favicons/favicon.svg" alt="GitHub" class="card-logo" style="float:left;margin-bottom:1em">
    <h4>Swiss AI on GitHub</h4>
    <p>Training code, evaluation scripts, and model artifacts for Apertus and related projects.</p>
  </a>
</div>

<div class="card-grid">
  <a href="https://github.com/swiss-ai/pretrain-data" class="card" style="text-decoration: none;">
    <h4>Pre-Training Data</h4>
    <p>Scripts for reconstructing the training dataset.</p>
  </a>
  <a href="https://github.com/swiss-ai/apertus-format" class="card" style="text-decoration: none;">
    <h4>Apertus Format</h4>
    <p>Python library for the custom chat format.</p>
  </a>
</div>

<p class="section-intro">
    Fully open training infrastructure. We look forward to hear about your efforts to build on our models!
</p>

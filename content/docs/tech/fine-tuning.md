---
title: "Fine-tuning"
description: "Some notes on creating new models based on Apertus"
icon: "rocket_launch"
date: "2026-02-11T11:11:00+01:00"
lastmod: "2026-02-17T15:11:45+01:00"
toc: true
tags: ["Technical"]
categories: [""]
author: "Apertus Project"
---


Fine-tuning is essential to adapt Apertus to your domain or task, whether you are working on specialized knowledge, improving performance on a certain dataset, or creating a custom application. We have crafted this guide to make the process accessible and flexible.

Our team has prepared recipes in the a fine-tuning repository, which in the future will also provide more references from the community, and diverse configurations for training to ensure you can tailor Apertus to your specific use cases efficiently. 

### 🚀 Getting Started

To begin, ensure you have a compatible environment set up:

```bash
# Download the recipes
git clone https://github.com/swiss-ai/apertus-finetuning-recipes.git
cd apertus-finetuning-recipes
# Create and activate the environment
uv venv apertus --python 3.10 && source apertus/bin/activate
# Install required dependencies
uv pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/test/cu128
uv pip install -r requirements.txt
# Launch LoRA training on a single GPU
python sft_train.py --config configs/sft_lora.yaml
```

This initial setup allows you to fine-tune the 8B model on a single GPU using LoRA (Low-Rank Adaptation of Large Language Models), a method that reduces computational cost and memory requirements. 

### 📚 Model Selection

You can choose between the 8B and 70B model sizes by specifying the `model_path` in the config file. For instance:

```yaml
model_path: "swiss-ai/Apertus-8B-Instruct-2509"
```

or

```yaml
model_path: "swiss-ai/Apertus-70B-Instruct-2509"
```

Both models work with all provided scripts, with the 8B variant being suitable for a single GPU setup and the 70B for multi-GPU clusters.

### Training Approaches

LoRA Fine-tuning (for 8B model on 1 GPU):

```bash
python sft_train.py --config configs/sft_lora.yaml
```

Full-parameter fine-tuning (for 70B model on 4+ GPUs):

```bash
# With standard attention
accelerate launch --config_file configs/zero3.yaml sft_train.py --config configs/sft_full.yaml
# With FlashAttention3 (optional)
accelerate launch --config_file configs/zero3.yaml sft_train.py \
    --config configs/sft_full.yaml \
    --attn_implementation kernels-community/vllm-flash-attn3
```

Multi-node training (3 nodes x 4 GPUs):

```bash
bash submit_multinode.sh
```


### Customization

- **Datasets**: Modify the `dataset_name` in your config file or via command-line arguments to use a custom dataset.
- **Hyperparameters**: Override any parameter in the config files or use CLI flags, such as `--batch_size`, `--learning_rate`, etc.

After training, your fine-tuned models are saved in:

**LoRA**: `Apertus-FT/output/apertus_lora/`
  - `adapter_model.safetensors`: LoRA weights
  - `adapter_config.json`: LoRA configuration
  - `training_args.bin`, `trainer_state.json`, etc.


**Full Fine-tuning**: `Apertus-FT/output/apertus_full/`
  - Full model and tokenizer files, along with logs.

### Using Your Fine-tuned Models

For LoRA Adapters:

```python
from transformers import AutoModelForCausalLM, AutoTokenizer
from peft import PeftModel

# Load base model and tokenizer
base_model = AutoModelForCausalLM.from_pretrained("swiss-ai/Apertus-8B-Instruct-2509")
tokenizer = AutoTokenizer.from_pretrained("swiss-ai/Apertus-8B-Instruct-2509")

# Load LoRA adapter
model = PeftModel.from_pretrained("Apertus-FT/output/apertus_lora/")
# Merge the adapter (optional for deployment)
model = model.merge_and_unload()
```

For Full Fine-tuned Models:

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

# Load your fine-tuned model and tokenizer directly
model = AutoModelForCausalLM.from_pretrained("Apertus-FT/output/apertus_full/")
tokenizer = AutoTokenizer.from_pretrained("Apertus-FT/output/apertus_full/")
```

---

### 🔐 Safety Considerations

- Avoid training on legally unclear, harmful or sensitive data.
- Use the [official guidelines](/pages/documentation) for responsible AI development.

### 🚀 Next Steps

- Explore the `configs/` directory for pre-configured examples.
- Try [Peft](https://github.com/bigscience-workshop/peft) for LoRA fine-tuning.
- **Experiment with custom datasets** and report back if you find improvements!

For more details, explore the [Apertus Fine-Tuning Recipes](https://github.com/swiss-ai/apertus-finetuning-recipes). This repository is designed for active collaboration. Share your results, suggest improvements, or request features. 


---
title: "Uploading"
description: "On preparing datasets for use with LLMs"
toc: true
tags: ["Technical"]
categories: [""]
author: "Apertus Project"
---


By sharing non-personal datasets, you can help advance research, innovation, and collaboration in machine learning and data science. This guide will walk you through the process of preparing, documenting, and publishing **AI-Ready Data**.

We are building these instructions on the foundation of the official [Adding Datasets](https://huggingface.co/docs/hub/datasets-adding) docs from Hugging Face, and [other publications](#references) in the community. 

Please note: __following this guide does not guarantee__ that your dataset will be included in Apertus training runs.

# Open Dataset Publishing Guide

---

## 1. Before You Start

Why publish open, AI-ready data?

- **Visibility**: Your dataset will be discoverable by researchers and developers on frontier efforts.
- **Impact**: Contribute to open science, accelerate innovation, get recognition on a global level.
- **Support**: Tap into the open data community as a valuable resources for assistance and impact.

What you will need:

- Your dataset in a **standard format** (CSV, JSON, Parquet, etc.).
- A **README file** with clear documentation (see below for details).
- An account on **Hugging Face** ([sign up here](https://huggingface.co/join)) or an [alternative platform](#alternative-platforms).

## 2. Preparing Your Dataset

Organize Your Data

- Ensure your dataset is **clean, well-structured, and consistent**.
- Use **standard file formats** (e.g., CSV for tabular data, JSON for structured data).
- If your dataset is large, consider **splitting it into smaller files** for easier download.

Quality Control

- Review your dataset for **biases, errors, or missing values**.
- Ensure it is **ethically sourced** and is **reproducible**:  
  1. The dataset should be version-controlled (e.g., hosted on Hugging Face, GitHub, or Zenodo) to enable traceability.  
     - If not version-controlled, use tools like Xet to version your data (e.g., add a version in the filename if using hf download).
  2. Check for legal compliance:  
     - Ensure the dataset has a permissive license (e.g., Apache 2.0, MIT, CC-BY, CC0).  
     - Avoid datasets with restrictive licenses (SA, NC) unless explicitly permitted. Check Apertus license guidelines.  
     - If the dataset is not compliant, consider republishing it under an open license or reaching out to the authors.
  3. Ethical Considerations:  
      - Your datasets complies with Swiss/EU privacy regulations (e.g., DVSO, GDPR) and moral codes.
      - Apertus [technical reports](/pages/research) and data guidelines have further recommendations, such as respecting opt-outs.  

Prepare for use in LLMs

- Follow best practices for **tokenization and formatting** (e.g., JSONL for text data).
- Consider filtering and splitting needs, provide guidance for machine learning use.
- Include **metadata** such as language, domain, and intended use cases.

## 3. Documenting Your Dataset

A well-documented dataset is **easier to use and more impactful**.

Visit the [@swiss-ai organization](https://huggingface.co/swiss-ai) that hosts the Apertus project for example datasets.

Include the following in your `README.md` file:

### Required Sections

1. **Dataset Name**: A clear, descriptive title.
2. **Description**: What the dataset contains, the method of its creation, and its purpose.
3. **License**: Specify the license - we recommend [data licenses](https://book.the-turing-way.org/reproducible-research/licensing/licensing-data/).
4. **Citation**: How users should cite your dataset - see [CESSDA](https://datacitation.cessda.eu/The-Anatomy-of-a-Data-Citation).
5. **Usage**: Example code or instructions for loading the dataset.
6. **Metadata**: File structure, column descriptions, and data types.

### Example README Template

```markdown
# [Dataset Name]

**Description**:
[Briefly describe the dataset, its purpose, and potential use cases.]

**License**:
[Specify the license, e.g., MIT License.]

**Citation**:
[Provide a citation format, e.g., APA or BibTeX.]

**Usage**:
[Code snippet or any usage notes.]

**Metadata**:

| File          | Description         | Size |
| ------------- | ------------------- | ---- |
| data.csv      | Main dataset file   | 10MB |
| metadata.json | Additional metadata | 1MB  |

**Contact**:  
[Your email or link to a support channel.]
```

---

## 4. Publishing on Hugging Face

### 4.1: Create a New Repository

1. Log in to your **Hugging Face account**.
2. Navigate to your personal account or organization.
3. Click **"New Dataset"** and select **"Create Dataset"**.

### 4.2: Upload Your Dataset

1. **Drag and drop** your dataset files into the repository.
2. Upload your `README.md` and any additional documentation.
3. Use the **Hugging Face Dataset Viewer** to preview your dataset.

### 4.3: Add Metadata

1. Fill in the **dataset card** with a description, tags, and license.
2. Specify the **dataset structure** (e.g., train/validation/test splits).
3. Add **example code** for loading the dataset.

### 4.4: Publish

1. Click **"Publish"** to make your dataset publicly available.
2. Share the link with your network and the Apertus community!

---

## 5. Alternative Platforms

You can also publish your dataset on other platforms:

- **[Zenodo](https://zenodo.org/)**: Ideal for long-term archiving and DOI assignment.
- **[Kaggle](https://www.kaggle.com/datasets)**: Great for competitions and collaborative projects.
- **[DataHub](https://datahub.io/)**: Open data platform with versioning.

If you publish elsewhere, let us know! If we use your results in a release, we will **link to your dataset** from the Apertus publications.

## 6. Sharing is Caring

There are many ways to announce your dataset, from metadata portals, to blogs and social media. The ideal format for open research is a formal publication or pre-print. You can also propose developing on it at a hackathon.

If your dataset is AI-Ready, it should be possible to apply fine-tuning techniques, optimizing the ability of LLMs to "speak the language", distinguish concepts or languages contained in it. See [our guide](/docs/tech/finetuning) for some more detail.

Please use the [contacts on our website](/contact) to share a link with our community.

Thank you for contributing to open science with Apertus! 🚀

---

# References

- [Apertus data transparency](https://github.com/swiss-ai/apertus-data-transparency/#apertus-data-transparency) (GitHub)
- [AI-Ready Data Whitepaper](https://www.swissdataalliance.ch/en/publikationen/whitepaper-ai-ready-data)
- [huggingface-rdf library](https://pypi.org/project/huggingface-rdf/)
- [The Turing Way](https://book.the-turing-way.org/)
- [CESSDA Data Citation Guide](https://datacitation.cessda.eu/)
- [A survey on dataset quality in machine learning](https://www.sciencedirect.com/science/article/pii/S0950584923001222) (Gong et al. 2023)
- [Llama-Ready Data Guide](https://www.marktechpost.com/2025/04/17/uploading-datasets-to-hugging-face-a-step-by-step-guide/)
- [Jianglong’s Blog Post](https://jaron-u.github.io/posts/upload_dataset_huggingface/)

# Apertus Open Dataset Publishing Guide

Thank you for your interest in contributing to the Apertus open dataset collection. By sharing your dataset, you help advance research, innovation, and collaboration in machine learning and data science. This guide will walk you through the process of preparing, documenting, and publishing your dataset on **Hugging Face**.

We are building these instructions on the foundation of the official [Adding Datasets - Hugging Face](https://huggingface.co/docs/hub/datasets-adding) docs, as well as tutorials in the community. Please send us any questions or [feedback](/feedback). 

1. [Before You Start](#before-you-start)
2. [Preparing Your Dataset](#preparing-your-dataset)
3. [Documenting Your Dataset](#documenting-your-dataset)
4. [Publishing on Hugging Face](#publishing-on-hugging-face)
5. [Alternative Platforms](#alternative-platforms)
6. [Need Help?](#need-help)

---

## 1. Before You Start

Why publish open AI-ready data?

- **Visibility**: Your dataset will be discoverable by researchers and developers on frontier efforts.
- **Impact**: Contribute to open science, accelerate innovation, get recognition on a global level.
- **Support**: Tap into our community as a valuable resources for assistance and impact.

What you’ll need:

- A **Hugging Face account** ([sign up here](https://huggingface.co/join)).
- Your dataset in a **standard format** (CSV, JSON, Parquet, etc.).
- A **README file** with clear documentation (see below for details).

## 2. Preparing Your Dataset

### Step 1: Organize Your Data

- Ensure your dataset is **clean, well-structured, and consistent**.
- Use **standard file formats** (e.g., CSV for tabular data, JSON for structured data).
- If your dataset is large, consider **splitting it into smaller files** for easier download.

### Step 2: Check for Bias and Quality

- Review your dataset for **biases, errors, or missing values**.
- Ensure it is **ethically sourced** and complies with privacy regulations (e.g., GDPR).

### Step 3: Prepare for LLMs (Optional)

If your dataset is intended for **Large Language Models (LLMs)**:

- Follow best practices for **tokenization and formatting** (e.g., JSONL for text data).
- Include **metadata** such as language, domain, and intended use cases.
- For inspiration, see:
  - [Llama-Ready Data Guide](https://www.marktechpost.com/2025/04/17/uploading-datasets-to-hugging-face-a-step-by-step-guide/)
  - [Jianglong’s Blog Post](https://jaron-u.github.io/posts/upload_dataset_huggingface/)

## 3. Documenting Your Dataset

A well-documented dataset is **easier to use and more impactful**. Include the following in your `README.md` file:

### Required Sections

1. **Dataset Name**: A clear, descriptive title.
2. **Description**: What the dataset contains and its purpose.
3. **License**: Specify the license (e.g., MIT, Apache 2.0, CC-BY-SA).
4. **Citation**: How users should cite your dataset.
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

### Step 1: Create a New Repository

1. Log in to your **Hugging Face account**.
2. Navigate to your personal account or organization.
3. Click **"New Dataset"** and select **"Create Dataset"**.
4. Name your dataset (e.g., `apertus/[your-dataset-name]`).

Visit the [@swiss-ai organization](https://huggingface.co/swiss-ai) that hosts the Apertus project for example datasets.

### Step 2: Upload Your Dataset

1. **Drag and drop** your dataset files into the repository.
2. Upload your `README.md` and any additional documentation.
3. Use the **Hugging Face Dataset Viewer** to preview your dataset.

### Step 3: Add Metadata

1. Fill in the **dataset card** with a description, tags, and license.
2. Specify the **dataset structure** (e.g., train/validation/test splits).
3. Add **example code** for loading the dataset.

### Step 4: Publish

1. Click **"Publish"** to make your dataset publicly available.
2. Share the link with your network and the Apertus community!

---

## 5. Alternative Platforms

You can also publish your dataset on other platforms:

- **[Zenodo](https://zenodo.org/)**: Ideal for long-term archiving and DOI assignment.
- **[Kaggle](https://www.kaggle.com/datasets)**: Great for competitions and collaborative projects.
- **[DataHub](https://datahub.io/)**: Open data platform with versioning.

If you publish elsewhere, let us know! If we use your results in a release, we will **link to your dataset** from the Apertus publications.

## 6. Need Help?

### Contact Us

Have questions or need assistance? Reach out to us:

- **Email**: [llm-requests@swiss-ai.org](mailto:support@apertus.ai)
- **Discord**: [Join our community](https://discord.gg/apertvs)
- **GitHub Issues**: [Open an issue](https://github.com/apertus-ai/datasets/issues)

### Feedback

Your feedback helps us improve! Let us know:
- What worked well.
- What could be clearer.
- Any features you’d like to see.

## Next Steps

1. **Prepare your dataset** using the guidelines above.
2. **Document it thoroughly** with a `README.md`.
3. **Publish it** on Hugging Face or another platform.
4. **Share it** with the Apertus community!

---

Thank you for contributing to open science with Apertus! 🚀

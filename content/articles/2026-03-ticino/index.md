---
title: "Apertus for Ticino"
date: 2026-03-17T10:00:00+01:00
draft: false
author: "ETH EPFL CSCS"
tags: ["Apertus", "Community", "Announcement"]
featured_image: "alps_computer.webp"
summary: "Fine-tuned model powers in-house AI translation"
comments: false
---

##### Multiple languages, sensitive documents, and strict privacy rules: for employees of the Canton of Ticino, automated translation is not trivial. The administration is now in the process of introducing an in-house translation system based on Switzerland's Apertus---keeping both sensitive data and the underlying technology under its own control.


Italian, German, French, Spanish, Romanian, Ukrainian---these are just some of the languages that employees of the Canton of Ticino need to translate between. As one of only two Swiss cantons with Italian as official language, Ticino handles a large volume of documents that require translation. Often, the information in these documents is sensitive and must remain confidential. For this reason, the cantonal administration is in the process of introducing a closed, secure in-house solution. The Computer Systems Centre [(Centro sistemi informativi, CSI)](https://www4.ti.ch/dfe/dr/csi/centro/) within the canton's Department of Finance and Economy has now installed an AI translation tool that uses Apertus as its underlying large language model (LLM).

[Apertus](https://apertvs.ai/), launched in September 2025, is Switzerland's transparent, fully open-source large language model. It was developed within the [Swiss AI Initiative](https://www.swiss-ai.org/) by [ETH Zürich](https://ethz.ch/en.html), [EPFL](https://www.epfl.ch/en/), and the [Swiss National Supercomputing Centre CSCS](https://www.cscs.ch/). The foundation model is freely available in two sizes---one with 8 billion parameters and the other with 70 billion. Both models are released under a permissive open-source license, allowing their use in education and research as well as in broader societal and commercial applications. A Ticino-based company, [Artificialy](https://www.artificialy.com/), has now built for the canton a fine-tuned version of the smaller model, Apertus 8b.

### **The smaller Apertus model performs**

To fine-tune the Apertus 8b foundation model, experts at Artificialy used language datasets designed to optimize the model for translation---not only for Switzerland's official languages but also, for example, for Romanian and Ukrainian. In a multi-stage process, increasingly curated datasets were used to step by step improve language correctness and specificity. In the final stage, the model is tailored for translation using manually translated example texts and human feedback.

"Notably, even though it's the smaller of the two released models, the fine-tuned Apertus 8b works very well for translation," says [Luca Gambardella](https://search.usi.ch/it/persone/2edd5d5695e5cb3ce10f0ead0fe15133/gambardella-luca-maria), CEO and co-founder of Artificialy and professor of AI at the [Università della Svizzera italiana (USI)](https://www.usi.ch/en). According to the company's benchmark tests comparing different models, Apertus 8b achieved a score of 90 percent before fine-tuning, while the larger Apertus 70b scored 92 percent. This compares to Meta's LLaMa Scout model, which scored 91.4 percent. After fine-tuning, the Apertus 8b-based tool reached 94 percent---a high value, as Luca Gambardella notes. "The good benchmark results were a prerequisite for our decision to use Apertus for our translation tool," he says.

### **Apertus-based solution: confidential, ethical, sovereign**

For the Canton of Ticino, this Apertus-based translation tool builds on an ongoing test project. "Many of our documents contain private, sensitive information, and we need to guarantee appropriate data protection," explains Rudi Belotti, head of systems and user support at the Cantonal Computer Systems Center (CSI). While the use of online AI translation tools such as DeepL is permitted for cantonal employees, they must strictly follow data protection regulations and anonymize documents manually when necessary. These complications disappear with the Apertus-based in-house AI translation tool. The solution is hosted in CSI's own data centre, where GPU processors have been installed to run local AI solutions.

Until now, Artificialy's translation tool within its "PrivateAI Suite" used Mixtral as the underlying LLM, a model from the French company Mistral AI. However, it is only partially open source: while the model weights are released, the full training data and training pipeline are not. In contrast, Apertus is fully open source, with model weights, training code, training data, and full training pipeline documented and publicly available.

![Screenshot of Apertus-based translation tool with example text.](https://www.cscs.ch/fileadmin/_processed_/e/4/csm_Tool-interface-example_10e2fb9b72.png) 

<small>Looks just like DeepL but provides confidentiality and data protection: Artificialy's translation tool for the Canton of Ticino based on Apertus. (Image: CSI TI)</small>

"This was another important reason why we decided to switch to Apertus," says Rudi Belotti. "As a public administration, we feel obligated to use ethical software applications. With Apertus we can be sure the model was trained in Switzerland and in accordance with the highest ethical standards, meaning it uses data that were not proprietary or copyright-protected but released for AI training." In addition, with this solution the canton gains sovereignty over its translation procedures, as both the hardware and the AI solution are located on-site rather than in data centres outside Switzerland.

### **From test to full availability**

In a first phase, the new Apertus-based tool will be tested by around 100 cantonal employees who require translations daily and have requested an automated AI translation solution to facilitate their work. Their feedback will be collected and, where possible, incorporated into the tool. "The test phase will start in the coming weeks," says Rudi Belotti. "Based on the benchmarks, it looks like a very promising solution. We look forward to making it available to our internal users---starting with employees who process confidential data that needs to remain within the perimeter of our data centre."

Author: Santina Russo, CSCS

Republished from [cscs.ch/science/computer-science-hpc/..](https://www.cscs.ch/science/computer-science-hpc/2026/apertus-powers-in-house-ai-translation-for-ticino)

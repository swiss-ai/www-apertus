---
title: "Open Source LLM Builders"
date: 2026-02-09T10:00:00+01:00
draft: false
author: "ETH EPFL CSCS"
tags: ["Apertus", "Community", "Event"]
featured_image: "IMG_2432.jpg"
summary: "What It Will Take to Enable Global Collaboration"
comments: false
---

##### The second [Open-Source LLM Builders Summit](https://luma.com/oj4e1fq8) concluded in Lausanne with a final plenary discussion led by the Scientific Committee (Antoine Bosselut, EPFL; Joost VandeVondele, CSCS; Imanol Schlag, ETHZ; Valentina Pyatkin, Ai2) synthesizing insights from the afternoon's breakout sessions. The goal of the workshop was to explore how meaningful global collaboration on open large language models (LLMs) could move from aspiration to implementation.

![](https://ai.epfl.ch/wp-content/uploads/LLMbuilderevent2026.webp)

How do we lower the frictions of collaboration and knowledge exchange, while preserving the diversity and healthy competition that drives innovation? What will it actually take to make global collaboration on open large language models work in practice? The discussion at the Open-Source LLM Builders Summit surfaced a set of concrete challenges and opportunities: accelerating knowledge exchange beyond traditional academic channels, moving from publishing models to building deployable systems, navigating regulatory fragmentation and infrastructure dependencies, and rethinking how teams and incentives are structured in AI research. What emerged from the discussion is a picture of what must technically, organizationally, and culturally change for open-source LLM development to scale further globally.

**Rethinking How We Share Knowledge and Practices**
---------------------------------------------------

One of the strongest threads running through the discussions was that open-source AI still relies heavily on traditional academic communication channels. Papers and conferences remain central, yet they may fail to capture the operational knowledge that actually determines progress: training configurations, data mixtures, filtering strategies, evaluation protocols, and, importantly, negative results are a key part of the process.

In practice, critical information often circulates informally, in private channels or internal documents, rather than in shared, structured formats. This slows collective learning and leads to duplicated effort.

The gap with closed-source players may not lie only in model architecture, but in how efficiently knowledge compounds. In order to be an advantage and not a constraint, open-source practices require better documenting of training data for reusability. Participants converged on the need for improved mechanisms to share practices --- not only code, but the full pipeline: how data is prepared and ingested, how post-training workflows are structured, how reinforcement learning environments are configured, and how evaluation is conducted at scale. 

Reproducibility, as noted by several participants, is not simply about releasing datasets; it requires sharing the entire provenance chain --- preprocessing decisions, mixing schedules, and engineering context.

While some participants flagged how maintaining such shared infrastructure requires sustained, coordinated effort and proper communication, others more optimistically suggested that current issues might be engineering problems solvable with better tooling and metadata standards. 

One area of consensus identified the improvement of communication to lower the cost of replication as a first enabling step for collaboration and competitiveness of open-source models.

![](https://ai.epfl.ch/wp-content/uploads/LLMbuilderevent2026_1.webp)

**From Open Models to Open Systems**
------------------------------------

Another clear message was that collaboration cannot stop at publishing weights.

The open community has demonstrated its ability to release high-quality models, competitive on widely used performance benchmarks. The next challenge is building systems around them, such as serving infrastructure, feedback loops, and user integration mechanisms that allow models to improve in real-world settings.

The Linux ecosystem provided a recurring example. The kernel alone did not drive adoption; it was the surrounding ecosystem -- packaging, support, documentation, industrial actors -- that transformed it into global infrastructure. A similar stabilizing layer is still emerging in the open-source AI ecosystem.

This raised practical questions: how can open initiatives gather structured user feedback? Should shared "AI factories" or deployment platforms be developed to collect real-world interaction data? How can reinforcement learning pipelines and tool-use environments be standardized so that post-training advances compound across initiatives rather than fragments?

There was no illusion that these challenges are simple. But there was growing clarity that the competitive frontier is increasingly defined by systems-level capabilities (distillation, RL workflows, evaluation engines) rather than pre-training alone.

**Geopolitical and Policies Barriers**
--------------------------------------

If collaboration is to scale, it must contend with a set of very real, practical, technical, regulatory, and geopolitical constraints.

The first layer concerns policy and regulation. Participants highlighted how differing legal environments, particularly around data usage, copyright, and compliance, shape what open initiatives can realistically do in different regions. A model that is feasible to train in one jurisdiction may face constraints in another. This fragmentation complicates cross-border collaboration and slows iteration.

Rather than accepting regulatory asymmetries as fixed, some participants raised the possibility of coordinated advocacy, a form of collective engagement with policymakers to clarify exemptions, harmonize frameworks, or ensure that open research remains viable. The idea surfaced that collaboration might extend beyond technical work to include strategic dialogue with regulators.

Closely related was the notion that sovereign AI does not necessarily imply one model per country. A coalition-based approach --- informally described during the discussion as a "NATO of AI" --- could align actors around shared values, governance principles, and long-term strategic objectives rather than strict national boundaries. Such a model would not centralize development under a single authority, but instead create a framework for coordination: pooling compute resources where appropriate, aligning on technical standards, sharing evaluation protocols, and ensuring interoperability across independently developed models. In this vision, resilience would come not from isolation, but from structured cooperation among partners committed to openness, democratic oversight, and scientific transparency. While exploratory, the concept reflected a growing awareness that governance architecture may be as decisive as model architecture in shaping the future of open AI.

**Infrastructure Dependencies and Ecosystem Resilience**
--------------------------------------------------------

Alongside policy fragmentation, the discussion addressed structural dependencies in infrastructure. Efficiency pressures at scale drive most initiatives towards dominant hardware ecosystems and highly optimized training frameworks offered by major private companies. These tools deliver unmatched throughput and are difficult to bypass without sacrificing performance. Yet long-term dependency on a narrow stack introduces strategic rigidity, especially when alternative accelerators or hardware paradigms emerge.

The issue is not immediate replacement, but resilience. Diversifying critical components in evaluation tooling, post-training engines, and distributed systems reduces fragility over time.

Even at a more operational level, friction persists. Modern AI workflows rely heavily on containerized environments to ensure reproducibility and portability, and many high-performance computing (HPC) facilities remain slow to adopt newer container technologies. Cultural inertia, security concerns, and legacy systems create bottlenecks that complicate distributed collaboration. The underlying technologies exist; aligning infrastructure practices across institutions remains the harder challenge.

![](https://ai.epfl.ch/wp-content/uploads/LLMbuilderevent2026_2.webp)

**The Human Factor: Engineering, Incentives, and Team Structure**
-----------------------------------------------------------------

Frontier LLM development is often described as a research challenge. In practice, it is equally a systems engineering discipline. It requires sustained expertise in distributed training, data pipelines, evaluation infrastructure, and reinforcement learning environments, areas that demand dedicated engineering effort.

This immediately raises an incentive tension. Academic career rewards publications. Building scalable training engines or maintaining shared infrastructure rarely translates directly into high-impact papers. As a result, the work most critical to collective progress can become structurally undervalued.

Participants drew parallels with computational science, where the recognition of research software engineers transformed the field's ability to build large-scale experiments. Open LLM collaboration may require a similar evolution: engineers not as support staff, but as core intellectual contributors.

Team structure was also debated. One view emphasized that pushing the frontier requires a small, tightly coordinated core team capable of making focused technical decisions at speed, particularly when training state-of-the-art models. Another perspective stressed the importance of broader distributed ecosystems, where many contributors can experiment, ablate, and extend capabilities in parallel. These positions are not mutually exclusive. They reflect a structural design choice: how to combine focused execution with open participation.

Ultimately, collaboration will depend not only on shared compute, but on how teams are structured, how engineering is funded, and how the right incentives are aligned with long-term system building.

**Toward Enabling Conditions for Collaboration**
------------------------------------------------

Taken together, the discussion made clear that global collaboration on open LLMs will not be unlocked by a single agreement or institutional design. It will emerge from strengthening the connective tissue of the ecosystem:

-   improving how knowledge and practices are shared
-   building systems that learn from real-world use
-   navigating regulatory and infrastructure constraints with coordination rather than isolation
-   investing in the engineering talent and team structures required to operate at scale. 

If these enabling layers are reinforced, collaboration becomes less an aspiration and more a natural consequence of interoperable, resilient, and collectively sustained open-source development. Given the speed at which the field is evolving, these conversations must become continuous rather than occasional. The [Swiss AI Initiative](https://www.swiss-ai.org/) will therefore continue to convene this community on a recurring basis, with the ambition of hosting the Open-Source LLM Builders Summit twice a year.

The next edition will aim to consolidate the themes identified in Lausanne into more concrete recommendations for policymakers and funders, and begin to align key model builders around shared communication channels and technical exchange formats. By lowering the friction of knowledge transfer --- from training practices to post-training workflows --- the community can move from episodic interaction to sustained, operational collaboration.

Further materials
-----------------

- All presentations are available on the [AI Center Channel](https://www.youtube.com/watch?v=DZT6S4CPu5M&list=PLVvh53KVINWPu6QQsLT3DECAq6IPa_ALS) (YouTube)
- Our community manager Oleg wrote some thoughts on a [personal blog](https://log.alets.ch/121/)

_This article was first published on [ai.epfl.ch](https://ai.epfl.ch/open-source-llm-builders-summit-what-it-will-take-to-enable-global-collaboration/) and [ai.ethz.ch](https://ai.ethz.ch/news-and-events/ai-center-news/2026/03/building-the-future-of-open-ai-insights-from-the-open-source-llm-builder-summit.html)._

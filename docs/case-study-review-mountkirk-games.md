# Case Study Review: Mountkirk Games
https://thecertsguy.com/bytes/case-study-review-mountkirk-games

Sep 20

Written By [Iman Ghanizada](http://thecertsguy.com/bytes?author=5fa84fe695301b2715979e31)

**Credit goes to** [**Indro Bhattacharya**](https://www.linkedin.com/in/indrajitbhattacharya/) **for this series of case study posts**

As most of you know by now, the Google PCA (Professional Cloud Architect) exam was revamped on May 1st, 2021. With the new version of the exam, and having cleared it myself last month, I noticed some significant changes. Some of the key changes from the previous version of the exam are:

- The questions are more conceptual than straightforward

- Introduction of new areas like Anthos and MLOps

- Longer questions

- Multiple services being tested on a question (like a true architect!)

- **All new case studies**


In this blog post, I will outline how I went about solving the new case studies. I will post the exact document I wrote, and which since May 14th 2021, over 240 Googlers across the globe have used as part of their exam prep. I want to thank the many Googlers who took time to comment and improve the document to get it to its current state. Big shout out to Iman for allowing me to post this on his amazing website. I hope this material will help in your prep as well.

If you haven’t already, please read the [exam deep dive](https://thecertsguy.com/bytes/how-to-pass-the-google-cloud-professional-cloud-architect-exam-in-30-days) to understand the overall strategy and key objectives to study for the Professional Cloud Architect exam.

All the best!

## Mountkirk Games

Mountkirk Games makes **online, session-based, multiplayer games** for **mobile** platforms. They have recently started **expanding to other platforms** after successfully **migrating** their **on-premises environments to Google Cloud**. Their most recent endeavor is to create a retro-style first-person shooter (FPS) game that **allows hundreds of simultaneous players to join** a **geo-specific** digital arena **from multiple platforms and locations.** A **real-time** digital banner will display a global **leaderboard** of all the top players across every active arena.

The existing environment was **recently migrated to Google Cloud**, and five games came across using **lift-and-shift virtual machine migrations**, with a few minor exceptions. **Each** new **game** exists in an **isolated Google Cloud project** nested below a **folder** that maintains most of the **permissions and network policies**. Legacy games with low traffic have been consolidated into a single project. There are also **separate environments for development and testing.**

### Solution concept

Mountkirk Games is building a new multiplayer game that they expect to be very popular. They plan to deploy the game’s **backend on Google Kubernetes Engine** so they can **scale** rapidly and use Google’s **global load balancer** to route players to the closest regional game arenas. In order to keep the global leader board in **sync**, they plan to use a **multi-region Spanner cluster.**

### Business Requirements

- Support multiple gaming platforms **(mobile, desktop, tablets)**

- Support multiple regions **(Clusters in multiple regions)**

- Support rapid iteration of game features **(CI/CD)**

- Minimize latency **(CDN)**

- Optimize for dynamic scaling

- Use managed services and pooled resources **(managed services)**

- Minimize costs **(managed services)**


### Technical requirements

- Dynamically scale based on game activity **(k8s)**

- Publish scoring data on a near real–time global leaderboard **(Memorystore)**

- Store game activity logs in **structured files** for future analysis **(Cloud Logging, GCS)**

- Use GPU processing to render graphics server-side for multi-platform support **(k8s)**

- Support eventual migration of legacy games to this new platform


### Executive statement

Our last game was the first time we used Google Cloud, and it was a tremendous success. We were able to analyze player behavior and game telemetry in ways that we never could before. This success allowed us to bet on a full migration to the cloud and to start building all-new games using cloud-native design principles. Our new game is our most ambitious to date and will open up doors for us to support more gaming platforms beyond mobile. **Latency is our top priority, although cost management is the next most important challenge**. As with our first cloud-based game, we have grown to expect the cloud to enable **advanced analytics capabilities (BQ)** so we can rapidly **iterate** on our **deployments** of bug fixes and new functionality **(Cloud Build)**

### Basic evaluation

**1) Client**

**Online game platform**

- Previous success on GCP has led to an ambitious project to create a new multiplayer game with the aim to support more gaming platforms


**2) Values**

- Already have a plan in place with general design for the infrastructure and some requirements.

- Wants to adopt CI/CD and expand to multiple gaming platforms.


**3) Immediate Goals**

- Reduce latency and support larger gamer footprint

- Keep costs low

- Enable advanced analytics

- Enable rapid deployments


### Technical evaluation

**1) Existing Environment**

- Games hosted on VMs

- Legacy games consolidated into a single project

- Separate Environments for dev and test


**2) Technical Watchpoints & Proposed Solutions**

- 100s of users joining simultaneously from across the globe

  - Create **K8 backends** in multiple regions (regional replication)
- Use **Global HTTPS LB** to balance traffic

  - Store user profiles in **Cloud Datastore**
- Global leaderboard to update in real time

  - **Cloud Memorystore** is the best solution for low latency storage of leaderboard results. Spanner is high cost and not required here.
- Store game activity logs in structured files for future analysis

  - Game activity logs can be logged using **Cloud Logging**, and then stored on **GCS**. These logs can later be analyzed using BQ when needed
- Support rapid iteration of game features

  - Mountkirk games should implement a DevOps process ( **Terraform &** **Cloud Build**) to increase developer productivity and enable rapid iteration of game features

Products: **k8s (Game server), Cloud datastore (User Profiles), Memorystore (leaderboards), Cloud Logging (Game activity logs), GCS (Log storage for analysis), Terraform & Cloud Build (rapid iterations on game features)﻿** ﻿

Stay tuned for case study reviews on each of the four business cases on the exam. In the meantime, don’t forget to check out the exam deep dive! Best of luck to you throughout your studies, you’ll do GREAT!

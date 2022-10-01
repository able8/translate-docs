# Case Study Review: EHR Healthcare

https://thecertsguy.com/bytes/case-study-review-ehr-healthcare

Dec 21

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

## EHR healthcare

EHR Healthcare is a leading **provider of electronic health record software** to the medical industry. EHR Healthcare provides their software as a service to multi-national medical offices, hospitals, and insurance providers.

### Solution concept

Due to rapid changes in the healthcare and insurance industry, EHR Healthcare’s **business** has been **growing exponentially** year over year. They need to be able to **scale their environment (k8s)**, **adapt their disaster recovery plan**, and **roll out new continuous deployment capabilities (CD)** to update their software at a fast pace. Google Cloud has been chosen to replace their **current colocation facilities**

### Existing Technical Environment

EHR’s software is **currently hosted in multiple colocation facilities**. The **lease** on **one** of the data centers is **about to expire**. Customer-facing **applications** are **web-based**, and many have **recently been containerized** to run on a group of **Kubernetes clusters**. **Data** is **stored** in a **mixture of relational and NoSQL databases** (MySQL, MS SQL Server, Redis, and MongoDB).

EHR is **hosting several legacy file- and API-based integrations** with insurance providers **on-premises**. These systems are scheduled to be replaced over the next several years. There is **no plan to upgrade or move these systems at the current time**. **Users are managed via Microsoft Active Directory (ADFS, GCDS)**. **Monitoring** is **currently** being **done via** various **open source tools (Cloud Monitoring, open source, so Prometheus...etc)**. **Alerts** are sent **via email** and are often **ignored**.

### Business Requirements

- On-board new insurance providers as **quickly as possible**

- Provide a minimum **99.9% availability** for all customer-facing systems

- Provide **centralized visibility** and proactive action **on system performance and usage (Cloud Operations/Stackdriver)**

- **Increase** ability to **provide insights** **into** healthcare **trends (Bigquery, AI services)**

- **Reduce latency** to all customers **(CDN, Multi-Region Deployments)**

- **Maintain** **regulatory compliance (HIPAA, Tokenize PII Data)**

- **Decrease** infrastructure **administration** **costs (Managed Services)**

- **Make predictions and generate reports** on industry trends based on provider data **(BigQuery, BQML, Datastudio, Looker)**


### Technical requirements

- **Maintain** legacy interfaces to insurance providers with **connectivity to both on-premises systems and cloud providers (VPN/Interconnect, BYOIP)**

- Provide a consistent way to **manage customer-facing applications** that are **container-based (Anthos)**

- Provide a **secure and high-performance connection** between on-premises systems and Google Cloud **(Dedicated Interconnect)**

- Provide **consistent logging, log retention, monitoring, and alerting capabilities (Anthos Service Mesh and Cloud Operations Suite)**

- Maintain and **manage** **multiple container-based environments (Anthos)**

- **Dynamically scale and provision** new environments **(IaC- Terraform)**

- Create **interfaces** to **ingest** and **process** **data** from new providers **(Pub/Sub, Dataflow)**


### Executive statement

Our **on-premises strategy** has worked for years but has required a **major investment of time** and **money** in **training** our team on distinctly **different systems**, managing similar but **separate environments**, and **responding to outages**. Many of these outages have been a result of **misconfigured systems**, **inadequate capacity** to manage spikes in traffic, and **inconsistent monitoring practices**. We want to use Google Cloud to leverage a **scalable, resilient** platform that can **span multiple environments** seamlessly and **provide a consistent and stable user experience** that positions us for future growth.﻿

### Basic evaluation

**1) Client**

- Leading provider of EHR **software services** to multi-national medical offices, hospitals, and insurance providers. Currently infrastructure is on-prem, but **co-lo is ending on one DC.**


**2) Values**

- Already refactored application on-prem and containerized them

- Significant investment in managing services on prem


**3) Immediate Goals**

- Reduce management of on-prem services

- Rapidly scale

- Improve SRE

- Introduce DevOps


### Technical evaluation

**Existing Environment**

**Application Servers:**

Customer-facing applications are web-based, and many have recently been containerized to run on a group of Kubernetes clusters.

**Technical Watchpoints**

- hosting several legacy file- and API-based integrations

- Customer-facing applications are web-based, and many have recently been containerized


**Proposed Solution**

- Migrate legacy apps which are not containerized using **Migrate for Anthos**

- Use **Anthos Service Mesh** and **Anthos Config Management** to view and manage all containers through a single pane of glass

- **Cloud CDN** for reduced latency


**Existing Environment**

**Databases**:

Mix of relational & Non-relational DBs

**Technical Watchpoints**

- MySQL

- MS SQL Server

- Redis

- MongoDB


**Proposed Solution**

- **Cloud SQL** (both MySQL and SQL Server) - If they have unique MS SQL Server requirements like using SSMS, they may need to run that on a GCE instance. Remember Cloud SQL parity isn't always 1:1

- **Cloud Memorystore**

- **Datastore** or managed instance of MongoDB through **Cloud marketplace**


**Existing Environment**

**Analytics & Reporting:**

Nothing mentioned but has future needs

**Technical Watchpoints**

- Needs ingestion and processing capabilities from multiple provider sources

- Needs ability to make predictions and create reports


**Proposed Solution**

- **Pub/Sub** and **Cloud Dataflow** for ingestion and processing

- **BigQuery** for predictions and **Datastudio/Looker** for reporting

- **AI Platform** can be used for advanced modeling and insights generation


**Existing Environment**

**Identity**:

Users currently on Microsoft AD

**Proposed Solution**

- Use **ADFS** and **GCDS** to sync users to the Google ecosystem


**Existing Environment**

**Monitoring & SRE:**

Open Source tools used

**Technical Watchpoints**

- Several issues in the past due to misconfigured systems, inadequate capacity to manage spikes in traffic, and inconsistent monitoring practices


**Proposed Solution**

- **Cloud Logging** to capture Audit and Network logs

- Store logs in **Cloud Storage** for compliance purposes and apply ACLs

- **Cloud Monitoring** using workspaces

- **Anthos Service Mesh** to monitor multi platform container environment

- Develop SRE practices such as **Chaos Engineering** and **Istio Fault Injection** to test out systems

- Refactor 4 golden “ **LETS**” metrics: Latency, Errors, Traffic, Saturation with percentiles to **reduce alert fatigue**


**Existing Environment**

**Dynamic Scaling & Provisioning of features**

**Technical Watchpoints**

- Need to roll out new continuous deployment capabilities to update their software at a fast pace


**Proposed Solution**

- Create a CI/CD pipelines with **Jenkins** (CI) and **Spinnaker** (CD)

- Use **Terraform** to create Infrastructure as Code. **Cloud Build** for provisioning infra.


Products: **Anthos (managing containerized apps both on-prem and on cloud), Migrate for Anthos (for moving their existing on-prem apps),  Cloud CDN (low latency), Cloud SQL (RDBMS), Memorystore (managed Redis), Pub/Sub (ingest data for asynchronous processing), Dataflow (ETL), BQ (analysis and predictions), Datastudio/Looker (reporting), AI Platform (insight generation), ADFS/GCDS (for AD identity sync), Cloud Operations Suite and Anthos Service Mesh and Cloud Operations (logging & monitoring), Terraform (IaC), Jenkins (CI), Spinnaker (CD)**

Stay tuned for case study reviews on each of the four business cases on the exam. In the meantime, don’t forget to check out the exam deep dive! Best of luck to you throughout your studies, you’ll do GREAT!

[Iman Ghanizada](http://thecertsguy.com/bytes?author=5fa84fe695301b2715979e31)

Iman is an Author & Cloud Security Dude at Google Cloud.

https://thecertsguy.com/bytes/case-study-review-ehr-healthcare

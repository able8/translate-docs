# Case Study Review: TerramEarth

https://thecertsguy.com/bytes/case-study-review-terramearth


Dec 21

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

## TerramEarth

TerramEarth **manufactures** heavy equipment for the mining and agricultural industries. They currently have over **500 dealers** and **service centers** in **100 countries.** Their mission is to build products that make their customers more productive.

### Solution concept

There are **2 million** TerramEarth vehicles in operation currently, and we see **20% yearly growth**. Vehicles collect **telemetry** **data** **from** many **sensors** during operation. A **small subset of critical data** is **transmitted** **from** the **vehicles** in **real time** to facilitate fleet management. The **rest** of the **sensor data** is **collected**, **compressed**, and **uploaded** **daily** when the vehicles return to home base. **Each vehicle** usually **generates** **200 to 500 megabytes of data per day**.

### Existing Technical Environment

TerramEarth’s vehicle **data aggregation and analysis infrastructure resides in Google Cloud** and serves **clients from all around the world**. A **growing** **amount** of **sensor data** is captured **(IoT Core)** from their two main manufacturing plants and **sent to private data centers** that contain their legacy inventory and logistics management systems. The **private data centers have multiple network interconnects** configured to Google Cloud. The **web frontend** for dealers and customers is **running in Google Cloud** and allows **access to stock management and analytics.**

### Business Requirements

- **Predict** and **detect** vehicle **malfunction** and **rapidly** ship parts to dealerships for just-in time repair where possible. **(AI Platform)**

- **Decrease** cloud **operational costs** and **adapt to seasonality**. **(Managed Services)**

- **Increase** **speed** and **reliability** of **development workflow**. **(CI/CD)**

- Allow **remote** **developers** to be **productive** without compromising code or **data** **security**. **(Private Google Access, IAP with signed headers)**

- Create a flexible and scalable platform for developers to create **custom API services** for dealers and partners. **(Apigee)**


### Technical requirements

- Create a **new abstraction layer** for **HTTP API access** **to** their **legacy systems** to enable a gradual move into the cloud without disrupting operations. **(Apigee)**

- **Modernize all CI/CD pipelines** to allow developers to deploy container-based workloads in highly scalable environments. **(GKE)**

- Allow developers to **run experiments** without compromising security and governance requirements **(Separate Project/IAM)**

- Create a **self-service portal** for internal and partner developers to create new projects, request resources for data analytics jobs, and centrally manage access to the API endpoints. **(IAM, Apigee)**

- Use cloud-native solutions for **keys and secrets management** and optimize for identity based access **(Cloud KMS, Secret Manager)**

- Improve and standardize **tools** necessary for **application and network monitoring** and troubleshooting **(Cloud Operations)**


### Executive statement

Our competitive advantage has always been our focus on the customer, with our ability to provide excellent customer service and **minimize vehicle downtimes**. After moving multiple systems into Google Cloud, we are seeking new ways to provide **best-in-class online fleet management services** to our customers and **improve operations of our dealerships**. Our 5-year strategic plan is to create a partner ecosystem of new products by **enabling access to our data**, increasing **autonomous operation capabilities** of our vehicles, and creating a path to move the remaining legacy systems to the cloud.

### Basic evaluation

**Client**

TerramEarth **manufactures** heavy equipment for the mining and agricultural industries. They currently have over **500 dealers** and **service centers** in **100 countries.** Their mission is to build products that make their customers more productive.

**Values**

- Already on GCP

- Multiple network interconnects in place between OnPrem and GCP

- Web Front end running on GCP


**Immediate Goals**

- Minimize Vehicle Downtimes

- Provide best in class online Fleet management services

- Improve dealership


### Technical evaluation

**Requirements**

**Predict** and **detect** vehicle **malfunction** and **rapidly** ship parts to dealerships for just-in time repair

**Technical Watchpoints**

- The **web frontend** for dealers and customers is **running in Google Cloud** and allows **access to stock management and analytics.**


**Proposed Solution**

- Use **AI Platform** to create prediction models

- BigQuery for handling real time data to facilitate fleet management


**Requirements**

**Decrease** cloud **operational costs** and **adapt to seasonality**

**Proposed Solution**

- **IoT Core, Pub/Sub** and **Dataflow** as we need to decouple the messages ingestion and processing


**Requirements**

**Increase** **speed** and **reliability** of **development workflow**.

**Technical Watchpoints**

- Modernize all CI/CD pipelines

- **keys and secrets management** and optimize for identity based access

- Standardize **tools** necessary for **application and network monitoring** and troubleshooting


**Proposed Solution**

- Modernize CI/CD with **Cloud Build** and **Deployment Manager**

- **Cloud KMS, Secret Manager**

- **Cloud Operations to capture Audit and Network Logs (VPC Flow Logs) Network Intelligence to monitor performance and topology**


**Requirements**

Remote developer productivity

**Technical Watchpoints**

- Allow developers to **run experiments** without compromising security and governance requirements


**Proposed Solution**

- Use Identity Aware Proxies ( **IAP**) and host the sandbox project in a **separate folder** with appropriate **policies** in place (IAM and network policies)


**Requirements**

Create **custom API services** for dealers and partners

**Technical Watchpoints**

- Create a **new abstraction layer** for **HTTP API access** **to** their **legacy systems**

- **Self-service portal to create projects, request resources for analytics jobs, and centrally manage APIs**


**Proposed Solution**

- **Apigee** as central portal for API access management, self service and monetization.

- **GKE** as backend service aggregating On-Prem data and Analytic data, requesting analytics jobs


**Requirements**

Interconnect with private data center

**Proposed Solution**

- **Cloud Router + Interconnect(Partner or Dedicated) for interconnect with private datacenter**


Products: **AI Platform (Predictions), IoT Core (for managing devices and creating a bridge to stream data), Pub/Sub (as endpoint to ingest streaming data from IoT devices) Dataflow (for processing), Terraform and Deployment Manager (CI/CD), Cloud KMS & Secrets Manager (reliability of dev workflow), Cloud Operations Suite (for monitoring), Cloud IAM and IAP (Remote dev productivity), Apigee (API layer for access to Legacy systems and self service portal)**

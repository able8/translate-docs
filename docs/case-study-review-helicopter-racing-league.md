# Case Study Review: Helicopter Racing League

https://thecertsguy.com/bytes/case-study-review-helicopter-racing-league

Dec 21

Written By [Iman Ghanizada](http://thecertsguy.com/bytes?author=5fa84fe695301b2715979e31)

**Credit goes to** [**Indro Bhattacharya**](https://www.linkedin.com/in/indrajitbhattacharya/) **& countless Googlers for this series of case study posts**

As most of you know by now, the Google PCA (Professional Cloud Architect) exam was revamped on May 1st, 2021. With the new version of the exam, and having cleared it myself last month, I noticed some significant changes. Some of the key changes from the previous version of the exam are:

- The questions are more conceptual than straightforward

- Introduction of new areas like Anthos and MLOps

- Longer questions

- Multiple services being tested on a question (like a true architect!)

- **All new case studies**


In this blog post, I will outline how I went about solving the new case studies. I will post the exact document I wrote, and which since May 14th 2021, over 240 Googlers across the globe have used as part of their exam prep. I want to thank the many Googlers who took time to comment and improve the document to get it to its current state. Big shout out to Iman for allowing me to post this on his amazing website. I hope this material will help in your prep as well.

If you haven’t already, please read the [exam deep dive](https://thecertsguy.com/bytes/how-to-pass-the-google-cloud-professional-cloud-architect-exam-in-30-days) to understand the overall strategy and key objectives to study for the Professional Cloud Architect exam.

All the best!

## Helicopter Racing League

Helicopter Racing League (HRL) is a **global** sports league for competitive helicopter racing. Each year HRL holds the world championship and several **regional** league competitions where teams compete to earn a spot in the world championship. HRL offers a **paid service** to **stream** the races **all over the world** with **live telemetry** and **predictions (real time analytics)** **throughout** **each** race.

[https://www.youtube.com/watch?v=cwTiYOtmQbI&list=PLne\_-oJR60mOciEvtX1AfER8lBafH8Lox&index=3](https://www.youtube.com/watch?v=cwTiYOtmQbI&list=PLne_-oJR60mOciEvtX1AfER8lBafH8Lox&index=3)’

### Solution concept

HRL wants to **migrate** their existing service to a new platform to **expand** their **use of managed AI and ML services (AI Platform)** to facilitate race predictions. Additionally, as new fans engage with the sport, particularly in emerging regions, they want to move the **serving of their content, both real-time and recorded (GCS), closer to their users (CDN) .**

### Existing Technical Environment

HRL is a **public cloud-first company**; the core of their mission-critical applications runs on their current public cloud provider. **Video recording and editing** is performed **at** the race **tracks**, and the **content is encoded and transcoded**, where needed, **in the cloud**.

**Enterprise-grade connectivity** and **local compute** is provided by **truck-mounted mobile data centers**. Their **race prediction services** are **hosted** exclusively on their **existing public cloud provider.** Their existing technical environment is as follows:

- Existing **content** is **stored in an object storage service** on their existing public cloud provider.

- Video **encoding** and **transcoding** is **performed on VMs** created for **each** job **(too many VMs, not sustainable)**

- Race **predictions** are performed using **TensorFlow** **running on VMs** **(Custom Code, no managed services)** in the current public cloud provider.


### Business Requirements

HRL’s owners want to **expand** their **predictive capabilities** and **reduce latency** for their viewers in emerging markets. Their requirements are:

- Support ability to **expose** the predictive **models to partners (Apigee)**.

- **Increase** predictive capabilities **during and before races**: **(AI Platform)**

  - Race results

  - Mechanical failures

  - Crowd sentiment
- **Increase telemetry** and create **additional** **insights** **(IoT Core, BQ, Looker/Datastudio)**

- **Measure** fan **engagement** with new predictions **(AI Platform)**

- **Enhance** global **availability** and **quality** of the broadcasts **(CDNs)**

- **Increase** the number of **concurrent** viewers **(CDNs, Global LB)**

- **Minimize** operational complexity **(Managed services)**

- Ensure **compliance** with regulations **(GDPR,PII)**

- Create a merchandising revenue stream **(Separate API?)**


### Technical requirements

- **Maintain** or **increase** **prediction** **throughput** and **accuracy (BQ ML function can predict)**

- **Reduce** viewer **latency (CDN)**

- **Increase** **transcoding** performance **(Transcoder API)**

- Create **real-time analytics** of viewer consumption patterns and engagement **(Trucks/On prem to PubSub,Dataflow,Bigtable, Looker)**

- Create a **data mart** to enable processing of large volumes of race data **(GCS,Dataflow,BQ,AI Platform - for managed tensorflow VMs)- Dataplex - New offering**


### Executive statement

Our CEO, S. Hawke, wants to bring high-adrenaline racing to fans all around the world. We listen to our fans, and they want enhanced video streams that include predictions of events within the race (e.g., overtaking). Our **current platform** **allows** us to **predict race outcomes** but **lacks** the facility to **support real-time predictions** during races **and** the **capacity** to **process season-long results**.

### Basic evaluation

**Client**

Helicopter Racing League (HRL) is a **global** sports league for competitive helicopter racing. HRL offers a **paid service** to **stream** the races **all over the world** with **live telemetry** and **predictions throughout** **each** race.

**Values**

- Video recording and editing is performed at the race tracks, and the content is encoded and transcoded, where needed, in the cloud.


- Real time predictions to users


**Immediate Goals**

- Large data processing (season long)

- Increase fan base and provide low latency viewing experience

- Increase telemetry and create additional insights


### Technical evaluation

**Requirements**

Content Serving from regions closer to the viewer.

**Technical Watchpoints**

- Both real time and recorded


**Proposed Solution**

- Use **GCS** to store the raw and  transcoded content (Multi-Region)

- Use **CDN** to deliver content


**Requirements**

Datamart for batch data processing

**Technical Watchpoints**

- Season long data (possibly petabyte scale)


**Proposed Solution**

- **GCS** can be used to store the unstructured video data. After transcoding, the data can be uploaded back to GCS and then to **AI Platform** for predictions (Offline Mode)


**Requirements**

Real time analytics and predictions

**Technical Watchpoints**

- Increased telemetry

- Real time predictions and analytics


**Proposed Solution**

- **Pub/Sub** and **Cloud Dataflow** for ingestion and processing

- **AI Platform** can be used for modeling and real time insights generation (also an option since they are already using Tensorflow)


**Requirements**

Real time Video Analysis and Transcoding performance

**Technical Watchpoints**

- Today the transcoding & encoding is done as required on the cloud


**Proposed Solution**

- **GCS** to store the videos

- **Cloud Function** is triggered each time content is uploaded

- CF calls the **Transcoder API**

- Transcoded video stored in GCS in new bucket

- **Video Intelligence API (for video analysis) triggered  via CF**


**Requirements**

Expose Prediction Model to Partners

**Proposed Solution**

- Use **Apigee** or **Cloud Endpoints** to expose prediction model to partners, **Apigee** if it requires monetization


**Requirements**

Minimize Operational Complexity

**Technical Watchpoints**

- Using trucks to stream data from race location and performing trans/encoding services on the cloud


**Proposed Solution**

- **Cloud Operations suite** for SRE


Products: **Cloud Operations, APigee, Video Intelligence API, Transcoder API, Cloud Functions, GCS, Pub/Sub, Dataflow, AI Platform, CDN**

Stay tuned for case study reviews on each of the four business cases on the exam. In the meantime, don’t forget to check out the exam deep dive! Best of luck to you throughout your studies, you’ll do GREAT!

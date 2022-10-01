# Case Study Review: Helicopter Racing League

# 案例研究回顾：直升机赛车联盟

https://thecertsguy.com/bytes/case-study-review-helicopter-racing-league

https://thecertsguy.com/bytes/case-study-review-helicopter-racing-league

Dec 21

12月21日

Written By [Iman Ghanizada](http://thecertsguy.com/bytes?author=5fa84fe695301b2715979e31)

由 [Iman Ghanizada](http://thecertsguy.com/bytes?author=5fa84fe695301b2715979e31) 撰写

**Credit goes to** [**Indro Bhattacharya**](https://www.linkedin.com/in/indrajitbhattacharya/) **& countless Googlers for this series of case study posts**

**功劳归于** [**Indro Bhattacharya**](https://www.linkedin.com/in/indrajitbhattacharya/) **& 无数 Google 员工为这一系列案例研究发帖**

As most of you know by now, the Google PCA (Professional Cloud Architect) exam was revamped on May 1st, 2021. With the new version of the exam, and having cleared it myself last month, I noticed some significant changes. Some of the key changes from the previous version of the exam are:

正如你们大多数人现在所知，Google PCA（专业云架构师）考试于 2021 年 5 月 1 日进行了改版。随着新版本的考试，并且上个月我自己通过了考试，我注意到了一些重大变化。与以前版本的考试相比，一些主要变化是：

- The questions are more conceptual than straightforward

- 这些问题比简单的更概念化

- Introduction of new areas like Anthos and MLOps

- 引入 Anthos 和 MLOps 等新领域

- Longer questions

- 更长的问题

- Multiple services being tested on a question (like a true architect!)

- 在一个问题上测试多个服务（就像一个真正的建筑师！）

- **All new case studies**

- **所有新案例研究**

In this blog post, I will outline how I went about solving the new case studies. I will post the exact document I wrote, and which since May 14th 2021, over 240 Googlers across the globe have used as part of their exam prep. I want to thank the many Googlers who took time to comment and improve the document to get it to its current state. Big shout out to Iman for allowing me to post this on his amazing website. I hope this material will help in your prep as well.

在这篇博文中，我将概述我是如何解决新案例研究的。我将发布我写的确切文件，自 2021 年 5 月 14 日以来，全球 240 多名 Google 员工已将其用作考试准备的一部分。我要感谢许多花时间评论和改进文档以使其达到当前状态的 Google 员工。向伊曼大喊大叫，让我将其发布在他令人惊叹的网站上。我希望这些材料也能帮助你做好准备。

If you haven't already, please read the [exam deep dive](https://thecertsguy.com/bytes/how-to-pass-the-google-cloud-professional-cloud-architect-exam-in-30-days) to understand the overall strategy and key objectives to study for the Professional Cloud Architect exam.

如果您还没有，请阅读 [考试深入了解](https://thecertsguy.com/bytes/how-to-pass-the-google-cloud-professional-cloud-architect-exam-in-30-天)了解学习专业云架构师考试的总体策略和关键目标。

All the best!

一切顺利！

## Helicopter Racing League

## 直升机竞速联赛

Helicopter Racing League (HRL) is a **global** sports league for competitive helicopter racing. Each year HRL holds the world championship and several **regional** league competitions where teams compete to earn a spot in the world championship. HRL offers a **paid service** to **stream** the races **all over the world** with **live telemetry** and **predictions (real time analytics)** **throughout** ** each** race.

直升机竞速联盟 (HRL) 是一个**全球**的竞技直升机竞速体育联盟。每年 HRL 都会举办世界锦标赛和几场 ** 地区** 联赛比赛，各支球队争夺世界冠军的席位。 HRL 提供**付费服务**，通过**实时遥测**和**预测（实时分析）****全程****直播**世界各地的比赛**每场**比赛。

[https://www.youtube.com/watch?v=cwTiYOtmQbI&list=PLne\_-oJR60mOciEvtX1AfER8lBafH8Lox&index=3](https://www.youtube.com/watch?v=cwTiYOtmQbI&list=PLne_-oJR60mOciEvtX1AfER8lBafH8Lox&index=3)’

[https://www.youtube.com/watch?v=cwTiYOtmQbI&list=PLne\_-oJR60mOciEvtX1AfER8lBafH8Lox&index=3](https://www.youtube.com/watch?v=cwTiYOtmQbI&list=PLne_-oJR60mOciEvtX1AfER8lBafH8Lox&index=3)’

### Solution concept

### 解决方案概念

HRL wants to **migrate** their existing service to a new platform to **expand** their **use of managed AI and ML services (AI Platform)** to facilitate race predictions. Additionally, as new fans engage with the sport, particularly in emerging regions, they want to move the **serving of their content, both real-time and recorded (GCS), closer to their users (CDN) .**

HRL 希望将他们现有的服务**迁移**到一个新平台，以**扩展**他们对托管 AI 和 ML 服务（AI 平台）的**使用，以促进比赛预测。此外，随着新粉丝参与这项运动，尤其是在新兴地区，他们希望将 ** 实时和录制 (GCS) 内容服务更靠近用户 (CDN)。**

### Existing Technical Environment

### 现有技术环境

HRL is a **public cloud-first company**; the core of their mission-critical applications runs on their current public cloud provider. **Video recording and editing** is performed **at** the race **tracks**, and the **content is encoded and transcoded**, where needed, **in the cloud**.

HRL 是一家**公共云优先的公司**；他们的关键任务应用程序的核心在他们当前的公共云提供商上运行。 **视频录制和编辑**在**在**比赛**赛道**进行，**内容在需要时被编码和转码**，在需要时，**在云中**。

**Enterprise-grade connectivity** and **local compute** is provided by **truck-mounted mobile data centers**. Their **race prediction services** are **hosted** exclusively on their **existing public cloud provider.** Their existing technical environment is as follows:

**企业级连接**和**本地计算**由**车载移动数据中心**提供。他们的**比赛预测服务**是**托管**在他们的**现有的公共云提供商上。**他们现有的技术环境如下：

- Existing **content** is **stored in an object storage service** on their existing public cloud provider.

- 现有的**内容**存储在其现有公共云提供商上的对象存储服务**中。

- Video **encoding** and **transcoding** is **performed on VMs** created for **each** job **(too many VMs, not sustainable)**

- 视频**编码**和**转码**在为**每个**作业创建的虚拟机**上执行**（虚拟机太多，不可持续）**

- Race **predictions** are performed using **TensorFlow** **running on VMs** **(Custom Code, no managed services)** in the current public cloud provider.

- 竞赛**预测**是使用 **TensorFlow** **在 VM 上运行** **（自定义代码，无托管服务）** 在当前公共云提供商中执行的。

### Business Requirements

###  业务需求

HRL’s owners want to **expand** their **predictive capabilities** and **reduce latency** for their viewers in emerging markets. Their requirements are:

HRL 的所有者希望为新兴市场的观众**扩展**他们的**预测能力**并**减少延迟**。他们的要求是：

- Support ability to **expose** the predictive **models to partners (Apigee)**.

- 支持向合作伙伴 (Apigee) **公开**预测模型**。

- **Increase** predictive capabilities **during and before races**: **(AI Platform)**

   - **提高**预测能力**在比赛期间和之前**：**（AI平台）**

  - Race results

   - 比赛结果

  - Mechanical failures

   - 机械故障

  - Crowd sentiment
- **Increase telemetry** and create **additional** **insights** **(IoT Core, BQ, Looker/Datastudio)**

- 人群情绪
- **增加遥测**并创建**附加** **见解** **（IoT Core、BQ、Looker/Datastudio）**

- **Measure** fan **engagement** with new predictions **(AI Platform)**

- **衡量**粉丝**参与度**与新预测**（AI平台）**

- **Enhance** global **availability** and **quality** of the broadcasts **(CDNs)**

- **增强**全球**可用性**和广播的**质量****（CDN）**

- **Increase** the number of **concurrent** viewers **(CDNs, Global LB)**

- **增加****并发**观众数量**（CDN，全球LB）**

- **Minimize** operational complexity **(Managed services)** 

- **最小化**操作复杂性**（托管服务）**

- Ensure **compliance** with regulations **(GDPR,PII)**

- 确保**遵守**法规**（GDPR，PII）**

- Create a merchandising revenue stream **(Separate API?)**

- 创建销售收入流**（单独的 API？）**

### Technical requirements

###  技术要求

- **Maintain** or **increase** **prediction** **throughput** and **accuracy (BQ ML function can predict)**

- **保持**或**增加** **预测** **吞吐量**和**准确度（BQ ML函数可以预测）**

- **Reduce** viewer **latency (CDN)**

- **减少**查看器**延迟（CDN）**

- **Increase** **transcoding** performance **(Transcoder API)**

- **提高** **转码**性能**（Transcoder API）**

- Create **real-time analytics** of viewer consumption patterns and engagement **(Trucks/On prem to PubSub,Dataflow,Bigtable, Looker)**

- 创建观众消费模式和参与度的**实时分析**（Trucks/On prem to PubSub、Dataflow、Bigtable、Looker）**

- Create a **data mart** to enable processing of large volumes of race data **(GCS,Dataflow,BQ,AI Platform - for managed tensorflow VMs)- Dataplex - New offering**

- 创建一个**数据集市**以支持处理大量竞赛数据**（GCS、Dataflow、BQ、AI 平台 - 用于托管 tensorflow VM）- Dataplex - 新产品**

### Executive statement

### 执行声明

Our CEO, S. Hawke, wants to bring high-adrenaline racing to fans all around the world. We listen to our fans, and they want enhanced video streams that include predictions of events within the race (e.g., overtaking). Our **current platform** **allows** us to **predict race outcomes** but **lacks** the facility to **support real-time predictions** during races **and** the **capacity ** to **process season-long results**.

我们的首席执行官 S. Hawke 希望将高肾上腺素赛车带给全世界的车迷。我们倾听粉丝的心声，他们想要增强的视频流，包括对比赛中事件的预测（例如，超车）。我们的**当前平台** **允许**我们**预测比赛结果**但**缺乏**在比赛期间**支持实时预测**的设施**和**容量** 到 ** 处理整个赛季的结果**。

### Basic evaluation

### 基本评价

**Client**

**客户**

Helicopter Racing League (HRL) is a **global** sports league for competitive helicopter racing. HRL offers a **paid service** to **stream** the races **all over the world** with **live telemetry** and **predictions throughout** **each** race.

直升机竞速联盟 (HRL) 是一个**全球**的竞技直升机竞速体育联盟。 HRL 提供**付费服务**，通过**实时遥测**和**全程预测****每场**比赛**流**世界各地的比赛**。

**Values**

**价值观**

- Video recording and editing is performed at the race tracks, and the content is encoded and transcoded, where needed, in the cloud.

- 在赛道上进行视频录制和编辑，并根据需要在云端对内容进行编码和转码。

- Real time predictions to users

- 对用户的实时预测

**Immediate Goals**

**即时目标**

- Large data processing (season long)

- 大数据处理（季节长）

- Increase fan base and provide low latency viewing experience

- 增加粉丝群并提供低延迟观看体验

- Increase telemetry and create additional insights

- 增加遥测并创造额外的见解

### Technical evaluation

### 技术评测

**Requirements**

**要求**

Content Serving from regions closer to the viewer.

从更接近观众的地区提供内容。

**Technical Watchpoints**

**技术观察点**

- Both real time and recorded

- 实时和记录

**Proposed Solution**

**建议的解决方案**

- Use **GCS** to store the raw and  transcoded content (Multi-Region)

- 使用 **GCS** 存储原始和转码内容（多区域）

- Use **CDN** to deliver content

- 使用 **CDN** 传送内容

**Requirements**

**要求**

Datamart for batch data processing

用于批量数据处理的 Datamart

**Technical Watchpoints**

**技术观察点**

- Season long data (possibly petabyte scale)

- 季节长数据（可能是 PB 级）

**Proposed Solution**

**建议的解决方案**

- **GCS** can be used to store the unstructured video data. After transcoding, the data can be uploaded back to GCS and then to **AI Platform** for predictions (Offline Mode)

- **GCS** 可用于存储非结构化视频数据。转码后数据可以上传回GCS，再上传到**AI平台**进行预测（离线模式）

**Requirements**

**要求**

Real time analytics and predictions

实时分析和预测

**Technical Watchpoints**

**技术观察点**

- Increased telemetry

- 增加遥测

- Real time predictions and analytics

- 实时预测和分析

**Proposed Solution**

**建议的解决方案**

- **Pub/Sub** and **Cloud Dataflow** for ingestion and processing

- **Pub/Sub** 和 **Cloud Dataflow** 用于摄取和处理

- **AI Platform** can be used for modeling and real time insights generation (also an option since they are already using Tensorflow)

- **AI Platform** 可用于建模和实时洞察生成（也是一个选项，因为他们已经在使用 Tensorflow）

**Requirements**

**要求**

Real time Video Analysis and Transcoding performance

实时视频分析和转码性能

**Technical Watchpoints**

**技术观察点**

- Today the transcoding & encoding is done as required on the cloud

- 今天转码和编码是在云端按要求完成的

**Proposed Solution**

**建议的解决方案**

- **GCS** to store the videos

- **GCS** 存储视频

- **Cloud Function** is triggered each time content is uploaded

- 每次上传内容都会触发**云功能**

- CF calls the **Transcoder API**

- CF 调用**Transcoder API**

- Transcoded video stored in GCS in new bucket

- 转码后的视频存储在新存储桶中的 GCS 中

- **Video Intelligence API (for video analysis) triggered  via CF**

- **通过 CF 触发的视频智能 API（用于视频分析）**

**Requirements**

**要求**

Expose Prediction Model to Partners

向合作伙伴公开预测模型

**Proposed Solution**

**建议的解决方案**

- Use **Apigee** or **Cloud Endpoints** to expose prediction model to partners, **Apigee** if it requires monetization

- 使用 **Apigee** 或 **Cloud Endpoints** 向合作伙伴公开预测模型，如果需要获利，则使用 **Apigee**

**Requirements**

**要求**

Minimize Operational Complexity

最大限度地降低运营复杂性

**Technical Watchpoints**

**技术观察点**

- Using trucks to stream data from race location and performing trans/encoding services on the cloud

- 使用卡车从比赛地点传输数据并在云端执行转码/编码服务

**Proposed Solution**

**建议的解决方案**

- **Cloud Operations suite** for SRE

- **适用于 SRE 的云操作套件**

Products: **Cloud Operations, APigee, Video Intelligence API, Transcoder API, Cloud Functions, GCS, Pub/Sub, Dataflow, AI Platform, CDN**

产品：**云操作、APIgee、视频智能 API、转码器 API、云函数、GCS、Pub/Sub、Dataflow、AI 平台、CDN**

Stay tuned for case study reviews on each of the four business cases on the exam. In the meantime, don’t forget to check out the exam deep dive! Best of luck to you throughout your studies, you’ll do GREAT! 

请继续关注有关考试中四个业务案例的案例研究评论。同时，不要忘记检查考试深度潜水！祝你在整个学习过程中好运，你会做得很好！


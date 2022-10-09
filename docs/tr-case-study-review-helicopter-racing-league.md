# Case Study Review: Helicopter Racing League

# 案例研究回顾：直升机赛车联盟

https://thecertsguy.com/bytes/case-study-review-helicopter-racing-league

Dec 21

## Helicopter Racing League

## 直升机竞速联赛

Helicopter Racing League (HRL) is a global sports league for competitive helicopter racing. Each year HRL holds the world championship and several regional league competitions where teams compete to earn a spot in the world championship. HRL offers a paid service to stream the races all over the world with live telemetry and predictions (real time analytics) throughout  each race.

直升机竞速联盟 (HRL) 是一个全球的竞技直升机竞速体育联盟。每年 HRL 都会举办世界锦标赛和几场  地区 联赛比赛，各支球队争夺世界冠军的席位。 HRL 提供付费服务，通过实时遥测和预测（实时分析）全程直播世界各地的比赛每场比赛。

# Solution concept

### 解决方案概念

HRL wants to migrate their existing service to a new platform to expand their use of managed AI and ML services (AI Platform) to facilitate race predictions. Additionally, as new fans engage with the sport, particularly in emerging regions, they want to move the serving of their content, both real-time and recorded (GCS), closer to their users (CDN) .

HRL 希望将他们现有的服务迁移到一个新平台，以扩展他们对托管 AI 和 ML 服务（AI 平台）的使用，以促进比赛预测。此外，随着新粉丝参与这项运动，尤其是在新兴地区，他们希望将  实时和录制 (GCS) 内容服务更靠近用户 (CDN)。

### Existing Technical Environment

### 现有技术环境

HRL is a public cloud-first company; the core of their mission-critical applications runs on their current public cloud provider. Video recording and editing is performed at the race tracks, and the content is encoded and transcoded, where needed, in the cloud.

HRL 是一家公共云优先的公司；他们的关键任务应用程序的核心在他们当前的公共云提供商上运行。 视频录制和编辑在在比赛赛道进行，内容在需要时被编码和转码，在需要时，在云中。

Enterprise-grade connectivity and local compute is provided by truck-mounted mobile data centers. Their race prediction services are hosted exclusively on their existing public cloud provider. Their existing technical environment is as follows:

企业级连接和本地计算由车载移动数据中心提供。他们的比赛预测服务是托管在他们的现有的公共云提供商上。他们现有的技术环境如下：

- Existing content is stored in an object storage service on their existing public cloud provider.

- 现有的内容存储在其现有公共云提供商上的对象存储服务中。

- Video encoding and transcoding is performed on VMs created for each job (too many VMs, not sustainable)

- 视频编码和转码在为每个作业创建的虚拟机上执行（虚拟机太多，不可持续）

- Race predictions are performed using TensorFlow running on VMs (Custom Code, no managed services) in the current public cloud provider.

- 竞赛预测是使用 TensorFlow 在 VM 上运行 （自定义代码，无托管服务） 在当前公共云提供商中执行的。

### Business Requirements

###  业务需求

HRL’s owners want to expand their predictive capabilities and reduce latency for their viewers in emerging markets. Their requirements are:

HRL 的所有者希望为新兴市场的观众扩展他们的预测能力并减少延迟。他们的要求是：

- Support ability to expose the predictive models to partners (Apigee).
- 支持向合作伙伴 (Apigee) 公开预测模型。
- Increase predictive capabilities during and before races: (AI Platform)

   - 提高预测能力在比赛期间和之前：（AI平台）
  - Race results
  - 比赛结果
  - Mechanical failures
  - 机械故障
  - Crowd sentiment
  - 人群情绪
- Increase telemetry and create additional insights (IoT Core, BQ, Looker/Datastudio)
- 增加遥测并创建附加 见解 （IoT Core、BQ、Looker/Datastudio）
- Measure fan engagement with new predictions (AI Platform)
- 衡量粉丝参与度与新预测（AI平台）
- Enhance global availability and quality of the broadcasts (CDNs)
- 增强全球可用性和广播的质量（CDN）
- Increase the number of concurrent viewers (CDNs, Global LB)
- 增加并发观众数量（CDN，全球LB）
- Minimize operational complexity (Managed services) 
- 最小化操作复杂性（托管服务）
- Ensure compliance with regulations (GDPR,PII)
- 确保遵守法规（GDPR，PII）
- Create a merchandising revenue stream (Separate API?)
- 创建销售收入流（单独的 API？）

### Technical requirements

###  技术要求

- Maintain or increase prediction throughput and accuracy (BQ ML function can predict)

- 保持或增加 预测 吞吐量和准确度（BQ ML函数可以预测）

- Reduce viewer latency (CDN)

- 减少查看器延迟（CDN）

- Increase transcoding performance (Transcoder API)

- 提高 转码性能（Transcoder API）

- Create real-time analytics of viewer consumption patterns and engagement (Trucks/On prem to PubSub,Dataflow, Bigtable, Looker)

- 创建观众消费模式和参与度的实时分析（Trucks/On prem to PubSub、Dataflow、Bigtable、Looker）

- Create a data mart to enable processing of large volumes of race data (GCS,Dataflow,BQ,AI Platform - for managed tensorflow VMs)- Dataplex - New offering

- 创建一个数据集市以支持处理大量竞赛数据（GCS、Dataflow、BQ、AI 平台 - 用于托管 tensorflow VM）- Dataplex - 新产品

### Executive statement

### 执行声明

Our CEO, S. Hawke, wants to bring high-adrenaline racing to fans all around the world. We listen to our fans, and they want enhanced video streams that include predictions of events within the race (e.g., overtaking). Our current platform allows us to predict race outcomes but lacks the facility to support real-time predictions during races and the capacity  to process season-long results.

我们的首席执行官 S. Hawke 希望将高肾上腺素赛车带给全世界的车迷。我们倾听粉丝的心声，他们想要增强的视频流，包括对比赛中事件的预测（例如，超车）。我们的当前平台 允许我们预测比赛结果但缺乏在比赛期间支持实时预测的设施和容量 到  处理整个赛季的结果。

### Basic evaluation

### 基本评价

Client

客户

Helicopter Racing League (HRL) is a global sports league for competitive helicopter racing. HRL offers a paid service to stream the races all over the world with live telemetry and predictions throughout each race.

直升机竞速联盟 (HRL) 是一个全球的竞技直升机竞速体育联盟。 HRL 提供付费服务，通过实时遥测和全程预测每场比赛流世界各地的比赛。

Values

价值观

- Video recording and editing is performed at the race tracks, and the content is encoded and transcoded, where needed, in the cloud.

- 在赛道上进行视频录制和编辑，并根据需要在云端对内容进行编码和转码。

- Real time predictions to users

- 对用户的实时预测

Immediate Goals

即时目标

- Large data processing (season long)

- 大数据处理（季节长）

- Increase fan base and provide low latency viewing experience

- 增加粉丝群并提供低延迟观看体验

- Increase telemetry and create additional insights

- 增加遥测并创造额外的见解

### Technical evaluation

### 技术评测

Requirements

要求

Content Serving from regions closer to the viewer.

从更接近观众的地区提供内容。

Technical Watchpoints

技术观察点

- Both real time and recorded

- 实时和记录

Proposed Solution

建议的解决方案

- Use GCS to store the raw and  transcoded content (Multi-Region)

- 使用 GCS 存储原始和转码内容（多区域）

- Use CDN to deliver content

- 使用 CDN 传送内容

Requirements

要求

Datamart for batch data processing

用于批量数据处理的 Datamart

Technical Watchpoints

技术观察点

- Season long data (possibly petabyte scale)

- 季节长数据（可能是 PB 级）

Proposed Solution

建议的解决方案

- GCS can be used to store the unstructured video data. After transcoding, the data can be uploaded back to GCS and then to AI Platform for predictions (Offline Mode)

- GCS 可用于存储非结构化视频数据。转码后数据可以上传回GCS，再上传到AI平台进行预测（离线模式）

Requirements

要求

Real time analytics and predictions

实时分析和预测

Technical Watchpoints

技术观察点

- Increased telemetry

- 增加遥测

- Real time predictions and analytics

- 实时预测和分析

Proposed Solution

建议的解决方案

- Pub/Sub and Cloud Dataflow for ingestion and processing

- Pub/Sub 和 Cloud Dataflow 用于摄取和处理

- AI Platform can be used for modeling and real time insights generation (also an option since they are already using Tensorflow)

- AI Platform 可用于建模和实时洞察生成（也是一个选项，因为他们已经在使用 Tensorflow）

Requirements

要求

Real time Video Analysis and Transcoding performance

实时视频分析和转码性能

Technical Watchpoints

技术观察点

- Today the transcoding & encoding is done as required on the cloud

- 今天转码和编码是在云端按要求完成的

Proposed Solution

建议的解决方案

- GCS to store the videos

- GCS 存储视频

- Cloud Function is triggered each time content is uploaded

- 每次上传内容都会触发云功能

- CF calls the Transcoder API

- CF 调用Transcoder API

- Transcoded video stored in GCS in new bucket

- 转码后的视频存储在新存储桶中的 GCS 中

- Video Intelligence API (for video analysis) triggered  via CF

- 通过 CF 触发的视频智能 API（用于视频分析）

Requirements

要求

Expose Prediction Model to Partners

向合作伙伴公开预测模型

Proposed Solution

建议的解决方案

- Use Apigee or Cloud Endpoints to expose prediction model to partners, Apigee if it requires monetization

- 使用 Apigee 或 Cloud Endpoints 向合作伙伴公开预测模型，如果需要获利，则使用 Apigee

Requirements

要求

Minimize Operational Complexity

最大限度地降低运营复杂性

Technical Watchpoints

技术观察点

- Using trucks to stream data from race location and performing trans/encoding services on the cloud

- 使用卡车从比赛地点传输数据并在云端执行转码/编码服务

Proposed Solution

建议的解决方案

- Cloud Operations suite for SRE

- 适用于 SRE 的云操作套件

Products: Cloud Operations, Apigee, Video Intelligence API, Transcoder API, Cloud Functions, GCS, Pub/Sub, Dataflow, AI Platform, CDN

产品：云操作、APIgee、视频智能 API、转码器 API、云函数、GCS、Pub/Sub、Dataflow、AI 平台、CDN


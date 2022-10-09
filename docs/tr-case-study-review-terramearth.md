# Case Study Review: TerramEarth

# 案例研究回顾：TerramEarth

https://thecertsguy.com/bytes/case-study-review-terramearth

Dec 21

## TerramEarth

TerramEarth manufactures heavy equipment for the mining and agricultural industries. They currently have over 500 dealers and service centers in 100 countries. Their mission is to build products that make their customers more productive.

TerramEarth 制造用于采矿和农业的重型设备。目前，他们在 100 个国家/地区拥有超过 500 家经销商和服务中心。他们的使命是打造能够让客户提高生产力的产品。

### Solution concept

### 解决方案概念

There are 2 million TerramEarth vehicles in operation currently, and we see 20% yearly growth. Vehicles collect telemetry data from many sensors during operation. A small subset of critical data is transmitted from the vehicles in real time to facilitate fleet management. The rest of the sensor data is collected, compressed, and uploaded daily when the vehicles return to home base. Each vehicle usually generates 200 to 500 megabytes of data per day.

目前有 200 万辆 TerramEarth 车辆在运营，我们看到 20% 的年增长率。车辆在运行期间收集遥测 数据 来自许多传感器。 一小部分关键数据从车辆实时传输，以促进车队管理。 传感器数据的其余在车辆返回基地时收集、压缩和上传每天。 每辆车通常产生每天 200 到 500 兆字节的数据。

### Existing Technical Environment

### 现有技术环境

TerramEarth’s vehicle data aggregation and analysis infrastructure resides in Google Cloud and serves clients from all around the world. A growing amount of sensor data is captured (IoT Core) from their two main manufacturing plants and sent to private data centers that contain their legacy inventory and logistics management systems. The private data centers have multiple network interconnects configured to Google Cloud. The web frontend for dealers and customers is running in Google Cloud and allows access to stock management and analytics.

TerramEarth 的车辆数据聚合和分析基础架构位于 Google Cloud 中，并为来自世界各地的客户提供服务。 传感器数据量增长量从他们的两个主要制造工厂捕获（物联网核心），并发送到私人数据中心，其中包含他们的遗留库存和物流管理系统。 私有数据中心有多个网络互连配置到 Google Cloud。面向经销商和客户的 Web 前端 在 Google Cloud 中运行，并允许访问库存管理和分析。

### Business Requirements

###  业务需求

- Predict and detect vehicle malfunction and rapidly ship parts to dealerships for just-in time repair where possible. (AI Platform)

- 预测和检测车辆故障和快速将零件运送到经销商处，以便在可能的情况下进行及时维修。 （人工智能平台）

- Decrease cloud operational costs and adapt to seasonality. (Managed Services)

- 降低云运营成本和适应季节性。 （管理服务）

- Increase speed and reliability of development workflow. (CI/CD)

- 提高速度和开发工作流程的可靠性。 (CI/CD)

- Allow remote developers to be productive without compromising code or data security. (Private Google Access, IAP with signed headers)

- 允许 remote developers 在不影响代码或 data security 的情况下实现 productive。 （私人 Google 访问，带有签名标头的 IAP）

- Create a flexible and scalable platform for developers to create custom API services for dealers and partners. (Apigee)

- 为开发人员创建一个灵活且可扩展的平台，为经销商和合作伙伴创建自定义 API 服务。 （Apigee）

### Technical requirements

###  技术要求

- Create a new abstraction layer for HTTP API access to their legacy systems to enable a gradual move into the cloud without disrupting operations. (Apigee)

- 为 HTTP API 访问创建一个新的抽象层到他们的遗留系统，以便在不中断操作的情况下逐步迁移到云中。 （Apigee）

- Modernize all CI/CD pipelines to allow developers to deploy container-based workloads in highly scalable environments. (GKE)

- 现代化所有 CI/CD 管道，以允许开发人员在高度可扩展的环境中部署基于容器的工作负载。 (GKE)

- Allow developers to run experiments without compromising security and governance requirements (Separate Project/IAM) 

- 允许开发人员在不影响安全和治理要求的情况下运行实验（独立项目/IAM）

- Create a self-service portal for internal and partner developers to create new projects, request resources for data analytics jobs, and centrally manage access to the API endpoints. (IAM, Apigee)

- 为内部和合作伙伴开发人员创建一个自助服务门户，以创建新项目、为数据分析作业请求资源并集中管理对 API 端点的访问。 （IAM，Apigee）

- Use cloud-native solutions for keys and secrets management and optimize for identity based access (Cloud KMS, Secret Manager)

- 使用云原生解决方案进行密钥和机密管理并针对基于身份的访问进行优化（Cloud KMS、Secret Manager）

- Improve and standardize tools necessary for application and network monitoring and troubleshooting (Cloud Operations)

- 改进和标准化应用程序和网络监控和故障排除（云操作）所需的工具

### Executive statement

### 执行声明

Our competitive advantage has always been our focus on the customer, with our ability to provide excellent customer service and minimize vehicle downtimes. After moving multiple systems into Google Cloud, we are seeking new ways to provide best-in-class online fleet management services to our customers and improve operations of our dealerships. Our 5-year strategic plan is to create a partner ecosystem of new products by enabling access to our data, increasing autonomous operation capabilities of our vehicles, and creating a path to move the remaining legacy systems to the cloud.

我们的竞争优势始终是我们以客户为中心，我们有能力提供出色的客户服务并最大限度地减少车辆停机时间。在将多个系统迁移到 Google Cloud 之后，我们正在寻找新的方法来为我们的客户提供一流的在线车队管理服务，并改善我们经销商的运营。我们的 5 年战略计划是通过允许访问我们的数据、提高我们车辆的自主操作能力以及创建将剩余的遗留系统迁移到云。

### Basic evaluation

### 基本评价

Client 客户

TerramEarth manufactures heavy equipment for the mining and agricultural industries. They currently have over 500 dealers and service centers in 100 countries. Their mission is to build products that make their customers more productive.

TerramEarth 制造用于采矿和农业的重型设备。目前，他们在 100 个国家/地区拥有超过 500 家经销商和服务中心。他们的使命是打造能够让客户提高生产力的产品。

Values 价值观

- Already on GCP

- 已经在 GCP 上

- Multiple network interconnects in place between OnPrem and GCP

- OnPrem 和 GCP 之间存在多个网络互连

- Web Front end running on GCP

- 在 GCP 上运行的 Web 前端

Immediate Goals 

即时目标

- Minimize Vehicle Downtimes

- 最大限度地减少车辆停机时间

- Provide best in class online Fleet management services

- 提供一流的在线车队管理服务

- Improve dealership

- 改善经销商

### Technical evaluation

### 技术评测

Requirements

要求

Predict and detect vehicle malfunction and rapidly ship parts to dealerships for just-in time repair

预测和检测车辆故障和快速将零件运送到经销商以便及时维修

Technical Watchpoints

技术观察点

- The web frontend for dealers and customers is running in Google Cloud and allows access to stock management and analytics.

- 面向经销商和客户的网络前端在 Google Cloud 中运行，并允许访问库存管理和分析。

Proposed Solution

建议的解决方案

- Use AI Platform to create prediction models

- 使用 AI Platform 创建预测模型

- BigQuery for handling real time data to facilitate fleet management

- BigQuery 用于处理实时数据以促进车队管理

Requirements

要求

Decrease cloud operational costs and adapt to seasonality

降低云运营成本和适应季节性

Proposed Solution

建议的解决方案

- IoT Core, Pub/Sub and Dataflow as we need to decouple the messages ingestion and processing

- IoT Core、Pub/Sub 和 Dataflow 因为我们需要解耦消息摄取和处理

Requirements

要求

Increase speed and reliability of development workflow.

提高开发工作流程的速度和可靠性。

Technical Watchpoints

技术观察点

- Modernize all CI/CD pipelines

- 现代化所有 CI/CD 管道

- keys and secrets management and optimize for identity based access

- 密钥和秘密管理并针对基于身份的访问进行优化

- Standardize tools necessary for application and network monitoring and troubleshooting

- 标准化应用程序和网络监控和故障排除所需的工具

Proposed Solution

建议的解决方案

- Modernize CI/CD with Cloud Build and Deployment Manager

- 使用 Cloud Build 和 Deployment Manager 实现 CI/CD 现代化

- Cloud KMS, Secret Manager

- 云 KMS、秘密管理器

- Cloud Operations to capture Audit and Network Logs (VPC Flow Logs) Network Intelligence to monitor performance and topology

- 云操作捕获审计和网络日志（VPC 流日志）网络智能以监控性能和拓扑

Requirements

要求

Remote developer productivity

远程开发人员生产力

Technical Watchpoints

技术观察点

- Allow developers to run experiments without compromising security and governance requirements

- 允许开发人员在不影响安全和治理要求的情况下运行实验

Proposed Solution

建议的解决方案

- Use Identity Aware Proxies ( IAP) and host the sandbox project in a separate folder with appropriate policies in place (IAM and network policies)

- 使用 Identity Aware Proxies (IAP) 并将沙盒项目托管在单独的文件夹中，并配备适当的策略（IAM 和网络策略）

Requirements

要求

Create custom API services for dealers and partners

为经销商和合作伙伴创建自定义 API 服务

Technical Watchpoints

技术观察点

- Create a new abstraction layer for HTTP API access to their legacy systems

- 为 HTTP API 访问 创建一个新的抽象层 到他们的 legacy 系统

- Self-service portal to create projects, request resources for analytics jobs, and centrally manage APIs

- 用于创建项目、请求分析作业资源和集中管理 API 的自助服务门户

Proposed Solution

建议的解决方案

- Apigee as central portal for API access management, self service and monetization.

- Apigee 作为 API 访问管理、自助服务和货币化的中央门户。

- GKE as backend service aggregating On-Prem data and Analytic data, requesting analytics jobs

- GKE 作为后端服务聚合本地数据和分析数据，请求分析工作

Requirements

要求

Interconnect with private data center

与私有数据中心互联

Proposed Solution

建议的解决方案

- Cloud Router + Interconnect(Partner or Dedicated) for interconnect with private datacenter 

- 云路由器+互连（合作伙伴或专用）用于与私有数据中心互连

Products: AI Platform (Predictions), IoT Core (for managing devices and creating a bridge to stream data), Pub/Sub (as endpoint to ingest streaming data from IoT devices) Dataflow (for processing), Terraform and Deployment Manager ( CI/CD), Cloud KMS & Secrets Manager (reliability of dev workflow), Cloud Operations Suite (for monitoring), Cloud IAM and IAP (Remote dev productivity), Apigee (API layer for access to Legacy systems and self service portal)

产品：AI 平台（预测）、IoT Core（用于管理设备和创建流数据的桥梁）、Pub/Sub（作为从 IoT 设备摄取流数据的端点）Dataflow（用于处理）、Terraform 和部署管理器（ CI/CD）、Cloud KMS 和 Secrets Manager（开发工作流的可靠性）、Cloud Operations Suite（用于监控）、Cloud IAM 和 IAP（远程开发效率）、Apigee（用于访问旧系统和自助服务门户的 API 层)


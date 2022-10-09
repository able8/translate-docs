# Case Study Review: EHR Healthcare

# 案例研究回顾：EHR 医疗保健

https://thecertsguy.com/bytes/case-study-review-ehr-healthcare

Dec 21

As most of you know by now, the Google PCA (Professional Cloud Architect) exam was revamped on May 1st, 2021. With the new version of the exam, and having cleared it myself last month, I noticed some significant changes. Some of the key changes from the previous version of the exam are:

正如你们大多数人现在所知，Google PCA（专业云架构师）考试于 2021 年 5 月 1 日进行了改版。随着新版本的考试，并且上个月我自己通过了考试，我注意到了一些重大变化。与以前版本的考试相比，一些主要变化是：

- The questions are more conceptual than straightforward
- Introduction of new areas like Anthos and MLOps
- Longer questions
- Multiple services being tested on a question (like a true architect!)
- All new case studies
- 这些问题比简单的更概念化
- 引入 Anthos 和 MLOps 等新领域
- 更长的问题
- 在一个问题上测试多个服务（就像一个真正的建筑师！）
- 所有新案例研究

In this blog post, I will outline how I went about solving the new case studies. I will post the exact document I wrote, and which since May 14th 2021, over 240 Googlers across the globe have used as part of their exam prep. I want to thank the many Googlers who took time to comment and improve the document to get it to its current state. Big shout out to Iman for allowing me to post this on his amazing website. I hope this material will help in your prep as well.

在这篇博文中，我将概述我是如何解决新案例研究的。我将发布我写的确切文件，自 2021 年 5 月 14 日以来，全球 240 多名 Google 员工已将其用作考试准备的一部分。我要感谢许多花时间评论和改进文档以使其达到当前状态的 Google 员工。向伊曼大喊大叫，让我将其发布在他令人惊叹的网站上。我希望这些材料也能帮助你做好准备。

If you haven't already, please read the [exam deep dive](https://thecertsguy.com/bytes/how-to-pass-the-google-cloud-professional-cloud-architect-exam-in-30-days) to understand the overall strategy and key objectives to study for the Professional Cloud Architect exam.

如果您还没有，请阅读 [考试深入了解](https://thecertsguy.com/bytes/how-to-pass-the-google-cloud-professional-cloud-architect-exam-in-30-天)了解学习专业云架构师考试的总体策略和关键目标。

All the best!

一切顺利！

## EHR healthcare

## EHR 医疗保健

EHR Healthcare is a leading provider of electronic health record software to the medical industry. EHR Healthcare provides their software as a service to multi-national medical offices, hospitals, and insurance providers.

EHR Healthcare 是医疗行业领先的电子健康记录软件提供商。 EHR Healthcare 将他们的软件作为服务提供给跨国医疗办公室、医院和保险提供商。

### Solution concept

### 解决方案概念

Due to rapid changes in the healthcare and insurance industry, EHR Healthcare’s business has been growing exponentially year over year. They need to be able to scale their environment (k8s), adapt their disaster recovery plan, and roll out new continuous deployment capabilities (CD) to update their software at a fast pace. Google Cloud has been chosen to replace their current colocation facilities

由于医疗保健和保险行业的快速变化，EHR Healthcare 的业务一直呈指数级增长年复一年。他们需要能够扩展他们的环境 (k8s)，调整他们的灾难恢复计划，并推出新的持续部署功能 (CD) 以快速更新他们的软件。已选择 Google Cloud 替换其当前的托管设施

### Existing Technical Environment

### 现有技术环境

EHR’s software is currently hosted in multiple colocation facilities. The lease on one of the data centers is about to expire. Customer-facing applications are web-based, and many have recently been containerized to run on a group of Kubernetes clusters. Data is stored in a mixture of relational and NoSQL databases (MySQL, MS SQL Server, Redis, and MongoDB).

EHR 的软件当前托管在多个托管设施中。 一个数据中心的租约即将到期。面向客户的应用程序是基于网络的，许多最近被容器化以在一组Kubernetes集群上运行。 数据存储在混合的关系数据库和 NoSQL 数据库（MySQL、MS SQL Server、Redis 和 MongoDB）中。

EHR is hosting several legacy file- and API-based integrations with insurance providers on-premises. These systems are scheduled to be replaced over the next several years. There is no plan to upgrade or move these systems at the current time. Users are managed via Microsoft Active Directory (ADFS, GCDS). Monitoring is currently being done via various open source tools (Cloud Monitoring, open source, so Prometheus...etc). Alerts are sent via email and are often ignored.

EHR 与保险提供商在本地托管多个基于文件和 API 的旧集成。这些系统计划在未来几年内更换。 目前没有升级或移动这些系统的计划。 用户通过 Microsoft Active Directory（ADFS、GCDS）进行管理。 监控目前正在通过各种开源工具（云监控、开源、普罗米修斯等）完成。 警报通过电子邮件发送，并且经常被忽略。

### Business Requirements

###  业务需求

- On-board new insurance providers as quickly as possible

- 尽快尽快加入新的保险提供商

- Provide a minimum 99.9% availability for all customer-facing systems

- 为所有面向客户的系统提供至少 99.9% 的可用性

- Provide centralized visibility and proactive action on system performance and usage (Cloud Operations/Stackdriver)

- 提供集中的可见性和主动行动对系统性能和使用（云操作/Stackdriver）

- Increase ability to provide insights into healthcare trends (Bigquery, AI services)

- 提高 提供洞察力 进入医疗保健 趋势（Bigquery、AI 服务）

- Reduce latency to all customers (CDN, Multi-Region Deployments)

- 为所有客户减少延迟 （CDN，多区域部署）

- Maintain regulatory compliance (HIPAA, Tokenize PII Data)

- 保持 监管合规性（HIPAA，标记 PII 数据）

- Decrease infrastructure administration costs (Managed Services)

- 减少基础设施 管理 成本（托管服务）

- Make predictions and generate reports on industry trends based on provider data (BigQuery, BQML, Datastudio, Looker)

- 根据提供商数据对行业趋势进行预测并生成报告（BigQuery、BQML、Datastudio、Looker）

### Technical requirements

###  技术要求

- Maintain legacy interfaces to insurance providers with connectivity to both on-premises systems and cloud providers (VPN/Interconnect, BYOIP)

- 维护与保险提供商的旧接口，连接到本地系统和云提供商（VPN/互连，BYOIP）

- Provide a consistent way to manage customer-facing applications that are container-based (Anthos)

- 提供一致的方式来管理面向客户的应用程序，这些应用程序基于容器 (Anthos)

- Provide a secure and high-performance connection between on-premises systems and Google Cloud (Dedicated Interconnect) 

- 在本地系统和 Google Cloud 之间提供安全和高性能的连接（专用互连）

- Provide consistent logging, log retention, monitoring, and alerting capabilities (Anthos Service Mesh and Cloud Operations Suite)

- 提供一致的日志记录、日志保留、监控和警报功能（Anthos Service Mesh 和 Cloud Operations Suite）

- Maintain and manage multiple container-based environments (Anthos)

- 维护和管理 多个基于容器的环境（Anthos）

- Dynamically scale and provision new environments (IaC-Terraform)

- 动态扩展和配置新环境（IaC-Terraform）

- Create interfaces to ingest and process data from new providers (Pub/Sub, Dataflow)

- 创建 interfaces 到 ingest 和 process data from new provider (Pub/Sub, Dataflow)

### Executive statement

### 执行声明

Our on-premises strategy has worked for years but has required a major investment of time and money in training our team on distinctly different systems, managing similar but separate environments, and responding to outages. Many of these outages have been a result of misconfigured systems, inadequate capacity to manage spikes in traffic, and inconsistent monitoring practices. We want to use Google Cloud to leverage a scalable, resilient platform that can span multiple environments seamlessly and provide a consistent and stable user experience that positions us for future growth.﻿

我们的本地策略已经行之有效，但需要大量时间和金钱来培训我们的团队在明显不同的系统上，管理类似的但是独立的环境和响应中断。其中许多中断是由于配置错误的系统、管理流量高峰的能力不足以及不一致的监控实践造成的。我们希望使用 Google Cloud 来利用一个可扩展、有弹性的平台，该平台可以无缝地跨越多个环境并提供一致且稳定的用户体验，从而为我们未来的发展做好准备。



---

### Basic evaluation

### 基本评价

1) Client

1) 客户

- Leading provider of EHR software services to multi-national medical offices, hospitals, and insurance providers. Currently infrastructure is on-prem, but co-lo is ending on one DC.

- 为跨国医疗机构、医院和保险提供商提供 EHR 软件服务的领先供应商。目前基础架构在本地，但 co-lo 将在一个 DC 上结束。

2) Values

2) 价值观

- Already refactored application on-prem and containerized them

- 已经在本地重构了应用程序并将它们容器化

- Significant investment in managing services on prem

- 对本地管理服务进行大量投资

3) Immediate Goals

3) 即时目标

- Reduce management of on-prem services

- 减少本地服务的管理

- Rapidly scale

- 快速扩展

- Improve SRE

- 提高 SRE

- Introduce DevOps

- 介绍 DevOps

### Technical evaluation

### 技术评测

Existing Environment

现有环境

Application Servers:

应用服务器：

Customer-facing applications are web-based, and many have recently been containerized to run on a group of Kubernetes clusters.

面向客户的应用程序是基于 Web 的，许多最近已经被容器化以在一组 Kubernetes 集群上运行。

Technical Watchpoints

技术观察点

- hosting several legacy file- and API-based integrations

- 托管几个遗留的基于文件和 API 的集成

- Customer-facing applications are web-based, and many have recently been containerized

- 面向客户的应用程序是基于 Web 的，并且许多最近已被容器化

Proposed Solution

建议的解决方案

- Migrate legacy apps which are not containerized using Migrate for Anthos

- 使用 Migrate for Anthos 迁移未容器化的旧版应用程序

- Use Anthos Service Mesh and Anthos Config Management to view and manage all containers through a single pane of glass

- 使用 Anthos Service Mesh 和 Anthos Config Management 通过单一管理平台查看和管理所有容器

- Cloud CDN for reduced latency

- 云 CDN 以减少延迟

Existing Environment

现有环境

Databases:

数据库：

Mix of relational & Non-relational DBs

关系数据库和非关系数据库的混合

Technical Watchpoints

技术观察点

- MySQL

- MySQL

- MS SQL Server

- 微软 SQL 服务器

- Redis

- 雷迪斯

- MongoDB

- MongoDB

Proposed Solution

建议的解决方案

- Cloud SQL (both MySQL and SQL Server) - If they have unique MS SQL Server requirements like using SSMS, they may need to run that on a GCE instance. Remember Cloud SQL parity isn't always 1:1

- Cloud SQL（MySQL 和 SQL Server）- 如果他们有独特的 MS SQL Server 要求，例如使用 SSMS，他们可能需要在 GCE 实例上运行。请记住 Cloud SQL 奇偶校验并不总是 1:1

- Cloud Memorystore

- 云记忆库

- Datastore or managed instance of MongoDB through Cloud marketplace

- 数据存储或通过云市场托管的 MongoDB 实例

Existing Environment

现有环境

Analytics & Reporting:

分析和报告：

Nothing mentioned but has future needs

没有提到，但有未来的需求

Technical Watchpoints

技术观察点

- Needs ingestion and processing capabilities from multiple provider sources

- 需要来自多个供应商来源的摄取和处理能力

- Needs ability to make predictions and create reports

- 需要做出预测和创建报告的能力

Proposed Solution

建议的解决方案

- Pub/Sub and Cloud Dataflow for ingestion and processing

- Pub/Sub 和 Cloud Dataflow 用于摄取和处理

- BigQuery for predictions and Datastudio/Looker for reporting

- BigQuery 用于预测，Datastudio/Looker 用于报告

- AI Platform can be used for advanced modeling and insights generation

- AI Platform 可用于高级建模和洞察生成

Existing Environment

现有环境

Identity:

身份：

Users currently on Microsoft AD

当前使用 Microsoft AD 的用户

Proposed Solution

建议的解决方案

- Use ADFS and GCDS to sync users to the Google ecosystem

- 使用 ADFS 和 GCDS 将用户同步到 Google 生态系统

Existing Environment

现有环境

Monitoring & SRE:

监控和 SRE：

Open Source tools used

使用的开源工具

Technical Watchpoints

技术观察点

- Several issues in the past due to misconfigured systems, inadequate capacity to manage spikes in traffic, and inconsistent monitoring practices

- 过去由于系统配置错误、管理流量高峰的能力不足以及监控实践不一致而导致的几个问题

Proposed Solution

建议的解决方案

- Cloud Logging to capture Audit and Network logs

- Cloud Logging 用于捕获审核和网络日志

- Store logs in Cloud Storage for compliance purposes and apply ACLs

- 出于合规目的将日志存储在云存储中并应用 ACL

- Cloud Monitoring using workspaces

- 云监控使用工作区

- Anthos Service Mesh to monitor multi platform container environment

- Anthos Service Mesh 监控多平台容器环境

- Develop SRE practices such as Chaos Engineering and Istio Fault Injection to test out systems

- 开发 SRE 实践，例如 Chaos Engineering 和 Istio Fault Injection 以测试系统

- Refactor 4 golden “ LETS” metrics: Latency, Errors, Traffic, Saturation with percentiles to reduce alert fatigue

- 重构 4 个黄金“LETS”指标：延迟、错误、流量、带百分位数的饱和度以减少警报疲劳

Existing Environment

现有环境

Dynamic Scaling & Provisioning of features

动态扩展和配置功能

Technical Watchpoints

技术观察点

- Need to roll out new continuous deployment capabilities to update their software at a fast pace

- 需要推出新的持续部署功能以快速更新他们的软件

Proposed Solution

建议的解决方案

- Create a CI/CD pipelines with Jenkins (CI) and Spinnaker (CD)

- 使用 Jenkins (CI) 和 Spinnaker (CD) 创建 CI/CD 管道

- Use Terraform to create Infrastructure as Code. Cloud Build for provisioning infra. 

- 使用 Terraform 创建基础设施即代码。 Cloud Build 用于配置基础设施。

Products: Anthos (managing containerized apps both on-prem and on cloud), Migrate for Anthos (for moving their existing on-prem apps),  Cloud CDN (low latency), Cloud SQL (RDBMS), Memorystore (managed Redis) , Pub/Sub (ingest data for asynchronous processing), Dataflow (ETL), BQ (analysis and predictions), Datastudio/Looker (reporting), AI Platform (insight generation), ADFS/GCDS (for AD identity sync), Cloud Operations Suite and Anthos Service Mesh and Cloud Operations (logging & monitoring), Terraform (IaC), Jenkins (CI), Spinnaker (CD)

产品：Anthos（管理本地和云端的容器化应用）、Migrate for Anthos（用于移动其现有的本地应用）、Cloud CDN（低延迟）、Cloud SQL (RDBMS)、Memorystore（托管 Redis） 、Pub/Sub（为异步处理提取数据）、Dataflow (ETL)、BQ（分析和预测）、Datastudio/Looker（报告）、AI Platform（洞察力生成）、ADFS/GCDS（用于 AD 身份同步）、云操作Suite 和 Anthos 服务网格和云操作（日志记录和监控）、Terraform (IaC)、Jenkins (CI)、Spinnaker (CD)


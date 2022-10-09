# Case Study Review: Mountkirk Games
https://thecertsguy.com/bytes/case-study-review-mountkirk-games

# 案例研究回顾：Mountkirk Games
Sep 20

## Mountkirk Games

## 芒特柯克游戏

Mountkirk Games makes online, session-based, multiplayer games for mobile platforms. They have recently started expanding to other platforms after successfully migrating their on-premises environments to Google Cloud. Their most recent endeavor is to create a retro-style first-person shooter (FPS) game that allows hundreds of simultaneous players to join a geo-specific digital arena from multiple platforms and locations.  A real-time digital banner will display a global leaderboard of all the top players across every active arena.

Mountkirk Games 为移动平台制作在线、基于会话的多人游戏。在成功将迁移他们的本地环境到 Google Cloud 后，他们最近开始扩展到其他平台。他们最近的努力是创建一款复古风格的第一人称射击 (FPS) 游戏，该游戏允许数百名玩家同时加入一个特定地理位置的数字竞技场来自多个平台和位置。 实时数字横幅将显示每个活跃竞技场中所有顶级玩家的全球排行榜。

The existing environment was recently migrated to Google Cloud, and five games came across using lift-and-shift virtual machine migrations, with a few minor exceptions. Each new game exists in an isolated Google Cloud project nested below a folder that maintains most of the permissions and network policies. Legacy games with low traffic have been consolidated into a single project. There are also separate environments for development and testing.

现有环境最近迁移到了 Google Cloud，有五款游戏使用直接迁移虚拟机迁移，但有一些小例外。 每个新的游戏都存在于一个隔离的 Google Cloud 项目中，该项目嵌套在一个文件夹下方，该文件夹维护着大部分权限和网络政策。低流量的传统游戏已合并到一个项目中。还有用于开发和测试的独立环境。

### Solution concept

### 解决方案概念

Mountkirk Games is building a new multiplayer game that they expect to be very popular. They plan to deploy the game’s backend on Google Kubernetes Engine so they can scale rapidly and use Google’s global load balancer to route players to the closest regional game arenas. In order to keep the global leader board in sync, they plan to use a multi-region Spanner cluster.

Mountkirk Games 正在开发一款新的多人游戏，他们预计该游戏会非常受欢迎。他们计划在 Google Kubernetes Engine 上部署游戏的后端，以便他们可以快速扩展并使用 Google 的全球负载均衡器将玩家路由到最近的区域游戏竞技场。为了让全球排行榜保持同步，他们计划使用多区域 Spanner 集群。

### Business Requirements

###  业务需求

- Support multiple gaming platforms (mobile, desktop, tablets)

- 支持多种游戏平台（手机、台式机、平板电脑）

- Support multiple regions (Clusters in multiple regions)

- 支持多区域（多区域集群）

- Support rapid iteration of game features (CI/CD)

- 支持游戏功能快速迭代(CI/CD)

- Minimize latency (CDN)

- 最小化延迟（CDN）

- Optimize for dynamic scaling

- 优化动态缩放

- Use managed services and pooled resources (managed services)

- 使用托管服务和资源池（托管服务）

- Minimize costs (managed services)

- 最小化成本（托管服务）

### Technical requirements

###  技术要求

- Dynamically scale based on game activity (k8s)

- 根据游戏活动动态扩展（k8s）

- Publish scoring data on a near real–time global leaderboard (Memorystore)

- 在近乎实时的全球排行榜上发布得分数据（Memorystore）

- Store game activity logs in structured files for future analysis (Cloud Logging, GCS)

- 将游戏活动日志存储在结构化文件中以供将来分析（Cloud Logging、GCS）

- Use GPU processing to render graphics server-side for multi-platform support (k8s)

- 使用 GPU 处理在服务器端渲染图形以支持多平台 (k8s)

- Support eventual migration of legacy games to this new platform

- 支持旧游戏最终迁移到这个新平台

### Executive statement 

### 执行声明

Our last game was the first time we used Google Cloud, and it was a tremendous success. We were able to analyze player behavior and game telemetry in ways that we never could before. This success allowed us to bet on a full migration to the cloud and to start building all-new games using cloud-native design principles. Our new game is our most ambitious to date and will open up doors for us to support more gaming platforms beyond mobile. Latency is our top priority, although cost management is the next most important challenge. As with our first cloud-based game, we have grown to expect the cloud to enable advanced analytics capabilities (BQ) so we can rapidly iterate on our deployments of bug fixes and new functionality (Cloud Build).

我们的上一款游戏是我们第一次使用 Google Cloud，它取得了巨大的成功。我们能够以前所未有的方式分析玩家行为和游戏遥测。这一成功让我们押注于全面迁移到云端，并开始使用云原生设计原则构建全新游戏。我们的新游戏是我们迄今为止最雄心勃勃的游戏，它将为我们打开大门，以支持除移动设备之外的更多游戏平台。 延迟是我们的首要任务，尽管成本管理是下一个最重要的挑战。与我们的第一个基于云的游戏一样，我们已经成长为期望云能够支持高级分析功能 (BQ)，以便我们可以快速迭代我们的错误修复和新功能的部署 （云构建）.

### Basic evaluation

### 基本评价

1) Client

1) 客户

Online game platform

网络游戏平台

- Previous success on GCP has led to an ambitious project to create a new multiplayer game with the aim to support more gaming platforms

- 之前在 GCP 上的成功促成了一个雄心勃勃的项目，即创建一款新的多人游戏，旨在支持更多游戏平台

2. Values

   价值观

- Already have a plan in place with general design for the infrastructure and some requirements.

- 已经制定了基础设施的总体设计和一些要求的计划。

- Wants to adopt CI/CD and expand to multiple gaming platforms.

- 想要采用 CI/CD 并扩展到多个游戏平台。

3. Immediate Goals

   即时目标

- Reduce latency and support larger gamer footprint

- 减少延迟并支持更大的游戏玩家足迹

- Keep costs low

- 保持低成本

- Enable advanced analytics

- 启用高级分析

- Enable rapid deployments

- 实现快速部署

### Technical evaluation

### 技术评测

1) Existing Environment

1) 现有环境

- Games hosted on VMs

- 在虚拟机上托管的游戏

- Legacy games consolidated into a single project

- 遗留游戏合并到一个项目中

- Separate Environments for dev and test

- 单独的开发和测试环境

2) Technical Watchpoints & Proposed Solutions

2) 技术观察点和建议的解决方案

- 100s of users joining simultaneously from across the globe
- 来自全球的 100 名用户同时加入
   
- Create K8 backends in multiple regions (regional replication)
- Use Global HTTPS LB to balance traffic

- 在多个区域创建K8 后端（区域复制）
- 使用全局 HTTPS LB 来平衡流量

- Store user profiles in Cloud Datastore
- Global leaderboard to update in real time

- 将用户配置文件存储在 Cloud Datastore 中
- 全球排行榜实时更新

- Cloud Memorystore is the best solution for low latency storage of leaderboard results. Spanner is high cost and not required here.
- Store game activity logs in structured files for future analysis

- Cloud Memorystore 是排行榜结果低延迟存储的最佳解决方案。扳手成本高，这里不需要。
- 将游戏活动日志存储在结构化文件中以供将来分析

- Game activity logs can be logged using Cloud Logging, and then stored on GCS. These logs can later be analyzed using BQ when needed
- Support rapid iteration of game features

- 可以使用 Cloud Logging 记录游戏活动日志，然后存储在 GCS 上。以后可以在需要时使用 BQ 分析这些日志
- 支持游戏功能的快速迭代

- Mountkirk games should implement a DevOps process ( Terraform & Cloud Build) to increase developer productivity and enable rapid iteration of game features

- Mountkirk 游戏应实施 DevOps 流程（Terraform & Cloud Build）以提高开发人员的生产力并实现游戏功能的快速迭代

Products: k8s (Game server), Cloud datastore (User Profiles), Memorystore (leaderboards), Cloud Logging (Game activity logs), GCS (Log storage for analysis), Terraform & Cloud Build (rapid iterations on game features)﻿  ﻿

产品：k8s（游戏服务器）、Cloud datastore（用户配置文件）、Memorystore（排行榜）、Cloud Logging（游戏活动日志）、GCS（用于分析的日志存储）、Terraform & Cloud Build（游戏功能的快速迭代）﻿ ﻿

Stay tuned for case study reviews on each of the four business cases on the exam. In the meantime, don’t forget to check out the exam deep dive! Best of luck to you throughout your studies, you’ll do GREAT! 

请继续关注有关考试中四个业务案例的案例研究评论。同时，不要忘记检查考试深度潜水！祝你在整个学习过程中好运，你会做得很好！


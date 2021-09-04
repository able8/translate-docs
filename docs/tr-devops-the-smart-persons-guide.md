# DevOps: A cheat sheet

# DevOps：备忘单

on November 2, 2018, 3:38 AM PST

This comprehensive guide covers DevOps, an increasingly popular organizational structure for delivering rapid software deployments in the enterprise.

本综合指南涵盖了 DevOps，这是一种日益流行的组织结构，用于在企业中交付快速软件部署。

Video: 3 things you should know about DevOps

视频：关于 DevOps 你应该知道的 3 件事

In the age of constant security vulnerabilities and subsequent security patches, and the rapid and frequent addition of features to software, DevOps--a workflow that emphasizes communication between software developers and IT pros managing production environments--is at the forefront when considering how to shape an IT department to fit an organization's internal needs and best serve its customers.

在安全漏洞和后续安全补丁层出不穷、软件功能快速频繁添加的时代，DevOps——一种强调软件开发人员和管理生产环境的 IT 专业人员之间沟通的工作流——在考虑如何塑造 IT 部门以适应组织的内部需求并最好地为客户服务。

### What's hot at TechRepublic

### TechRepublic 的热门话题

- [15 highest-paying certifications for 2021](https://www.techrepublic.com/article/the-15-highest-paying-certifications-for-2021/)
- [My life with Linux: A retrospective](https://www.techrepublic.com/article/my-life-with-linux-a-retrospective/)
- [Excel is still a security headache after 30 years because of this one feature](https://www.techrepublic.com/article/excel-is-still-a-security-headache-after-30-years-because-of-this-one-feature/)
- [How to clean up your Gmail inbox with this mass delete trick (free PDF)](https://www.techrepublic.com/resource-library/downloads/how-to-clean-up-your-gmail-inbox-with-this-mass-delete-trick-free-pdf/)
- [2021 年 15 项收入最高的认证](https://www.techrepublic.com/article/the-15-highest-paying-certifications-for-2021/)
- [我的 Linux 生活：回顾](https://www.techrepublic.com/article/my-life-with-linux-a-retrospective/)
- [Excel 30 年后仍然是一个安全问题，因为这个功能](https://www.techrepublic.com/article/excel-is-still-a-security-headache-after-30-years-because-一个功能/)
- [如何使用此批量删除技巧（免费 PDF）清理您的 Gmail 收件箱](https://www.techrepublic.com/resource-library/downloads/how-to-clean-up-your-gmail-inbox-with-this-mass-delete-trick-free-pdf/)

This cheat sheet is an introduction to DevOps, as well as a "living" guide that will be updated periodically as trends and methods in this field change.

这份备忘单是对 DevOps 的介绍，也是一份“活的”指南，会随着该领域的趋势和方法的变化而定期更新。

**SEE:** [**How iRobot used data science, cloud, and DevOps to design its next-gen smart home robots (cover story PDF)**](https://www.techrepublic.com/resource-library/whitepapers/how-irobot-used-data-science-cloud-and-devops-to-design-its-next-gen-smart-home-robots-cover-story-pdf/) **(TechRepublic)**

**参见：** [iRobot 如何使用数据科学、云和 DevOps 来设计其下一代智能家居机器人（封面故事 PDF）](https://www.techrepublic.com/resource-library/whitepapers/how-irobot-used-data-science-cloud-and-devops-to-design-its-next-gen-smart-home-robots-cover-story-pdf/) **(TechRepublic)**

## Executive summary

## 执行摘要

- **What is DevOps?** DevOps is an ethos centered around integration and communication between software developers and IT professionals who manage production operations.
- **Why does DevOps matter?** DevOps allows organizations to more rapidly deliver software and security updates internally and to customers.
- **Who does DevOps affect?** Because implementing DevOps requires a corporate culture change, it affects the entire company. The benefits have far outweighed the difficulty of a culture shift in companies that have adopted DevOps.
- **When did companies start using DevOps?** DevOps first gained traction in 2009. Large organizations such as Amazon, Walmart, and Adobe use DevOps.
- **How do I implement DevOps?** DevOps is not a switch that can simply be turned on; it requires careful and gradual implementation so as to not disrupt the functioning of your organization.
- **什么是 DevOps？** DevOps 是一种以软件开发人员和管理生产运营的 IT 专业人员之间的集成和沟通为中心的精神。
- **为什么 DevOps 很重要？** DevOps 允许组织更快速地在内部和向客户交付软件和安全更新。
- **DevOps 会影响谁？** 因为实施 DevOps 需要改变企业文化，它会影响整个公司。在采用 DevOps 的公司中，好处远远超过了文化转变的困难。
- **公司什么时候开始使用 DevOps？** DevOps 在 2009 年首次受到关注。亚马逊、沃尔玛和 Adobe 等大型组织使用 DevOps。
- **我如何实施 DevOps？** DevOps 不是一个可以简单打开的开关；它需要谨慎和循序渐进地实施，以免破坏您组织的运作。

**SEE: [Special report: Riding the DevOps revolution (free PDF)](https://www.techrepublic.com/resource-library/whitepapers/riding-the-devops-revolution-free-pdf/)(TechRepublic) **

**参见：[特别报告：驾驭 DevOps 革命（免费 PDF）](https://www.techrepublic.com/resource-library/whitepapers/riding-the-devops-revolution-free-pdf/)(TechRepublic) **

## What is DevOps?

## 什么是 DevOps？

DevOps (a combination of "Development" and "Operations") is an ethos that emphasizes the importance of communication and collaboration between software developers and production IT professionals while automating the deployment of software and infrastructure changes.

DevOps（“开发”和“运营”的组合）是一种精神，它强调软件开发人员和生产 IT 专业人员之间的沟通和协作的重要性，同时自动化软件和基础架构更改的部署。

**SEE: [Job description: DevOps engineer](http://www.techproresearch.com/downloads/job-description-devops-engineer/) (Tech Pro Research)**

**参见：[职位描述：DevOps 工程师](http://www.techproresearch.com/downloads/job-description-devops-engineer/) (Tech Pro Research)**

The goal of DevOps is to create a working environment in which building, testing, and deploying software can occur rapidly, frequently, and reliably. In turn, this allows for an organization to achieve its goals quicker, allowing for a faster turnaround time in the deployment of new features, security patches, and bug fixes.

DevOps 的目标是创建一个工作环境，在该环境中可以快速、频繁且可靠地构建、测试和部署软件。反过来，这使组织能够更快地实现其目标，从而缩短部署新功能、安全补丁和错误修复的周转时间。

DevOps is a common attribute of startups, as the limited headcount inherent in startups necessitates that programmers must be responsible for deployment of production implementation of software. While DevOps as a resource management strategy is also present in some larger organizations, it has also been criticized for being an inefficient use of human capital.

DevOps 是初创公司的一个共同属性，因为初创公司固有的有限人数要求程序员必须负责部署软件的生产实施。虽然 DevOps 作为一种资源管理策略也存在于一些较大的组织中，但它也因人力资本使用效率低下而受到批评。

**SEE: [All of TechRepublic's cheat sheets and smart person's guides](https://www.techrepublic.com/topic/smart-persons-guides/)**

**参见：[TechRepublic 的所有备忘单和聪明人指南](https://www.techrepublic.com/topic/smart-persons-guides/)**

DevOps encompasses the already popular programming concepts of agile development, continuous integration, and continuous delivery, and extends that ethos into the social aspect of IT by placing a premium on the importance of tearing down walls that divide development, operations, support, and management teams .

DevOps 包含已经流行的敏捷开发、持续集成和持续交付的编程概念，并通过重视拆除分隔开发、运营、支持和管理团队的墙壁的重要性，将这种精神扩展到 IT 的社会方面.

In the same vein, DevOps is a descriptive--not a prescriptive--concept. There is no single product or silver bullet that can fix existing problems in an organization; the purpose of DevOps is to increase collaboration.

同样，DevOps 是一个描述性的——而不是规定性的——概念。没有单一的产品或灵丹妙药可以解决组织中存在的问题； DevOps 的目的是增加协作。

**Additional resources:**

**其他资源：**

- [Six steps to DevOps success, analyzed](https://www.zdnet.com/article/six-steps-to-successful-devops-success-analyzed/)(ZDNet)
- [DevOps decoded: What it is and why you should care](http://www.zdnet.com/article/devops-decoded-guru-explains-what-it-is-and-why-you-should-care/) (ZDNet)
- [10 best practices for DevOps](https://www.techrepublic.com/blog/10-things/10-best-practices-for-devops/)(TechRepublic)
- [How to tell if your DevOps is delivering business value](http://www.zdnet.com/article/how-you-can-tell-if-your-devops-is-successful/)(ZDNet)
- [Video: Tips for how to become a DevOps engineer](https://www.techrepublic.com/videos/tips-for-how-to-become-a-devops-engineer/)(TechRepublic)
- [Quick glossary: DevOps](http://www.techproresearch.com/downloads/quick-glossary-devops/) (Tech Pro Research)
- [DevOps 成功的六个步骤，分析](https://www.zdnet.com/article/six-steps-to-successful-devops-success-analyzed/)(ZDNet)
- [DevOps 解码：它是什么以及为什么你应该关心](http://www.zdnet.com/article/devops-decoded-guru-explains-what-it-is-and-why-you-should-care/) (ZDNet)
- [DevOps 的 10 个最佳实践](https://www.techrepublic.com/blog/10-things/10-best-practices-for-devops/)(TechRepublic)
- [如何判断您的 DevOps 是否提供业务价值](http://www.zdnet.com/article/how-you-can-tell-if-your-devops-is-successful/)(ZDNet)
- [视频：如何成为 DevOps 工程师的技巧](https://www.techrepublic.com/videos/tips-for-how-to-become-a-devops-engineer/)(TechRepublic)
- [快速词汇表：DevOps](http://www.techproresearch.com/downloads/quick-glossary-devops/) (Tech Pro Research)

## Why does DevOps matter?

## 为什么 DevOps 很重要？

To put it simply, DevOps makes the entire software lifecycle faster, from code commit to production deployment.

简而言之，DevOps 使整个软件生命周期更快，从代码提交到生产部署。

A [survey](https://puppet.com/resources/white-paper/2016-state-devops-report/) of 4,600 IT professionals by [Puppet](http://www.puppet.com/) in June 2016 found that IT departments with a robust DevOps workflow deploy software 200 times more frequently than low-performing IT departments. In addition, they have 24 times faster recovery times, and three times lower rates of change failure, while spending 50% less time overall addressing security issues, and 22% less time on unplanned work.

[调查](https://puppet.com/resources/white-paper/2016-state-devops-report/)[Puppet](http://www.puppet.com/) 6 月份对 4,600 名 IT 专业人员进行的调查2016 年发现，拥有强大 DevOps 工作流程的 IT 部门部署软件的频率是低绩效 IT 部门的 200 倍。此外，他们的恢复时间缩短了 24 倍，更改失败率降低了三倍，同时总体解决安全问题的时间减少了 50%，计划外工作的时间减少了 22%。

While the concept of continuous delivery--and by extension, DevOps--may be counterintuitive to some, the end goal of frequent software deployments is to make the process so routine as to be a non-event, as opposed to a disruptive major rollout .

虽然持续交付的概念——进而扩展，DevOps——对某些人来说可能有悖常理，但频繁软件部署的最终目标是使流程如此常规，使其成为非事件，而不是破坏性的重大部署.

**Additional resources:**

**其他资源：**

- [Gap between DevOps-savvy and non-savvy companies is huge, survey finds](http://www.zdnet.com/article/gap-between-devops-savvy-and-non-savvy-companies-is-huge-survey-finds/) (ZDNet)
- [The road to digital bliss is paved with service thinking and DevOps](http://www.zdnet.com/article/the-road-to-digital-bliss-is-paved-with-service-thinking-and-devops/) (ZDNet)
- [Enterprise cloud performs best with DevOps, software-defined networks](http://www.zdnet.com/article/enterprise-cloud-performs-best-with-devops-software-defined-networks/)(ZDNet)
- [7 critical lessons businesses learn when implementing DevOps](https://www.techrepublic.com/article/7-critical-lessons-businesses-learn-when-implementing-devops/)(TechRepublic)
- [DevOps isn't a matter of speed, it's all about software quality](https://www.techrepublic.com/article/devops-isnt-a-matter-of-speed-its-all-about-software-quality/) (TechRepublic)
- [Where does DevOps start? Not where you think](https://www.zdnet.com/article/where-does-devops-start-not-where-you-think/)(ZDNet)
- [50+% of DevOps pros handle security, but they lack proper knowledge and skills](https://www.techrepublic.com/article/50-of-devops-pros-handle-security-but-they-lack-proper-knowledge-and-skills/) (TechRepublic)
- [DevOps-savvy 和 non-savvy 公司之间的差距很大，调查发现](http://www.zdnet.com/article/gap-between-devops-savvy-and-non-savvy-companies-is-huge-调查发现/）（ZDNet)
- [数字幸福之路是用服务思维和DevOps铺成的](http://www.zdnet.com/article/the-road-to-digital-bliss-is-paved-with-service-thinking-and-devops/) (ZDNet)
- [企业云在 DevOps、软件定义网络中表现最佳](http://www.zdnet.com/article/enterprise-cloud-performs-best-with-devops-software-defined-networks/)(ZDNet)
- [企业在实施 DevOps 时学到的 7 个关键教训](https://www.techrepublic.com/article/7-critical-lessons-businesses-learn-when-implementing-devops/)(TechRepublic)
- [DevOps 不在于速度，而在于软件质量](https://www.techrepublic.com/article/devops-isnt-a-matter-of-speed-its-all-about-software-质量/)(TechRepublic)
- [DevOps 从哪里开始？不是你想的地方](https://www.zdnet.com/article/where-does-devops-start-not-where-you-think/)(ZDNet)
- [50+% 的 DevOps 专业人员处理安全问题，但他们缺乏适当的知识和技能](https://www.techrepublic.com/article/50-of-devops-pros-handle-security-but-they-lack-适当的知识和技能/）（TechRepublic)

## Who does DevOps affect?

## DevOps 会影响谁？

Any organization with an in-house IT department can benefit from adopting a DevOps communication culture. For the consumer side of the equation, DevOps allows for a significantly reduced time to market, allowing organizations to deliver new features and security patches to customers more quickly and efficiently.

任何拥有内部 IT 部门的组织都可以从采用 DevOps 沟通文化中受益。对于等式的消费者方面，DevOps 可以显着缩短上市时间，使组织能够更快、更高效地向客户提供新功能和安全补丁。

**SEE: [How to build a successful career as a DevOps engineer (free PDF)](https://www.techrepublic.com/resource-library/whitepapers/how-to-build-a-successful-career-as-a-devops-engineer/) (TechRepublic)**

**参见：[如何建立成功的 DevOps 工程师职业生涯（免费 PDF）](https://www.techrepublic.com/resource-library/whitepapers/how-to-build-a-successful-career-as-a-devops-engineer/) (TechRepublic)**

DevOps can also increase job satisfaction among IT professionals. The initiative allows for a much-needed conversation change from "How can we reduce cost?" to "How can we increase speed?" which is more likely to be a successful long-term strategy.

DevOps 还可以提高 IT 专业人员的工作满意度。该计划允许从“我们如何降低成本？”进行急需的对话改变。到“我们怎样才能提高速度？”这更有可能是一个成功的长期战略。

The aforementioned Puppet survey also notes that employees in high-performing DevOps teams were 2.2 times more likely to "recommend their organization as a great place to work."

上述 Puppet 调查还指出，高绩效 DevOps 团队中的员工“推荐他们的组织作为工作的好地方”的可能性是其他人的 2.2 倍。

**Additional resources:**

**其他资源：**

- [How to become a DevOps engineer: A cheat sheet](https://www.techrepublic.com/article/how-to-become-a-devops-engineer-a-cheat-sheet/)(TechRepublic)
- [How to become a software engineer: A cheat sheet](https://www.techrepublic.com/article/how-to-become-a-software-engineer-a-cheat-sheet/)(TechRepublic)
- [How to become a DevOps manager: 5 tips](https://www.techrepublic.com/article/how-to-become-a-devops-manager-5-tips/)(TechRepublic)
- [10 DevOps experts to follow on Twitter](https://www.techrepublic.com/article/10-devops-experts-to-follow-on-twitter/)(TechRepublic)
- [10 questions DevOps engineers can expect to be asked in a job interview](https://www.techrepublic.com/article/10-questions-devops-engineers-can-expect-to-be-asked-in-a-job-interview/) (TechRepublic)
- [DevSecOps teams securing cloud-based assets: Why collaboration is key](https://www.techrepublic.com/article/devsecops-teams-securing-cloud-based-assets-why-collaboration-is-key/) ( TechRepublic)
- [Video: What's in store for the next generation of software development](https://www.techrepublic.com/videos/video-whats-in-store-for-the-next-generation-of-software-development/) (TechRepublic)
- [DevOps double case study: CloudBees and Perfecto Mobile](http://www.techproresearch.com/article/devops-double-case-study-cloudbees-and-perfecto-mobile/) (Tech Pro Research)
- [DevOps: Where it's going and how to make the most of it](http://www.zdnet.com/article/devops-where-its-going-and-how-to-make-the-most-of-it/) (ZDNet)
- [Why leading DevOps may get you a promotion](http://www.zdnet.com/article/will-leading-devops-get-you-promoted/)(ZDNet)
- [Bring on the DevOps, say IT support managers](http://www.zdnet.com/article/devops-to-bring-it-support-teams-into-software-development-loop/)(ZDNet)
- [Why every DevOps team needs a contrarian thinker](http://www.zdnet.com/article/how-to-put-together-a-devops-team/)(ZDNet)
- [Executives love continuous integration more than developers, for now](https://www.zdnet.com/article/world-embraces-cloud-computing-containers-cross-the-chasm/)(ZDNet)
- [如何成为 DevOps 工程师：备忘单](https://www.techrepublic.com/article/how-to-become-a-devops-engineer-a-cheat-sheet/)(TechRepublic)
- [如何成为软件工程师：备忘单](https://www.techrepublic.com/article/how-to-become-a-software-engineer-a-cheat-sheet/)(TechRepublic)
- [如何成为 DevOps 经理：5 个技巧](https://www.techrepublic.com/article/how-to-become-a-devops-manager-5-tips/)(TechRepublic)
- [在 Twitter 上关注 10 位 DevOps 专家](https://www.techrepublic.com/article/10-devops-experts-to-follow-on-twitter/)(TechRepublic)
- [DevOps 工程师在求职面试中可能会被问到的 10 个问题](https://www.techrepublic.com/article/10-questions-devops-engineers-can-expect-to-be-asked-in-a-job-interview/) (TechRepublic)
- [DevSecOps 团队保护基于云的资产：为什么协作是关键](https://www.techrepublic.com/article/devsecops-teams-securing-cloud-based-assets-why-collaboration-is-key/)(科技共和国)
- [视频：为下一代软件开发准备了什么](https://www.techrepublic.com/videos/video-whats-in-store-for-the-next-generation-of-software-development/）（科技共和国)
- [DevOps 双案例研究：CloudBees 和 Perfecto Mobile](http://www.techproresearch.com/article/devops-double-case-study-cloudbees-and-perfecto-mobile/) (Tech Pro Research)
- [DevOps：它的发展方向以及如何充分利用它](http://www.zdnet.com/article/devops-where-its-going-and-how-to-make-the-most-of-it/) (ZDNet)
- [为什么领先的 DevOps 可能会让你升职](http://www.zdnet.com/article/will-leading-devops-get-you-promoted/)(ZDNet)
- [带来 DevOps，说 IT 支持经理](http://www.zdnet.com/article/devops-to-bring-it-support-teams-into-software-development-loop/)(ZDNet)
- [为什么每个 DevOps 团队都需要一个逆向思维者](http://www.zdnet.com/article/how-to-put-together-a-devops-team/)(ZDNet)
- [目前，高管们比开发人员更喜欢持续集成](https://www.zdnet.com/article/world-embraces-cloud-computing-containers-cross-the-chasm/)(ZDNet)

## When did companies start using DevOps?

## 公司什么时候开始使用 DevOps？

DevOps as an idea is an outgrowth of Agile Infrastructure, effectively extrapolated to the entire enterprise rather than just the IT department. The concept gained a lot of traction with the [first devopsdays conference in Belgium in 2009](https://legacy.devopsdays.org/events/2009-ghent/).

DevOps 作为一个想法是敏捷基础设施的产物，有效地外推到整个企业，而不仅仅是 IT 部门。 [2009 年在比利时举行的第一次 devopsdays 会议](https://legacy.devopsdays.org/events/2009-ghent/)，这个概念获得了很大的关注。

While traditional technology companies like Amazon, Adobe, and Netflix have been early adopters of the strategy, DevOps also enjoys popularity in the retail space with Target, Walmart, and Nordstrom utilizing the communication model.

虽然亚马逊、Adobe 和 Netflix 等传统科技公司是该战略的早期采用者，但 DevOps 在零售领域也很受欢迎，Target、沃尔玛和 Nordstrom 利用通信模型。

**Additional resources:**

**其他资源：**

- [Research: DevOps adoption rates, associated hiring and retraining, and outcomes after implementation](http://www.techproresearch.com/downloads/research-devops-adoption-rates-associated-hiring-and-retraining-and-outcomes-after-implementation/) (Tech Pro Research)
- [CIO Jury: 50% of tech leaders are implementing DevOps](https://www.techrepublic.com/article/cio-jury-50-of-tech-leaders-are-implementing-devops/)(TechRepublic)
- [How one company's DevOps success got them the green light to hire 1000 developers](https://www.techrepublic.com/article/how-one-companys-devops-success-got-them-the-greenlight-to-hire-1000-developers/) (TechRepublic)
- [研究：DevOps 采用率、相关招聘和再培训以及实施后的结果](http://www.techproresearch.com/downloads/research-devops-adoption-rates-related-hiring-and-retraining-and-outcomes-实施后/）（技术专业研究)
- [CIO 评审团：50% 的技术领导者正在实施 DevOps](https://www.techrepublic.com/article/cio-jury-50-of-tech-leaders-are-implementing-devops/)(TechRepublic)
- [一家公司的 DevOps 成功如何让他们获得雇用 1000 名开发人员的绿灯](https://www.techrepublic.com/article/how-one-companys-devops-success-got-them-the-greenlight-to-雇用 1000 名开发人员/) (TechRepublic)
- [DevOps: Chef offers enterprise-wide analytics with Automate tool](http://www.zdnet.com/article/devops-chef-offers-enterprise-wide-analytics-with-automate-tool/)(ZDNet)
- [Why the DevOps faithful keep pulling away from their competitors](https://www.techrepublic.com/article/why-the-devops-faithful-keep-pulling-away-from-their-competitors/)(TechRepublic)
- [Kubernetes 1.4: One DevOps tool to rule all the containers](http://www.zdnet.com/article/kubernetes-1-4-one-devops-tool-to-rule-all-the-containers/)(ZDNet)
- [Video: How AI chatbots are shaping DevOps workflows](https://www.techrepublic.com/videos/video-how-ai-chatbots-are-shaping-devops-workflows-interview/)(TechRepublic)
- [Video: How ChatOps enables the next wave of DevOps in the enterprise](https://www.techrepublic.com/videos/video-how-chatops-enables-the-next-wave-of-devops-in-the-enterprise/) (TechRepublic)
- [DevOps, machine learning dominate technology opportunities this year](https://www.zdnet.com/article/devops-machine-learning/)(ZDNet)
- [A proactive flavor of DevOps grows at Google](https://www.zdnet.com/article/googles-proactive-approach-to-devops/)(ZDNet)
- [AI-powered DevOps is how CA wants to reinvent software development and itself](https://www.zdnet.com/article/ai-powered-devops-is-how-ca-wants-to-reinvent-software-development-and-itself/) (ZDNet)
- [DevOps：Chef 通过自动化工具提供企业范围的分析](http://www.zdnet.com/article/devops-chef-offers-enterprise-wide-analytics-with-automate-tool/)(ZDNet)
- [为什么 DevOps 忠实者不断远离他们的竞争对手](https://www.techrepublic.com/article/why-the-devops-faithful-keep-pulling-away-from-their-competitors/)(TechRepublic)
- [Kubernetes 1.4: 一个 DevOps 工具来统治所有容器](http://www.zdnet.com/article/kubernetes-1-4-one-devops-tool-to-rule-all-the-containers/)(ZDNet)
- [视频：AI 聊天机器人如何塑造 DevOps 工作流程](https://www.techrepublic.com/videos/video-how-ai-chatbots-are-shaping-devops-workflows-interview/)(TechRepublic)
- [视频：ChatOps 如何在企业中实现下一波 DevOps](https://www.techrepublic.com/videos/video-how-chatops-enables-the-next-wave-of-devops-in-the-企业/) (TechRepublic)
- [DevOps，机器学习主导今年的技术机会](https://www.zdnet.com/article/devops-machine-learning/)(ZDNet)
- [Google 发展出一种积极主动的 DevOps](https://www.zdnet.com/article/googles-proactive-approach-to-devops/)(ZDNet)
- [AI 驱动的 DevOps 是 CA 希望重塑软件开发及其自身的方式](https://www.zdnet.com/article/ai-powered-devops-is-how-ca-wants-to-reinvent-software-开发本身/)(ZDNet)

## How do I implement DevOps?

## 我如何实施 DevOps？

Because DevOps is, at its core, a cultural and procedural adjustment from the way things have been done in the past, it is not possible to implement DevOps overnight. The steps needed to implement DevOps is really dependent on the existing IT infrastructure and corporate structure of a given organization--groups already using cloud infrastructure and agile development practices are several steps ahead of groups not using those systems.

由于 DevOps 的核心是对过去工作方式的文化和程序调整，因此不可能在一夜之间实施 DevOps。实施 DevOps 所需的步骤实际上取决于给定组织的现有 IT 基础设施和企业结构——已经使用云基础设施和敏捷开发实践的团队比不使用这些系统的团队提前了几步。

[Gene Kim](http://itrevolution.com/authors/gene-kim/), the founder of TripWire, and author of various books on development, [advocates for an incremental approach for adopting DevOps](http://cdn.inedo.com/documents/Inedo-Incremental-DevOps.pdf) (PDF link). While this is a long-term change that will likely take 1-2 years for established organizations, results from the pilot candidate can be seen in a matter of weeks.

[Gene Kim](http://itrevolution.com/authors/gene-kim/)，TripWire 的创始人，以及各种开发书籍的作者，[提倡采用增量方法来采用 DevOps](http://cdn.inedo.com/documents/Inedo-Incremental-DevOps.pdf)（PDF 链接)。虽然这是一个长期的变化，对于成熟的组织来说可能需要 1-2 年的时间，但试点候选人的结果可以在几周内看到。

**Additional resources:**

**其他资源：**

- [10 steps to DevOps success in the enterprise](https://www.techrepublic.com/article/10-steps-to-devops-success-in-the-enterprise/)(TechRepublic)
- [Top 10 challenges to DevOps implementation](https://www.techrepublic.com/article/top-10-challenges-to-devops-implementation/)(TechRepublic)
- [10 books to add to your DevOps reading list](https://www.techrepublic.com/article/10-books-to-add-to-your-devops-reading-list/)(TechRepublic)
- [Ebook: IT leader's guide to making DevOps work](http://www.techproresearch.com/downloads/it-leader-s-guide-to-making-devops-work/) (Tech Pro Research)
- [DevOps enters the mainstream: Here's how to make it work for you](http://www.zdnet.com/article/devops-enters-the-mainstream-heres-how-to-make-it-work-for-you/) (ZDNet)
- [It takes more than one champion: Getting DevOps working for you](http://www.zdnet.com/article/it-takes-more-than-one-champion-getting-devops-working-for-you/) (ZDNet)
- [Agile and DevOps: How to rethink your software development culture](https://www.zdnet.com/article/agile-and-devops-how-to-rethink-your-software-development-culture/) (ZDNet )
- [10 ways to improve time-to-market for your applications](https://www.techrepublic.com/blog/10-things/10-ways-to-improve-time-to-market-for-your-applications/) (TechRepublic)
- [Shifting to DevOps? Put your ducks in a row first](https://www.techrepublic.com/article/shifting-to-devops-put-your-ducks-in-a-row-first/)(TechRepublic)
- [Standardizing DevOps tools requires culture change and careful evaluation](http://www.techproresearch.com/article/standardizing-devops-tools-requires-culture-change-and-careful-evaluation/) (Tech Pro Research)
- [企业 DevOps 成功的 10 个步骤](https://www.techrepublic.com/article/10-steps-to-devops-success-in-the-enterprise/)(TechRepublic)
- [DevOps 实施的 10 大挑战](https://www.techrepublic.com/article/top-10-challenges-to-devops-implementation/)(TechRepublic)
- [添加到您的 DevOps 阅读列表中的 10 本书](https://www.techrepublic.com/article/10-books-to-add-to-your-devops-reading-list/)(TechRepublic)
- [电子书：让 DevOps 发挥作用的 IT 领导者指南](http://www.techproresearch.com/downloads/it-leader-s-guide-to-making-devops-work/) (Tech Pro Research)
- [DevOps 进入主流：这里是如何让它为你工作](http://www.zdnet.com/article/devops-enters-the-mainstream-heres-how-to-make-it-work-for-你/）（ZDNet)
- [需要不止一个冠军：让 DevOps 为你工作](http://www.zdnet.com/article/it-takes-more-than-one-champion-getting-devops-working-for-you/) (ZDNet)
- [敏捷和 DevOps：如何重新思考你的软件开发文化](https://www.zdnet.com/article/agile-and-devops-how-to-rethink-your-software-development-culture/) (ZDNet )
- [缩短应用上市时间的 10 种方法](https://www.techrepublic.com/blog/10-things/10-ways-to-improve-time-to-market-for-your-应用程序/)(TechRepublic)
- [转向 DevOps？先把你的鸭子排成一排](https://www.techrepublic.com/article/shifting-to-devops-put-your-ducks-in-a-row-first/)(TechRepublic)
- [标准化 DevOps 工具需要文化变革和仔细评估](http://www.techproresearch.com/article/standardizing-devops-tools-requires-culture-change-and-careful-evaluation/) (Tech Pro Research)

# SaaS: Single Tenant vs Multi-Tenant - What's the Difference?

# SaaS：单租户与多租户 - 有什么区别？

by [Chris Brook](https://digitalguardian.com/author/chris-brook) on Tuesday December 1, 2020

作者：[Chris Brook](https://digitalguardian.com/author/chris-brook)，2020 年 12 月 1 日，星期二

What are the advantages of a multi-tenancy SaaS architecture? How does it differ from single tenant instances? We break down the differences and highlight the benefits of implementing a SaaS solution for data protection in this blog.

多租户 SaaS 架构有哪些优势？它与单租户实例有何不同？我们在此博客中分解差异并强调实施 SaaS 解决方案以保护数据的好处。

### Multi-Tenancy – A Core Benefit of SaaS

### 多租户 – SaaS 的核心优势

In the early days of the cloud, organizations were reluctant to adopt cloud strategies. Few organizations considered applying policies, technologies, and controls to protect data across the cloud. The last several years, which has seen the proven effectiveness of cloud deployments in scalability, cost, and security, has changed that however. Now we’re seeing the rapid adoption of cloud platforms by organizations of all shapes and sizes.

在云计算的早期，组织不愿意采用云计算策略。很少有组织考虑应用策略、技术和控制来保护云中的数据。过去几年，云部署在可扩展性、成本和安全性方面的有效性已经得到证实，但这种情况发生了变化。现在，我们看到各种形式和规模的组织都在迅速采用云平台。

[Digital Guardian's Data Protection Platform](https://digitalguardian.com/products/threat-aware-data-protection-platform) leverages software as a service, or SaaS, to provide data protection in a package that results in superior security, better economics, and reduced overhead. One of the ways we do this is through multi-tenant architecture.

[Digital Guardian 的数据保护平台](https://digitalguardian.com/products/threat-aware-data-protection-platform) 利用软件即服务或 SaaS 以包的形式提供数据保护，从而实现卓越的安全性，更好的经济性，并减少开销。我们这样做的方法之一是通过多租户架构。

### Single Tenant vs Multi-Tenant – Learn the Difference

### 单租户与多租户 - 了解差异

Single Tenant – A single instance of the software and supporting infrastructure serve a single customer. With single tenancy, each customer has his or her own independent database and instance of the software. Essentially, there is no sharing happening with this option.

单一租户——软件和支持基础设施的单一实例服务于单一客户。对于单一租户，每个客户都有自己独立的数据库和软件实例。本质上，此选项不会发生共享。

Potential benefits of single-tenant include:

单租户的潜在好处包括：

- Security: A single customer and a single server is often contained on secure hardware being used by a limited number of people.
- Dependability: With an entire environment dedicated to one client, resources are abundant and available anytime.
- Customization: Control over the entire environment allows for customization and added functionality, if desired.

- 安全性：单个客户和单个服务器通常包含在由少数人使用的安全硬件上。
- 可靠性：整个环境专用于一个客户，资源丰富且随时可用。
- 定制：如果需要，对整个环境的控制允许定制和添加功能。

Potential drawbacks of single-tenant:

单租户的潜在缺点：

- Maintenance: Single-tenant typically means more tasks and regular maintenance to keep things running smoothly and efficiently.
- Setup/Management: By comparison, SaaS multi-tenant environments are quick to setup and manage.
- Cost: Single-tenant typically allows for more resources, but at a premium price given that there is only one customer for the entire environment.

- 维护：单租户通常意味着更多的任务和定期维护，以保持事情顺利高效地运行。
- 设置/管理：相比之下，SaaS 多租户环境可以快速设置和管理。
- 成本：单租户通常允许更多资源，但鉴于整个环境只有一个客户，因此价格较高。

Multi-Tenant – Multi-tenancy means that a single instance of the software and its supporting infrastructure serves multiple customers. Each customer shares the software application and also shares a single database. Each tenant’s data is isolated and remains invisible to other tenants.

多租户 - 多租户意味着软件的单个实例及其支持基础设施服务于多个客户。每个客户共享软件应用程序并共享一个数据库。每个租户的数据都是隔离的，对其他租户不可见。

Potential benefits of multi-tenant:

多租户的潜在好处：

- Affordable Cost: Multiple customers means that the cost for the environment is shared, and those savings (from the SaaS vendor) are typically transferred to the cost of the software.
- Integrations: Cloud environments allow for easier integration with other applications through the use of APIs.
- “Hands-free” Maintenance: The server technically belongs to the SaaS vendor, meaning that a certain level of database maintenance is handled by the vendor, instead of you maintaining the environment yourself.

- 负担得起的成本：多个客户意味着环境成本是共享的，而这些节省（来自 SaaS 供应商）通常会转移到软件成本中。
- 集成：云环境允许通过使用 API 更轻松地与其他应用程序集成。
- “免提”维护：服务器在技术上属于 SaaS 供应商，这意味着一定程度的数据库维护由供应商处理，而不是您自己维护环境。

Potential drawbacks of multi-tenant:

多租户的潜在缺点：

- Limited Management/Customization: While you do have added integration benefits, custom changes to the database aren’t typically an option.
- Security: Other tenants won’t see your data. However, multiple users (not associated with your organization) are allowed on the same database. This broader access reduces control of security.
- Updates/Changes: If you’re reliant on integrations with other SaaS products and one updates their system, it may cause issues with those connecting apps.

- 有限的管理/自定义：虽然您确实增加了集成优势，但对数据库的自定义更改通常不是一种选择。
- 安全性：其他租户不会看到您的数据。但是，同一数据库允许多个用户（与您的组织无关）。这种更广泛的访问减少了对安全性的控制。
- 更新/更改：如果您依赖与其他 SaaS 产品的集成并且其中一个更新了他们的系统，则可能会导致连接应用程序出现问题。

datasheets

数据表

Digital Guardian SaaS Solution Sheet

Digital Guardian SaaS 解决方案表

[VIEW RESOURCE](https://info.digitalguardian.com/rs/768-OQW-145/images/DG-saas-solution-sheet.pdf)

[查看资源](https://info.digitalguardian.com/rs/768-OQW-145/images/DG-saas-solution-sheet.pdf)

### Benefits of SaaS Multi-Tenant Architecture

### SaaS 多租户架构的好处

- Lower costs through economies of scale: With multi-tenancy, scaling has far fewer infrastructure implications than with a single-tenancy-hosted solution because new users get access to the same basic software. 

- 通过规模经济降低成本：多租户与单租户托管解决方案相比，扩展对基础架构的影响要小得多，因为新用户可以访问相同的基本软件。

- Shared infrastructure leads to lower costs: SaaS allows companies of all sizes to share infrastructure and data center operational costs. There is no need to add applications and more hardware to their environment. Not having to provision or manage any infrastructure or software above and beyond internal resources enables businesses to focus on everyday tasks.
- Ongoing maintenance and updates: Customers don’t need to pay costly maintenance fees to keep their software up to date. Vendors roll out new features and updates. These are often included with a SaaS subscription.
- Configuration can be done while leaving the underlying codebase unchanged: Single-tenant-hosted solutions are often customized, requiring changes to an application’s code. This customization can be costly and can make upgrades time-consuming because the upgrade might not be compatible with your environment.

- 共享基础设施可降低成本：SaaS 允许各种规模的公司共享基础设施和数据中心运营成本。无需向其环境添加应用程序和更多硬件。无需在内部资源之外配置或管理任何基础设施或软件，企业就可以专注于日常任务。
- 持续的维护和更新：客户无需支付昂贵的维护费用来保持他们的软件是最新的。供应商推出新功能和更新。这些通常包含在 SaaS 订阅中。
- 可以在保持底层代码库不变的情况下进行配置：单租户托管解决方案通常是定制的，需要更改应用程序的代码。由于升级可能与您的环境不兼容，因此这种定制可能成本高昂并且会使升级变得耗时。

Multi-tenant solutions are designed to be highly configurable so that businesses can make the application perform the way they want. There is no changing the code or data structure, making the upgrade process easy.

多租户解决方案被设计成高度可配置的，因此企业可以让应用程序按照他们想要的方式运行。无需更改代码或数据结构，使升级过程变得容易。

Multi-tenancy architecture also allows Digital Guardian to efficiently service everyone from small customers, whose scale may not warrant dedicated infrastructure, to large enterprises that need access to the cloud’s virtually unlimited compute resources. Software development and maintenance costs are shared, driving down expenditures, resulting in savings that are passed onto you, the customers.

多租户架构还允许 Digital Guardian 高效地为每个人提供服务，从规模可能无法保证专用基础设施的小客户到需要访问云计算资源几乎无限的大型企业。软件开发和维护成本是分摊的，从而降低了支出，从而将节省的费用转嫁给了您，即客户。

### Additional Benefits of SaaS

### SaaS 的其他好处

Multi-tenancy is just one of multiple benefits of SaaS. Download this white paper – [7 Reasons to Move to SaaS Data Protection](https://info.digitalguardian.com/whitepaper-seven-reasons-to-move-to-saas-data-protection.html) to learn:

多租户只是 SaaS 的众多优势之一。下载此白皮书 - [转向 SaaS 数据保护的 7 个理由](https://info.digitalguardian.com/whitepaper-seven-reasons-to-move-to-saas-data-protection.html) 以了解：

- The 7 reasons why moving to SaaS data protection enables you to manage risk more effectively
- How Digital Guardian’s cloud architecture is built with the latest tools and methodologies
- How we can help offset resource constraints with our Managed Security Program

- 迁移到 SaaS 数据保护的 7 个原因使您能够更有效地管理风险
- 如何使用最新的工具和方法构建 Digital Guardian 的云架构
- 我们如何通过我们的托管安全计划帮助抵消资源限制

To learn more about the benefits of SaaS, watch the clip below from our webinar, [Benefits of Implementing a SaaS Cybersecurity Solution](https://info.digitalguardian.com/on-demand-webinar-implementing-saas-cybersecurity-solution.html), which is presented by Andras Cser, VP Principal Analyst at Forrester. You can watch the full webinar [here](https://info.digitalguardian.com/on-demand-webinar-implementing-saas-cybersecurity-solution.html).

要了解有关 SaaS 优势的更多信息，请观看我们网络研讨会的以下剪辑，[实施 SaaS 网络安全解决方案的优势](https://info.digitalguardian.com/on-demand-webinar-implementing-saas-cybersecurity-solution.html)，由 Forrester 副首席分析师 Andras Cser 提出。您可以观看完整的网络研讨会 [此处](https://info.digitalguardian.com/on-demand-webinar-implementing-saas-cybersecurity-solution.html)。

Tags: [Data Protection](http://digitalguardian.com/blog/search/data-protection), [Company Information](http://digitalguardian.com/blog/search/company-information)

标签：[数据保护](http://digitalguardian.com/blog/search/data-protection)、[公司信息](http://digitalguardian.com/blog/search/company-information)

[Whitepaper **Considering SaaS Data Protection?** Learn why you should move to SaaS for superior security and reduced overhead. READ NOW](https://info.digitalguardian.com/whitepaper-seven-reasons-to-move-to-saas-data-protection.html?utm_source=blog) 

[白皮书**考虑 SaaS 数据保护？** 了解为什么应该迁移到 SaaS 以获得卓越的安全性并减少开销。立即阅读](https://info.digitalguardian.com/whitepaper-seven-reasons-to-move-to-saas-data-protection.html?utm_source=blog)


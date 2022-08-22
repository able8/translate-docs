# Single-Tenant Vs. Multi-Tenant Cloud: Which Should You Use?

# 单租户VS。多租户云：您应该使用哪个？

Understand the differences between single-tenant versus multi-tenant cloud architecture so you can build cost-effective applications.

了解单租户与多租户云架构之间的区别，以便您可以构建具有成本效益的应用程序。

June 11, 2021 From: https://www.cloudzero.com/blog/single-tenant-vs-multi-tenant

[Cloud Cost Optimization](https://www.cloudzero.com/blog/tag/cloud-cost-optimization)

[云成本优化](https://www.cloudzero.com/blog/tag/cloud-cost-optimization)

Is your current cloud cost tool giving you the cost intelligence you need? Most tools are manual, clunky, and inexact. Discover how CloudZero takes a new  approach to organizing your cloud spend. Click here to learn more.

您当前的云成本工具是否为您提供所需的成本情报？大多数工具都是手动的、笨重的和不精确的。了解 CloudZero 如何采用新方法来组织您的云支出。点击这里了解更多。

When operating in the cloud, one of the key decisions to make is about which type of architecture to adopt for your business and customer data. This is because choosing cost-effective architecture is key to building profitable SaaS software.

在云中运行时，要做出的关键决策之一是为您的业务和客户数据采用哪种类型的架构。这是因为选择具有成本效益的架构是构建可盈利的 SaaS 软件的关键。

Single-tenant and multi-tenant cloud environments are the options to consider. Both types of architecture have security and privacy implications. There’s also the issue of cost, which differs significantly depending on the architecting model you adopt.

单租户和多租户云环境是需要考虑的选项。两种类型的架构都具有安全和隐私含义。还有成本问题，这取决于您采用的架构模型。

In this article, we’ll compare multi-tenant vs. single-tenant cloud architecture and how to build more cost-effective applications irrespective of the model you operate.

在本文中，我们将比较多租户与单租户云架构，以及如何构建更具成本效益的应用程序，而不管您运营的模型如何。

[![Related Article: The 11 Best Cloud Cost Management Tools In 2022](https://no-cache.hubspot.com/cta/default/2983524/2ea8634d-2f73-45cc-b72b-32b354bc0790.png)](https://cta-redirect.hubspot.com/cta/redirect/2983524/2ea8634d-2f73-45cc-b72b-32b354bc0790)

://cta-redirect.hubspot.com/cta/redirect/2983524/2ea8634d-2f73-45cc-b72b-32b354bc0790)

## **Defining Multi-Tenant Vs. Single-Tenant Cloud Architecture**

## **定义多租户与。单租户云架构**

What is a multi-tenant application vs. single tenant?

什么是多租户应用程序与单租户？

**Single-tenant cloud architecture** is one where a single software instance and its supporting [infrastructure](https://www.cloudzero.com/blog/cloud-infrastructure)/database serve only one customer. In a single-tenant environment, all customer data and interactions are separate from every other customer. Customer data is not housed in the same database and there’s no sharing of data in any way.

**单租户云架构**是单个软件实例及其支持的 [基础架构](https://www.cloudzero.com/blog/cloud-infrastructure)/数据库仅服务于一个客户的架构。在单租户环境中，所有客户数据和交互都与其他所有客户分开。客户数据没有存放在同一个数据库中，也没有以任何方式共享数据。

**A multi-tenant architecture** is one where a single software instance and database serves multiple customers (i.e. tenants).

**多租户架构**是一个软件实例和数据库为多个客户（即租户）服务的架构。

The real estate analogy is often used to explain single-tenant vs. multi-tenant cloud architecture. In a single-tenant cloud, each customer lives alone in a single apartment building which has its own security system and facilities and is completely isolated from neighboring buildings.

房地产类比通常用于解释单租户与多租户云架构。在单租户云中，每个客户单独住在一个公寓楼中，该公寓楼拥有自己的安全系统和设施，并与邻近的建筑物完全隔离。

In multi-tenant cloud architecture, tenants live in different apartments inside a single apartment building. They share the same security system and communal facilities. But each tenant has a key to their respective apartments so their privacy is guaranteed within their apartment. However, the activities of fellow tenants are more likely to have an impact on their comfort in the building.

在多租户云架构中，租户住在单个公寓楼内的不同公寓中。他们共享相同的安全系统和公共设施。但是每个租户都有各自公寓的钥匙，因此他们的隐私在他们的公寓内得到保证。然而，其他租户的活动更有可能影响他们在大楼中的舒适度。

Most startups opt for a multi-tenant setup, which means having one big database that houses all the customer data. With the right security systems in place, customer data stays private. While customers cannot see each other’s data, they do live in the same database, and the same computer processes all of the data; so, it’s not segmented in any way.

大多数初创公司选择多租户设置，这意味着拥有一个包含所有客户数据的大型数据库。有了正确的安全系统，客户数据就会保持私密。虽然客户无法看到彼此的数据，但他们确实存在于同一个数据库中，并且同一台计算机处理所有数据；所以，它没有以任何方式分割。

## **Benefits And Drawbacks Of Single-tenant Cloud Architecture**

## **单租户云架构的优缺点**

The benefits associated with single-tenant cloud architecture include:

与单租户云架构相关的好处包括：

- **Data security** \- Each customer’s data is isolated, minimizing the risk of a data breach that affects multiple customers.

- **数据安全** \- 每个客户的数据都是隔离的，从而将影响多个客户的数据泄露风险降至最低。

- **Customization** \- More customization options are available in a single-tenant system compared to a multi-tenant system because each customer has dedicated software and hardware. 

- **定制** \- 与多租户系统相比，单租户系统提供了更多定制选项，因为每个客户都有专用的软件和硬件。

- **Portability** \- It can also be easier to migrate data from a single-tenant architecture. The most obvious reason is that there is simply less data to handle compared to a multi-tenant setup. Also, since the data store contains only data for a single customer, the team has less worry about complicated migration scripts and mixing customer data on migration. Lastly, most single-tenant architecture replicate the same infrastructure for each customer; therefore, migrating data is often similar to setting up a new customer..

- **可移植性** \- 从单租户架构迁移数据也更容易。最明显的原因是与多租户设置相比，需要处理的数据更少。此外，由于数据存储仅包含单个客户的数据，因此团队不必担心复杂的迁移脚本和迁移时混合客户数据。最后，大多数单租户架构为每个客户复制相同的基础架构；因此，迁移数据通常类似于设置新客户。

Some disadvantages of single-tenant cloud architecture are:

单租户云架构的一些缺点是：

- **Complex setup and management** \- You have to deploy an instance and set up a database for each customer. As the clientele grows, it becomes more difficult to manage multiple applications. Single-tenant architecture is particularly unsuitable for small startup teams.

- **复杂的设置和管理** \- 您必须为每个客户部署一个实例并设置一个数据库。随着客户的增长，管理多个应用程序变得更加困难。单租户架构特别不适合小型初创团队。

- **High costs** \- Running a single-tenant set-up is more expensive than using shared resources. If every customer had their own database and compute, you would need a lot of resources to manage it.

- **高成本** \- 运行单租户设置比使用共享资源更昂贵。如果每个客户都有自己的数据库和计算，您将需要大量资源来管理它。

- **Inefficient resource usage** \- Resources are also likely to be underutilized, leading to inefficiencies. Since customers do not have the same usage patterns, allocating the same amount of physical resources to each customer will result in some customers using all (or even needing more) of their allocated resources while others use little or none of their resources.

- **资源使用效率低** \- 资源也可能未被充分利用，导致效率低下。由于客户没有相同的使用模式，因此为每个客户分配相同数量的物理资源将导致一些客户使用所有（甚至需要更多）他们分配的资源，而另一些客户则很少或不使用他们的资源。

## **Benefits And Drawbacks Of Multi-tenant Cloud Architecture**

## **多租户云架构的优缺点**

Some of the benefits associated with multi-tenant cloud architecture are:

与多租户云架构相关的一些好处是：

- **Easy deployability**\- Consider a team of two people building a new startup. It would be too much work for a team of two to serve multiple clients using a single-tenant approach. While it is doable, it will require a lot of effort. If they adopt a multi-tenant system, however, they can reach for tools provided by the cloud services provider, such as AWS, and build their application faster. Multi-tenancy is at the heart of effective SaaS operations because it makes it easy to build and deploy applications faster and to scale those applications quickly.
- **Efficient resource usage**\- Unlike the single-tenant model where resources are likely to be underutilized, available resources in a multi-tenant environment are used maximally because they are shared by multiple users.

- **易于部署**\- 考虑一个由两人组成的团队建立一个新的创业公司。对于一个两个人的团队来说，使用单租户方法为多个客户提供服务将是太多的工作。虽然可行，但需要付出很多努力。但是，如果他们采用多租户系统，他们可以使用 AWS 等云服务提供商提供的工具，并更快地构建他们的应用程序。多租户是有效 SaaS 运营的核心，因为它可以更轻松地更快地构建和部署应用程序并快速扩展这些应用程序。
- **高效的资源使用**\- 与资源可能未被充分利用的单租户模型不同，多租户环境中的可用资源得到最大程度的利用，因为它们由多个用户共享。

- **Reduced costs** \- Multiple customers share the cost of the environment in multi-tenancy, so the application is cheaper to build and maintain. Maintenance and management costs are shared by all customers. For example, a single DynamoDB Table can easily hold all application data for millions of customers (as is the case with Amazon itself).

- **降低成本** \- 多个客户在多租户中分摊环境成本，因此应用程序的构建和维护成本更低。维护和管理成本由所有客户分摊。例如，单个 DynamoDB 表可以轻松保存数百万客户的所有应用程序数据（亚马逊本身就是这种情况）。

While multi-tenant cloud architecture is usually the best approach for most consumer-facing applications, it is not without its disadvantages, which include:

虽然多租户云架构通常是大多数面向消费者的应用程序的最佳方法，但它并非没有缺点，包括：

- **Greater security risk** \- In a multi-tenant system, the risks are higher because resources are shared by multiple customers. If one customer’s data is compromised, it is more likely that it will affect other customers, unlike in a single-tenant cloud where security incidents are isolated to a single client.
- **Lack of cost visibility**\- In comparison to a single-tenant system where each customer has their own database, it is much harder to separate costs in the multi-tenant system.

- **更大的安全风险** \- 在多租户系统中，风险更高，因为资源由多个客户共享。如果一个客户的数据被泄露，它更有可能影响其他客户，这与将安全事件隔离到单个客户的单租户云不同。
- **缺乏成本可见性**\- 与每个客户都有自己的数据库的单租户系统相比，在多租户系统中分离成本要困难得多。

[You Might Also Like: The Definitive DevOps Tools List: 55 Tools For 2022](https://cta-redirect.hubspot.com/cta/redirect/2983524/8d3f05b2-9b44-42f4-b4b5-ed6f9ce2d436)

[您可能还喜欢：权威 DevOps 工具列表：2022 年的 55 种工具](https://cta-redirect.hubspot.com/cta/redirect/2983524/8d3f05b2-9b44-42f4-b4b5-ed6f9ce2d436)

## **When To Use Single-tenant Vs. Multi-tenant Architecture**

## **何时使用单租户与。多租户架构**

A single tenant architecture may be suited for specific industries or sectors where there are strong privacy and security concerns around customer data. Good examples are the healthcare and finance industries.

单一租户架构可能适用于对客户数据存在强烈隐私和安全问题的特定行业或部门。很好的例子是医疗保健和金融行业。

In the healthcare industry, for example, applications must meet HIPAA requirements when dealing with patient information. So each hospital may need to have its own data center onsite to ensure compliance. The same applies to certain types of financial data. 

例如，在医疗保健行业，应用程序在处理患者信息时必须满足 HIPAA 要求。因此，每家医院可能都需要在现场拥有自己的数据中心以确保合规性。这同样适用于某些类型的财务数据。

Most consumer-facing applications are best built as multi-tenant applications. The cloud itself is multi-tenant. Cloud service providers such as AWS use the same hardware for various customers under the covers. While each customer has a different Amazon account, the same computers process customer information.

大多数面向消费者的应用程序最好构建为多租户应用程序。云本身是多租户的。 AWS 等云服务提供商在幕后为各种客户使用相同的硬件。虽然每个客户都有不同的亚马逊账户，但相同的计算机处理客户信息。

## **Understanding Cost Per Customer In A Multi-Tenant Architecture**

## **了解多租户架构中每位客户的成本**

One of the major trade-offs with multi-tenancy is cost visibility. Since all your customer data is housed in one database, it’s difficult to figure out the cost of servicing each customer individually.

多租户的主要权衡之一是成本可见性。由于您的所有客户数据都保存在一个数据库中，因此很难计算单独为每个客户提供服务的成本。

A lack of visibility into your costs per tenant (i.e., per customer) makes it difficult to make critical pricing decisions because you have no idea what your costs and margins are for each customer or customer segment.

对每个租户（即每个客户）的成本缺乏可见性，因此很难做出关键的定价决策，因为您不知道每个客户或客户群的成本和利润是多少。

So, how do you tease apart your multi-tenant architecture to understand your costs per tenant? This is where [CloudZero](https://www.cloudzero.com) comes in.

那么，您如何梳理多租户架构以了解每个租户的成本？这就是 [CloudZero](https://www.cloudzero.com) 的用武之地。

CloudZero’s cloud cost intelligence platform lets you have the best of both worlds when managing customer data in a multi-tenant environment. You do not have to manage all your customers separately but you can still decipher your cost per customer.

CloudZero 的云成本智能平台可让您在多租户环境中管理客户数据时两全其美。您不必单独管理所有客户，但您仍然可以破译每位客户的成本。

[![CloudZero aligns cloud costs to key business metrics, such as cost per  customer or product feature.Our Cost Per Customer report allows teams to see  how individual customers drive their cloud spend and how much specific  customers cost their business. With cloud cost intelligence, companies can make  informed engineering, business, and pricing that ensure profitability.Click  here to learn more.](https://no-cache.hubspot.com/cta/default/2983524/690c9dfe-5370-428d-84f3-46586422d4d0.png)](https://cta-redirect.hubspot.com/cta/redirect/2983524/690c9dfe-5370-428d-84f3-46586422d4d0)

我们的每客户成本报告允许团队查看单个客户如何推动他们的云支出以及特定客户的业务成本。借助云成本智能，公司可以制定明智的工程、业务和定价，以确保盈利。单击此处了解更多信息。](https://no-cache.hubspot.com/cta/default/2983524/690c9dfe-5370-428d-84f3-46586422d4d0.png)](https://cta-redirect.hubspot.com/cta/redirect/2983524/690c9dfe-5370-428d-84f3-46586422d4d0)

CloudZero automatically allocates cost per tenant and delivers granular metrics for your SaaS business. CloudZero pulls in your AWS events, normalizes the data, and allows you to correlate cost with other types of metrics.

CloudZero 自动为每个租户分配成本，并为您的 SaaS 业务提供精细的指标。 CloudZero 提取您的 AWS 事件，标准化数据，并允许您将成本与其他类型的指标相关联。

[Request a demo](https://www.cloudzero.com/demo) to see how CloudZero helps you break down costs in a multi-tenant architecture, giving you the insight you need to run your business better.

[请求演示](https://www.cloudzero.com/demo) 了解 CloudZero 如何帮助您在多租户架构中分解成本，为您提供更好地运营业务所需的洞察力。

[![](https://no-cache.hubspot.com/cta/default/2983524/ef571aa9-fdfe-408e-9d42-0aa3a43dcdef.png)](https://cta-redirect.hubspot.com/cta/redirect/2983524/ef571aa9-fdfe-408e-9d42-0aa3a43dcdef)

/redirect/2983524/ef571aa9-fdfe-408e-9d42-0aa3a43dcdef)

#### STAY IN THE LOOP

#### 留在循环中

### Join thousands of engineers who already receive the best AWS and cloud cost intelligence content. 

### 加入数千名已经收到最佳 AWS 和云成本情报内容的工程师的行列。


# An Introduction to the Cloud Native Landscape

#### 21 Jul 2020 9:52am,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini")

If you’ve researched cloud native applications and technologies, you’ve probably come across the [Cloud Native Computing Foundation (CNCF)](https://landscape.cncf.io/) cloud native landscape map. Unsurprisingly, the sheer scale of it can be overwhelming. So many categories and so many technologies. How do you make sense of it?

As with anything else, if you break it down and analyze it one piece at the time, you’ll find it’s not that complex and makes a lot of sense. In fact, the map is neatly organized by functionality, and once you understand what each category represents, navigating it becomes a lot easier.

In this article, the first in a series, we’ll break this mammoth landscape down and provide a high-level overview of the entire landscape, its layers, columns and categories. In follow up articles, we’ll zoom into each layer and column and provide more details on what each category is, what problem it solves, and how.

![CNCF landscape](https://cdn.thenewstack.io/media/2020/07/6b2f1a36-screen-shot-2020-07-17-at-8.17.23-am.png)

## The Four Layers of the Cloud Native Landscape

First, let’s strip all individual technologies from the landscape and look at the categories. There are different “rows” reflecting architectural layers each with its own set of subcategories. In the first layer, you have tools to provision infrastructure, that’s your foundation. Then you start adding tooling needed to run and manage apps such as the runtime and orchestration layer. At the very top you have tools to define and develop your application, such as databases, image building, and CI/CD tools (we’ll discuss each of these below).

![CNCF landscape categories](https://cdn.thenewstack.io/media/2020/07/794834eb-screen-shot-2020-07-17-at-2.39.02-pm.png)

For now, what you should remember is that the landscape starts with the infrastructure and, with each layer, moves closer to the actual app. That’s what these layers represent (we’ll address the two “columns” running across those layers later). Let’s explore each layer at a time, starting with the bottom.

### 1\. The Provisioning Layer

Provisioning refers to the tools involved in creating and hardening the foundation on which cloud native applications are built. It covers everything from automating the creation, management, and configuration of infrastructure to scanning, signing and storing container images. Provisioning even extends into the security space by providing tools that allow you to set and enforce policies, build authentication and authorization into your apps and platforms, and handle secrets distribution.

In the provisioning layer, you’ll find:

- **Automation and configuration tooling** to help engineers build computing environments without human intervention.
- **Container registries** store executable files of the apps.
- **Security and** **compliance** frameworks address different security areas.
- **Key management** solutions help with encryption to ensure only authorized users have access to the application.

These tools allow engineers to codify all infrastructure specifics, so that the system can spin new environments up and down as needed, ensuring they are consistent and secure.

### 2\. The Runtime Layer

Next, is the runtime layer. Runtime is one of those terms that can be confusing. Like many terms in IT, there is no strict definition and it can be used differently, depending on the context. In a narrow sense, runtime is a sandbox on a specific machine prepared to run an app — the bare minimum an app needs. In the widest of senses, runtime is any tool the app needs to run.

In the CNCF cloud native landscape, runtime is defined somewhere in between focusing on the components that matter for the containerized apps in particular: what they need to run, remember, and communicate. They include:

- **Cloud native storage** provides virtualized disks or persistence for containerized apps.
- **Container runtime** delivers the constraint, resource, and security considerations for containers and executes the files with the codified app.
- **Cloud native networking**, the network over which nodes (machines or processes) of a [distributed system](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/) are connected and communicate.

### 3\. The Orchestration and Management Layer

Once you automate infrastructure provisioning following security and compliance standards (provisioning layer) and set up the tools the app needs to run (runtime layer), engineers must figure out how to orchestrate and manage their apps. The orchestration and management layer deals with how all containerized services (app components) are managed as a group. They need to identify other services, communicate with one another, and coordinate. Inherently scalable, cloud native apps rely on automation and resilience, enabled by this layer.

In this layer you’ll find:

- **Orchestration and scheduling** to deploy and manage container clusters ensuring they are resilient, loosely-coupled, and scalable. In fact, the orchestration tool, in most cases [Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/), is what makes a cluster by managing containers and the operating environment
- **Coordination and service discovery** so services (app components) can locate and communicate with one another.
- **Remote procedure call (RPC)**, a technique enabling a service on one node to communicate with a service on a different node connected through a network.
- **Service proxy** is an intermediary placed between services through which they communicate. The sole purpose of the proxy is to exert more control over service communication, it doesn’t add anything to the communication itself. These proxies are crucial to service meshes mentioned below.
- **API gateway**, an abstraction layer through which external applications can communicate.
- **Service mesh** is similar to the API gateway in the sense that it’s a dedicated infrastructure layer through which apps communicate, but it provides policy-driven _internal_ service-to-service communication. Additionally, it may include everything from traffic encryption to service discovery, to application observability.

### 4\. The Application Definition and Development Layer

Now let’s move to the top layer. As the name suggests, the application definition and development layer focus on the tools that enable engineers to build apps and allow them to function.  Everything discussed above was related to building a reliable, secure environment and providing all needed app dependencies.

Under this category you’ll see:

- **Databases** enabling apps to collect data in an organized manner.
- **Streaming and messaging** enable apps to send and receive messages (events and streams). It’s not a networking layer, but rather a tool to queue and process messages.
- **Application definition and image build** are services that help configure, maintain, and run container images (the executable files of an app).
- **Continuous integration and delivery (CI/CD)** allow developers to automatically test that their code works with the codebase (the rest of the app) and, if their team is mature enough, even automate deployment into production.

## Tools Running Across All Layers

Going back to the category overview, we’ll explore the two columns running across all layers. Observability and analysis are tools that monitor all layers. Platforms, on the other hand, bundle multiple technologies within these layers into one solution, including observability and analysis.

### ![cncf landscape rows](https://cdn.thenewstack.io/media/2020/07/50c10db4-screen-shot-2020-07-17-at-2.39.23-pm.png)

### Observability and Analysis

To limit service disruption and help drive down MRRT (meantime to resolution), you’ll need to monitor and analyze every aspect of your application so any anomaly gets detected and rectified right away. Failures _will_ occur in complex environments and these tools help make them less impactful by helping identify and resolve failures as quickly as possible. Since this category runs across and monitors all layers, it’s on the side and not embedded in a specific layer.

Here you’ll find:

- **Logging** tools to collect event logs (info about processes).
- **Monitoring** solutions to collect metrics (numerical system parameters, such as RAM availability).
- **Tracing** goes one step further than monitoring and monitors the propagation of user requests. This is relevant in the context of service meshes.
- **Chaos engineering** are tools to test software in production to identify weaknesses and fix them before they impact service delivery.

### Platforms

As we’ve seen, each of these modules solves a particular problem. Storage alone does not provide all you need to manage your app. You’ll need an orchestration tool, container runtime, service discovery, networking, API gateway, etc. Covering multiple layers, platforms bundle different tools together solving a larger problem.

Configuring and fine-tuning different modules so they are reliable and secure and ensuring all the technologies it leverages are updated and vulnerabilities patched is no easy task. With platforms, users don’t have to worry about these details — a real value add.

You’ll probably notice, the categories all revolve around Kubernetes. That’s because Kubernetes, while only one piece of the puzzle, is at the core of the cloud native stack. The CNCF, by the way, was created with Kubernetes as its first seeding project; all other projects followed later.

Platforms can be categorized in four groups:

- **Kubernetes distributions** take the unmodified, open source code (although some modify it) and add additional features their market needs around it.
- **Hosted Kubernetes** (aka managed Kubernetes) is similar to a distribution but it’s managed by your provider on their or on your own infrastructure.
- **Kubernetes installers** are exactly that, they automate the installation and configuration process of Kubernetes.
- **PaaS / container services** are similar to hosted Kubernetes, but include a broad set of application deployment tools (generally a subset from the cloud native landscape).

## Conclusion

In each category, there are different tools aimed at solving the same or similar problems. Some are pre-cloud native technologies adapted to the new reality, while others are completely new. Differences lie in their implementation and design approaches. There is no perfect technology that checks all the boxes. In most cases technology is limited by design and architectural choices — there is always a tradeoff.

When selecting the stack, engineers must carefully consider each capability and tradeoff to identify the best option for their use case. While this brings additional complexity, it’s never been more feasible to choose a data storage, infrastructure management, messaging system, etc. that best fits the application’s needs. Architecting systems today is a lot easier than in a pre-cloud native world. And, if architected appropriately, cloud native technologies offer powerful and much-needed flexibility. In today’s fast-changing technology ecosystem, that is likely one of the most important capabilities.

We hope this quick overview was helpful. Stay tuned for our follow up articles to learn more about each layer and column!

**_As always, thanks to [Oleg Chunihkin](https://www.linkedin.com/in/olegch/) for all his input and also to [Jason Morgan](https://www.linkedin.com/in/jasonmorgan2/), my co-author for the upcoming more detailed landscape articles (very excited about that!). And a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._**

Feature image by [Huper by Joshua Earle](https://unsplash.com/@huper?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)

At this time, The New Stack does not allow comments directly on this website. We invite all readers who wish to discuss a story to visit us on [Twitter](https://twitter.com/thenewstack) or [Facebook](https://www.facebook.com/thenewstack/). We also welcome your news tips and feedback via email: [feedback@thenewstack.io](mailto:feedback@thenewstack.io).

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Real.

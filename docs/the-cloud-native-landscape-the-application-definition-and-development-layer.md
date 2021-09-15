# The Cloud Native Landscape: The Application Definition and Development Layer

#### 26 Jan 2021 6:00am,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack.io/author/jason-morgan/ "Posts by Jason Morgan")

![](https://cdn.thenewstack.io/media/2021/01/4fd4a803-cncf-landscape.png)

_This post is part of an ongoing series from [Cloud Native Computing Foundation Business Value Subcommittee](https://lists.cncf.io/g/cncf-business-value) co-chairs [Catherine Paganini](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) and [Jason Morgan](https://thenewstack.io/author/jason-morgan/) that focuses on explaining each category of the cloud native landscape to a non-technical audience as well as engineers just getting started with cloud native computing._


When looking at the [Cloud Native Landscape](https://landscape.cncf.io/), you’ll note a few distinctions:

- Projects in large boxes are CNCF-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- Projects in small white boxes are open source projects.
- Products in gray boxes are proprietary.

Please note that even during the time of this writing, we saw new projects becoming part of the CNCF so always refer to the actual landscape — things are moving fast!

## Database [![databases](https://cdn.thenewstack.io/media/2021/01/ef82aef4-screen-shot-2021-01-24-at-5.50.05-pm.png)](https://landscape.cncf.io/)

### What It Is

A database management system is an application through which other apps can efficiently store and retrieve data.

It ensures data gets stored, only authorized users are able to access it, and allows users to retrieve it via specialized requests. While there are numerous different types of databases with different approaches, they ultimately all have these same overarching goals.

### Problem It Addresses

Most applications need an effective way to store and retrieve data while keeping that data safe. Databases do this in a structured way with proven technology (there is quite a bit of complexity that goes into doing this well).

### How It Helps

Databases provide a common interface for storing and retrieving data for applications. Developers use these standard interfaces and, in most cases a simple query language, to store, query, and retrieve information in a database. At the same time, databases allow users to continuously back up and save data as well as encrypt and regulate access to it.

### Technical 101

We’ve established that a database management system is an application that stores and retrieves data. It does so using a common language and interface that can be easily used by a number of different languages and frameworks.

In general, we see two common types of databases, structured query language (SQL) databases, and no-SQL databases. Which database a particular application uses should be driven by its needs and constraints.

With the rise of Kubernetes and its ability to support stateful applications, we’ve seen a new generation of databases that leverage containerization. These new cloud native databases aim to bring the scaling and availability benefits of Kubernetes to databases. Tools like [YugaByte](https://landscape.cncf.io/?selected=yuga-byte-db) and [Couchbase](https://landscape.cncf.io/?selected=couchbase) are examples of cloud native databases, although more traditional databases like MySQL and Postgres run successfully and effectively in Kubernetes clusters.

[Vitess](https://landscape.cncf.io/?selected=vitess) and [TiKV](https://landscape.cncf.io/?selected=ti-kv) are CNCF projects in this space.

### Sidenote

If you look at this category, you’ll notice multiple names ending in DB (e.g. MongoDB, CockroachDB, FaunaDB) which, as you may guess, stands for database. You’ll also see various names ending in SQL (e.g. MySQL or memSQL) — they are still relevant. Some are “old school” databases that have been adapted to a cloud native reality. There are also some databases that are no-SQL but SQL compatible, such as YugaByte and Vitess.

BuzzwordsPopular Projects

- SQL
- DB
- Persistence

- Postgres
- MySQL
- Redis

## Streaming and Messaging

[![streaming and messaging ](https://cdn.thenewstack.io/media/2021/01/2d3a7aa5-screen-shot-2021-01-24-at-5.52.28-pm.png)](https://landscape.cncf.io/)

### What It Is

Streaming and messaging tools enable service-to-service communication by transporting messages (i.e. events) between systems. Individual services connect to the messaging service to either publish events, read messages from other services, or both. This dynamic creates an environment where individual apps are either publishers, meaning they write events, or subscribers that read events, or more likely both.

### Problem It Addresses

As services proliferate, application environments become increasingly complex, making the orchestration of communication between apps more challenging. A streaming or messaging platform provides a central place to publish and read all the events that occur within a system, allowing applications to work together without necessarily knowing anything about one another.

### How It Helps

When a service does something other services should know about, it “publishes” an event to the streaming or messaging tool. Services that need to know about these types of events subscribe and watch the streaming or messaging tool. That’s the essence of a publish-subscribe, or just pub-sub, approach and is enabled by these tools.

By introducing a “go-between” layer that manages all communication, we are decoupling services from one another. They simply watch for events, take action, and publish a new one.

Here’s an example. When you first sign up for Netflix, the signup service publishes a “new signup event” to a messaging platform with further details (e.g. name, email address, subscription level, etc.). The account creator service, which subscribes to signup events, will see the event and create your account. A customer communication service that also subscribes to new signup events will add your email address to the customer mailing list and generate a welcome email, and so on.

This allows for a highly decoupled architecture where services can collaborate without needing to know about one another. This decoupling enables engineers to add new functionality without updating downstream apps (consumers) or sending a bunch of queries. The more decoupled a system is, the more flexible and amenable to change. And that is exactly what engineers strive for in a system.

### Technical 101

Messaging and streaming tools have been around long before cloud native became a thing. To centrally manage business-critical events, organizations have built large enterprise service buses. But when we talk about messaging and streaming in a cloud native context, we’re generally referring to tools like NATS, RabbitMQ, Kafka, or cloud provided message queues.

What these systems have in common, is the architectural patterns they enable. Application interactions in a cloud native environment are either orchestrated or choreographed. You can find a lot more details about the topic in Sam Newton’s book [Building Microservices](https://samnewman.io/books/building_microservices/), as well as brief summaries [in this StackOverflow post](https://stackoverflow.com/questions/4127241/orchestration-vs-choreography) and [this great Medium blog by Chen Chen](https://medium.com/ingeniouslysimple/choreography-vs-orchestration-a6f21cfaccae). To simplify things, let’s just say that orchestrated refers to systems that are centrally managed and choreographed are those that allow individual components to act independently.

Messaging and streaming systems provide a central place for choreographed systems to communicate. The message bus provides a common place where all apps can go to, to both tell others what they’re doing by publishing messages, or see what things are going on by subscribing to messages.

The [NATS](https://landscape.cncf.io/?selected=nats) and [Cloudevents](https://landscape.cncf.io/?selected=cloud-events) projects are both incubating projects in this space with NATS providing a mature messaging system and Cloudevents being an effort to standardize message formats between systems. [Strimzi](https://landscape.cncf.io/?selected=strimzi), [Pravega](https://landscape.cncf.io/?selected=pravega), and [Tremor](https://landscape.cncf.io/?selected=tremor) are sandbox projects with each being tailored to a unique use case around streaming and messaging.

BuzzwordsPopular Projects

- Choreography
- Steaming
- MQ
- Message Bus

- Spark
- Kafka
- RabbitMQ
- Nats

## Application Definition and Image Build

[![app definition and image build](https://cdn.thenewstack.io/media/2021/01/1165999f-screen-shot-2021-01-24-at-5.54.25-pm.png)](https://landscape.cncf.io/)

### What It Is

App definition and image build is a broad category that can be broken down into two main subgroups. First, dev-focused tools that help build application code into containers and/or Kubernetes. And second, ops-focused tools that deploy apps in a standardized way. Whether to speed up or simplify your development environment, provide a standardized way to deploy third-party apps, or simplify the process of writing a new Kubernetes extension, this category serves as a catch-all for a number of projects and products that optimize the Kubernetes developer and operator experience.

### Problem It Addresses

Kubernetes, and containerized environments more generally, are incredibly flexible and powerful. With that flexibility also comes complexity, mainly in form of numerous configuration options for various often new use cases. Developers must containerize their code and have the ability to develop in production-like environments. And with a rapid release schedule, operators need a standardized way to deploy apps into container environments.

### How It Helps

Tools in this space aim at solving some of these developer or operator challenges. On the developer side, there are tools that simplify the process of extending Kubernetes to build, deploy, and connect applications. A number of projects and products help to store or deploy pre-packaged apps. These allow operators to quickly deploy a streaming service like Kafka or install a service mesh like Linkerd.

Developing cloud native applications brings a whole new set of challenges calling for a large set of diverse tools to simplify application build and deployments. As you start addressing operational and developer concerns in your environment, look for tools in this category.

### Technical 101

App definition and build tools encompass a huge range of functionality. From extending [Kubernetes](https://landscape.cncf.io/?selected=kubernetes) to virtual machines with [KubeVirt](https://landscape.cncf.io/?selected=kube-virt), to speeding app development by allowing you to port your development environment into Kubernetes with tools like Telepresence, and everything in between. At a high level, tools in this space solve either developer-focused concerns, like how to correctly write, package, test, or run custom apps, or operations-focused concerns, such as deploying and managing applications.

Helm, the only graduated project in this category, underpins many app deployment patterns. Allowing Kubernetes users to deploy and customize many popular third-party apps, it has been adopted by other projects like the [Artifact Hub](https://landscape.cncf.io/?selected=artifact-hub), a CNCF sandbox project, and [Bitnami](https://landscape.cncf.io/?selected=bitnami-tanzu-application-catalog) to provide curated catalogs of apps. [Helm](https://landscape.cncf.io/?selected=helm) is also flexible enough to allow users to customize their own app deployments.

The [Operator Framework](https://landscape.cncf.io/?selected=operator-framework) is an incubating project aimed at simplifying the process of building and deploying operators. Operators are out of scope for this article but let’s note here that they help deploy and manage apps, similar to Helm (you can read more about operators [here](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)). [Cloud Native Buildpacks](https://landscape.cncf.io/?selected=buildpacks), another incubating project, aims to simplify the process of building application code into containers.

There’s a lot more in this space and exploring it all would require a dedicated article. But research these tools further if you want to make Kubernetes easier for developers and operators. You’ll likely find something that meets your needs.

BuzzwordsPopular Projects

- Package Management
- Charts
- Operators

- Helm
- Buildpacks
- Tilt
- Okteto

## Continuous Integration and Continuous Delivery

[![CICD](https://cdn.thenewstack.io/media/2021/01/836e47cf-screen-shot-2021-01-24-at-5.55.02-pm.png)](https://landscape.cncf.io/)

### What It Is

Continuous integration (CI) and continuous delivery (CD) tools enable fast and efficient development with embedded quality assurance (we cover [CI/CD in detail in this primer](https://thenewstack.io/a-primer-continuous-integration-and-continuous-delivery-ci-cd/)). CI automates code changes by immediately building and testing the code, ensuring it produces a deployable artifact. CD goes one step further and pushes the artifact through the deployment phases.

Mature CI/CD systems watch source code for changes, automatically build and test the code, then begin moving it from development to production where it has to pass a variety of tests or validation to determine if the process should continue or fail. Tools in this category enable such an approach.

### Problem It Addresses

Building and deploying applications is a difficult and error-prone process. Particularly, when it involves a lot of human intervention and manual steps. The longer a developer works on a piece of software without integrating it into the codebase, the longer it will take to identify an error and the more difficult it is to fix. By integrating code on a regular basis, errors are caught early and are easier to troubleshoot. After all, finding an error in a few lines of code is a lot easier than doing so in a few hundred lines of code.

While tools like Kubernetes offer great flexibility for running and managing apps, they also create new challenges and opportunities for CI/CD tooling. Cloud native CI/CD systems are able to leverage Kubernetes itself to build, run, and manage the CI/CD process, often referred to as [pipelines](https://www.redhat.com/en/topics/devops/what-cicd-pipeline). Kubernetes also provides information about the health of our apps enabling cloud native CI/CD tools to more easily determine if a given change was successful or needs to be rolled back.

### How It Helps

CI tools ensure that any code change or updates developers introduce are built, validated, and integrated with other changes automatically and continuously. Each time a developer adds an update, automated testing is triggered ensuring only good code makes it into the system. CD extends CI to include pushing the result of the CI process into production-like and production environments.

Let’s say a developer changes the code for a web app. The CI system sees the code change, and then builds and tests a new version of that web app. The CD system takes that new version and deploys it into a dev, test, pre-production, and finally production environment. It does that while testing the deployed app after each step in the process. All together these systems represent a CI/CD pipeline for that web app.

### Technical 101

Over time, a number of tools have been built to help with the process of moving code from a repository, where the code is stored, to production, where the finished app runs. Like most other areas of computing, the advent of cloud native development has changed CI/CD systems. Some traditional tools like [Jenkins](https://landscape.cncf.io/?selected=jenkins), probably the most widely-used CI tool on the market, has been [overhauled](https://jenkins-x.io/) entirely to better fit into the Kubernetes ecosystem. Others, like [Flux](https://landscape.cncf.io/?selected=flux) and [Argo](https://landscape.cncf.io/?selected=argo) have pioneered a new way of doing continuous delivery called GitOps.

In general, you’ll find projects and products in this space are either (1) CI systems, (2) CD systems, (3) tools that help the CD system decide if the code is ready to be pushed into production, or (4), in the case of [Spinnaker](https://landscape.cncf.io/?selected=spinnaker) and Argo, all three. Argo and Brigade are the only CNCF projects in this space but you can find many more options hosted by the [Continuous Delivery Foundation](https://cd.foundation/). Look for tools in this space to help your organization automate your path to production.

BuzzwordsPopular Projects

- CI/CD
- Continuous integration
- Continuous delivery
- Blue/Green
- Canary Deploy

- Argo
- Flagger
- Spinnaker
- Jenkins

## Conclusion

As we’ve seen, tools in the application definition and development layer enable engineers to build cloud native apps. You’ll find databases to store and retrieve data or streaming and messaging tools allowing for decoupled, choreographed architectures. App definition and image build tools encompass a variety of technologies that improve the developer and operator experience, and CI/CD ensures code is at a deployable state and helps engineers catch any errors early on, ensuring better code quality.

This article concludes the layers of the CNCF landscape. In our next article, we’ll focus on cloud native platforms, the first column running across all these layers. Configuring tools across the layers we discussed so far so they work well together is no easy task. Platforms bundle them together easing adoption, but more to that in our next article.

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

The Cloud Native Computing Foundation is a sponsor of The New Stack.

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Bit.

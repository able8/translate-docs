# The Cloud Native Landscape: The Orchestration and Management Layer

#### 15 Dec 2020 11:54am,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack.io/author/jason-morgan/ "Posts by Jason Morgan")

![](https://cdn.thenewstack.io/media/2020/12/57e7bc17-meteora-3717220_640.jpg)

_This post is part of an ongoing series from_ [_CNCF Business Value Subcommittee_](https://lists.cncf.io/g/cncf-business-value) _co-chairs_ [_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _that focuses on explaining each category of the cloud native landscape to a non-technical audience as well as engineers just getting started with cloud native._

[![](https://cdn.thenewstack.io/media/2020/12/5984c027-screen-shot-2020-12-08-at-8.51.01-am.png)\
\
Catherine Paganini\
\
Catherine is Head of Marketing at Buoyant, the creator of Linkerd. A marketing leader turned cloud native evangelist, Catherine is passionate about educating business leaders on the new stack and the critical flexibility it provides.](https://www.linkedin.com/in/catherinepaganini/en/)

The orchestration and management layer is the third layer in the [Cloud Native Computing Foundation’s cloud native landscape](https://landscape.cncf.io). Before tackling tools in this category, engineers have presumably already automated infrastructure provisioning following security and compliance standards ( [provisioning layer](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)) and set up the runtime for the application ( [runtime layer](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)). Now they must figure out how to orchestrate and manage all app components as a group. Components must identify one another to communicate and coordinate to accomplish a common goal. Inherently scalable, cloud native apps rely on automation and resilience, enabled by these tools.

### **_Sidenote:_**

[![](https://cdn.thenewstack.io/media/2020/08/572c2815-jason.png)\
\
Jason Morgan\
\
Jason Morgan, a Solutions Engineer with VMware, focuses on helping customers build and mature microservices platforms. Passionate about helping others on their cloud native journey, Jason enjoys sharing lessons learned with the broader developer community.](https://blog.59s.io/)

When looking at the [Cloud Native Landscape](https://landscape.cncf.io), you’ll note a few distinctions:

- **Projects in large boxes** are CNCF-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- **Projects in small white boxes** are open source projects.
- **Products/projects in gray boxes** are proprietary.

Please note that even during the time of this writing, we saw new projects becoming part of the CNCF so always refer to the actual landscape — things are moving fast!

## Orchestration and Scheduling

### What It Is

Orchestration and scheduling refer to running and managing containers, a new-ish way to package and ship applications, across a cluster. A cluster is a group of machines, physical or virtual, connected over a network.

Container orchestrators (and schedulers) are somewhat similar to the operating system (OS) on your laptop which manages all your apps (e.g. Microsoft 360, Slack, Zoom, etc.). The OS executes the apps you want to use and schedules when which app gets to use your laptop’s CPU and other hardware resources.

While running everything on a single machine is great, most applications today are a lot bigger than one computer can possibly handle. Think Gmail or Netflix. These massive apps are distributed across multiple machines forming a [distributed application](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/). Most modern-day applications are distributed, and that requires software that is able to manage all components running across these different machines. In short, you need a “cluster OS.” That’s where orchestration tools come in.

If you’ve read our previous articles, you probably noticed that containers come up time and again. Their ability to run apps in many different environments is key. And so are container orchestrators which, in most cases, is [Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/). Containers and Kubernetes are both central to cloud native architectures, which is why we hear so much about them.

### Problem It Addresses

In cloud native architectures, applications are broken down into small components, or services, each placed in a container. You may have heard of them referred to as microservices. Instead of having one large application, you now have multiple small services each in need of resources, monitoring, and fixing if a problem occurs. While it is feasible to do those things manually for a single service, you’ll need automated processes when you have hundreds of containers.

### How It Helps

Container orchestrators automate container management. But what does that mean in practice? Let’s answer that for Kubernetes specifically since it’s the de facto container orchestrator.

Kubernetes does something called desired state reconciliation: it matches the _current_ state of containers within a cluster to the _desired_ state. The desired state is specified by the engineer in a file (e.g. ten instances of service A running on three nodes, i.e. machines, with access to database B, etc.) and continuously compares against the actual state. If the desired and actual state don’t match, Kubernetes reconciles them by creating or destroying objects (e.g. if a container crashes, it will spin a new one up).

In short, Kubernetes allows you to treat a cluster as one computer. It focuses only on what that environment should look like and handles the implementation details for you.

### Technical 101

Kubernetes lives in the orchestration and scheduling section along with other container orchestrators like Docker Swarm and Mesos. Its basic purpose is to allow you to manage a number of disparate computers as a single pool of resources. On top of that, it allows you to manage them in a declarative way, i.e. instead of telling Kubernetes how to do something you provide a definition of what you want to be done. This allows you to maintain the desired state in one or more YAML files and apply it more or less unchanged to any Kubernetes cluster. The orchestrator itself then creates anything that’s missing or deletes anything that should no longer exist. While Kubernetes isn’t the only orchestrator that the CNCF hosts as a project (both Crossplane and Volcano are incubating projects) it is the most commonly used and actively maintained.

**Buzzwords**

**Popular Projects**

- Cluster
- Scheduler
- Orchestration

- Kubernetes
- Docker Swarm
- Mesos

![Scheduling and orchestration](https://cdn.thenewstack.io/media/2020/12/cf616dae-screen-shot-2020-12-08-at-8.21.04-am.png)

## Coordination and Service Discovery

### What It Is

As we’ve seen, modern applications are composed of multiple individual services that need to collaborate to provide value to the end user. To collaborate, they communicate over a network (discussed in our [runtime layer article](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)). And to communicate, they must first locate one another. Service discovery is the process of figuring out how to do that.

### Problem It Addresses

Cloud native architectures are dynamic and fluid, meaning they are constantly changing. When a container crashes on one node, a new container is spun up on a different node to replace it. Or, when an app scales, replicas are spread out throughout the network. There is no one place where a particular service is; the location of everything is constantly changing. Tools in this category keep track of services within the network so services can find one another when needed.

### How It Helps

Service discovery tools address this problem by providing a common place to find and potentially identify individual services. There are basically two types of tools in this category: (1) Service discovery engines are database-like tools that store info about what services exist and how to locate them. And (2) name resolution tools (e.g. Core DNS) receive service location requests and return network address information.

### **_Sidenote:_**

In Kubernetes, to make a pod reachable a new abstraction layer called “service”is introduced. Services provide a single stable address for a dynamically changing group of pods.

Please note that “service” may have different meanings in different contexts, which can be quite confusing. The term “services” generally refers to the service placed inside a container/pod. It’s the app component or microservice with a specific function within the actual app (e.g. face recognition algorithm of your iPhone).

A Kubernetes service, on the other hand, is the abstraction that helps pods find and connect to each other. It is an entry point for a service (functionality) as a collection of processes or pods. In Kubernetes, when you create a service (abstraction), you create a group of pods that together provide a service (functionality) with a single endpoint (entry) which is the Kubernetes service.

### Technical 101

As distributed systems became more and more prevalent, traditional DNS processes and traditional load balancers were often unable to keep up with changing endpoint information. To make up for these shortcomings, service discovery tools were created to handle individual application instances rapidly registering and deregistering themselves. Some options such as etcd and CoreDNS are Kubernetes native, others have custom libraries or tools to allow services to operate effectively. CoreDNS and etcd are CNCF projects and are built into Kubernetes.

**Buzzwords**

**Popular Projects**

- DNS
- Service Discovery

- CoreDNS
- etcd
- Zookeeper
- Eureka

![Coordination and service discovery](https://cdn.thenewstack.io/media/2020/12/673ee53a-screen-shot-2020-12-08-at-8.23.02-am.png)

## Remote Procedure Call

### What It Is

Remote Procedure Call (RPC) is a particular technique enabling applications to talk to each other. It represents one of a few ways for applications to structure their communications with one another. This section is not particularly relevant to non-developers.

### Problem It Addresses

Modern apps are composed of numerous individual services that must communicate in order to collaborate. RPC is one option for handling the communication between applications.

### How It Helps

RPC provides a tightly coupled and highly opinionated way of handling communication between services. It allows for bandwidth-efficient communications and many languages enable RPC interface implementations. RPC is not the only nor the most common way to address this problem.

### Technical 101

RPC provides a highly structured and tightly coupled interface for communications between services. There are a lot of potential benefits with RPC: It makes coding the connection easier, it allows for extremely efficient use of the network layer and well-structured communications between services. RPC has also been criticized for creating brittle connection points and forcing users to do coordinated upgrades for multiple services. gRPC is a particularly popular RPC implementation and has been adopted by the CNCF.

**Buzzwords**

**Popular Projects**

- gRPC

- gRPC

![Remote procedure call](https://cdn.thenewstack.io/media/2020/12/1e4e9de8-screen-shot-2020-12-08-at-8.26.32-am.png)

## Service Proxy

### What It Is

A service proxy is a tool that intercepts traffic to or from a given service, applies some logic to it, then forwards that traffic to another service. It essentially acts as a “go-between” that collects information about network traffic and/or applies rules to it. This can be as simple as serving as a load balancer that forwards traffic to individual applications or as complex as an interconnected mesh of proxies running side by side with individual containerized applications handling all network connections.

While a service proxy is useful in and of itself, especially when driving traffic from the broader network into a Kubernetes cluster, service proxies are also building blocks for other systems, such as API gateways or service meshes, which we’ll discuss below.

### Problem It Addresses

Applications should send and receive network traffic in a controlled manner. To keep track of the traffic and potentially transform or redirect it, we need to collect data. Traditionally, the code enabling data collection and network traffic management was embedded within each application.

A service proxy allows us to “externalize” this functionality. No longer does it need to live within the apps. Instead, it’s now embedded into the platform layer (where your apps run). This is incredibly powerful because it allows developers to fully focus on writing application logic, your value generating code, while the universal task of handling traffic is managed by the platform team (whose responsibility it is in the first place). By centralizing the distribution and management of globally needed service functionality (e.g. routing or TLS termination) from a single, common location, communication between services is more reliable, secure, and performant.

### How It Helps

Proxies act as gatekeepers between the user and services or between different services. With this unique positioning, they provide insight into what type of communication is happening. Based on their insight, they determine where to send a particular request or even deny it entirely.

Proxies gather critical data, manage routing (spreading traffic evenly among services or rerouting if some services break down), encrypt connections, and cache content (reducing resource consumption).

### Technical 101

Service proxies work by intercepting traffic between services, performing some logic on them, then potentially allowing traffic to move on. By putting a centrally controlled set of capabilities into this proxy, administrators are able to accomplish several things. They can gather detailed metrics about inter-service communication, protect services from being overloaded, and apply other common standards to services, like mutual TLS. Service proxies are fundamental to other tools like service meshes as they provide a way to enforce higher-level policies to all network traffic.

Please note, the CNCF includes load balancers and ingress providers into this category. Envoy, Contour, and BFE are all CNCF projects.

**Buzzwords**

**Popular Projects**

- Service proxy
- Ingress

- Envoy
- Contour
- NGINX

![Service proxy](https://cdn.thenewstack.io/media/2020/12/ed807753-screen-shot-2020-12-08-at-8.28.02-am.png)

## API Gateway

### What It Is

While humans generally interact with computer programs via a GUI (graphical user interface) such as a webpage or a (desktop) application, computers interact with each other through APIs (application programming interfaces). But an API shouldn’t be confused with an API gateway.

An API gateway allows organizations to move key functions, such as authorizing or limiting the number of requests between applications, to a centrally managed location. It also functions as a common interface to (often external) API consumers.

Through API gateways, organizations can centrally control (limit or enable) interactions between apps and keep track of them, enabling things like chargeback, authentication, and protecting services from being overused (aka rate-limiting).

### Example

Take Amazon store cards. To offer them, Amazon partners with a bank that will issue and manage all Amazon store cards. In return, the bank will keep, let’s say, $1 per transaction. To authorize the retailer to request new cards, keep track of transactions, and maybe even restrict the number of requested cards per minute, the bank will use an API gateway. All that functionality is encoded into the gateway, not the services using it. Services just worry about issuing cards.

### Problem It Addresses

While most containers and core applications have an API, an API gateway is more than just an API. An API gateway simplifies how organizations manage and apply rules to all interactions.

API gateways allow developers to write and maintain less custom code (the system functionality is encoded into the API gateway, remember?). They also enable teams to see and control the interactions between application users and the applications themselves.

### How It Helps

An API gateway sits between the users and the application. It acts as a go-between that takes the messages (requests) from the users and forwards them to the appropriate service. But before handing the request off, it evaluates whether the user is allowed to do what they’re trying to do and records details about who made the request and how many requests they’ve made.

Put simply, an API gateway provides a single point of entry with a common user interface for app users. It also enables you to handoff tasks otherwise implemented within the app to the gateway, saving developer time and money.

### Technical 101

Like many categories in this layer, an API gateway takes custom code out of our apps and brings it into a central system. The API gateway works by intercepting calls to backend services, performing some kind of value add activity like validating authorization, collecting metrics, or transforming requests, then performing whatever action it deems appropriate. API gateways serve as a common entry point for a set of downstream applications while at the same time providing a place where teams can inject business logic to handle authorization, rate limiting, and chargeback. They allow application developers to abstract away changes to their downstream APIs from their customers and offload tasks like onboarding new customers to the gateway.

**Buzzwords**

**Popular Projects/Products**

- API gateway

- Kong
- Mulesoft
- Ambassador

![API gateway](https://cdn.thenewstack.io/media/2020/12/0550d18a-screen-shot-2020-12-08-at-8.29.46-am.png)

## Service Mesh

### What It Is

If you’ve been reading a little bit about cloud native, the term service mesh probably rings a bell. It’s been getting quite a bit of attention lately. According to long-time TNS contributor Janakiram MSV, “after Kubernetes, the service mesh technology has become the most critical component of the cloud native stack.” So pay attention to this one — you’ll likely hear a lot more about it.

Service meshes manage traffic (i.e. communication) between services. They enable platform teams to add reliability, observability, and security features uniformly across all services running within a cluster without requiring any code changes.

### Problem It Addresses

In a cloud native world, we are dealing with multiple services all in need to communicate. This means a lot more traffic is going back and forth on an inherently unreliable and often slow network. To address this new set of challenges, engineers must implement additional functionality. Prior to the service mesh, that functionality had to be encoded into every single application. This custom code often became a source of technical debt and provided new avenues for failures or vulnerabilities.

### How It Helps

Service meshes add reliability, observability, and security features uniformly across all services on a platform layer without touching the app code. They are compatible with any programming language, allowing development teams to focus on writing business logic.

### **_Sidenote:_**

Since traditionally, these service mesh features had to be coded into each service, each time a new service was released or updated, the developer had to ensure these features were functional, too, providing a lot of room for human error. And here’s a dirty little secret, developers prefer focusing on business logic (value-generating functionalities) rather than building reliability, observability, and security features. For the platform owners, on the other hand, these are core capabilities and central to everything they do. Making developers responsible for adding features that platform owners need is inherently problematic. This, by the way, also applies to general-purpose proxies and API gateways mentioned above. Service meshes and API gateways solve that very issue as they are implemented by the platform owners and applied universally across all services.

### Technical 101

Service meshes bind all services running on a cluster together via service proxies creating a mesh of services, hence service mesh. These are managed and controlled through the service mesh control plane. Service meshes allow platform owners to perform common actions or collect data on applications without having developers write custom logic.

In essence, a service mesh is an infrastructure layer that manages inter-service communications by providing command and control signals to a network, or mesh, of service proxies. Its power lies in its ability to provide key system functionality without having to modify the applications.

Some service meshes use a general-purpose service proxy (see above) for their data plane. Others use a dedicated proxy; Linkerd, for example, uses the [Linkerd2-proxy “micro proxy”](https://linkerd.io/) to gain an advantage in performance and resource consumption. These proxies are uniformly attached to each service through so-called sidecars. Sidecar refers to the fact that the proxy runs in its own container but lives in the same pod. Just like a motorcycle sidecar, it’s a separate module attached to the motorcycle, following it wherever it goes.

### Example:

Take circuit breaking. In microservice environments, individual components often fail or begin running slowly. Without a service mesh, developers would have to write custom logic to handle downstream failures gracefully and potentially set cooldown timers to avoid that upstream services continually request responses from degraded or failed downstream services. With a service mesh, that logic is handled at a platform level.

Service meshes provide many useful features, including the ability to surface detailed metrics, encrypt all traffic, limit what operations are authorized by what service, provide additional plugins for other tools, and much more. For more detailed information, check out the [service mesh interface](https://smi-spec.io/) specification.

**Buzzwords**

**Popular Projects**

- Service mesh
- Sidecar
- Data plane
- Control plane

- Linkerd
- Consul
- Istio

![Service mesh](https://cdn.thenewstack.io/media/2020/12/458f49aa-screen-shot-2020-12-08-at-8.30.08-am.png)

## Conclusion

As we’ve seen, tools in this layer deal with how all these independent containerized services are managed as a group. Orchestration and scheduling tools are some sort of cluster OS managing containerized applications across your cluster. Coordination and service discovery, service proxies, and service meshes ensure services can find each other and communicate effectively in order to collaborate as one cohesive app. API gateways are an additional layer providing even more control over service communication, in particular between external applications.In our next article, we’ll discuss the application definition and development layer — the last layer of the CNCF landscape. It covers databases, streaming and messaging, application definition, and image build, as well as continuous integration and delivery.

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

The Cloud Native Computing Foundation and VMware are sponsors of The New Stack.

Feature image [Antonios Ntoumas](https://pixabay.com/fr/users/atlantios-4957810/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3717220) de [Pixabay](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3717220).

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: MADE, Docker, Ambassador, Prevalent, Bit.

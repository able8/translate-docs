# The Cloud Native Landscape: The Runtime Layer Explained

#### 29 Sep 2020 1:24pm,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack.io/author/jason-morgan/ "Posts by Jason Morgan")

_This post is part of an ongoing series from_ [_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _that focuses on explaining each category of the cloud native landscape to a non-technical audience as well as engineers just getting started with cloud native._

Jason Morgan, a Solutions Engineer with VMware, focuses on helping customers build and mature microservices platforms. Passionate about helping others on their cloud native journey, Jason enjoys sharing lessons learned with the broader developer community.](https://blog.59s.io/)

In our previous article, we explored the [provisioning layer of the Cloud Native Computing Foundation’s cloud native landscape](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/) which focuses on building the foundation of your cloud native platforms and applications. This article zooms into the runtime layer encompassing everything a container needs in order to run in a cloud native environment. That means the code used to start a container, referred to as the runtime engine; the tools to make persistent storage available to containers; and those that manage the container environment networks.

But note, these resources shouldn’t be confused with the networking and storage work handled by the infrastructure and provisioning layer concerned with getting the container platform running. Tools in this category are used by the containers directly to start/stop, store data, and talk to each other.

![runtime layer ](https://cdn.thenewstack.io/media/2020/09/756d2023-screen-shot-2020-09-23-at-6.34.09-pm.png)

When looking at the [Cloud Native Landscape](https://landscape.cncf.io), you’ll note a few distinctions:

- **Projects in large boxes** are CNCF-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- **Projects in small white boxes** are open source projects.
- **Products/projects in gray boxes** are proprietary.

Please note that even during the time of this writing, we saw new projects becoming part of the CNCF so always refer to the actual landscape — things are moving fast!

## Cloud Native Storage

### What It Is

Storage is where the persistent data of an app is stored, often referred to as persistent volume. Easy access to it is critical for the app to function reliably. Generally, when we say persistent data we mean storing things like databases, messages, or any other information we want to ensure doesn’t disappear when an app gets restarted.

### Problem It Addresses

Cloud native architectures are fluid, flexible, and elastic, making persisting data between restarts challenging. To scale up and down or self-heal, containerized apps are continuously created and deleted, changing physical location over time. Therefore, cloud native storage must be provided in a node-independent fashion. To store data, however, you’ll need hardware — a disk to be specific — and disks, just like any other hardware, are infrastructure-bound. That’s the first big challenge.

Then there is the actual storage interface which can change significantly between datacenters (in the old world, each infrastructure had their own storage solution with its own interface), making portability really tough. And lastly, to benefit from the elasticity of the cloud, storage must be provisioned in an automated fashion as manual provisioning and autoscaling aren’t compatible.

Cloud native storage is tailored to this new cloud native reality.

### How It Helps

The tools in this category help either a) provide cloud native storage options for containers, b) standardize the interfaces between containers and storage providers or c) provide data protection through backup and restore operations. The former means storage that uses a cloud native compatible container storage interface (aka tools in the second category) and which can be provisioned automatically, enabling autoscaling and self-healing by eliminating the human bottleneck.

### Technical 101

Cloud native storage is largely made possible by the Container Storage Interface (CSI) which allows a standard API for providing file and block storage to containers. There are a number of tools in this space, both open source and vendor-provided that leverage the CSI to provide on-demand storage to containers. In addition to that extremely important functionality, we have a number of other tools and technologies which aim to solve storage problems in the cloud native space. Minio is a popular project that, among other things, provides an S3-compatible API for object storage. Tools like Velero help simplify the process of backing up and restoring both the Kubernetes clusters themselves as well as persistent data used by the applications.

**Buzzwords**

**Popular Projects/Products**

- CSI
- Storage API
- Backup and Restore

- **Minio**
- **CSI**
- **Ceph + Rook**
- **Velero**

![cloud native storage](https://cdn.thenewstack.io/media/2020/09/3160a300-screen-shot-2020-09-23-at-6.39.03-pm.png)

## Container Runtime

### What It Is

As discussed in the [provisioning layer article](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/), a container is a set of **compute constraints** used to execute (that’s tech-speak for launch) an application. Containerized apps believe they are running on their own dedicated computer and are oblivious that they are sharing resources with other processes (similar to virtual machines).

The container _runtime_ is the software that executes containerized (or “constrained”) applications. Without the runtime, you only have the container image, the file specifying how the containerized app should look like. The runtime will start an app within a container and provide it with the needed resources.

### Problem It Addresses

Container images (the files with the application specs) must be launched in a standardized, secure, and isolated way. **Standardized** because you need standard operating rules no matter where they are running. **Secure**, well, because you don’t want anyone who shouldn’t access it to do so. And **isolated** because you don’t want the app to affect or be affected by other apps (for instance, if a co-located application crashes). Isolation basically functions as protection. Additionally, the application must be provided resources, from CPU to storage to memory.

### How It Helps

The container runtime does all that. It launches apps in a standardized fashion across all environments and sets security boundaries. The latter is where some of these tools differ. Runtimes like CRI-O or gVisor have hardened their security boundaries. The runtime also sets resource limits for the container. Without it, the app could consume resources as needed, potentially taking resources away from other apps, so you always need to set limits.

### Technical 101

Not all tools in this category are created equal. Containerd (part of the famous Docker product) and CRI-O are standard container runtime implementations. Then there are tools that expand the use of containers to other technologies, such as Kata which allows you to run containers as VMs. Others aim at solving a specific container-related problem such as gVisor which provides an additional security layer between containers and the OS.

**Buzzwords**

**Popular Projects/Products**

- Container
- MicroVM

- **Containerd**
- **CRI-O**
- **Kata**
- **gVisor**
- **Firecracker**

![Container runtime](https://cdn.thenewstack.io/media/2020/09/6d332514-screen-shot-2020-09-23-at-6.48.22-pm.png)

## Cloud Native Networking

### What It Is

Containers talk to each other and to the infrastructure layer through a cloud native network. [Distributed applications](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/) have multiple components that use the network for different purposes.Tools in this category overlay a virtual network on top of existing networks specifically for apps to communicate, referred to as an _overlay_ network.

### Problem It Addresses

While it’s common to refer to the code running in a container as an app, the reality is that most containers hold only a small specific set of functionalities of a larger application. Modern applications such as Netflix or Gmail are actually composed of a number of these smaller components each running in its own container. For all these independent pieces to function as a cohesive application, containers need to communicate with each other privately. Tools in this category provide that private communication network.

Additionally, messages exchanged between these containers may be private, sensitive, or extremely important. This leads to additional requirements such as providing isolation for the various components and the ability to inspect traffic to identify network issues. In some use cases, you may want to extend these networks and network policies (e.g. firewall and access rules) so your app can connect to virtual machines or services running externally to our container network.

### How It Helps

Projects and products in this category use the CNCF project – Container Network Interface (CNI) to provide networking functionalities to containerized applications. Some tools, like Flannel, are rather minimalistic providing bare-bones connectivity to containers. Others, such as NSX-T provide a full software-defined networking layer creating an isolated virtual network for every Kubernetes namespace.

At a minimum, a container network needs to assign IP addresses to pods (that’s where containerized apps run in Kubernetes), that allows other processes to access it.

### Technical 101

Similar to storage, the variety and innovation in this space is largely made possible by the CNCF project CNI (Container Networking Interface) which standardizes how network layers provide functionalities to pods.Selecting the right container network for your Kubernetes environment is critical and you’ve got a number of tools to choose from. Weave Net, Antrea, Calico, and Flannel all provide effective open source networking layers. Their functionalities vary widely and your choice should be ultimately driven by your specific needs.

Additionally,there are many vendors ready to support and extend your Kubernetes networks with Software Defined Networking (SDN) tools that allow you to gain additional insights into network traffic, enforce network policies, and even extend your container networks and policies to your broader datacenter.

**Buzzwords**

**Popular Projects/Products**

- SDN
- Network Overlay
- CNI

- **Calico**
- **Weave Net**
- **Flannel**
- **Antrea**
- **NSX-T**

![cloud native network](https://cdn.thenewstack.io/media/2020/09/01e7d965-screen-shot-2020-09-23-at-6.51.02-pm.png)

This concludes our overview of the runtime layer which provides all the tools containers need to run in a cloud native environment. From storage that gives apps easy and fast access to data needed to run reliably, to the container runtime which executes the application code, to the network over which containerized apps communicate. In our next article, we’ll focus on the orchestration and management layer which deals with how all these containerized apps are managed as a group.

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

The Cloud Native Computing Foundation and VMware are sponsors of The New Stack.

Feature Image by [Candid\_Shots](https://pixabay.com/users/Candid_Shots-11873433/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=5582775) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=5582775).

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: MADE, Docker, Famous.

# 3 Years of Kubernetes in Production–Here’s What We Learned

## Key takeaways from our Kubernetes journal


[Sep 9, 2020](https://betterprogramming.pub/3-years-of-kubernetes-in-production-heres-what-we-learned-44e77e1749c8?source=post_page-----44e77e1749c8--------------------------------) · 7 min read

We started out building our first Kubernetes cluster in 2017, version  1.9.4. We had two clusters, one that ran on bare-metal RHEL VMs, and  another that ran on AWS EC2.

Today, our Kubernetes infrastructure fleet consists of over 400 virtual  machines spread across multiple data-centres. The platform hosts  highly-available mission-critical software applications and systems, to  manage a massive live network with nearly four million active devices.

Kubernetes eventually made our lives easier, but the journey was a hard one, a  paradigm shift. There was a complete transformation in not just our  skillset and tools, but also our design and thinking. We had to embrace  multiple new technologies and invest massively to upscale and upskill  our teams and infrastructure.

Looking back, after three years of running Kubernetes in production, here are key lessons from our journal.

[How a Faulty Visa System is Killing the Technology MarketWoes of immigrant software engineers on work visasmedium.com](https://medium.com/digital-diplomacy/how-a-faulty-visa-system-is-killing-the-technology-market-adcd8588d880)

# 1. The Curious Case of Java Apps

When it comes to microservices and containerization, engineers tend to steer clear of using Java, primarily due to its notorious memory management.  However, things have changed now and Java’s container compatibility has  improved over the years. After all, ubiquitous systems like `Apache Kafka`and `Elasticsearch`run on Java.

Back in 2017–18, we had a few apps that ran on Java version 8. These often  struggled to understand container environments like Docker and crashed  from heap memory issues and unusual garbage collection trends. We  learned that these were [caused by JVM’s inability ](https://developers.redhat.com/blog/2017/03/14/java-inside-docker/)to honor Linux `cgroups`and `namespaces` , that are at the core of containerization technology.

However, since then, Oracle has been continuously improving Java’s compatibility in the container world. Even Java 8's subsequent patches introduced  experimental JVM flags to tackle these problems, `XX:+UnlockExperimentalVMOptions` and `XX:+UseCGroupMemoryLimitForHeap`

But despite all the improvements, there is no denying that Java still has a bad reputation for hogging memory and its slow startup time compared to its peers like Python or Go. Its primarily caused by JVM’s memory  management and class-loader.

Today, if we *have* to choose Java, we ensure that it’s version 11 or above. And our  Kubernetes memory limits are set to 1GB on top of JVM max heap memory (`-Xmx`) for headroom. That is, if JVM uses 8GB for heap memory, our Kubernetes  resources limits for the app would be 9GB. With that, life has been  better.

[Why Java Is DyingWhat does the future hold for Java?medium.com](https://medium.com/better-programming/why-java-is-dying-b02b5fd44db9)

# 2. Kubernetes Lifecycle Upgrades

Kubernetes lifecycle management such as upgrades or enhancements is cumbersome, especially if you’ve built your own cluster on [bare metal or VMs](https://platform9.com/blog/where-to-install-kubernetes-bare-metal-vs-vms-vs-cloud/). For upgrades, we’ve realized that the easiest way is to build a new  cluster with the latest version and transition workloads from old to  new. The effort and the planning that goes into in-place node upgrades  are just not worth it.

Kubernetes has multiple moving parts that need to align with an upgrade. From  Docker to CNI plugins like Calico or Flannel, you need to carefully  piece it all together for it to work. Although projects like Kubespray,  Kubeone, Kops, and Kubeaws make it easier, they all come with  shortcomings.

We built our clusters using Kubespray on RHEL VMs. Kubespray was great, it had playbooks for building, adding and removing new nodes, upgrading  version, and pretty much everything we needed for operating Kubernetes  in production. But however, the upgrade playbook came with a disclaimer  that prevented us from skipping even minor versions. So one would have  to go through all intermediate versions to reach the target version.

The takeaway is that if you plan to use Kubernetes or are already using  one, think about lifecycle activities and how your solution addresses  that. It’s relatively easier to build and run the cluster, but lifecycle maintenance is a whole new game with multiple moving parts.

# 3. Build and Deployment

Be prepared to redesign your entire build and deployment pipelines. Our  build process and deployment had to go through a complete transformation for the Kubernetes world. There was a lot of restructuring in not just  Jenkins pipelines, but using new tools like Helm, strategizing new git  flows and builds, tagging docker images, and versioning helm deployment  charts.

You would need a strategy to maintain not just code, but Kubernetes  deployment files, Docker files, Docker images, Helm charts, and design a way to link it all together.

After several iterations, we settled on the following design.

- Application code and its helm charts reside in separate git repositories. This allows us to version them separately. ([semantic versioning](https://semver.org/))
- We then save a map of chart version with the app version and use that for tracking a release. So for example, `app-1.2.0`deployed with `charts-1.1.0` . If only Helm values file were to change, then only the patch version of the chart would change. (e.g. from `1.1.0` to `1.1.1`). All these versions were dictated by release notes in each repository, `RELEASE.txt`.
- System apps like Apache Kafka or Redis whose code we did not build or modify,  worked differently. That is, we did not have two git repositories as  Docker tag was simply part of Helm chart’s versioning. If we ever  changed the docker tag for an upgrade, we would bump up the major  version in the chart’s tag.

# 4. Liveliness and Readiness Probes (the Double-Edged Sword)

Kubernetes’ liveliness and readiness probes are excellent features to combat system problems autonomously. They can restart containers on failures and  divert traffic away from unhealthy instances. But in certain failure  conditions, these probes can become a double-edged sword and affect your application’s startup and recovery, particularly, stateful applications like messaging platforms or databases.

Our Kafka system was a victim of this. We ran a `3 Broker 3 Zookeeper` stateful set with a `replicationFactor of 3 `and a `minInSyncReplica of 2.` The issue occurred when Kafka started up after accidental system  failures or crashes. This caused it to run additional scripts during  startup to fix corrupted indices, which took anywhere between 10 to 30  mins depending on the severity. Because of this added time, the  liveliness probes would constantly fail, issuing a kill signal to Kafka  to restart. This prevented Kafka from ever fixing up the indices and  starting up altogether.

The only solution is to configure `**initialDelaySeconds**` in liveliness probe settings to delay probe evaluations after the  container startup. But the problem, of course, is that its hard to put a number to this. Some recoveries even take even an hour, and we need to  provide enough headroom to account for this. But the more you increase `**initialDelaySeconds**` , the slower your resilience, as it would take longer for Kubernetes to restart your container during startup failures.

So the middle ground is to assess a value for the `**initialDelaySeconds**` field such that it better balances between the resilience you seek in  Kubernetes and the time taken by the app to successfully start in all  fault conditions (disk failures, network failures, system crashes, etc.)

> ***Update\****: If you are on the last few latest releases,* [*Kubernetes has introduced a third probe-type called, ‘Startup Probe,’ to tackle this problem*](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)*. It is available in* `*alpha from 1.16* `*and* `*beta from 1.18*` *onwards.*
>
> *A startup probe disables readiness and liveliness checks until the  container has started up, making sure the application’s startup isn’t  interrupted.*

# 5. Exposing External IPs

We learned that exposing services using static external IP takes a huge  toll on your kernel’s connection tracking mechanism. It simply breaks  down at scale unless planned thoroughly.

Our cluster runs on `Calico for CNI `and `BGP` as our routing protocol inside Kubernetes and also to peer with edge routers. For Kubeproxy, we use `IP Tables`mode. We host a massive service in our Kubernetes exposed via external IP  that handles millions of connections every day. Because of all the SNAT  and masquerading that comes from software-defined networks, Kubernetes  needs a mechanism to track all these logical flows. To achieve this, it  uses the Kernel’s `Conntrack and netfilter` tools to manage these external connections to the static IP, which then translates to internal service IP and then to your pod IP. This is all  done through the`conntrack` table and IP Tables.

This `conntrack` table, however, has its limits. Once you hit the limit, your Kubernetes cluster (the OS Kernel underneath) will not be able to accept new  connections any more. On RHEL, you can check it this way.

```
$  sysctl net.netfilter.nf_conntrack_count net.netfilter.nf_conntrack_maxnet.netfilter.nf_conntrack_count = 167012
net.netfilter.nf_conntrack_max = 262144
```

Some ways to work around this is to peer multiple nodes with edge routers so that the incoming connection to your static IP is sprayed across your  cluster. So if your cluster has a large fleet of machines, cumulatively  you could have a large `conntrack` table to handle massive incoming connections.

Back when we started in 2017, this completely threw us off, but recently, a  detailed study on this was published by Calico in 2019, aptly titled “[Why conntrack is no longer your friend](https://www.projectcalico.org/when-linux-conntrack-is-no-longer-your-friend/).”

[4 Simple Kubernetes Terminal Customizations to Boost Your ProductivityThis is what I use for managing large-scale Kubernetes clusters in productionmedium.com](https://medium.com/better-programming/4-simple-kubernetes-terminal-customizations-to-boost-your-productivity-deda60a19924)

# Do you absolutely need Kubernetes?

Three years in and we still continue to discover and learn something new  every day. It is a complex platform with its own set of challenges,  particularly, the overhead in building and maintaining the environment.  It will change your design, thinking, architecture, and will require  upskilling and upscaling your teams to meet the transformation.

However, if you are on the cloud and are able to use Kubernetes as a “service,”  it can relieve you from most of that overhead that comes with platform  maintenance like “How do I expand my internal network CIDR?” or “How do I upgrade my Kubernetes version?”

Today, we’ve realized that the first question you need to ask yourself is “Do you *absolutely* need Kubernetes?” This can help assess the problem you have and how significantly or not, Kubernetes addresses it.

Kubernetes transformation is not cheap. The price you pay for it must really  justify ‘your’ use case and how it leverages the platform. If it does,  then Kubernetes can immensely boost your productivity.

> Remember, technology for the sake of technology is meaningless.

[Better Programming](https://betterprogramming.pub/?source=post_sidebar--------------------------post_sidebar-----------)

Advice for programmers.

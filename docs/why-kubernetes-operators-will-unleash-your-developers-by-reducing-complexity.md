# Why Kubernetes Operators Will Unleash Your Developers by Reducing Complexity

#### 17 Aug 2020 8:56am,   by [Rob Szumski](https://thenewstack.io/author/rob-szumski/ "Posts by Rob Szumski")

Rob is a Product Manager for OpenShift at Red Hat](https://www.linkedin.com/in/robszumski/)

Kubernetes and the many projects under the Cloud Native Computing Foundation (CNCF) umbrella are advancing so quickly that it can sometimes feel as though the benefits of the open hybrid cloud are, perhaps, not yet evenly distributed. While systems administrators and operators are having a blast modernizing legacy applications and automating highly-scalable, container-based systems, sometimes developers can feel a bit left out in the cold.

That’s because there’s a disconnect between the velocity increase that IT is experiencing through the adoption of containers, and the agility increase offered to developers. The benefits that Kubernetes offers to IT operators are in service of simplifying the lives of their developers. Developers and IT administrators want different things, right? Or do they?

Back in 2016 when I was at CoreOS, we saw this disconnect beginning to form. We worried that while enabling developers was job one for IT, it wasn’t necessarily easy when developers had to relearn their entire stacks from the ground up in order to adopt containers. It was also tough for IT administrators to provision compliant, governed services at the pace expected in cloud development.

To fix this, CoreOS introduced the concept of Kubernetes Operators. Soon after, we produced a set of tools known as the Operator Framework to help users build, ship and discover their own Kubernetes Operators. While a lot has changed for myself and CoreOS since 2016, most notably the fact that we are now both a part of the Red Hat family, Operators remain an important part of our strategy for helping developers and administrators focus on using Kubernetes — rather than tweaking its various knobs and configurations.

**Sponsor Note**

![sponsor logo](https://cdn.thenewstack.io/media/2016/05/20f61f0a-red-hat-openshift@2x.png)

OpenShift is Red Hat’s container application platform that allows developers to quickly develop, host and scale applications in a cloud environment. With OpenShift, you have a choice of offerings, including online, on-premise and hosted service offerings.

Naturally, all of the work on Operators was done as open source, and as a result we recently contributed the Operators Framework to the CNCF, to become an incubated project. We did this for a few reasons, including that it was the right thing to do in this open ecosystem. But we also did it so that Operators will become a ubiquitous path to distributing servers into Kubernetes clusters. With Operators available to every Kubernetes distribution, the entire open hybrid cloud ecosystem will benefit from greater compatibility, simplicity and easier management.

Operators are a big part of helping us reach this automated, on-demand, container-based future. Operators are operational procedures and best practices that are codified into software. They make automated day two operations possible on Kubernetes, and model the complexities of today’s distributed systems.

For example, there isn’t a concept of data rebalancing in Kubernetes, but that can be built on top of the Kubernetes APIs with the Operator Framework. These types of applications are required for running what we call the third wave of applications on Kubernetes: complex distributed systems.

[![](https://cdn.thenewstack.io/media/2020/08/ce7b9fcc-image3.png)](https://cdn.thenewstack.io/media/2020/08/ce7b9fcc-image3.png)

The Operator Framework comes with all the tools required for developers to build this software, and everything that cluster administrators need to safely install, upgrade and manage their Operators. These tools tie into other CNCF projects — like KubeBuilder, Helm, kuttl — and popular open source software like Ansible.

The framework is loosely coupled, so if you have a favorite testing tool, you can keep using that; and if you want to build your Operator outside of the SDK, that’s fine too, you can still run it with the Operator Lifecycle Manager.

Operator Framework has three flavors of SDK today, with more to come in the future. Each one addresses a different type of author, from IT administrator to traditional developer to hardcore Kubernetes experts.

[![](https://cdn.thenewstack.io/media/2020/08/5484c998-image1.png)](https://cdn.thenewstack.io/media/2020/08/5484c998-image1.png)

## Help for Building

Testing your Operator is critical, and that’s why I advise most Operator authors to utilize one of our SDKs, or at least join our [community](https://github.com/operator-framework/community-operators). Our experts in the community have modeled many types of applications as Operators, so we can help you save time and avoid some bugs. Plus, it’s been really fun to see all the projects coming into our community and to help people meet their goals.

Once you’ve written your Operator, you need to hand it off to your users and actually run it. This is where the Operator Lifecycle Manager comes into play. There are actually a lot of tricky problems here that you might not even know you have yet, like collision detection for CRDs. Imagine you have a database managed by a Custom Resource Definition (CRD), but there’s another Operator that also wants to manage that database. That’s no good. The lifecycle of a CRD itself is also important. The Operator can manage the CRD as part of the upgrade process.

We’re deeply committed to Operators, and we have a lot planned for the future of these powerful abstractions. For example, we’re currently working on the concept of Bundles, which will allow Operators to be cataloged together inside clusters. Administrators will be able to use a new tool we’re working on called OPM, which would allow them to better curate those in-cluster catalogs. We’re also working on a new Operator API, which is designed to provide an easier way to access Operators through Kubernetes APIs.

[![](https://cdn.thenewstack.io/media/2020/08/39abe664-image2.png)](https://cdn.thenewstack.io/media/2020/08/39abe664-image2.png)

While much of the workaround setting up an Operator inside a cluster will be done by the IT staff, it is the developer who really benefits. For example, IT can use the [Crunchy Data PostgreSQL Operator](https://operatorhub.io/operator/postgresql) to pre-configure database RBAC and security controls, while also pre-wiring backup and replication services. After this is done, any blessed user of the cluster can then deploy a fresh instance of PostgreSQL on-demand, eliminating the time required to provision servers for new application development work.

This benefit also extends to automated on-demand provisioning of test and build environments. These environments can be pre-configured to adhere to corporate governance and data control policies through Operators properly configured by IT administrators. The best part is that they only have to do this once, and the deployment according to these rules is automated across the cluster.

Operators have been available on OpenShift since version four was released last summer. We’ve been working hard with many developers across the enterprise and open source landscapes to help flesh out [Operatorhub.io](https://operatorhub.io/), our public repository for Kubernetes Operators. So the Operator ecosystem is already quite vibrant and ready for exploration.

We’re excited to see how much larger this community can grow, now that it is fully a part of the CNCF’s open source processes. This was always what we intended to do with the Operator Framework, and now after four years of hard work, we’re very proud to see the project take the next step and run free as an open source project inside the CNCF. With your help, Operators will save everyone in the Kubernetes community — both administers and developers — a lot of time and worry.

The Cloud Native Computing Foundation is a sponsor of The New Stack.

Feature image via [Pixabay](https://pixabay.com/illustrations/complex-fractal-chaos-grid-clock-664440/).

At this time, The New Stack does not allow comments directly on this website. We invite all readers who wish to discuss a story to visit us on [Twitter](https://twitter.com/thenewstack) or [Facebook](https://www.facebook.com/thenewstack/). We also welcome your news tips and feedback via email: [feedback@thenewstack.io](mailto:feedback@thenewstack.io).

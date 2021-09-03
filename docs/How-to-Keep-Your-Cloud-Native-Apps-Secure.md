# How to Keep Your Cloud-Native Apps Secure

June 28, 2021

An increasing number of organizations are adopting cloud-native apps, and new research from [Snyk](https://snyk.io/state-of-cloud-native-application-security/) reveals that along with that adoption, 60 percent of organizations have increased security concerns.

We got together with Jörg Müller, principal consultant at INNOQ, to hear his tips on how to help with the security and operational challenges companies often face when adopting cloud-native apps.

## **How do you secure the network traffic?**

**You should encrypt everything in your environment, but you need to choose on what level. If you have encryption on the highest level, you might not need additional encryption lower down that would just cost you performance.**

In the past you had something that was called security-on-perimeter, which meant that you had a firewall around your internal network and everything inside your network was trusted. Services could communicate to each other without encryption and without authorization.

This was never a very good practice, but it was a very common one, and for most people it was enough. Today it’s not, and there are two main reasons for that:

1\. It’s not that complicated to have encrypted authorized network traffic inside your own network these days, and there’s no performance overhead you have to take care of

2\. Zero trust networks are a far better security practice, which is used today

### There’s no reason for you to take an on-perimeter security approach

It’s interesting to look at where you encrypt authorized network traffic inside your own network, and on what level. If you’re looking at a modern deployment of an application, then the network stack alone has several layers:

**Level 1: The physical network**

The one where the cables are running and where the physical network card is found.  Some network cards already have the capabilities to run encryption on this level.

**Level 2: Layer of virtualization**

Above the physical network you usually have a layer of virtualization, which means that the virtual machines have their own IP addresses that don’t match the addresses of any network cards, and the communication between the virtual machines might already be encrypted and authorized.

**Level 3: Overlay network**

If you’re using Kubernetes, you’ll have an overlay network for Kubernetes which is doing the same again. If you’re using Cilium or something similar, then you’re doing encryption and authorization there again.

**Level 4: Service mesh**

You get a level higher if you’re using service meshes — you have an encryption on the service level, running encryption between services.

Finding the right combination of that stack in your company is complicated.

It might be enough if you’re using an encrypted overlay network on a Kubernetes layer, or it might be necessary to have something like a service mesh to encrypt between services and authorize on that level because your applications need to know what your other services are communicating with me and need to authorize them, which is something you can’t do on the lower levels because you’re missing the knowledge there.

## **Container isolation**

**Where do you get the containers? And are they secure? What privileges do they need? What are they containing, and are they containing any vulnerabilities?**

If you’re talking about cloud native, let’s presume you’re using containers, since it’s usually the way people deploy today.

From a security standpoint, a container does not run in a complete virtualization but usually as an isolated process. It’s a process on your operating system, on your kernel, that is isolated from the other processes, isolated from the network level, isolated on the process level itself (so it can’t see the other ones), and it’s not allowed to do anything.

But sometimes it needs to do some things, and sometimes there are some privileges that slip through because they can call the kernel, and if you don’t configure the container in the right way, then you have some issues that can break out of your container.

There are mechanisms in Kubernetes like pod security policies where Kubernetes controls what features or privileges a container tries to get and can prevent a container from just starting up.

This is something you need to take care of, otherwise you run the risk of running a rogue container from taking over something in your network, having some privileges that you don’t want it to have.

That’s even worse if you’re using containers that are not inside your organization, so that’s another level of security you have to take care of: where are the containers coming from and are they secure?

### VM for every container

If you really want really strong container isolation, you can use a container runtime that builds a virtual machine for every container.

There are container runtimes that create new virtual machines, or that are very strict about the privileges, adding another layer of security to a Kubernetes cluster, which helps you a lot in keeping things isolated. This is a very high-level approach.

### Scanning containers

There’s a lot of stuff you can do around automatically scanning a container before you even deploy it. That’s a big advantage compared to the old way of setting up a virtual machine and maintaining it where you would need to run your own updates. Today, you can just build the container and run it through a scanner.

Docker has some stuff built in, but there are plenty of scanners out there which will run your container against the CVE database and tell you if the container you’re using contains any software with known vulnerabilities. Once scanned you can then integrate that into your CD process, and you’ll have a much more secure container rollout and content of your containers.

This is completely different to how it was done in the past where the main thing was to ensure your virtual machine and operating system were both up to date. Today it’s a completely different approach of building a container, scanning a container and rolling it out.

### Service mesh

Service meshes usually do a thing which is called MTLS (mutual TLS). TLS is something you know from your normal web browser, if you’re using any website which has https then your browser can validate that the certificate for a website you’re accessing really is the certificate for this website through certificate chains.

Mutual TLS does that in both directions. So the server will validate if you are who you say you are, which is usually not the case in normal web environments, but very useful if you’re doing this in a serveless or Kubernetes environment because then the service that’s providing you the service knows it’s been called by someone who’s allowed to call it and not someone else. This is a very high level of security and was unheard of before the days of cloud native.

## **Speed of updating all infrastructure components**

**One of the operational challenges is having to update all your components on a continuous basis.**

Kubernetes releases a new version every three months, Linkerd as a service mesh releases around every three months, and many other tools and products are doing a lot of high-frequency releases.

In the past you usually had companies or system admins that tried to keep a very stable environment with a long-term supported OS version like RedHat, which was supported for eight years with regular updates. You were not updating your OS very often, and that was very normal.

Today, it’s completely different. You have to update your Kubernetes cluster every three to six months. If you skip too many versions it will be very difficult for you to update them again.

This stacked release cycle also leads to APIs deprecating very quickly — within just half a year to a year of a managed deprecation process for an API, the stuff you would have been using would now be gone.

The update cycle these days runs very fast, but this can be seen as a good thing from a security perspective as you’re forced to have the latest versions of software, you can’t just go by with software that is four to five years old — a Kubernetes environment wouldn’t even run.

### Use the stuff that everybody else uses

You can manage a lot of that by not doing it yourself but letting your cloud provider do it. Use managed Kubernetes instances from Google, AWS, etc. You need a very good reason not to run a managed Kubernetes cluster.

There might be reasons to do that in your own data center. Some reasons might be costs, but that should be one of the last reasons. Security or GDPR rules could be another reason, because most of the cloud providers are American in Europe.

In general, using the managed offerings is the way to go. It wasn’t the case two or three years ago — many customers still set up Kubernetes clusters of their own, they just used it and went on to updating things on their own.

## **Operational Issues**

### Dev team responsibilities

More often than not, security is often not a very strong focus point in development teams. Bringing more management responsibilities into the development teams helps alleviate this problem.

The conflict between stability from the operations team and the new features and progress from the development team itself is something that we try to solve in DevOps, but I think you can only solve that by working together and by having most of the responsibility with the development team itself, even for security.

This is something which should be done and requires a lot of training from the developers side. Not every developer wants that but it makes sense, and I know that there are a lot of developers out there that agree.

### Clusters

How large is your cluster? How are you approaching Kubernetes clusters?

There are two extremes of ways to approach clusters:

1. Having one cluster per application, or even per server

2. Having a large cluster which includes everything and tries to isolate stuff with namespaces

You have to find a balance in your operational department, somewhere in between those two extremes. How many applications, how many teams are you? Are you accessing a single cluster? How many clusters do you have? Will you try to consolidate everything into one or many?

There are a lot of advantages and disadvantages on each side.

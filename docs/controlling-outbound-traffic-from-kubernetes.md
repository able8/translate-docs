# Controlling outbound traffic from Kubernetes

[Read the article](http://monzo.com#article)

At Monzo, the Security Team's highest priority is to keep your money and data safe. And to achieve this, we're always adding and refining security controls across our banking platform.

Late last year, we wrapped up a major networking project which let us control internal traffic in our platform (read about it [here](https://monzo.com/blog/we-built-network-isolation-for-1-500-services)). This gave us a lot of confidence that malicious code or an intruder compromising an individual microservice wouldn't be able to hurt our customers.

Since then, we've been thinking about how we can add similar security to network traffic _leaving_ our platform. A lot of attacks begin with a compromised platform component 'phoning home' — that is, communicating with a computer outside of Monzo that is controlled by the attacker. Once this communication is established, the attacker can control the compromised service and attempt to infiltrate deeper into our platform. We knew that if we could block that communication, we'd stand a better chance at stopping an attack in its tracks.

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/4SIwnE4zIp90oqeLFVRJLj/26b0213a3b8b7b83a3d49390757a6835/Untitled__2_.png?w=1280&q=90)

In this blog, we'll cover the journey to build our own solution for controlling traffic leaving our platform ( [you can read more about our platform here](https://monzo.com/blog/2016/09/19/building-a-modern-bank-backend)). This project formed part of our wider effort to move towards a platform with less trust, where applications need to be granted specific permissions instead of being allowed to do whatever they like. Our project contributed to this by allowing us to say that a specific service can communicate to an external provider like GitHub, but others can't.

Check out [other posts](https://monzo.com/blog/authors/jack-kleeman/) from Security Team for more projects we've completed to get closer to zero trust.

## We started by identifying where traffic leaves the platform

With a few exceptions, we previously didn't have filtering on what outbound traffic could send. This means that we started in the dark about which services needed to talk to the internet and which didn't. This is similar to where we started our previous network isolation project, when we didn't to know what services talk to what other services.

But unlike when we started the network isolation project, tools and processes we built for that project are now at our disposal: a combination of [calico-accountant](https://github.com/monzo/calico-accountant) and packet logging allowed us to rapidly identify which of our microservices actually talk to the internet, and what external IPs each talks to.

Even with IP information, it wasn't always trivial to find out what domains each service talked to. This is because many external providers our services talk to use CDNs like AWS CloudFront, where a single IP serves many websites. For simpler cases, we'd simply read the code for the service to find out what hostnames are used. If this isn't possible, we'd log the outgoing packets of the service carefully to identify their destinations.

Once we put together a (surprisingly small) spreadsheet on what services used what external hostnames and ports, we began a design process to think about what an ideal long term solution would look like.

We had a few requirements:

- **We wanted to be able to write policies for outgoing traffic from specific source services to specific destination hostnames.** We have thousands of services running in our platform, each serving a small but distinctive purpose; and we cannot just have one universal list of allowed destinations, such as saying 'all services can talk to GitHub'. Some destinations on the internet used by our services would allow a certain degree of attacker control — anyone can host something on GitHub, for example. Therefore we want to limit 'dangerous' destinations to only the services that need them.

- **It must be possible to allow a specific DNS name.** Allowing traffic to IPs is super easy already with Calico — our Kubernetes networking stack — but over time IPs change for most DNS names. For a long-term solution, we need to be able to allow traffic to a domain name like `google.com`, and at any point in time thereafter, traffic to resolved IPs of `google.com`, should just work.

- **We wanted to be able to reliably alert when we detect packets whose sources or destinations aren't allowed,** like we can for our internal network controls.


![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/6Oe5lLcMnqoUBZ3FQB3ihE/d1e3247c2695c9968a1cabb63bb593d1/Untitled__3_.png?w=1280&q=90)

We realised almost no drop-in solution on the market could check all the boxes, and it'd take a fairly long process to implement our ideal solution. So we decided to ship a simple solution first, and iterate on it.

## We started with port based filtering

Without building anything new, we realised that with what's already running in our platform, we could implement simple rules on destination ports without support for hostnames. That's to say, we can make statements like this:

_"Allow service.github to make outbound connections to the public internet on port 443"_

The service will be able to reach any domain on that port, which isn't a very tight control. But more importantly, services that shouldn't be making any outbound connections at all won't be allowed to. The vast majority of our services don't need to talk to the public internet, so we actually get a ton of benefit from this simple approach. We also knew that a few services would always need to be allowed to reach a very large number of domains on certain ports, so it was useful to allow that use case early on.

The way we implemented this was through a Kubernetes NetworkPolicy for each port that we needed to allow. If a pod of a service is labelled with `external-egress.monzo.com/443: true`, then it will match a policy which allows traffic to all public IPs on port 443. Any traffic not allowed by one of these policies is logged (a feature of Calico), then dropped.

![apiVersion: networking.k8s.io/v1 kind: NetworkPolicy metadata:   name: egress-external-tcp-443 spec:   egress: 	- to: 		- ipBlock: 				# allow the whole internet 				cidr: 0.0.0.0/0 				except: # private IP addresses         - 0.0.0.0/8         - 10.0.0.0/8         - ... 		ports: 			- port: 443   podSelector:  		matchLabels: 			external-egress.monzo.com/443: "true"   policyTypes:   - Egress](https://images.ctfassets.net/ro61k101ee59/FbZ4BTweom8d0ptxRbN5M/af6ef10de11bb4dc7bab4de25e78ea11/Screenshot_2020-04-06_at_12.16.34.png?w=1280&q=90)

To label pods, we used a similar approach to the one we used for internal network controls: 'rule files' which are part of a service's code. To start with, we added files like `service.github/manifests/egress/external/443.rule`, and updated our deployment pipeline to automatically convert these files into the right labels.

## We investigated existing solutions

Having implemented a simple solution supporting only ports, we then set off looking for a full solution supporting hostname filtering. There are two approaches that we saw in the open source and enterprise software communities:

### SNI and host header inspection

One approach, used by [Istio](https://istio.io/), is to run an egress proxy inside Kubernetes. This egress proxy would inspect the traffic sent to it, figure out the destination, determine whether the destination was allowed (given the source) and then pass it on. To determine destinations for encrypted traffic, it has to use the SNI ( [Server Name Indication](https://en.wikipedia.org/wiki/Server_Name_Indication)) field for encrypted (TLS) traffic. For unencrypted web traffic — which is becoming increasingly rare — it can simply inspect the HTTP Host header.

Because Istio is simply configuring Envoy to achieve this functionality, and Monzo already had a fairly advanced Envoy service mesh, we were confident we could build something very similar. And because this solution would run inside Kubernetes, we could enforce policies per service based on their source pod IPs, which is more challenging to do outside Kubernetes, as we run a virtual network for pods. We could also set up alerting for rogue packets fairly easily, as we'd control the application which determines what's allowed.

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/4G7BCJKxdfei8vM8dDatKm/9cac731554e252e39f242a2ff43c2d07/Untitled__4_.png?w=1280&q=90)

But this proxy would only work for TLS or HTTP traffic. Monzo's platform runs many services, with lots of different protocols to talk to external providers. In particular, we deal with quite a lot of outbound SFTP traffic, which wouldn't be supported by this approach. We weren't sure if it would've been possible to modify Envoy to support determining SSH destinations in a reliable way.

Some enterprise solutions take a similar approach to Istio. Instead of a proxy, they'll set themselves up as the NAT gateway for all outbound traffic in your AWS cluster, then inspect traffic to determine what's allowed. Again, we didn't expect SSH to be supported.

### DNS inspection

A more common approach to domain name based filtering is based on the observation that, for an application to reach an IP for `github.com`, it must have (at some point) resolved `github.com` via DNS (DNS is how computers turn domain names into IP addresses). Well behaved applications will do this regularly, so that their IPs are never older than the expiry of the DNS record. As a result, by watching DNS traffic, we can figure out key information: firstly, that a given application wants to reach `github.com`, and secondly that it will try to do so on a specific set of IPs.

Many implementations of egress filtering use this principle to allow (for a limited period of time) traffic to the returned IPs, but only for the source IP that requested them. We can choose to only allow this for a certain set of domain names, effectively allowing outgoing traffic to these domains only. Our firewall will then have a set of dynamically changing rules. Critically, our firewall must update itself to accept traffic to the new IPs _before_ these IPs are returned to the application — otherwise subsequent requests from the application could fail.

We found quite a few enterprise solutions that use this approach and would sit as a box inside our platform. Generally, the box would carry all outbound traffic, and in some cases it would also act as a DNS resolver, or otherwise outbound DNS traffic must all pass through it. It would then be able to adjust firewall rules for newly resolved IPs as they are returned in DNS responses.

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/4qahu4EL7RPvOQujSwXeqt/04127ba6942ebdb17c2e5a8dbe650459/Untitled__5_.png?w=1280&q=90)

The main issue was that the enterprise solutions couldn't run inside Kubernetes. So they wouldn't be able to let specific services access specific destinations, as the source IP would just look like the IP of a cloud instance running Kubernetes, since our cluster runs a virtual network for pods. They will however allow us to control egress from our non-Kubernetes workloads in the same way; but in most cases we can control that traffic just as easily with IP-based rules on AWS.

### **Building our own egress firewall**

We considered how complicated it'd be to build a similar approach. We'd need to create an application inside Kubernetes that we routed all outbound traffic through, as well as DNS traffic. Our main concern was availability - to make this application reliable, we'd want to run many instances of it, with outgoing traffic balanced evenly between instances.

But this creates a state consistency problem where when one instance discovers a new IP that needs to be allowed, it'd need to very rapidly communicate this to other instances - ideally, before the DNS query completes. This is possible, but adds a lot of complexity to the system.

Another approach, taken by the enterprise solutions, would be to use a single 'leader' which handles all egress traffic, but one which can fail over to a set of 'followers', ideally waiting for state to transfer in the process. But since we can't redirect routed traffic in our cloud environment instantly, a short but significant delay would be inevitable before the gateway's available again, leading to potential failures in applications that need to reach external services.

## We came up with a hybrid approach

We saw the merits of both approaches:

- Egress proxies which carry outgoing traffic are easy to run in a highly available way, but inspection doesn't work universally

- DNS inspection is a fairly universal way to figure out where traffic is going, but it's tricky to build a reliable distributed system to take advantage of that


We had an idea: **what if we just have a proxy for** **_each_** **domain we needed to reach?** This might sound complex and wasteful, but we can easily configure and run these proxies inside our Kubernetes cluster at scale. We call these proxies _egress gateways_.

For example, there's a gateway for `github.com`, with a specific load balancing internal IP address. This gateway accepts traffic and always sends it to `github.com`, doing its own DNS resolution. The gateway just operates on TCP (it is a [Layer 4](https://en.wikipedia.org/wiki/Transport_layer) proxy), so Git over SSH is supported as well as almost everything else.

To actually get the traffic flowing through these gateways, we'd need to "hijack" DNS responses for domains they proxy to: when a microservice asks for `github.com`, we need to respond with the correct gateway IP address, instead of the real IP addresses of `github.com`. This'll cause this service to use our gateway to talk to `github.com` automatically, without any changes to the client.

So once all services start using gateways exclusively for their outgoing traffic, we can block all traffic to the public internet, except from the gateways. We consider it safer to have many uniform TCP proxies able to reach the public internet than a diverse set of microservices. And more importantly, we can control which applications can talk to the gateway, using our existing internal network isolation controls, taking advantage of the same tools, the same alerting, and the same understanding we already have.

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/2Ovajc6p7Cpn9Bc41uZtkI/51f5a73a7475f55d5e6ce1aeac2cdb99/Untitled__6_.png?w=1280&q=90)

### Implementing egress gateways in Kubernetes

We manage each egress gateway in our cluster with some Kubernetes building blocks:

1. **A** [**Deployment**](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) **of** [**Envoy**](https://www.envoyproxy.io/) **pods.** We usually run three pods for high availability. The envoy pods use a local file to get their configuration.

2. **A** [**ConfigMap**](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) **to configure Envoy.** The config determines what ports the gateway listens on and what DNS name traffic is destined for (e.g. `github.com`).

3. **A** [**HorizontalPodAutoscaler**](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) **to increase the number of pods dynamically based on CPU usage.** TCP proxying is a good use case for horizontal scaling — lots of pods running in parallel usually handle concurrent connection loads as well as a few powerful pods.

4. **A** [**Service**](https://kubernetes.io/docs/concepts/services-networking/service/) **which gives the pods a single internal IP address.** This means that we can set up DNS to resolve `github.com` to a single internal IP, but each new connection is load balanced to different gateway pods.

5. **A** [**NetworkPolicy**](https://kubernetes.io/docs/concepts/services-networking/network-policies/) **which controls access to the gateways.** If the gateway is named `github`, then access to it is restricted to pods which are labelled with `egress.monzo.com/allowed-github: true`.


To set this all up, we just needed to set a few configuration files for each gateway. But we actually went one step further and automated the process of managing these objects with an [Kubernetes operator](https://coreos.com/blog/introducing-operators.html). This operator accepts a very small set of information (in this case the domain names that we need to allow), and creates the above resources. If we change how we want to set up those resources, all gateway configurations will be updated automatically. This lets us manage hundreds of gateways with minimal manual work.

We also needed to build a [CoreDNS](https://coredns.io/) plugin which allowed us to hijack DNS traffic for `github.com` and send it to the appropriate gateway. This turned out to be fairly easy — we just needed to [watch](https://kubernetes.io/docs/reference/using-api/api-concepts/#efficient-detection-of-changes) Kubernetes Service objects to figure out which gateways exist and for what domains, and then wrap the existing CoreDNS `rewrite` plugin to rewrite queries for `github.com` into a query for `github.egress-operator-system.svc.cluster.local`, which is the DNS name for the GitHub egress gateway.

Both the operator and its CoreDNS plugin we built can be found open-source [here](https://github.com/monzo/egress-operator).

From the perspective of the end user (an engineer building microservices in our cluster), a service can contain a file `service.github/manifests/egress/external/github.com:443.rule`, and our deployment pipeline will convert this into the appropriate pod label, by determining which egress gateway allows traffic to that destination.

From the perspective of the Security team maintaining the egress gateways, we manage a file like this for each destination:

![apiVersion: egress.monzo.com/v1 kind: ExternalService metadata:   name: github spec:   dnsName: github.com   hijackDns: true   ports:   - port: 443](https://images.ctfassets.net/ro61k101ee59/2yaNN8oWWm5ZQVQvRl2X5T/fb56160ae22edeb589a6467c48036107/Screenshot_2020-04-06_at_12.16.55.png?w=1280&q=90)

We've been enforcing our new firewall for some time now and we're really pleased with the outcome. Egress gateways provide us with:

- Very granular control of outgoing traffic to each external domain used by each service

- Great inspectability via Envoy as well as our existing network isolation tools

- The ability to use a single firewall (Kubernetes network policies) for both internal and external traffic.


We can definitely see us keeping this approach as a long-term solution.
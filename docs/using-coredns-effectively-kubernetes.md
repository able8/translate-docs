# How to use CoreDNS Effectively with Kubernetes

May 5th, 2021 From: https://www.infracloud.io/blogs/using-coredns-effectively-kubernetes/

### Background story

We were increasing HTTP requests for one of our applications, hosted on the Kubernetes cluster, which resulted in a spike of 5xx errors.
The application is a GraphQL server calling a lot of external APIs and then returning an aggregated response.
Our initial response was to increase the number of replicas for the application to see if it improves the performance and reduces errors.
As we drilled down further with the application developer, found most of the failures were related to DNS resolution.
That’s where we started drilling down DNS resolution in Kubernetes.

This post highlights our learnings related to using CoreDNS effectively with Kubernetes, as we did deep dive in process of configuring and troubleshooting.

### CoreDNS Metrics

DNS server stores record in its database and answers domain name query using the database.
If the DNS server doesn’t have this data, it tries to find for a solution from other DNS servers.

CoreDNS became the [default DNS service](https://kubernetes.io/blog/2018/12/03/kubernetes-1-13-release-announcement/#coredns-is-now-the-default-dns-server-for-kubernetes) for Kubernetes 1.13+ onwards.
Nowadays, when you are using a managed Kubernetes cluster or you are self-managing a cluster for your application workloads, you often focus on tweaking your application but not much on the services provided by Kubernetes or how you are leveraging them.
DNS resolution is the basic requirement of any application so you need to ensure it’s working properly.
We would suggest looking at [dns-debugging-resolution](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/) troubleshooting guide and ensure your CoreDNS is configured and running properly.

By default, when you provision a cluster you should always have a dashboard to observe for key CoreDNS metrics.
For getting CoreDNS metrics you should have [Prometheus plugin](https://coredns.io/plugins/metrics/) enabled as part of the CoreDNS config.

Below sample config using `prometheus` plugin to enable metrics collection from CoreDNS instance.

```
.:53 {
    errors
    health
    kubernetes cluster.local in-addr.arpa ip6.arpa {
      pods verified
      fallthrough in-addr.arpa ip6.arpa
    }
    prometheus :9153
    forward . /etc/resolv.conf
    cache 30
    loop
    reload
    loadbalance
}

```

Following are the key coreDNS metrics, we would suggest to have in your dashboard:
If you are using Prometheus, DataDog, Kibana etc, you may find ready to use dashboard template from community/provider.

- **Cache Hit percentage:** Percentage of requests responded using CoreDNS cache
- **DNS requests latency**  - CoreDNS: Time taken by CoreDNS to process DNS request
  - Upstream server: Time taken to process DNS request forwarded to upstream
- **Number of requests forwarded to upstream servers**
- **[Error codes](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6) for requests**  - NXDomain: Non-Existent Domain
  - FormErr: Format Error in DNS request
  - ServFail: Server Failure
  - NoError: No Error, successfully processed request
- **CoreDNS resource usage:** Different resources consumed by server such as memory, CPU etc.

We were using DataDog for specific application monitoring. Following is just a sample dashboard I built with DataDog for my analysis.

![datadog-coredns-dashboard](https://d33wubrfki0l68.cloudfront.net/26c0fc0512f17a51d420ed77746900ff5d0c278e/b52b1/assets/img/blog/using-coredns-effectively-with-kubernetes/dd-dashboard.png)

### How to reduce CoreDNS errors?

As we started drilling down more into how the application is making requests to CoreDNS, we observed most of the outbound requests happening through the application to an external API server.

This is typically how resolv.conf looks in the application deployment pod.

```
nameserver 10.100.0.10
search kube-namespace.svc.cluster.local svc.cluster.local cluster.local us-west-2.compute.internal
options ndots:5

```

If you understand how Kubernetes tries to resolve an FQDN, it tries to DNS lookup at different levels.

Considering the above DNS config, when the DNS resolver sends a query to the CoreDNS server, it tries to search the domain considering the search path.

If we are looking for a domain boktube.io, it would make the following queries and it receives a successful response in the last query.

```
botkube.io.kube-namespace.svc.cluster.local  <= NXDomain
botkube.io.svc.cluster.local <= NXDomain
boktube.io.cluster.local <= NXDomain
botkube.io.us-west-2.compute.internal <= NXDomain
botkube.io <= NoERROR

```

As we were making too many external lookups, we were getting a lot of NXDomain responses for DNS searches.
To optimize this we customized `spec.template.spec.dnsConfig` in the Deployment object.
This is how change looks like:

```
     <span class="na">dnsPolicy</span><span class="pi">:</span> <span class="s">ClusterFirst</span>
     <span class="na">dnsConfig</span><span class="pi">:</span>
       <span class="na">options</span><span class="pi">:</span>
       <span class="pi">-</span> <span class="na">name</span><span class="pi">:</span> <span class="s">ndots</span>
         <span class="na">value</span><span class="pi">:</span> <span class="s2">"</span><span class="s">1"</span>

```

With the above change, resolve.conf on pods changed.
The search was being performed only for an external domain.
This reduced number of queries to DNS servers.
This also helped in reducing 5xx errors for an application.
You can notice the difference in the NXDomain response count in the following graph.

![coredns-rcode-nxdomain-reduction](https://d33wubrfki0l68.cloudfront.net/59e1f460286518c33b5841818c0eb7dae16992cc/15811/assets/img/blog/using-coredns-effectively-with-kubernetes/coredns-rcode-nxdomain.png)

A better solution for this problem is [Node Level Cache](https://kubernetes.io/docs/tasks/administer-cluster/nodelocaldns/) which is introduced Kubernetes 1.18+.

### How to customize CoreDNS to your needs?

We can customize CoreDNS by using plugins.
Kubernetes supports a different kind of workload and the standard CoreDNS config may not fit all your needs.
CoreDNS has a couple of in-tree and external plugins.
Depending on the kind of workloads you are running on your cluster, let’s say applications intercommunicating with each other or standalone apps that are interacting outside your Kubernetes cluster, the kind of FQDN you are trying to resolve might vary.
We should try to adjust the knobs of CoreDNS accordingly.
Suppose you are running Kubernetes in a particular public/private cloud and most of the DNS backed applications are in the same cloud.
In that case, CoreDNS also provides particular cloud-related or generic plugins which can be used to extend DNS zone records.

If you are interested in customizing DNS behaviour for your needs, we would suggest going through the book “Learning CoreDNS” by Cricket Liu and John Belamaric.
Book provides a detailed overview of different CoreDNS plugins with their use cases.
It also covers CoreDNS + Kubernetes integration in depth.

If you are running an appropriate number of instances of CoreDNS in your Kubernetes cluster or not is one of the critical factors to decide.
It’s recommended to run a least two instances of the CoreDNS servers for a better guarantee of DNS requests being served.
Depending upon the number of requests being served, nature of requests, number of workloads running on the cluster, and size of cluster you may need to add extra instances of CoreDNS or configure HPA (Horizontal Pod Autoscaler) for your cluster.
Factors like the number of requests being served, nature of requests, number of workloads running on the cluster and cluster size should help you in deciding the number of CoreDNS instances.
You may need to add extra instances of CoreDNS or configure HPA (Horizontal Pod Autoscaler) for your cluster.

### Summary

This blog post tries to highlight the importance of the DNS request cycle in Kubernetes, and many times you would end up in a situation where you start with “it’s not DNS” but end up “it’s always DNS !”.
So be careful of these landmines.

Enjoyed the article? Let’s start a conversation on [Twitter](https://twitter.com/sanketsudake) and share your “It’s always DNS” stories and how it was resolved.

Love Cloud Native? We do too ❤️
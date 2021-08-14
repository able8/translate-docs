# [Handling Client Requests Properly with Kubernetes](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/)

From [Kubernetes in Action](https://www.manning.com/books/kubernetes-in-action) 

It goes without saying that we want client requests to be handled  properly. We obviously don’t want to see broken connections when pods  are starting up or shutting down. By itself, Kubernetes doesn’t ensure  this never happens. Your app needs to follow a few rules to prevent  broken connections. This article discusses those rules.

**Making sure all client requests are handled properly**

Let’s look at the pod’s lifecycle from the perspective of the pod’s  clients (clients consuming the service the pod is providing). We want to make sure that our clients’ requests are properly handled, because if  connections start breaking when pods start up or shut down, we’re in  trouble. Kubernetes by itself doesn’t guarantee that this won’t happen,  so let’s see what we need to do to prevent it from happening.

**Preventing broken client connections when a pod is starting up**

Ensuring each connection is handled properly at pod startup is fairly simple, if you understand how services and service endpoints work. When a pod is started, it’s added as an endpoint to all the services, whose  label selector matches the pod’s labels. The pod also needs to signal to Kubernetes that it’s ready. Until it is, it won’t become a service  endpoint and therefore won’t receive any requests from clients.

If you don’t specify a readiness probe in your pod spec, the pod is  considered ready all the time. This means it’ll start receiving requests almost immediately – as soon as the first Kube-Proxy updates the  iptables rules on its node and the first client pod tries to connect to  the service. If your app isn’t ready to accept connections by then,  clients will see “connection refused” types of errors.

All you need to do is make sure that your readiness probe returns  success only when your app is ready to properly handle incoming  requests. A good first step is to add an HTTP GET readiness probe and  point it to the base URL of your app. In a lot of cases, that gets you  far enough and saves you from having to implement a special readiness  endpoint in your app.

**Preventing broken connections during pod shutdown**

Now let’s see what happens at the other end of a pod’s life – when  the pod is deleted and its containers are terminated. The pod’s  containers should start shutting down cleanly as soon as they receive a  SIGTERM signal (or even before that – when its pre-stop hook is  executed), but does this ensure all client requests will be handled  properly?

How should the app behave when it receives a termination signal?  Should it continue accepting requests? What about requests that have  already been received, but haven’t completed yet? What about persistent  HTTP connections, which may be in between requests, but are open  (there’s no active request on the connection)? Before we can answer  those questions, we need to take a detailed look at the chain of events  that unfolds across the cluster when a pod is deleted.

**Understanding the sequence of events occurring at pod deletion**

You need to always keep in mind that the components in a Kubernetes  cluster run as separate processes on multiple machines. They aren’t part of a single big monolithic process. It takes time for all the  components to be on the same page regarding the state of the cluster.  Let’s explore this fact by looking at what happens across the cluster  when a pod is deleted.

When a request for a pod deletion is received by the API server, it first modifies the state in *etcd* and then notifies its watchers of the deletion. Among those watchers  are the Kubelet and the Endpoints controller. The two sequences of  events, which happen in parallel (marked with either A or B), are shown  in figure 1.

------

![img](https://freecontent.manning.com/wp-content/uploads/Luksa_HCRPwK_01.png)

Figure 1 Sequence of events that occurs when a pod is deleted

------

In the A sequence of events, you’ll see that as soon as the Kubelet  receives the notification that the pod should be terminated, it  initiates the shutdown sequence (run the pre-stop hook, send SIGTERM,  wait for some time and then forcibly kill the container if it hasn’t yet terminated on its own). If the app responds to the SIGTERM by  immediately ceasing to receive client requests, any client trying to  connect to it receives a Connection refused error. The time it takes for this to happen from the time the pod is deleted is relatively short,  because of the direct path from the API server to the Kubelet.

Now, let’s look at what happens in the other sequence of events – the one leading up to the pod being removed from the iptables rules  (sequence B in the figure). When the Endpoints controller (which runs in the Controller Manager in the Kubernetes Control Plane) receives the  notification of the pod being deleted, it removes the pod as an endpoint in all services that the pod is a part of. It does this by modifying  the Endpoints API object by sending a REST request to the API server.  The API server then notifies everyone watching the Endpoints object.  Among those watchers are the Kube-Proxies running on the worker nodes.  Each of these proxies updates the *iptables* rules on its node,  which is what prevents new connections from being forwarded to the  terminating pod. An important detail here is that removing the iptables  rules has no effect on existing connections – clients who’re already  connected to the pod can still send additional requests to the pod  through those existing connections.

Both of these sequences of events happen in parallel. Most likely,  the time it takes to shut down the app process in the pod is slightly  shorter than the time required for the *iptables* rules to be  updated. This is because the chain of events that leads to iptables  rules being updated is considerably longer (see figure 2), because the  event must first reach the Endpoints controller, which sends a new  request to the API server, and then the API server must notify the  Kube-Proxy, before the proxy finally modifies the *iptables*  rules. This means there’s a high probability of the SIGTERM signal being sent well before the iptables rules are updated on all nodes.

------

![img](https://freecontent.manning.com/wp-content/uploads/Luksa_HCRPwK_02.png)

**Figure 2** Timeline of events when pod is deleted

------

The result is that the pod may still receive client requests after it has received the termination signal. If the app stops accepting  connections immediately, it causes clients to receive “connection  refused” types of errors (like what happens at pod startup if your app  isn’t capable of accepting connections immediately and you don’t define a readiness probe for it).

Solving the problem

Googling for solutions to this problem makes it seem that adding a  readiness probe to your pod solves this problem. Supposedly, all you  need to do is make the readiness probe start failing as soon as the pod  receives the SIGTERM. This is supposed to cause the pod to be removed as the endpoint of the service. The removal would happen only after the  readiness probe fails for a few consecutive times (this is configurable  in the readiness probe spec). And, obviously, the removal then still  needs to reach the Kube-Proxy before the pod is removed from *iptables* rules.

In reality, the readiness probe has absolutely no bearing on the  whole process at all. The Endpoints controller removes the pod from the  service endpoints as soon as it receives notice of the pod being deleted (when the *deletionTimestamp* field in the pod’s spec is no longer null). From that point on, the result of the readiness probe is irrelevant.

What’s the proper solution to the problem? How can we make sure all requests are handled fully?

Well, it’s clear that the pod needs to keep accepting connections  even after it receives the termination signal, up until all the  Kube-proxies have finished updating the iptables rules. Well, it’s not  only the Kube-Proxies. There may also be Ingress controllers or load  balancers forwarding connections to the pod directly, without going  through the service (iptables). This also includes clients using  client-side load-balancing. To ensure none of the clients experience  broken connections, you’d have to wait until all of them somehow notify  you they’ll no longer forward connections to the pod.

This is impossible, because all those components are distributed  across many different computers. Even if you knew the location of each  and every one of them and could wait until all of them say it’s ok to  shut down the pod, what would you do if one of them doesn’t respond? How long do you wait for the response? Remember, during that time, you’re  holding up the shutdown process.

The only reasonable thing you can do is wait for a long enough time  to ensure all of the proxies have done their job. But how long is long  enough? A few seconds should be enough in most situations, but  obviously, there’s no guarantee it’ll suffice every time. When the API  server or the Endpoints controller are overloaded, it may take longer  for the notification to reach the Kube-Proxy. It’s important to  understand that you can’t solve the problem perfectly, but even a five  or ten second delay should improve the user experience considerably. You can use a longer delay, but don’t go overboard, because the delay  prevents the container from shutting down promptly and causes the pod to still be shown in lists long after it has been deleted, which is always frustrating to the user deleting the pod.

 

Properly shutting down an application includes these steps:

- Wait for a few seconds, then stop accepting new connections,
- Close all keep-alive connections that aren’t in the middle of a request,
- Wait for all active requests to finish, and then
- Shut down completely.

 

To understand what’s happening with the connections and requests during this process, examine figure 3 carefully.

------

![img](https://freecontent.manning.com/wp-content/uploads/Luksa_HCRPwK_03.png)

**Figure 3** Properly handling existing and new connections after receiving a termination signal

------

Not as simple as exiting the process immediately upon receiving the  termination signal, right? Is it worth going through all this? That’s  for you to decide. But the least you can do is add a pre-stop hook that  waits a few seconds. Something like this, perhaps:

```
    lifecycle:                   
      preStop:                   
        exec:                    
          command:               
          - sh
          - -c
          - "sleep 5"
 
```

This way, you don’t need to modify the code of your app at all. If  your app already ensures all in-flight requests are processed  completely, this pre-stop delay may be all you need.

That’s all for this article.

------

For more, check out the entire book on liveBook [here](https://livebook.manning.com/#!/book/kubernetes-in-action/)*.*

------

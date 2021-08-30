# Graceful shutdown and zero downtime deployments in Kubernetes

Published in August 2020

**TL;DR:** *In this article, you will learn how to prevent broken connections when a  Pod starts up or shuts down. You will also learn how to shut down  long-running tasks gracefully.*

![Graceful shutdown and zero downtime deployments in Kubernetes](https://learnk8s.io/a/55d21503055aaf9ef8a04d5e595ed505.png)

> You can [download this handy diagram as a PDF here](https://learnk8s.io/a/graceful-shutdown-and-zero-downtime-deployments-in-kubernetes/graceful-shutdown.pdf).

In Kubernetes, creating and deleting Pods is one of the most common tasks.

Pods are created when you execute a rolling update, scale deployments, for every new release, for every job and cron job, etc.

But Pods are also deleted and recreated after evictions — when you mark a node as unschedulable for example.

*If the nature of those Pods is so ephemeral, what happens when a Pod is in the middle of responding to a request but it's told to shut down?*

*Is the request completed before shutdown?*

*What about subsequent requests, are those redirected somewhere else?*

Before discussing what happens when a Pod is deleted, it's necessary to talk to about what happens when a Pod is created.

Let's assume you want to create the following Pod in your cluster:

pod.yaml

```
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: web
      image: nginx
      ports:
        - name: web
          containerPort: 80
```

You can submit the YAML definition to the cluster with:

bash

```
kubectl apply -f pod.yaml
```

As soon as you enter the command, kubectl submits the Pod definition to the Kubernetes API.

*This is where the journey begins.*

## Saving the state of the cluster in the database

The Pod definition is received and inspected by the API and subsequently stored in the database — etcd.

The Pod is also added to [the Scheduler's queue.](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#scheduling-cycle-binding-cycle)

The Scheduler:

1. inspects the definition
2. collects details about the workload such as CPU and memory requests and then
3. decides which Node is best suited to run it [(through a process called Filters and Predicates).](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#extension-points)

At the end of the process:

- The Pod is marked as *Scheduled* in etcd.
- The Pod has a Node assigned to it.
- The state of the Pod is stored in etcd.

**But the Pod still does not exist.**

- ![When you submit a Pod with kubectl apply -f the YAML is sent to the Kubernetes API.](https://learnk8s.io/a/54a28f4c41dfd3abb594848af5f71eaf.svg)

  1/3 When you submit a Pod with `kubectl apply -f` the YAML is sent to the Kubernetes API. Next 

The previous tasks happened in the control plane, and the state is stored in the database.

*So who is creating the Pod in your Nodes?*

## The kubelet — the Kubernetes agent

**It's the kubelet's job to poll the control plane for updates.**

You can imagine the kubelet relentlessly asking to the master node: *"I look after the worker Node 1, is there any new Pod for me?".*

When there is a Pod, the kubelet creates it.

*Sort of.*

The kubelet doesn't create the Pod by itself. Instead, it delegates the work to three other components:

1. **The Container Runtime Interface (CRI)** — the component that creates the containers for the Pod.
2. **The Container Network Interface (CNI)** — the component that connects the containers to the cluster network and assigns IP addresses.
3. **The Container Storage Interface (CSI)** — the component that mounts volumes in your containers.

In most cases, the Container Runtime Interface (CRI) is doing a similar job to:

bash

```
docker run -d <my-container-image>
```

The Container Networking Interface (CNI) is a bit more interesting because it is in charge of:

1. Generating a valid IP address for the Pod.
2. Connecting the container to the rest of the network.

As you can imagine, there are several ways to connect the container to the network and assign a valid IP address (you could choose between IPv4 or IPv6 or maybe assign multiple IP addresses).

As an example, [Docker creates virtual ethernet pairs and attaches it to a bridge](https://archive.shivam.dev/docker-networking-explained/), whereas [the AWS-CNI connects the Pods directly to the rest of the Virtual Private Cloud (VPC).](https://itnext.io/kubernetes-is-hard-why-eks-makes-it-easier-for-network-and-security-architects-ea6d8b2ca965)

When the Container Network Interface finishes its job, the Pod is connected  to the rest of the network and has a valid IP address assigned.

*There's only one issue.*

**The kubelet knows about the IP address (because it invoked the Container Network Interface), but the control plane does not.**

No one told the master node that the Pod has an IP address assigned and it's ready to receive traffic.

As far the control plane is concerned, the Pod is still being created.

It's the job of **the kubelet to collect all the details of the Pod such as the IP address and report them back to the control plane.**

You can imagine that inspecting etcd would reveal not just where the Pod is running, but also its IP address.

- ![The Kubelet polls the control plane for updates.](https://learnk8s.io/a/b8bdb7bf659fbea3d2949930093d56b1.svg)

  1/5 The Kubelet polls the control plane for updates.  Next 

If the Pod isn't part of any Service, this is the end of the journey.

The Pod is created and ready to use.

*When the Pod is part of the Service, there are a few more steps needed.*

## Pods and Services

When you create a Service, there are usually two pieces of information that you should pay attention to:

1. The selector, which is used to specify the Pods that will receive the traffic.
2. The `targetPort` — the port used by the Pods to receive traffic.

A typical YAML definition for the Service looks like this:

service.yaml

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  ports:
  - port: 80
    targetPort: 3000
  selector:
    name: app
```

When you submit the Service to the cluster with `kubectl apply`, Kubernetes finds all the Pods that have the same label as the selector (`name: app`) and collects their IP addresses — but only if they passed the [Readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-tcp-liveness-probe).

Then, for each IP address, it concatenates the IP address and the port.

If the IP address is `10.0.0.3` and the `targetPort` is `3000`, Kubernetes concatenates the two values and calls them an endpoint.

```
IP address + port = endpoint
---------------------------------
10.0.0.3   + 3000 = 10.0.0.3:3000
```

The endpoints are stored in etcd in another object called Endpoint.

*Confused?*

Kubernetes refers to:

- endpoint (in this article and the Learnk8s material this is referred to as a lowercase `e` endpoint) is the IP address + port pair (`10.0.0.3:3000`).
- Endpoint (in this article and the Learnk8s material this is referred to as an uppercase `E` Endpoint) is a collection of endpoints.

The Endpoint object is a real object in Kubernetes and for every Service Kubernetes automatically creates an Endpoint object.

You can verify that with:

bash

```
kubectl get services,endpoints
NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)
service/my-service-1   ClusterIP   10.105.17.65   <none>        80/TCP
service/my-service-2   ClusterIP   10.96.0.1      <none>        443/TCP

NAME                     ENDPOINTS
endpoints/my-service-1   172.17.0.6:80,172.17.0.7:80
endpoints/my-service-2   192.168.99.100:8443
```

The Endpoint collects all the IP addresses and ports from the Pods.

*But not just one time.*

The Endpoint object is refreshed with a new list of endpoints when:

1. A Pod is created.
2. A Pod is deleted.
3. A label is modified on the Pod.

So you can imagine that every time you create a Pod and after the kubelet  posts its IP address to the master Node, Kubernetes updates all the  endpoints to reflect the change:

bash

```
kubectl get services,endpoints
NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)
service/my-service-1   ClusterIP   10.105.17.65   <none>        80/TCP
service/my-service-2   ClusterIP   10.96.0.1      <none>        443/TCP

NAME                     ENDPOINTS
endpoints/my-service-1   172.17.0.6:80,172.17.0.7:80,172.17.0.8:80
endpoints/my-service-2   192.168.99.100:8443
```

Great, the endpoint is stored in the control plane, and the Endpoint object was updated.

- ![In this picture, there's a single Pod deployed in your cluster. The Pod belongs to a Service. If you were to inspect etcd, you would find the Pod's details as well as Service.](https://learnk8s.io/a/5ec899bd0f7067f1e01bb1accacb35ac.svg)

  1/8

  In this picture, there's a single Pod deployed in your cluster. The Pod  belongs to a Service. If you were to inspect etcd, you would find the  Pod's details as well as Service.

  Next 

*Are you ready to start using your Pod?*

**There's more.**

A lot more!

## Consuming endpoints in Kubernetes

**Endpoints are used by several components in Kubernetes.**

Kube-proxy uses the endpoints to set up iptables rules on the Nodes.

So every time there is a change to an Endpoint (the object), kube-proxy  retrieves the new list of IP addresses and ports and write the new  iptables rules.

- ![Let's consider this three-node cluster with two Pods and no Services. The state of the Pods is stored in etcd.](https://learnk8s.io/a/50e4746cbbda956a4550077f2954dd7a.svg)

  1/6

  Let's consider this three-node cluster with two Pods and no Services. The state of the Pods is stored in etcd.

  Next 

The Ingress controller uses the same list of endpoints.

The Ingress controller is that component in the cluster that routes external traffic into the cluster.

When you set up an Ingress manifest you usually specify the Service as the destination:

ingress.yaml

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
spec:
  rules:
  - http:
      paths:
      - backend:
          service:
            name: my-service
            port:
              number: 80
        path: /
        pathType: Prefix
```

*In reality, the traffic is not routed to the Service.*

Instead, the Ingress controller sets up a subscription to be notified every time the endpoints for that Service change.

**The Ingress routes the traffic directly to the Pods skipping the Service.**

As you can imagine, every time there is a change to an Endpoint (the  object), the Ingress retrieves the new list of IP addresses and ports  and reconfigures the controller to include the new Pods.

- ![In this picture, there's an Ingress controller with a Deployment with two replicas and a Service.](https://learnk8s.io/a/175161d4127b2daf602b2a240190771d.svg)

  1/9

  In this picture, there's an Ingress controller with a Deployment with two replicas and a Service.

  Next 

There are more examples of Kubernetes components that subscribe to changes to endpoints.

CoreDNS, the DNS component in the cluster, is another example.

If you use [Services of type Headless](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services), CoreDNS will have to subscribe to changes to the endpoints and reconfigure itself every time an endpoint is added or removed.

The same endpoints are consumed by service meshes such as Istio or Linkerd, [by cloud providers to create Services of `type:LoadBalancer`](https://thebsdbox.co.uk/2020/03/18/Creating-a-Kubernetes-cloud-doesn-t-required-boiling-the-ocean/) and countless operators.

You must remember that several components subscribe to change to endpoints  and they might receive notifications about endpoint updates at different times.

*Is it enough, or is there something happening after you create a Pod?*

**This time you're done!**

A quick recap on what happens when you create a Pod:

1. The Pod is stored in etcd.
2. The scheduler assigns a Node. It writes the node in etcd.
3. The kubelet is notified of a new and scheduled Pod.
4. The kubelet delegates creating the container to the Container Runtime Interface (CRI).
5. The kubelet delegates attaching the container to the Container Network Interface (CNI).
6. The kubelet delegates mounting volumes in the container to the Container Storage Interface (CSI).
7. The Container Network Interface assigns an IP address.
8. The kubelet reports the IP address to the control plane.
9. The IP address is stored in etcd.

And if your Pod belongs to a Service:

1. The kubelet waits for a successful Readiness probe.
2. All relevant Endpoints (objects) are notified of the change.
3. The Endpoints add a new endpoint (IP address + port pair) to their list.
4. Kube-proxy is notified of the Endpoint change. Kube-proxy updates the iptables rules on every node.
5. The Ingress controller is notified of the Endpoint change. The controller routes traffic to the new IP addresses.
6. CoreDNS is notified of the Endpoint change. If the Service is of type Headless, the DNS entry is updated.
7. The cloud provider is notified of the Endpoint change. If the Service is of `type: LoadBalancer`, the new endpoint are configured as part of the load balancer pool.
8. Any service mesh installed in the cluster is notified of the Endpoint change.
9. Any other operator subscribed to Endpoint changes is notified too.

*Such a long list for what is surprisingly a common task — creating a Pod.*

The Pod is *Running*. It is time to discuss what happens when you delete it.

## Deleting a Pod

You might have guessed it already, but when the Pod is deleted, you have to follow the same steps but in reverse.

First, the endpoint should be removed from the Endpoint (object).

This time the Readiness probe is ignored, and the endpoint is removed immediately from the control plane.

That, in turn, fires off all the events to kube-proxy, Ingress controller, DNS, service mesh, etc.

Those components will update their internal state and stop routing traffic to the IP address.

Since the components might be busy doing something else, **there is no guarantee on how long it will take to remove the IP address from their internal state.**

For some, it could take less than a second; for others, it could take more.

- ![If you're deleting a Pod with kubectl delete pod, the command reaches the Kubernetes API first.](https://learnk8s.io/a/336567c1d80c8853cb6900bdf6fd30d9.svg)

  1/5

  If you're deleting a Pod with `kubectl delete pod`, the command reaches the Kubernetes API first.

  Next 

At the same time, the status of the Pod in etcd is changed to *Terminating*.

The kubelet is notified of the change and delegates:

1. Unmounting any volumes from the container to the Container Storage Interface (CSI).
2. Detaching the container from the network and releasing the IP address to the Container Network Interface (CNI).
3. Destroying the container to the Container Runtime Interface (CRI).

In other words, Kubernetes follows precisely the same steps to create a Pod but in reverse.

- ![If you're deleting a Pod with kubectl delete pod, the command reaches the Kubernetes API first.](https://learnk8s.io/a/dbfd8be2cd6fbc3984dbb12cfece608d.svg)

  1/3

  If you're deleting a Pod with `kubectl delete pod`, the command reaches the Kubernetes API first.

  Next 

However, there is a subtle but essential difference.

**When you terminate a Pod, removing the endpoint and the signal to the kubelet are issued at the same time.**

When you create a Pod for the first time, Kubernetes waits for the kubelet  to report the IP address and then kicks off the endpoint propagation.

**However, when you delete a Pod, the events start in parallel.**

And this could cause quite a few race conditions.

*What if the Pod is deleted before the endpoint is propagated?*

- ![Deleting the endpoint and deleting the Pod happen at the same time.](https://learnk8s.io/a/e17a49eb08f03f2c2ff02b91409314fb.svg)

  1/3

  Deleting the endpoint and deleting the Pod happen at the same time.

  Next 

## Graceful shutdown

When a Pod is terminated before the endpoint is removed from kube-proxy or  the Ingress controller, you might experience downtime.

And, if you think about it, it makes sense.

Kubernetes is still routing traffic to the IP address, but the Pod is no longer there.

The Ingress controller, kube-proxy, CoreDNS, etc. didn't have enough time to remove the IP address from their internal state.

Ideally, Kubernetes should wait for all components in the cluster to have an  updated list of endpoints before the Pod is deleted.

*But Kubernetes doesn't work like that.*

Kubernetes offers robust primitives to distribute the endpoints (i.e. the Endpoint object and more advanced abstractions such as [Endpoint Slices](https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/)).

However, Kubernetes does not verify that the components that subscribe to  endpoints changes are up-to-date with the state of the cluster.

*So what can you do avoid this race conditions and make sure that the Pod is deleted after the endpoint is propagated?*

**You should wait.**

**When the Pod is about to be deleted, it receives a SIGTERM signal.**

Your application can capture that signal and start shutting down.

Since it's unlikely that the endpoint is immediately deleted from all components in Kubernetes, you could:

1. Wait a bit longer before exiting.
2. Still process incoming traffic, despite the SIGTERM.
3. Finally, close existing long-lived connections (perhaps a database connection or WebSockets).
4. Close the process.

*How long should you wait?*

**By default, Kubernetes will send the SIGTERM signal and wait 30 seconds before force killing the process.**

So you could use the first 15 seconds to continue operating as nothing happened.

Hopefully, the interval should be enough to propagate the endpoint removal to kube-proxy, Ingress controller, CoreDNS, etc.

And, as a consequence, less and less traffic will reach your Pod until it stops.

After the 15 seconds, it's safe to close your connection with the database  (or any persistent connections) and terminate the process.

If you think you need more time, you can stop the process at 20 or 25 seconds.

However, you should remember that Kubernetes will forcefully kill the process after 30 seconds [(unless you change the `terminationGracePeriodSeconds` in your Pod definition).](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#hook-handler-execution)

*What if you can't change the code to wait longer?*

You could invoke a script to wait for a fixed amount of time and then let the app exit.

Before the SIGTERM is invoked, Kubernetes exposes a `preStop` hook in the Pod.

You could set the `preStop` to hook to wait for 15 seconds.

Let's have a look at an example:

pod.yaml

```
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: web
      image: nginx
      ports:
        - name: web
          containerPort: 80
      lifecycle:
        preStop:
          exec:
            command: ["sleep", "15"]
```

The `preStop` hook is one of the [Pod LifeCycle hooks](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/).

*Is a 15 seconds delay the recommended amount?*

It depends, but it could be a sensible way to start testing.

Here's a recap of what options you have:

- ![You already know that, when a Pod is deleted, the kubelet is notified of the change.](https://learnk8s.io/a/bdaa1da0be0fa3e9fe022cbf2b22bd1d.svg)

  1/5

  You already know that, when a Pod is deleted, the kubelet is notified of the change.

  Next 

## Grace periods and rolling updates

Graceful shutdown applies to Pods being deleted.

*But what if you don't delete the Pods?*

Even if you don't, Kubernetes deletes Pods all the times.

In particular, Kubernetes creates and deletes Pods every time you deploy a newer version of your application.

When you change the image in your Deployment, Kubernetes rolls out the change incrementally.

pod.yaml

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  replicas: 3
  selector:
    matchLabels:
      name: app
  template:
    metadata:
      labels:
        name: app
    spec:
      containers:
      - name: app
        # image: nginx:1.18 OLD
        image: nginx:1.19
        ports:
          - containerPort: 3000
```

If you have three replicas and as soon as you submit the new YAML resources Kubernetes:

- Creates a Pod with the new container image.
- Destroys an existing Pod.
- Waits for the Pod to be ready.

And it repeats the steps above until all the Pods are migrated to the newer version.

Kubernetes repeats each cycle only after the new Pod is ready to receive traffic  (in other words, it passes the Readiness check).

*Does Kubernetes wait for the Pod to be deleted before moving to the next one?*

**No.**

If you have 10 Pods and the Pod takes 2 seconds to be ready and 20 to shut down this is what happens:

1. The first Pod is created, and a previous Pod is terminated.
2. The new Pod takes 2 seconds to be ready after that Kubernetes creates a new one.
3. In the meantime, the Pod being terminated stays terminating for 20 seconds

After 20 seconds, all new Pods are live (10 Pods, *Ready* after 2 seconds) and all 10 the previous Pods are terminating (the first *Terminated* Pod is about to exit).

In total, you have double the amount of Pods for a short amount of time (10 *Running*, 10 *Terminating*).

![Rolling update and graceful shutdown](https://learnk8s.io/a/7934a67ff44a183254acf81a763e6c2f.svg)

The longer the graceful period compared to the Readiness probe, the more Pods you will have *Running* (and *Terminating*) at the same time.

*Is it bad?*

Not necessarily since you're careful not dropping connections.

## Terminating long-running tasks

*And what about long-running jobs?*

*If you are transcoding a large video, is there any way to delay stopping the Pod?*

Imagine you have a Deployment with three replicas.

Each replica is assigned a video to transcode, and the task could take several hours to complete.

When you trigger a rolling update, the Pod has 30 seconds to complete the task before it's killed.

*How can you avoid delaying shutting down the Pod?*

You could increase the `terminationGracePeriodSeconds` to a couple of hours.

**However, the endpoint of the Pod is unreachable at that point.**

![Unreachable Pod](https://learnk8s.io/a/f6777b220dd43ad627b02073c2ffde30.svg)

If you expose metrics to monitor your Pod, your instrumentation won't be able to reach your Pod.

*Why?*

**Tools such as Prometheus rely on Endpoints to scrape Pods in your cluster.**

However, as soon as you delete the Pod, the endpoint deletion is propagated in the cluster — even to Prometheus!

**Instead of increasing the grace period, you should consider creating a new Deployment for every new release.**

When you create a brand new Deployment, the existing Deployment is left untouched.

The long-running jobs can continue processing the video as usual.

Once they are done, you can delete them manually.

If you wish to delete them automatically, you might want to set up an  autoscaler that can scale your deployment to zero replicas when they run out of tasks.

An example of such Pod autoscaler is [Osiris — a general-purpose, scale-to-zero component for Kubernetes.](https://github.com/deislabs/osiris)

The technique is sometimes referred to as **Rainbow Deployments** and is useful every time you have to keep the previous Pods *Running* for longer than the grace period.

*Another excellent example is WebSockets.*

If you are streaming real-time updates to your users, you might not want  to terminate the WebSockets every time there is a release.

If you are frequently releasing during the day, that could lead to several interruptions to real-time feeds.

**Creating a new Deployment for every release is a less obvious but better choice.**

Existing users can continue streaming updates while the most recent Deployment serves the new users.

As a user disconnects from old Pods, you can gradually decrease the replicas and retire past Deployments.

## Summary

You should pay attention to Pods being deleted from your cluster since their IP address might be still used to route traffic.

Instead of immediately shutting down your Pods, you should consider waiting a little bit longer in your application or set up a `preStop` hook.

The Pod should be removed only after all the endpoints in the cluster are  propagated and removed from kube-proxy, Ingress controllers, CoreDNS,  etc.

If your Pods run long-lived  tasks such as transcoding videos or serving real-time updates with  WebSockets, you should consider using rainbow deployments.

In rainbow deployments, you create a new Deployment for every release and  delete the previous one when the connection (or the tasks) drained.

You can manually remove the older deployments as soon as the long-running task is completed.

Or you could automatically scale your deployment to zero replicas to automate the process.
# Rolling Updates and Blue-Green Deployments with Kubernetes and HAProxy

Feb 11, 2020 From: https://www.haproxy.com/blog/rolling-updates-and-blue-green-deployments-with-kubernetes-and-haproxy/

_The HAProxy Kubernetes Ingress Controller supports two popular deployment patterns for updating applications in Kubernetes: rolling updates and blue-green deployments._

_This is the second post in a [series about HAProxy’s role](https://www.haproxy.com/blog/building-blocks-of-a-modern-proxy/) in building a modern systems architecture that relies on cloud-native technology such-as Docker containers and Kubernetes. Containers have revolutionized how software is deployed, allowing the microservice pattern to flourish and enabling self-healing, autoscaling applications. HAProxy is an intelligent load balancer that adds high performance, observability, security, and many other features to the mix._

**Learn more by registering for our webinar** [**_“HAProxy Skills Lab: Deployment Patterns in Kubernetes Using the HAProxy Kubernetes Ingress Controller_ _”._**](https://www.haproxy.com/blog/webinar-haproxy-skills-lab-deployment-patterns-in-kubernetes-using-the-haproxy-kubernetes-ingress-controller/)

So, you have deployed your application to Kubernetes and it’s running flawlessly. The next important question is, how should you deploy the next version of it safely? How can you replace the existing pods without disrupting traffic? Furthermore, how is it affected by routing traffic through the HAProxy Kubernetes Ingress Controller?

Kubernetes accommodates a wide range of deployment methods. We’ll cover two that guarantee a safe rollout while keeping the ability to revert if necessary:

- _Rolling updates_ have first-class support in Kubernetes and allow you to phase in a new version gradually;
- _Blue-green deployments_ avoid having two versions at play at the same time by swapping one set of pods for another.

The HAProxy Kubernetes Ingress Controller is powered by the world’s fastest and most widely used software load balancer. Known to provide the utmost performance, observability, and security, it is the most efficient way to route traffic into a Kubernetes cluster. It automatically detects changes within your Kubernetes infrastructure and ensures accurate distribution of traffic to healthy pods. Its design prevents downtime even when there are rapid configuration changes. It supports both deployment patterns and reliably exposes the correct pods to clients.

## Deploy the HAProxy Kubernetes Ingress Controller

In this blog post, I use [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) to start up a simple Kubernetes cluster on my workstation. Minikube requires a hypervisor, such as VirtualBox, to be installed. Once it’s up and running, you will be able to expose services running inside the Kubernetes cluster at the IP address **192.168.99.100**.

After installing and starting Minikube, deploy the [HAProxy Kubernetes Ingress Controller](https://www.haproxy.com/documentation/hapee/2-0r1/traffic-management/kubernetes-ingress-controller/), which is responsible for routing traffic into your Kubernetes cluster. You can either install the open-source version or the Enterprise version, which is built upon HAProxy Enterprise. It adds features such as a Web Application Firewall, which is essential for stopping application-layer attacks.

By default, the Ingress Controller assumes that you want to configure SSL. If you prefer to try things without SSL, then [download its YAML file](https://raw.githubusercontent.com/haproxytech/kubernetes-ingress/master/deploy/haproxy-ingress.yaml) and modify its `ConfigMap` so that `ssl-redirect` is _OFF_.

haproxy-ingress.yaml

## Rolling Updates

A **rolling update** offers a way to deploy the new version of your application gradually across your cluster. It replaces pods during several phases. For example, you may replace 25% of the pods during the first phase, then another 25% during the next, and so on until all are upgraded. Since the pods are not replaced all at once, this means that both versions will be live, at least for a short time, during the rollout.

**Did You Know?** Because a rolling update creates the potential for two versions of your application to be deployed simultaneously, make sure that any upstream databases and services are compatible with both versions.

This deployment model enjoys first-class support in Kubernetes with baked-in YAML configuration options. Here’s how it works:

1. Version 1 of your application is already deployed.
2. Push version 2 of your application to your container image repository.
3. Update the version number in the Deployment object’s definition.
4. Apply the change with`kubectl`.
5. Kubernetes staggers the rollout of the new version across your pods.
6. The HAProxy Kubernetes Ingress Controller detects when the new pods are live. It automatically updates its proxy configuration, routing traffic away from the old pods and towards the new ones.

A rolling update dodges downtime by replacing existing pods incrementally. If the new pods introduce an error that stops them from starting up, Kubernetes will pause the rollout. Also, a rolling update ensures that some pods are always up, so there’s no downtime. Kubernetes keeps a minimum number of pods running during the rollout. However, this requires that you’ve added a _readiness check_ to your pods so that Kubernetes knows when they are truly ready to receive traffic.

### Deploy the Original Application

Kubernetes enables rolling updates by default. An update begins when you change your `Deployment` resource’s YAML file and then use `kubectl` apply. Consider the following definition, which deploys version 1 of an application. Note that I am using the [errm/versions](https://hub.docker.com/r/errm/versions) Docker image because it displays the version of the application when you browse to its webpage, which makes it easy to see which version you’re running.

app.yaml

The `readinessProbe` section tells Kubernetes to send an HTTP request to the application five seconds after it has started, and then every five seconds thereafter. No traffic is sent to the pod until a successful response is returned. This is key to preventing downtime.

**Did You Know?** Consider tagging your container images with version numbers, rather than using a tag like _latest_. This allows you to keep track of the versions that are deployed and manage the release of new versions.

Next, define a `Service` object that will categorize the pods into a single group that the Ingress Controller will watch:

app-service.yaml

Next, define an `Ingress` object. This configures how the HAProxy Ingress Controller will route traffic to the pods:

ingress.yaml

Use `kubectl apply` to deploy the pods, service and ingress:

Version 1 of your application is now deployed. Run the following command to see which port the HAProxy Kubernetes Ingress Controller has mapped to port 80:

You can then see that the application is exposed on port 31179. You can see it by visiting the Minikube IP address **http://192.168.99.100:31179** in your browser.

![Version 1 web page](https://cdn.haproxy.com/wp-content/uploads/2020/02/version1.png)

Version 1 web page

Let’s see how to upgrade it to version 2 next.

### Upgrade Using a Rolling Update

After you have pushed a new version of your application to your container repository, trigger a rolling update by increasing the version number set on the `Deployment` definition’s `spec.template.spec.containers.image` property. This tells Kubernetes that the current, desired version of your application has changed. In our example, since we’re using a prebaked image, there’s already a version 2 set up in the Docker Hub repository.

app.yaml

Then, use `kubectl apply` to start the rollout:

You can check the status of the rollout by using the `kubectl rollout status` command:

Once completed, you can access the application again at the same URL, **http://192.168.99.100:31179**. It shows you a new web page signifying that version 2 has been deployed.

![Version 2 web page](https://cdn.haproxy.com/wp-content/uploads/2020/02/version2.png)

Version 2 web page

If you decide that the new version is faulty, you can revert to the previous one by using the `kubectl rollout undo` command, like this:

The HAProxy Kubernetes Ingress Controller detects pod changes quickly and can switch back and forth between versions without dropping connections. Rolling updates aren’t the only way to accomplish highly-available services, though. In the next section, you’ll learn about blue-green deployments, which update all pods simultaneously.

## Blue-Green Deployments

A **blue-green deployment** lets you replace an existing version of your application across all pods at once. The name, blue-green, [was coined](https://gitlab.com/snippets/1846041) in the book _Continuous Delivery_ by Jez Humble and David Farley. Here’s how it works:

1. Version 1 of your application is already deployed.
2. Push version 2 of your application to your container image repository.
3. Deploy version 2 of your application to a new group of pods. Both versions 1 and 2 pods are now running in parallel. However, only version 1 is exposed to external clients.
4. Run internal testing on version 2 and make sure it is ready to go live.
5. Flip a switch and the ingress controller in front of your clusters stops routing traffic to the version 1 pods and starts routing it to the version 2 pods.

This deployment pattern has a few advantages over a rolling update. For one, at no time are there ever two versions of your application accessible to external clients at the same time. So, all users will receive the same client-side Javascript files and be routed to a version of the application that supports the API calls within those files. It also simplifies upstream dependencies, such as database schemas.

Another advantage is that it gives you time to test the new version in a production environment before it goes live. You control how long to wait before making the switch. Meanwhile, you can verify that the application and its dependencies function normally.

On the other hand, a blue-green deployment is all-or-nothing. Unlike a rolling update, you aren’t able to gradually roll out the new version. All users will receive the update at the same time, although existing sessions will be allowed to finish their work on the old instances. So, the stakes are a bit higher that everything should work, once you do initiate the change. It also requires allocating more server resources, since you will need to run two copies of every pod.

Luckily, the rollback procedure is just as easy: You simply flip the switch again and the previous version is swapped back into place. That’s because the old version is still running on the old pods. It is simply that traffic is no longer being routed to them. When you’re confident that the new version is here to stay, you can decommission those pods.

You’ll need to set up your original application in a slightly different way when you expect to use a blue-green deployment. There is more emphasis on using Kubernetes metadata labels, which will become clear in the next section.

### Deploy the Original Application

Consider the following definition, which deploys version 1 of your application. Note its `spec.selector` section, which specifies a label called _version_:

app-v1.yaml

A `Deployment` object defines a `spec.selector` section that matches the `spec.template.metadata` section. This is how a Deployment tags pods and keeps track of them. This is the key to setting up a blue-green deployment. By using different labels, you can deploy multiple versions of the same application. Here, the `spec.selector.matchLabels` property is set to _run=app,version=0.0.1_. The version should match the version tag of your Docker image, for convenience and simplicity.

The following Service definition targets that same selector:

app-service-bg.yaml

Next, use the following `Ingress` definition to expose the version 1 pods to the world. It registers a route with the HAProxy Kubernetes Ingress Controller:

ingress.yaml

Apply everything using `kubectl`:

At this point, you can access the application at the HTTP port exposed by the Ingress Controller: **http://192.168.99.100:31179**. Now, let’s see how to use a blue-green deployment to upgrade the version.

### Upgrade Using a Blue-green Deployment

Now that the _blue_ version (i.e. version 1) is released, create a _green_ version of your `Deployment` object that will deploy version 2. The YAML will be the same, except that you increase the value of the _version_ label, as well as the Docker image tag. Also note that the name of the deployment is changed from _app-blue_ to _app-green_, since you cannot have two Deployments with the same name that target different pods.

app-v2.yaml

Apply it with `kubectl`:

At this point, both blue (version 1) and green (version 2) are deployed. Only the blue instance is receiving traffic, though. To make the switch, update your `Service` definition’s _version_ selector so that it points to the new version:

app-service.yaml

Apply it with `kubectl`:

Check the application again and you will see that the new version is live. If you need to roll back to the earlier version, simply change the `Service` definition’s selector back and reapply it. The HAProxy Kubernetes Ingress Controller detects these changes almost instantly and you can swap back and forth to your heart’s content. There’s no downtime during the cutover. Established TCP connections will finish normally on the instance where they began.

### Testing the New Pods

You can also test the new version before it’s released by registering a different ingress route that exposes the application at a new URL path. First, create another `Service` definition called _test-service_:

test-service.yaml

Note that we are including the path-rewrite annotation, which rewrites the URL **/test** to **/** before it reaches the pod. Then, add a new route to your existing `Ingress` object that exposes this service at the URL path **/test**, as shown:

ingress.yaml

This lets you check your application by visiting **/test** in your browser.

## Conclusion

The HAProxy Kubernetes Ingress Controller is powered by the legendary HAProxy. Known to provide the utmost performance, observability, and security, it features many benefits including SSL termination, rate limiting, and IP whitelisting. When you deploy the ingress controller into your cluster, it’s important to consider how your applications will be upgraded later. Two popular methods are rolling updates and blue-green deployments.

Rolling updates allow you to phase in a new version gradually and it has first-class support in Kubernetes. Blue-green deployments avoid the complexity of having two versions at play at the same time and give you a chance to test the change before going live. In either case, the HAProxy Kubernetes Ingress Controller detects these changes quickly and maintains uptime throughout.

If you enjoyed this post and want to see more like it, subscribe to this blog! You can also follow us on [Twitter](https://twitter.com/haproxy) and join the conversation on [Slack](https://slack.haproxy.com/).

The Enterprise version of the ingress controller combines HAProxy, the world’s fastest and most widely used open-source software load balancer and application delivery controller, with enterprise-class features, services and premium support. [Contact us](https://www.haproxy.com/contact-us/) to learn more and sign up for a [free trial](https://www.haproxy.com/downloads/hapee-trial/).

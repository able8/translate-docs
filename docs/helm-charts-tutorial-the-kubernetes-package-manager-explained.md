# Helm Charts Tutorial: The Kubernetes Package Manager Explained

![Sebastian Sigl](http://www.freecodecamp.org/news/content/images/size/w100/2020/01/profile_photo_edited.png)[Sebastian Sigl](http://www.freecodecamp.org/news/author/sesigl/)

![Helm Charts Tutorial: The Kubernetes Package Manager Explained](http://www.freecodecamp.org/news/content/images/size/w2000/2020/12/helm-blog-logo.jpg)

There are different ways of running production services at a high scale. One popular solution for running containers in production is Kubernetes. But interacting with Kubernetes directly comes with some caveats.

Helm tries to solve some of the challenges with useful features that increase productivity and reduce maintenance efforts of complex deployments.

In this post you will learn:

- What Helm is
- The most common use-cases of Helm
- How to configure and deploy a publicly available Helm package
- How to deploy a custom application using Helm

Every code example in this post requires a Kubernetes cluster. The easiest way to get a cluster to play with is to install Docker and activate its Kubernetes cluster feature. Also, you need to [install kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and [Helm](https://helm.sh/docs/intro/install/) to interact with your cluster.

_Please note: When you try the examples, be patient. If you are too fast then the containers are not ready. It might take a few minutes until the containers can receive requests._

## What is Helm?

Helm calls itself ”The Kubernetes package manager”. It is a command-line tool that enables you to create and use so-called Helm Charts.

A Helm Chart is a collection of templates and settings that describe a set of Kubernetes resources. Its power spans from managing a single node definition to a highly scalable multi-node cluster.

The architecture of Helm has changed over the last years. The current version of Helm communicates directly to your Kubernetes cluster via Rest. If you read something about Tiller in the context of Helm, then you're reading an old article. Tiller was removed in Helm 3.

Helm itself is stateful. When a Helm Chart gets installed, the defined resources are getting deployed and meta-information is stored in Kubernetes secrets.

## How to Deploy a Simple Helm Application

Let’s get our hands dirty and make sure Helm is ready to use.

First, we need to be connected to a Kubernetes cluster. In this example, I will concentrate on a Kubernetes cluster that comes with your Docker setup. So if you use some other Kubernetes cluster, configurations and outputs might differ.

```shell
$ kubectl config use-context docker-desktop

Switched to context "docker-desktop".

$ kubectl get node

NAME             STATUS   ROLES    AGE   VERSION
docker-desktop   Ready    master   20d   v1.19.3
```

Let’s deploy an Apache webserver using Helm. As a first step, we need to tell Helm what location to search by adding a Helm repository:

```shell
$ helm repo add bitnami https://charts.bitnami.com/bitnami
```

Let’s install the actual container:

```shell
$ helm install my-apache bitnami/apache --version 8.0.2
```

After a few minutes your deployment is ready. We can check the state of the containers using kubectl:

```shell
$ kubectl get pods

NAME                               READY   STATUS    RESTARTS   AGE
my-apahe-apache-589b8df6bd-q6m2n   1/1     Running   0          2m27s
```

Now, open [http://localhost](http://localhost) to see the default Apache exposed website locally. Also, Helm can show us information about current deployments:

```shell
$ helm list

NAME     	REVISION	STATUS  	CHART       	VERSION
my-apache	1		deployed  	apache-8.0.2	2.4.46
```

### How to Upgrade a Helm Application

We can upgrade our deployed application to a new version like this:

```shell
$ helm upgrade my-apache bitnami/apache --version 8.0.3

$ helm list

NAME     	REVISION	STATUS  	CHART      	VERSION
my-apache	2       	deployed	apache-8.0.3	2.4.46
```

The column Revision indicates that this is the 2nd version we've deployed.

### How to Rollback a Helm Application

So let’s try to rollback to the first deployed version:

```shell
$ helm rollback my-apache 1

Rollback was a success! Happy Helming!

$ helm list

NAME     	REVISION	STATUS  	CHART		VERSION
my-apache	3       	deployed	apache-8.0.2	2.4.46
```

This is a very powerful feature that allows you to roll back changes in production quickly.

I mentioned that Helm stores deployment information in secrets – here they are:

```shell
$ kubectl get secret

NAME		  		  TYPE        	   	  	DATA   AGE
default-token-nc4hn               kubernetes.io/sat		3      20d
sh.helm.release.v1.my-apache.v1   helm.sh/release.v1		1      1m
sh.helm.release.v1.my-apache.v2   helm.sh/release.v1		1      1m
sh.helm.release.v1.my-apache.v3   helm.sh/release.v1		1      1m
```

### How to Remove a Deployed Helm Application

Let’s clean up our Kubernetes by removing the my-apache release:

```shell
$ helm delete my-apache

release "my-apache" uninstalled
```

Helm gives you a very convenient way of managing a set of applications that enables you to deploy, upgrade, rollback and delete.

Now, we are ready to use more advanced Helm features that will boost your productivity!

## How to Access Production-Ready Helm Charts

You can search public hubs for Charts that enable you to quickly deploy your desired application with a customizable configuration.

A Helm Chart doesn't just contain a static set of definitions. Helm comes with capabilities to hook into any lifecycle state of a Kubernetes deployment. This means during the installation or upgrade of an application, various actions can be executed like creating a database update before updating the actual database.

This powerful definition of Helm Charts lets you share and improve an executable description of a deployment setup that spans from initial installation and version upgrades to rollback capabilities.

Helm might be heavy for a simple container like a single node web server, but it’s very useful for more complex applications. For example it works great for a distributed system like Kafka or Cassandra that usually runs on multiple distributed nodes on different datacenters.

We've already leveraged Helm to deploy a single Apache container. Now, we will deploy a production-ready WordPress application that contains:

- Containers that serve WordPress,
- Instances of MariaDB for persistence and
- Prometheus sidecar containers for each WordPress container to expose health metrics.

Before we deploy, it’s recommended to increase your Docker limits to at least 4GB of memory.

Setting everything up sounds like a job that would take weeks. To make it resilient and scale, probably a job that would take months. In these areas, Helm Charts can really shine. Due to the growing community, there might already be a Helm Chart that we can use.

### How to Deploy WordPress and MariaDB

There are different public hubs for Helm Charts. One of them is [artifacthub.io](https://artifacthub.io). We can search for “WordPress” and find an interesting [WordPress Chart](https://artifacthub.io/packages/helm/bitnami/wordpress).

On the right side, there is an install button. If you click it, you get clear instructions about what to do:

```shell
$ helm repo add bitnami https://charts.bitnami.com/bitnami

$ helm install my-wordpress bitnami/wordpress --version 10.1.4
```

You will also see some instructions that tell you how to access the admin interface and the admin password after installation.

Here is how you can get and decode the password for the **admin** user on Mac OS:

```shell
$ echo Username: user
$ echo Password: $(kubectl get secret --namespace default my-wordpress-3 -o jsonpath="{.data.wordpress-password}" | base64 --decode)

Username: user
Password: sZCa14VNXe
```

On windows, you can get the password for the **user** user in the powershell:

```powershell
$pw=kubectl get secret --namespace default my-wordpress -o jsonpath="{.data.wordpress-password}"
[System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($pw))
```

Our local development will be available at: [http://localhost](http://localhost).

Our admin interface will be available at: [https://localhost/admin](https://localhost/admin%20).

So we have everything to run it locally. But in production, we want to scale some parts of it to serve more and more visitors. We can scale the number of WordPress services. We also want to expose some health metrics like the usage of our CPU and memory.

We can [download the example configuration for production](https://raw.githubusercontent.com/bitnami/charts/master/bitnami/wordpress/values-production.yaml) from the maintainer of the WordPress Chart. The most important changes are:

```yaml
### Start 3 WordPress instances that will all receive
### requests from our visitors. A load-balancer will distribute calls
### to all containers evenly.
replicaCount: 3

### start a sidecar container that will expose metrics for your wordpress container
metrics:
enabled: true
image:
    registry: docker.io
    repository: bitnami/apache-exporter
    tag: 0.8.0-debian-10-r243
```

Let’s stop the default application:

```shell
$ helm delete my-wordpress

release "my-wordpress" uninstalled
```

### How to Start a Multi-instance WordPress and MariaDB Deployment

Deploy a new release using the production values:

```shell
$ helm install my-wordpress-prod bitnami/wordpress --version 10.1.4 -f values-production.yaml
```

This time, we have more containers running:

```shell
$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
my-wordpress-prod-5c9776c976-4bs6f   2/2     Running   0          103s
my-wordpress-prod-5c9776c976-9ssmr   2/2     Running   0          103s
my-wordpress-prod-5c9776c976-sfq84   2/2     Running   0          103s
my-wordpress-prod-mariadb-0          1/1     Running   0          103s

```

We see 4 lines: 1 line for MariaDB, and 3 lines for our actual WordPress pods.

A pod in Kubernetes is a group of containers. Each group contains 2 containers, one for WordPress and one for an exporter for Prometheus that exposes valuable metrics in a special format.

As in the default setup, we can [open localhost](http://localhost) and play with our WordPress application.

### How to Access Exposed Health Metrics

We can check the exposed health metrics by proxying to one of the running pods:

```shell
kubectl port-forward my-wordpress-prod-5c9776c976-sfq84 9117:9117
```

_Make sure to replace the pod-id with your own pod ID when you execute the port-forward command._

Now, we are connected to port 9117 of the WordPress Prometheus exporter and map the port to our local port 9117. Open [http://localhost:9117](http://localhost:9117/metrics) to check the output.

If you are not used to the Prometheus format, it might be a little bit confusing in the beginning. But it’s actually pretty easy to read. Each line without _‘#’_ contains a metric key and a value behind it:

```prometheus
apache_cpuload 1.2766
process_resident_memory_bytes 1.6441344e+07
```

If you are not used to such metrics, don't worry – you will get used to them quickly. You can Google each of the keys and find out what it means. After some time, you will identify what metrics are the most valuable for you and how they behave as soon as your containers receive more and more production traffic.

Let’s tidy up our setup by:

```shell
$ helm delete my-wordpress-prod

release "my-wordpress-prod" uninstalled
```

We touched on a lot of deployment areas and features. We deployed multiple WordPress instances and scaled it up to more containers for production. You could even go one step further and activate auto-scaling. Check out the documentation of the Helm Chart and play around with it!

### MariaDB Helm Chart

The persistence of the helm Chart for WordPress depends on MariaDB. It builds on another [Helm Chart for MariaDB](https://artifacthub.io/packages/helm/bitnami/mariadb) that you can configure and scale to your needs by, for example, starting multiple replicas.

The possibilities that you have when running containers in production using Kubernetes are enormous. The definition of the WordPress Chart is publicly available.

In the next section, we will create our own Helm Chart with a basic application to understand the fundamentals of creating a Helm Chart and to make a static container deployment more dynamic.

## How to Create a Template for Custom Applications

Helm adds a lot more flexibility to your Kubernetes deployment files. Kubernetes deployment files are static by their nature. This means, adjustments like

- desired container count,
- environment variables or
- CPU and memory limit

are not adjustable by using plain Kubernetes deployment files. Either you solve this by duplicating configuration files or you put placeholders in your Kubernetes deployment files that are replaced at deploy-time.

Both of these solutions require some additional work and will not scale well if you deploy a lot of applications with different variations.

But for sure, there is a smarter solution that is based on Helm that contains a lot of handy features from the Helm community. Let’s create a custom Chart for a blogging engine, this time for a NodeJS based blog called [ghost blog](https://ghost.org/).

### How to Start a Ghost Blog Using Docker

A simple instance can be started using pure Docker:

```shell
docker run --rm -p 2368:2368 --name my-ghost ghost
```

Our blog is available at: [http://localhost:2368](http://localhost:2368/).

Let's stop the instance to be able to launch another one using Kubernetes:

```shell
$ docker rm -f my-ghost

my-ghost
```

Now, we want to deploy the ghost blog with 2 instances in our Kubernetes cluster. Let’s set up a plain deployment first:

```yaml
# file 'application/deployment.yaml'

apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: 2
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost

          ports:
            - containerPort: 2368
```

and put a load balancer before it to be able to access our container and to distribute the traffic to both containers:

```yaml
# file 'application/service.yaml'

apiVersion: v1
kind: Service
metadata:
name: my-service-for-ghost-app
spec:
type: LoadBalancer
selector:
    app: ghost-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
```

We can now deploy both resources using kubectl:

```
$ kubectl apply -f ./appplication/deployment.yaml -f ./appplication/service.yaml

deployment.apps/ghost-app created
service/my-service-for-ghost-app created
```

The ghost application is now available via [http://localhost](http://localhost). Let's again stop the application:

```
$ kubectl delete -f ./appplication/deployment.yaml -f ./appplication/service.yaml

deployment.apps/ghost-app delete
service/my-service-for-ghost-app delete
```

So far so good, it works with plain Kubernetes. But what if we need different settings for different environments?

Imagine that we want to deploy it to multiple data centers in different stages (non-prod, prod). You will end up duplicating your Kubernetes files over and over again. It will be hell for maintenance. Instead of scripting a lot, we can leverage Helm.

Let’s create a new Helm Chart from scratch:

```shell
$ helm create my-ghost-app

Creating my-ghost-app
```

Helm created a bunch of files for you that are usually important for a production-ready service in Kubernetes. To concentrate on the most important parts, we can remove a lot of the created files. Let’s go through the only required files for this example.

We need a project file that is called Chart.yaml:

```yaml
# Chart.yaml

apiVersion: v2
name: my-ghost-app
description: A Helm chart for Kubernetes
type: application
version: 0.1.0
appVersion: 1.16.0
```

The deployment template file:

```yaml
# templates/deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: {{ .Values.replicaCount }}
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost
          ports:
            - containerPort: 2368
          env:
            - name: url
              {{- if .Values.prodUrlSchema }}
              value: http://{{ .Values.baseUrl }}
              {{- else }}
              value: http://{{ .Values.datacenter }}.non-prod.{{ .Values.baseUrl }}
              {{- end }}
```

It looks very similar to our plain Kubernetes file. Here, you can see different placeholders for the replica count, and an if-else condition for the environment variable called url. In the following files, we will see all the values defined.

The service template file:

```yaml
# templates/service.yaml

apiVersion: v1
kind: Service
metadata:
name: my-service-for-my-webapp
spec:
type: LoadBalancer
selector:
    app: ghost-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
```

Our Service configuration is completely static.

The values for the templates are the last missing parts of our Helm Chart. Most importantly, there is a default values file required called values.yaml:

```yaml
# values.yaml

replicaCount: 1
prodUrlSchema: false
datacenter: us-east
baseUrl: myapp.org
```

A Helm Chart should be able to run just by using the default values. Before you proceed, make sure that you have deleted:

- my-ghost-app/templates/tests/test-connection.yaml
- my-ghost-app/templates/serviceaccount.yaml
- my-ghost-app/templates/ingress.yaml
- my-ghost-app/templates/hpa.yaml
- my-ghost-app/templates/NOTES.txt.

We can get the final output that would be sent to Kubernetes by executing a “dry-run”:

```shell
$ helm template --debug my-ghost-app

install.go:159: [debug] Original chart version: ""
install.go:176: [debug] CHART PATH: /helm/my-ghost-app

---
# Source: my-ghost-app/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
name: my-service-for-my-webapp
spec:
type: LoadBalancer
selector:
    app: my-example-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
---
# Source: my-ghost-app/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: 1
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost
          ports:
            - containerPort: 2368
          env:
            - name: url
              value: us-east.non-prod.myapp.org

```

Helm inserted all the values and also set url to _`us-east.non-prod.myapp.org`_ because in the _`values.yaml`_, `prodUrlSchema` is set to false and the datacenter is set to us-east.

To get some more flexibility, we can define some override value files. Let’s define one for each datacenter:

```yaml
# values.us-east.yaml
datacenter: us-east
```

```yaml
# values.us-west.yaml
datacenter: us-west
```

and one for each stage:

```yaml
# values.nonprod.yaml
replicaCount: 1
prodUrlSchema: false
```

```yaml
# values.prod.yaml
replicaCount: 3
prodUrlSchema: true
```

We can now use Helm to combine them as we want and check the result again:

```shell
$ helm template --debug my-ghost-app -f my-ghost-app/values.nonprod.yaml  -f my-ghost-app/values.us-east.yaml

install.go:159: [debug] Original chart version: ""
install.go:176: [debug] CHART PATH: /helm/my-ghost-app

---
# Source: my-ghost-app/templates/service.yaml
# templates/service.yaml

apiVersion: v1
kind: Service
metadata:
name: my-service-for-my-webapp
spec:
type: LoadBalancer
selector:
    app: my-example-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
---
# Source: my-ghost-app/templates/deployment.yaml
# templates/deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: 1
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost
          ports:
            - containerPort: 2368
          env:
            - name: url
              value: http://us-east.non-prod.myapp.org

```

And for sure, we can do a final deployment:

```shell
$ helm install -f my-ghost-app/values.prod.yaml my-ghost-prod ./my-ghost-app/

NAME: my-ghost-prod
LAST DEPLOYED: Mon Dec 21 00:09:17 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

As before, our ghost blog is available via [http://localhost](http://localhost).

We can delete this deployment and deploy the application with us-east and non prod settings like this:

```shell
$ helm delete my-ghost-prod
release "my-ghost-prod" uninstalled

$ helm install -f my-ghost-app/values.nonprod.yaml -f my-ghost-app/values.us-east.yaml my-ghost-nonprod ./my-ghost-app
```

We finally clean up our Kubernetes deployment via Helm:

```shell
$ helm delete my-ghost-nonprod
```

So we can combine multiple override value files as we want. We can automate deployments in a flexible way that we need for many use-cases of deployment pipelines.

Especially for companies, this means defining Chart Skeletons once to ensure the required criteria are fulfilled. Later, you can copy them and adjust them to the needs of your application.

## Conclusion

The power of a great templating engine and the possibility of executing releases, upgrades, and rollbacks makes Helm great. On top of that comes the publicly available Helm Chart Hub that contains thousands of production-ready templates. This makes Helm a must-have tool in your toolbox if you work with Kubernetes on a bigger scale!

I hope you enjoyed this hands-on tutorial. Motivate yourself to Google around, check out other examples, deploy containers, connect them, and use them.

You will learn many cool features in the future that enable you to ship your application to production in an effortless, reusable and scalable way.

As always, I appreciate any feedback and comments. I hope you enjoyed the article. If you like it and feel the need for a round of applause, [follow me on Twitter](https://twitter.com/journerist).  I work at eBay Kleinanzeigen, one of the biggest classified companies globally. By the way, [we are hiring](https://jobs.ebayclassifiedsgroup.com/ebay-kleinanzeigen)!

References:

- [https://helm.sh/docs/chart\_template\_guide/getting\_started/](https://helm.sh/docs/chart_template_guide/getting_started/)

* * *

![Sebastian Sigl](http://www.freecodecamp.org/news/content/images/size/w100/2020/01/profile_photo_edited.png)[Sebastian Sigl](http://www.freecodecamp.org/news/author/sesigl/)

Software engineer who loves writing software and teaching people. Since 2018, I work for eBay Kleinanzeigen that is nowadays part of Adevinta.

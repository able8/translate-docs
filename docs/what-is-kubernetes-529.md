# What is Kubernetes?

If you know what is Kubernetes you may be interested in our previous blog “ [How Kubernetes Works](https://goglides.io/kubernetes-how-it-works/94/)“.

Designed and developed by Google, Kubernetes is an open-sourced container-orchestration system that helps in the automation of application deployment, scaling containerized applications, and monitoring such applications. Now maintained by the Cloud Native Computing Foundation, Kubernetes seeks to provide a platform for the automated deployment, scaling, and the operations of application containers across a cluster of hosts.

What Kubernetes does is that it provides you with a framework that allows you to run distributed systems smoothly and effectively. It looks after the scaling, looks after any application failovers, and immediately provides a deployment pattern.

The development of Kubernetes was influenced by Google’s work on its Borg System. And following the release of Kubernetes v1.0, Google partnered with the Linux Foundation and formed the Cloud Native Computing Foundation.

## Basic Kubernetes Terminologies

Before we learn what Kubernetes is and how it operates, let us first understand a few application program interfaces that are a part of Kubernetes.

1. **Node:** A node like a virtual machine is a machine entity or physical hardware. It acts as a worker instance within the Kubernetes Cluster. The nodes are responsible for performing the assigned tasks. Nodes are the physical infrastructure that allows your application to run on the Virtual Machine’s server.
2. **Pods:** Pods are a group of one or more containers that share network and storage with a Kubernetes configuration, which allows you to move containers within the cluster more efficiently. Each particular pod is assigned with an IP address, a hostname, IPC, which enable it to be accessed by other pods within a cluster.
3. **Service:** Service helps in separating the work definitions from the pods. Services are a sort of static pointers to the pods—services help in assuring that each pod receives an IP address.
4. **Replica Set:** The primary function of a replica set is to ensure that a given number of pods are running at a particular time. Replica Set is also a vital aspect of the Kubernetes autoscaling functionality. If a pod crashes or dies, the Replica Set makes sure a new one is created. In case there are too many pods, it cut down on them, and if there are few, then it creates new pods. Replica sets also make it much easier to target a particular set of pods.
5. **Kubelet:** A Kubelet is present inside each of the Nodes. Each node has a Kubelet. The primary responsibility of the Kubelet is to manage the nodes. It reads the container manifests and ensures that the particular containers are up and running.
6. **Kubectl:** A Kubectl is simply a command-line configuration interface for Kubernetes. It consists of an extensive list of available commands that can be used to manage Kubernetes.

## How Does Kubernetes Work?

When you first deploy Kubernetes, you get a Cluster. A Cluster consists of two parts, i.e., The Control Plane and Compute Machines. The control plane includes Master Nodes, and the Compute Machines consist of Worker Nodes. The worker nodes are responsible for running the Pods, which are made up of Containers. The Control Plane is responsible for managing the worker nodes as well as the pods within the cluster.

![kubernetes-diagram](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/01-kubernetes-diagram-1.png?resize=1024%2C636&ssl=1)kubernetes-diagram

[Source](https://www.redhat.com/cms/managed-files/kubernetes-diagram-2-824x437.png)

You can see the image above, Kubernetes operates on top of the operating system and then interacts with the pods of the container that are running on the nodes. The operator relays the command to the master node, which then distributes it to the secondary nodes. The node that is best suited for the particular task will be automatically decided, and the resources required for the task will be allocated to the particular pod required to fulfill the assigned work.

## What Services Does Kubernetes Provide You?

Some of the most basic features that Kubernetes provides you are:

1. Maintain and adapt containers across a wide range of hosts.
2. Self-Heal applications with Auto-Placement, Restart, Scaling, and Replication.
3. Makes sure the applications run as you intended it to run.
4. It helps you control the updates and application deployment.
5. Mount and add storage systems of your choice
6. Efficient use of hardware for maximum utilization of resources.

## Kubernetes Components & Architecture

Before studying the Kubernetes Components, let’s take a brief look at the principles and design that define Kubernetes.

The Kubernetes Cluster design is based on three simple principles, i.e.

1. Easy to use: The Cluster should be operable using minimum commands.
2. Extendable: It should be customizable.
3. Secure: Probably the most important, The Kubernetes Cluster should always follow the latest up-to-date security practices and principles.

Kubernetes consists of many components that all communicate with each other through the API server. As we talked about before, the Kubernetes cluster consists of The Control Plane, which holds the Master Node, and the Compute Machine, which holds the Worker Node.

![Components of Kubernetes](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/02-components-of-kubernetes.png?resize=1024%2C605&ssl=1)Components of Kubernetes

[Source:](https://d33wubrfki0l68.cloudfront.net/7016517375d10c702489167e704dcb99e570df85/7bb53/images/docs/components-of-kubernetes.png)

The above image shows the components inside the Kubernetes Control Plane and the components inside of the Compute Machine.

Now let’s take a closer look at the components.

## Control Plane Components

The Control Plane is the central part of the Kubernetes cluster. Inside the Control Plane is the Master Node, which holds all the components that control the Kubernetes Cluster. It makes the entire decision about the cluster. It also detects as well as response to any cluster events.

1. **kube-apiserver:** This is the front end of the Kubernetes Control Plane. It handles the internal as well as external requests. The kube-apiserver has been specifically designed to scale horizontally, which means that it is scaling up by deploying more instances. It also helps to communicate with all the other components inside the cluster.
2. **etcd:** The primary function of ‘etcd’ is to store the configuration information. The nodes can then use the information within the cluster. As the etcd contains sensitive information, it can only be accessed by the Kubernetes API server. It is a distributed key-value store responsible for implementing locks within the cluster.
3. **kube-scheduler:** A vital component of the Kubernetes, a kube-scheduler, is responsible for the distribution of the container and the workloads across multiple nodes. In simple terms, it allocates newly formed pods to the available nodes. The kube-scheduler is responsible for finding newly formed containers and assigning them to the available nodes.
4. **kube-controller-manager:** This is a component of the Master Node that is responsible for running the controller. The kube-controller-manager contains several controller functions compiled into a single binary to save time and reduce complexity.

The controller functions include:

- **Node controller:** It notices and responds whenever a particular node goes down.
- **Replication controller:** It is responsible for maintaining the correct number of pods.
- **Endpoints controller:** It is responsible for populating the endpoint objects.
- **Service accounts & token controller:** It is responsible for creating default accounts and API access tokens for the new namespaces.

1. **Cloud-controller-manager:** The cloud-controller-manager allows you to interact and link your cluster to your cloud provider’s API. It means that it allows you to talk to your cloud providers. The cloud-controller-manager also combines several different control functions into one.

The controller functions include:

- **Node Controller:** It is responsible for checking whether or not the node has been deleted from the cloud provider after it stops responding.
- **Route Controller:** It is responsible for setting up routes in the underlying cloud infrastructure.
- **Service Controller:** It is responsible for creating, updating, and deleting the cloud provider load balancers.

## Worker Node Components

The Master Node controls the worker node. The node components will run on every node while maintaining running pods as well as providing the Kubernetes runtime environment.

1. **kubelet:** We talked briefly about what kubelet is. A kubelet is present inside of the Node. Each node has a kubelet. The primary responsibility of the Kubelet is to manage the nodes. It reads the container manifests and ensures that the particular containers are up and running.
2. **kube-proxy:** A kube-proxy is a network proxy service that runs on each of the nodes within a cluster. It also helps the external host by making services available. It also assists in forwarding all the requests to the correct containers and is also capable of adequately performing first load balancing. A kube-proxy managed the pods within a node creates new containers’ health checkups, etc.
3. **container runtime:** A container runtime is responsible for running the containers. Kubernetes supports several different container runtimes, such as Docker, CRI-O, containers as well as any and all implementations of the Kubernetes Container Runtime Interface.

## Kubernetes & Docker

Docker is a tool that is used for building, distributing as well as running Docker containers. Docker provides the user with its clustering tools that can be used to arrange and schedule the containers on the clusters. Likewise, Kubernetes is a more extensive container orchestration system for Docker. It is meant for the effective coordination of the cluster of nodes efficiently. Although Kubernetes and Docker are different types of technologies, they coordinate well with one another and help in the management and the deployment of containers in a distributed architecture. When using Kubernetes with Docker, the automated system requests the Docker to do things such as launch a specified container, start and stop the containers, etc., in cases when it would have been done manually by an admin for all the containers.

## Why Do You Need Kubernetes?

To meet the demands of the ever-changing business world, you and your team must be able to adapt to the situation. You will be required to build new applications and services at a rapid pace.

An application requires multiple containers. The use of Kubernetes will allow you to build applications and services that will span across several different containers. It will also allow you to schedule those containers across clusters while also scaling those containers and managing the health of the containers over time. It will provide you with information and metrics regarding your containers and clusters.

The use of Kubernetes will assist in conserving resources more effectively and efficiently. It monitors all the clusters and makes decisions on where to effectively launch the containers with the available resources currently being utilized on the nodes. And if an application goes down, Kubernetes recovers it automatically.

## Installing & Getting Started with Kubernetes

The first thing that is required is for you to install Kubernetes on your hardware or one of the major cloud providers. Installing Kubernetes is quite tricky because there are many components involved. But there are plenty of tools, both open-source and paid solutions, in the market place, which makes the installation process more comfortable. There are many ways how you can install a Kubernetes cluster. For this tutorial, we will be using Minikube. For that, you will first need to start and run two things, i.e., Kubectl and Minikube.

Kubectl is a command-line interface (CLI) tool that will allow you to interact with the cluster.

Minikube is a binary that will deploy the cluster locally on your development machine.

With these tools, you can now start arranging your containerized applications to the cluster in a few short minutes. After you have installed Minikube, you can now run a single node cluster within your local machine.

To start the Minikube cluster, you can use the code.

```
minikube start

```

![minikube start command](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/minikube-start.png?resize=1024%2C264&ssl=1)minikube start

Likewise, if you want to interact with the Kubernetes Cluster, you will need to install Kubectl CLI. Once you installed it, you can use the following codes to interact with the Kubernetes cluster.

```
kubectl config get-contexts

```

![kubectl-config-get-contexts](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-config-get-contexts.png?resize=1024%2C136&ssl=1)

```
kubectl config set-contexts <context-name>

```

![kubectl-config-set-context-minikube](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-config-set-context-minikube.png?resize=1024%2C46&ssl=1)

```
kubectl config current-context

```

![kubectl-config-current-context](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-config-current-context.png?resize=982%2C76&ssl=1)

```
Kubectl config delete-context <context-name>

```

By now, you will have a local single Node Kubernetes cluster up and running on your machine. Now, if you are looking to deploy your first containerized application, then you run the following codes to Minikube. In this demonstration, we will be deploying a simple Hello World with an exposed endpoint on the Minikube IP address.

> Start with,

```
kubectl create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.10

```

![kubectl-create-deployment-hello-minikube](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-create-deployment-hello-minikube.png?resize=1024%2C56&ssl=1)

You will see that your deployment was successful. To view the deployment, you can use :

```
kubectl get deployments

```

![kubectl-get-deployment](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-get-deployment.png?resize=882%2C96&ssl=1)

After the deployment process, a Kubernetes Pod will have been created. To view the pods, you can use:

```
kubectl get pods

```

![kubectl-get-pods](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-get-pods.png?resize=996%2C92&ssl=1)

You will then need to expose the pod as a Kubernetes service before being able to hit the Hello World with an HTTP request from outside your cluster. To do so, you can use:

```
kubectl expose deployment hello-minikube --type=NodePort --port=8080

```

![kubectl-expose-deployment-hello-minikube](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-expose-deployment-hello-minikube.png?resize=1024%2C65&ssl=1)

The exposure will create a Kubernetes service. And to view the Service, you can use:

```
kubectl get services

```

![kubectl-get-services](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-get-services.png?resize=1024%2C111&ssl=1)

If you want to find out the URL that was used to access your containerized application, you can use:

```
minikube service hello-minikube --url

```

![minikube-serviec-hello-minikube-url](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/minikube-serviec-hello-minikube-url.png?resize=1024%2C104&ssl=1)

To test whether or not your exposed Service has reached the pod, you can curl the response from your terminal. To do that, you can use:

```
curl http://<minikube-ip>:<port>

```

![curl-url-command](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/curl-url-command.png?resize=1004%2C802&ssl=1)

The HTTP request has made via your Kubernetes Service. And once you check your logs, you will see the following,

![kubectl-logs-pods](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-logs-pods.png?resize=1024%2C268&ssl=1)

Now once you have followed the above steps, you should have a functioning Kubernetes pod and deployment that is running a simple Hello World application.

Likewise, if you are looking to start using Kubernetes, you can start by trying to build a Kubernetes Cluster. You can start by taking a look at different Managed Kubernetes offering provided by some top cloud providers.

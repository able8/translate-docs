# Using Telepresence 2 for Kubernetes debugging and local development

# 使用 Telepresence 2 进行 Kubernetes 调试和本地开发

[ Kostis Kapelonis ](https://codefresh.io/author/kostiscodefresh-io/)  · Apr 15, 2021

Telepresence 2 was [recently released](https://blog.getambassador.io/whats-new-in-telepresence-2-dbdff62b82d5) and (like Telepresence 1) it is a worthy addition to your Kubernetes  tool chest. Telepresence is one of those tools you cannot live without  after discovering how your daily workflow is improved.

Telepresence 2 [最近发布](https://blog.getambassador.io/whats-new-in-telepresence-2-dbdff62b82d5) 了并且（如 Telepresence 1）它是您的 Kubernetes 工具箱中的一个有价值的补充。 Telepresence 是在发现您的日常工作流程如何得到改进后，您离不开的工具之一。

So what is Telepresence? It is too hard to describe all the  functionalities of the tool in a single sentence, but for now I would  describe it as the “Kubernetes swiss army networking tool”. In this post we will see the major use cases that it covers but as time goes on and  more teams adopt Telepresence I am sure that more creative uses of it  will be discovered.

那么什么是Telepresence？用一句话描述该工具的所有功能太难了，但现在我将其描述为“Kubernetes 瑞士军队网络工具”。在这篇文章中，我们将看到它涵盖的主要用例，但随着时间的推移和越来越多的团队采用 Telepresence，我相信会发现它的更多创造性用途。

The major problems that Telepresence solves are:

  1. Kubernetes debugging and issue analysis. This is useful for both developers and Kubernetes operators
  2. Easy local development without a locally running Kubernetes cluster. This is mostly useful for developers
  3. Preview environment creation and real time collaboration in a team. This is great for large teams that have adopted Kubernetes.

Telepresence解决的主要问题是：
1、Kubernetes调试及问题分析。这对开发人员和 Kubernetes 操作员都很有用

2. 无需本地运行的 Kubernetes 集群即可轻松进行本地开发。这对开发人员最有用
3. 预览团队中的环境创建和实时协作。这对于采用 Kubernetes 的大型团队来说非常有用。

We will see these use cases in turn as each one builds on top of the previous one.

我们将依次看到这些用例，因为每个用例都建立在前一个用例之上。

Telepresence is a project in the growing discipline of Developer Experience, in which we have other tools such as [tilt.dev](https://tilt.dev), [garden.io](http://garden.io), [okteto ](http://okteto.com)and resources like [dex.dev](https://www.dex.dev/). It focuses specifically on helping developers work with containers and  Kubernetes (as opposed to tools that focus on administration and  management of clusters).

Telepresence 是不断发展的开发者体验学科中的一个项目，其中我们有其他工具，例如 [tilt.dev](https://tilt.dev)、[garden.io](http://garden.io)、 [okteto ](http://okteto.com) 和 [dex.dev](https://www.dex.dev/) 等资源。它专门帮助开发人员使用容器和 Kubernetes（而不是专注于集群管理的工具）。

By the way, if you are already familiar with Telepresence 1, then  Telepresence 2 is a complete rewrite (now in Go) with many more  improvements for reliability and extra features.

顺便说一下，如果您已经熟悉 Telepresence 1，那么 Telepresence 2 是完全重写的（现在在 Go 中），在可靠性和额外功能方面有更多改进。

## Beam yourself into your Kubernetes cluster (using Telepresence for Kubernetes debugging)

## 进入你的 Kubernetes 集群（使用 Telepresence 进行 Kubernetes 调试）

Let’s say that you are a Kubernetes operator. You are tasked with  deploying an application with many microservices (queue, auth, backend,  front-end etc). You deploy your manifests and while the pods seem to  start up ok, communications between the services are not working  correctly.

假设您是 Kubernetes 操作员。您的任务是部署具有许多微服务（队列、身份验证、后端、前端等）的应用程序。您部署了清单，虽然 pod 似乎可以正常启动，但服务之间的通信无法正常工作。

The front-end cannot seem to find the backend, the queue is  inaccessible from the backend and so on. You need a way to test  connectivity between services and how they respond. Your first impulse  would be to use the venerable [kubectl port-forward](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/) to make all these services available locally to your workstation. While this approach might work, it is not really what you want, because with  kubectl port-forward you only make a remote port available to a local  port.

前端好像找不到后端，后端无法访问队列等等。您需要一种方法来测试服务之间的连接性及其响应方式。您的第一个冲动是使用古老的 [kubectl port-forward](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/) 来制作所有这些服务本地可用到您的工作站。虽然这种方法可能有效，但它并不是您真正想要的，因为使用 kubectl port-forward 您只能将远程端口提供给本地端口。

What you really want is a way to see **INSIDE** the  cluster and see things from the perspective of the service that you are  interested in. You want to understand how the backend pod for example  can access other services from inside the cluster. Also if you want to  debug too many services at once, opening multiple terminals with kubectl exec or kubectl port-forward is a cumbersome process.

您真正想要的是一种查看 **INSIDE** 集群并从您感兴趣的服务的角度看待事物的方法。例如，您想了解后端 pod 如何从集群内部访问其他服务。另外如果你想一次调试太多服务，用 kubectl exec 或 kubectl port-forward 打开多个终端是一个繁琐的过程。

Wouldn’t it be great if you had a magic way that transferred your  laptop inside the Kubernetes cluster? This way you could run any kind of command directly in your terminal from the same perspective of a pod  and understand network connectivity with simple tools like curl, netcat, wget etc.

如果你有一种神奇的方式将你的笔记本电脑转移到 Kubernetes 集群中，那不是很好吗？通过这种方式，您可以从 pod 的同一角度直接在终端中运行任何类型的命令，并使用 curl、netcat、wget 等简单工具了解网络连接。

[![Your laptop inside the cluster](https://codefresh.io/wp-content/uploads/2021/04/magic.png)](https://codefresh.io/wp-content/uploads/2021/ 04/magic.png)Your laptop inside the cluster

Well this magic way is exactly what Telepresence does! As the name  suggests, telepresence moves your local workstation inside the  Kubernetes cluster and makes it look like your local environment is  inside a pod. All networking services and DNS names available to the pod are now available to your local workstation as well.

那么这种神奇的方式正是 Telepresence 所做的！顾名思义，远程呈现将您的本地工作站移动到 Kubernetes 集群内，并使您的本地环境看起来像是在一个 Pod 内。 Pod 可用的所有网络服务和 DNS 名称现在也可用于您的本地工作站。

Behind the scenes telepresence runs a small agent in the cluster and  routes traffic back and forth between a secure network tunnel. This is  also one of the big differences with Telepresence 1. Telepresence 2 has a global routing agent for the whole cluster and each intercepted service gets its own traffic manager (as we will see in the next use case)

在幕后，远程呈现在集群中运行一个小型代理，并在安全网络隧道之间来回路由流量。这也是与 Telepresence 1 的一大区别。 Telepresence 2 为整个集群提供了一个全局路由代理，每个被拦截的服务都有自己的流量管理器（我们将在下一个用例中看到）

[![安全隧道](https://codefresh.io/wp-content/uploads/2021/04/secure-tunner.png)](https://codefresh.io/wp-content/uploads/2021/04/secure-tunner.png) 安全隧道

The telepresence CLI modifies your local networking settings and  allows you to use your favorite tools on your local workstation as if  you were inside the cluster. 

Telepresence CLI 修改您的本地网络设置，并允许您在本地工作站上使用您喜欢的工具，就像您在集群中一样。

This scenario is even more powerful if you are a developer. Let’s say that you are adding a new feature to an existing application with many  microservices. Your tasks are

1. Inspecting existing tables on a MySQL database which is only accessible from within the cluster
 2. Loading test data to the mysql database with a script that you will develop
 3. Examining REST responses from another service that you are going to depend on
 4. Send messages to a queue that is also running inside the cluster

如果您是开发人员，这种情况会更加强大。假设您要向具有许多微服务的现有应用程序添加新功能。你的任务是

1. 检查 MySQL 数据库上的现有表，该数据库只能从集群内部访问
   2.使用您将开发的脚本将测试数据加载到mysql数据库
2. 检查来自您将依赖的另一个服务的 REST 响应
3. 将消息发送到也在集群内部运行的队列

Again, you could open 3-4 terminals with kubectl forward and have your tools connect to `localhost:port`.

同样，您可以使用 kubectl forward 打开 3-4 个终端，并将您的工具连接到 `localhost:port`。

Wouldn’t it be better to just open your Mysql administration tool and simply connect to `mysql:3306`?

直接打开你的Mysql管理工具，直接连接到`mysql:3306`不是更好吗？

This is exactly what telepresence does for you. If your database is  available in the “mysql” dns name inside the cluster, then by running  telepresence you make it available to your local workstation in the same manner. Your workstation essentially becomes part of the cluster.

这正是Telepresence为您所做的。如果您的数据库在集群内的“mysql”dns 名称中可用，那么通过运行远程呈现，您可以以相同的方式将其提供给本地工作站。您的工作站基本上成为集群的一部分。

## How to make your workstation part of the cluster

## 如何让你的工作站成为集群的一部分

This is the easiest way to get started with telepresence. [Download the telepresence cli](https://www.getambassador.io/docs/telepresence/latest/install/) (currently available for Mac and Linux) and run in your terminal (with kubectl access to your cluster)

这是开始使用远程呈现的最简单方法。 [下载远程呈现 cli](https://www.getambassador.io/docs/telepresence/latest/install/)（目前适用于 Mac 和 Linux）并在您的终端中运行（使用 kubectl 访问您的集群）

>telepresence connect Launching Telepresence Daemon v2.1.3 (api v3) Connecting to traffic manager... Connected to context mydemoAkscluster (https://mydemoakscluster-dns-8734f6ac.hcp.centralus.azmk8s.io:443)


Telepresence will install a small agent in your cluster and setup  local networking on your workstation to make any Kubernetes DNS names  available locally.

Telepresence 将在您的集群中安装一个小型代理并在您的工作站上设置本地网络，以使任何 Kubernetes DNS 名称在本地可用。

Now you can run your favorite tools like inside the cluster. If for  example your backend service is running with name “my-backend” on the  namespace “demo” at port 9999, you can simply write:

现在您可以像在集群内一样运行您喜欢的工具。例如，如果您的后端服务在端口 9999 的命名空间“demo”上以名称“my-backend”运行，您可以简单地编写：

curl http://my-backend.demo:9999

curl http://my-backend.demo:9999

And get back a response! It doesn’t get any easier than this. You can launch your favorite IDE and connect to debugging ports, your Database  administration tool and connect to your db, your internal dashboard for  your ingress controller (that is normally available only within the  cluster) and so on.

并得到回复！没有比这更容易的了。您可以启动您最喜欢的 IDE 并连接到调试端口、您的数据库管理工具并连接到您的数据库、您的入口控制器的内部仪表板（通常仅在集群内可用）等等。

By the way, this is one of the areas where telepresence 2 has seen  major improvements from the first version as the in-cluster agent allows for more reliable connections.

顺便说一下，这是 Telepresence 2 比第一个版本有了重大改进的领域之一，因为集群内代理允许更可靠的连接。

## Local Kubernetes development without a local Kubernetes cluster

##没有本地Kubernetes集群的本地Kubernetes开发

In the previous section we have seen how you can use Telepresence to  debug your application. So let’s say that you found the problem and you  want to fix it by making some code changes and redeploying the  container.

在上一节中，我们已经了解了如何使用 Telepresence 调试应用程序。因此，假设您发现了问题，并且希望通过更改一些代码并重新部署容器来修复它。

Normally you could use the usual dance of code, package image, push,  deploy to the cluster and that would work fine if you wish to spend your time by waiting for docker images to be built and pushed.

通常，您可以使用通常的代码、打包镜像、推送、部署到集群的方式，如果您希望通过等待构建和推送 docker 镜像来花时间，这会很好地工作。

But with telepresence you can skip this process completely! Remember  that we said that the communication tunnel between your workstation and  the cluster is two way. This means that you can simply launch your  application locally (outside of a cluster and outside of a container)  and simply tell telepresence to “intercept” or route all traffic to your local port.

但是通过远程呈现，您可以完全跳过这个过程！请记住，我们说过工作站和集群之间的通信隧道是双向的。这意味着您可以简单地在本地（集群外部和容器外部）启动您的应用程序，并简单地告诉远程呈现“拦截”或将所有流量路由到您的本地端口。

Here is how it works:

下面是它的工作原理：

[![Local Kubernetes development](https://codefresh.io/wp-content/uploads/2021/04/local-development.png)](https://codefresh.io/wp-content/uploads/2021/04/local-development.png)Local Kubernetes development

Telepresence installs a sidecar agent next to your existing  application. This agent captures all traffic requests that go in the  container and instead of sending them to the application inside the  cluster, it routes all traffic to your local workstation.

Telepresence 会在您现有的应用程序旁边安装一个 sidecar 代理。该代理捕获进入容器的所有流量请求，而不是将它们发送到集群内的应用程序，而是将所有流量路由到您的本地工作站。

Note that this is also a big difference between telepresence 1 and 2. Telepresence 1 used to replace the whole deployment in the cluster,  while 2 uses an agent (and both versions of the application local and  remote exist at the same time)

注意这也是Telepresence1和Telepresence2的一个很大的区别。Telepresence1用来代替集群中的整个部署，而2使用了一个代理（本地和远程两个版本的应用程序同时存在）

The end result is that you have an end-to-end hybrid workflow where:

 - The application that you are developing runs locally on your  workstation but thinks it is inside your cluster and can communicate  with all other services in a transparent way

 - All other services on the cluster (usually the services that are  used by your application) also think that they are talking to another  application inside the cluster, while in reality they are talking to an  application with your local workstation. 

最终结果是您拥有一个端到端的混合工作流程，其中：

- 您正在开发的应用程序在您的工作站本地运行，但认为它在您的集群内，并且可以以透明的方式与所有其他服务通信
- 集群上的所有其他服务（通常是您的应用程序使用的服务）也认为它们正在与集群内的另一个应用程序通信，而实际上它们正在与本地工作站的应用程序通信。

With this kind of setup, the development process is straightforward. You just make code changes locally and the application is instantly  updated. There is nothing to redeploy or repackage in a docker image. Any live-reload mechanism that your programming language supports can be used as is without any special changes. You can also launch debuggers,  tracers and other dev tools locally as before.

通过这种设置，开发过程很简单。您只需在本地更改代码，应用程序就会立即更新。在 docker 镜像中没有什么可以重新部署或重新打包的。您的编程语言支持的任何实时重新加载机制都可以按原样使用，无需任何特殊更改。您还可以像以前一样在本地启动调试器、跟踪器和其他开发工具。

Telepresence essentially cuts down the unnecessary parts of the [development loop](https://www.getambassador.io/docs/telepresence/latest/concepts/devloop/).

Telepresence 本质上减少了开发循环中不必要的部分。

[![Extra development steps](https://codefresh.io/wp-content/uploads/2021/04/broken-cycle.png)](https://codefresh.io/wp-content/uploads/2021/04/broken-cycle.png)Extra development steps



If you are accustomed to running a local Kubernetes cluster for  development (or use other tools that simply redeploy your image to a  remote cluster), you will already be familiar with the extra price that  comes with Kubernetes development. They are the steps shown in red in  the picture above. And Telepresence eliminates them!

如果您习惯于运行本地 Kubernetes 集群进行开发（或使用其他工具将您的镜像简单地重新部署到远程集群），那么您将已经熟悉 Kubernetes 开发带来的额外费用。它们是上图中红色显示的步骤。 Telepresence 消除了它们！

This is also another big difference between Telepresence and other  tools destined for Kubernetes development (such as Okteto, Tilt, and  Garden.io). With telepresence there is no code sync process, no smart  live reload mechanisms and no local build of any kind of image. The  application is running outside a Kubernetes cluster, on your laptop.

这也是 Telepresence 与其他专门用于 Kubernetes 开发的工具（如 Okteto、Tilt 和 Garden.io）之间的又一大区别。远程呈现没有代码同步过程，没有智能实时重新加载机制，也没有任何类型的图像的本地构建。该应用程序在 Kubernetes 集群之外，在您的笔记本电脑上运行。

The icing on the cake is that with this approach you no longer have  to maintain a different set of environment properties for your cluster  and for your local workstation. You can use a single property set for  all configurations and your application will work in the same way  regardless of where they run (your workstation or the cluster). Telepresence even saves for your all properties available inside the  cluster as we will see in the next section.

锦上添花的是，使用这种方法，您不再需要为集群和本地工作站维护一组不同的环境属性。您可以为所有配置使用单个属性集，并且您的应用程序将以相同的方式运行，无论它们在哪里运行（您的工作站或集群）。 Telepresence 甚至可以保存集群内所有可用的属性，我们将在下一节中看到。

## How to make your local application think it is inside the cluster

## 如何让你的本地应用程序认为它在集群内部

First launch your application locally and make sure that it exposes its port correctly (e.g. at `localhost:3000`). This is also the time to setup your hot-reload mechanism supported by  your programming language. Ideally each time that you make a change to  your editor, your local app should refresh automatically.

首先在本地启动您的应用程序并确保它正确公开其端口（例如在`localhost:3000`）。这也是设置您的编程语言支持的热重载机制的时候了。理想情况下，每次对编辑器进行更改时，您的本地应用程序都应自动刷新。

Note that you should launch your application on its own. No need for a local Kubernetes cluster or a docker image.

请注意，您应该自行启动您的应用程序。不需要本地 Kubernetes 集群或 docker 镜像。

Then run this command:
 telepresence intercept dataprocessingservice --port 3000 -n demo --env-file ~/example-service-intercept.env

然后运行这个命令：
远程呈现拦截数据处理服务 --port 3000 -n demo --env-file ~/example-service-intercept.env

This is taken from the [example application](https://www.getambassador.io/docs/telepresence/latest/quick-start/demo-node/) but feel free to change the command with our own service, port and kubernetes namespace.

这是从 [示例应用程序](https://www.getambassador.io/docs/telepresence/latest/quick-start/demo-node/) 中获取的，但您可以随意使用我们自己的服务、端口和 kubernetes 更改命令命名空间。

And that’s it! The agent running next to the real application in the  remote cluster will now intercept all requests to it and send them  locally to your workstation. The reverse is also true and all requests  made by your application with the cluster dns names will be routed to  the cluster.

就是这样！运行在远程集群中真实应用程序旁边的代理现在将拦截所有对它的请求，并将它们本地发送到您的工作站。反过来也是如此，您的应用程序使用集群 dns 名称发出的所有请求都将路由到集群。

You can now start developing as normal. Any code changes you make are visible on the cluster right away. No more waiting for docker builds or image pushes. And you don’t even need a local Kubernetes cluster.

您现在可以开始正常开发了。您所做的任何代码更改都会立即在集群上可见。不再等待 docker 构建或图像推送。而且您甚至不需要本地 Kubernetes 集群。

The .env file can be used by your IDE or your scripts to replicate  the values from all the properties available from the cluster to your  local workstation.

您的 IDE 或脚本可以使用 .env 文件将值从集群中所有可用的属性复制到您的本地工作站。

## Kubernetes collaboration in a team environment

## Kubernetes 在团队环境中的协作

One of the most important differences between Telepresence 1 and 2 is that Telepresence no longer swaps the deployment inside the cluster (a  method also followed by Okteto and similar tools) but instead installs a routing agent next to the existing service that takes care of all  network communication.

Telepresence 1 和 2 之间最重要的区别之一是 Telepresence 不再交换集群内部的部署（Okteto 和类似工具也遵循这种方法）而是在现有服务旁边安装一个路由代理来处理所有网络沟通。

The added advantage of this approach is the capability to add some  routing behavior in the agent. In the previous use case we have seen how the agent can route all calls to your local workstation. But we can be a bit smarter and instruct the agent to route only some calls to your  workstation while leaving the rest of the calls unaffected (so that they hit the original application inside the cluster).

这种方法的附加优势是能够在代理中添加一些路由行为。在前面的用例中，我们已经看到代理如何将所有呼叫路由到您的本地工作站。但是我们可以更聪明一点，并指示代理仅将一些呼叫路由到您的工作站，而其余呼叫不受影响（以便它们访问集群内的原始应用程序）。

And this exactly what Telepresence 2 offers in the form of preview environments! Here is how it works

这正是 Telepresence 2 以预览环境的形式提供的！下面是它的工作原理

[![Preview routing](https://codefresh.io/wp-content/uploads/2021/04/preview-routing.png)](https://codefresh.io/wp-content/uploads/2021/04/preview-routing.png)Preview routing



Ambassador offers a free service that creates preview URLs for you  anytime you do an intercept in a service. If you enable it (it is  completely optional), you will get a public URL every time you intercept an existing application in the cluster. 
大使提供了一项免费服务，可以在您在服务中进行拦截时为您创建预览 URL。如果启用它（完全可选），则每次拦截集群中的现有应用程序时都会获得一个公共 URL。

Then if somebody visits your preview URL, the agent will send all  their traffic to your workstation. If somebody else visits your service  NOT from the preview URL, then all of their traffic will go to the  existing application that runs inside the cluster.

然后，如果有人访问您的预览 URL，代理会将所有流量发送到您的工作站。如果其他人不是从预览 URL 访问您的服务，那么他们的所有流量都将转到在集群内运行的现有应用程序。

This allows you to quickly share your local environment with your team. Think of it as [ngrok ](https://ngrok.com/)for kubernetes. The beauty here is the fact that both applications are  running at the same time. This means that you are free to debug and  develop your own application for fixing something, while all existing  users of your application are still connecting to the production version oblivious to your changes.

这使您可以快速与您的团队共享您的本地环境。将其视为 kubernetes 的 [ngrok ](https://ngrok.com/)。这里的美妙之处在于两个应用程序同时运行。这意味着您可以自由地调试和开发自己的应用程序来修复某些问题，而应用程序的所有现有用户仍在连接到生产版本，而不会注意到您的更改。

Once you are ready to publish your changes to everybody you can of  course commit your code and let your CI/CD solution update all live  traffic. But during your development time there is no better way to work in a production like environment without actually affecting production.

一旦您准备好向所有人发布您的更改，您当然可以提交您的代码并让您的 CI/CD 解决方案更新所有实时流量。但是在您的开发期间，没有更好的方法可以在不实际影响生产的情况下在类似生产的环境中工作。

I see this use case is very important for teams that want to quickly  test hot-fixes without actually deploying anything but even basic  collaboration with another developer becomes very simple if you can  quickly share your feature with them and quickly iterate on it.

我认为这个用例对于想要在不实际部署任何东西的情况下快速测试修补程序的团队非常重要，但如果您可以快速与他们分享您的功能并快速迭代，即使与其他开发人员的基本协作也会变得非常简单。

Note that this feature requires an ingress in your cluster. And it  can be any kind of compliant ingress (and not just the Ambassador API  gateway/edge stack).

请注意，此功能需要您的集群中有一个入口。它可以是任何类型的合规入口（而不仅仅是大使 API 网关/边缘堆栈）。

The Ambassador Cloud dashboard also provides a nice UI for managing your preview environments:

Ambassador Cloud 仪表板还提供了一个很好的 UI 来管理您的预览环境：

[![Preview dashboard](https://codefresh.io/wp-content/uploads/2021/04/preview-env.png)](https://codefresh.io/wp-content/uploads/2021/04/preview-env.png)Preview dashboard

You can also decide if you want your environment to be accessible  only from the same GitHub/GitLab organisation as you or any other user.

您还可以决定是否希望您的环境只能从与您或任何其他用户相同的 GitHub/GitLab 组织访问。

## How to use preview environments with Telepresence and Ambassador cloud

## 如何在Telepresence 和 Ambassador cloud 使用预览环境

Creating a preview environment is not a separate command on its own. You use the normal intercept command as before but you need to login  into ambassador cloud first

创建预览环境本身并不是一个单独的命令。您像以前一样使用普通的拦截命令，但您需要先登录大使云

telepresence login

telepresence intercept dataprocessingservice --port 3000 -n demo

The first command will open the UI and ask you to signup with  Github/Google/Gitlab. Then the second command will intercept the service as before but also print out the URL of the preview environment.

第一个命令将打开 UI 并要求您注册 Github/Google/Gitlab。然后第二个命令将像以前一样拦截服务，但也会打印出预览环境的 URL。

You can share the preview environment via email/slack/messaging with  your colleagues. It is interesting to notice that Telepresence adds an  extra header to the traffic request when it works in intercept mode. You can use this header yourself in the application to make your  application behave differently according to the source of the request  (e.g. enable special debugging support if the request comes from a  preview environment).

您可以通过电子邮件/松弛/消息与您的同事共享预览环境。有趣的是，Telepresence 在拦截模式下工作时，会为流量请求添加一个额外的标头。您可以在应用程序中自己使用此标头，使您的应用程序根据请求的来源表现出不同的行为（例如，如果请求来自预览环境，则启用特殊调试支持）。

## Comparison with other tools

##与其他工具的比较

The biggest competitor to Telepresence 2 right now is the previous  version. There are some other tools that share some of the same  functionality.

Telepresence 2 目前最大的竞争对手是之前的版本。还有一些其他工具共享一些相同的功能。

##### Comparison with kubectl port-forward:

##### 与 kubectl port-forward 的比较：

The port-forward command is very simplistic as it only works on a  single port for a single service and it is a one way connection. It is  great for quick access to the cluster but Telepresence has many more  features. You can still use kubectl port forward for adhoc connection  issues, but for application development there are many better choices.

port-forward 命令非常简单，因为它仅适用于单个服务的单个端口，并且是单向连接。它非常适合快速访问集群，但 Telepresence 具有更多功能。您仍然可以使用 kubectl port forward 来解决临时连接问题，但对于应用程序开发，有很多更好的选择。

##### Comparison with kubefwd:

##### 与kubefwd 的比较：

[Kubefwd ](https://github.com/txn2/kubefwd)works similar  to Telepresence by making your local environment think it is inside the  cluster. The networking tunnel is one direction only. Telepresence is  much smarter as it also makes the other cluster applications think that  your local app is inside the same cluster. So with Kubefwd you only get  50% of what basic Telepresence offers. Telepresence also has [volume mounting support](https://www.getambassador.io/docs/telepresence/latest/reference/volume/) for more advanced scenarios.

[Kubefwd ](https://github.com/txn2/kubefwd) 的工作原理与 Telepresence 类似，让您的本地环境认为它在集群内部。网络隧道只有一个方向。 Telepresence 更智能，因为它还使其他集群应用程序认为您的本地应用程序在同一个集群内。因此，使用 Kubefwd，您只能获得基本 Telepresence 提供的 50%。 Telepresence 还有【卷挂载支持】（https://www.getambassador.io/docs/telepresence/latest/reference/volume/），适用于更高级的场景。

##### Comparison with Telepresence 1 

##### 与 Telepresence 1 的比较

Telepresence 2 is improved in every aspect compared to the previous version.The [networking architecture](https://www.getambassador.io/docs/telepresence/latest/reference/architecture/) is now completely redesigned. There is a global traffic manager on each cluster and each intercepted service has its own sidecar container.The  new routing agent is great for reliability and Telepresence 2 will work  better with spotty network connections. The ability to use the sidecar  router instead of fully swapping the deployment comes with its own  advantages (the preview environments). Also Telepresence 2 is a single  binary written in Go making installation much easier (Telepresence 1 was a multi file Python application)

Telepresence 2 相比之前的版本在各个方面都有所改进。 [网络架构](https://www.getambassador.io/docs/telepresence/latest/reference/architecture/) 现在完全重新设计。每个集群上都有一个全局流量管理器，每个被拦截的服务都有自己的 sidecar 容器。新的路由代理非常可靠，Telepresence 2 可以更好地处理不稳定的网络连接。使用 sidecar 路由器而不是完全交换部署的能力有其自身的优势（预览环境）。 Telepresence 2 也是一个用 Go 编写的二进制文件，使安装更容易（Telepresence 1 是一个多文件 Python 应用程序）

## Combining Telepresence with outer loop tools

## 将 Telepresence 与外环工具相结合

Telepresence is great for handling the so-called [inner-loop of development](https://www.getambassador.io/docs/telepresence/latest/concepts/devloop/). That is the part where you as a developer write your code and test  right away on your local machine. For this part of the process you want  the quickest feedback possible.

Telepresence 非常适合处理所谓的[开发内循环](https://www.getambassador.io/docs/telepresence/latest/concepts/devloop/)。这是您作为开发人员编写代码并立即在本地机器上测试的部分。对于流程的这一部分，您希望获得尽可能快的反馈。

However once you complete your feature, you should still deploy your  application to a cluster to verify its behavior and avoid the dreaded  “works on my machine” issues. Telepresence is great for the inner  development loop, but you need to remember that your application runs on your local workstation outside of a container and outside of a cluster.

然而，一旦你完成了你的功能，你仍然应该将你的应用程序部署到一个集群来验证它的行为并避免可怕的“在我的机器上工作”的问题。 Telepresence 非常适合内部开发循环，但您需要记住，您的应用程序在容器外部和集群外部的本地工作站上运行。

After you feel comfortable with your feature you can complete your  workflow with tools that address the outer development loop in the sense that they actually gather all your application dependencies in a  container and then deploy your application to a Kubernetes cluster.

在您对您的功能感到满意后，您可以使用解决外部开发循环的工具来完成您的工作流程，因为它们实际上将您的所有应用程序依赖项收集到一个容器中，然后将您的应用程序部署到 Kubernetes 集群。

We have already explored such tools in our previous blog posts. Popular tools on this category are:

 - [Okteto](https://okteto.com/)
 - [Tilt.dev](https://Tilt.dev)
 - [Garden.io](http://Garden.io)
 - [Skaffold](https://skaffold.dev/)

我们已经在之前的博文中探索过此类工具。此类别中的流行工具有：

- [Okteto](https://okteto.com/)
- [Tilt.dev](https://Tilt.dev)
- [Garden.io](http://Garden.io)
- [Skaffold](https://skaffold.dev/)

[Okteto ](https://okteto.com/)has a powerful sync process  that creates a development environment inside your cluster and runs the  application there. If you use specific Kubernetes features such as GPU  nodes or have complex networking and security requirements (IAM roles,  specific identities) then Okteto can help you quickly see how your  application behaves in the real cluster. Okteto also comes with the  Okteto Cloud/Okteto Enterprise option that has additional features in  the context of a team (e.g. namespace isolation, credential management,  build service, deploy from git, the Okteto Registry). See our [review of Okteto](https://codefresh.io/kubernetes-tutorial/okteto/) for more details.

[Okteto ](https://okteto.com/) 有一个强大的同步过程，可以在你的集群内创建一个开发环境并在那里运行应用程序。如果您使用特定的 Kubernetes 功能（例如 GPU 节点）或具有复杂的网络和安全要求（IAM 角色、特定身份），那么 Okteto 可以帮助您快速了解您的应用程序在真实集群中的行为。 Okteto 还附带 Okteto Cloud/Okteto Enterprise 选项，该选项在团队上下文中具有附加功能（例如命名空间隔离、凭证管理、构建服务、从 git 部署、Okteto 注册表）。有关更多详细信息，请参阅我们的 [Okteto 评论](https://codefresh.io/kubernetes-tutorial/okteto/)。

[Tilt.dev ](https://tilt.dev/)is another service for local Kubernetes development. Its main strength is its innovative UI that  groups all related microservices from an application plus any custom  resources that you want to create. Tilt has many different options for [local, hybrid or fully remote development](https://docs.tilt.dev/local_vs_remote.html). In the case of local development you still need [a local Kubernetes cluster](https://docs.tilt.dev/choosing_clusters.html).

[Tilt.dev ](https://tilt.dev/) 是本地 Kubernetes 开发的另一个服务。它的主要优势在于其创新的 UI，可将应用程序中的所有相关微服务以及您想要创建的任何自定义资源进行分组。 Tilt 为 [本地、混合或完全远程开发](https://docs.tilt.dev/local_vs_remote.html) 提供了许多不同的选项。在本地开发的情况下，您仍然需要[本地 Kubernetes 集群](https://docs.tilt.dev/choosing_clusters.html)。

The GUI of Tilt is specifically designed for applications with a  large number of microservices and is great for making changes to more  than one service at once. Tilt is very extensible, making it easy to  adapt to whatever setup the user has (as opposed to the user having to  adapt to the tool), and has a growing extensions ecosystem with many  third-party contributors (https://github.com /tilt-dev/tilt-extensions/). See our [review of tilt.dev](https://codefresh.io/kubernetes-tutorial/local-kubernetes-development-tilt-dev/) for more details. 
Tilt 的 GUI 是专门为具有大量微服务的应用程序设计的，非常适合一次对多个服务进行更改。 Tilt 具有很强的可扩展性，可以轻松适应用户的任何设置（而不是用户必须适应该工具），并且拥有越来越多的第三方贡献者的扩展生态系统 (https://github.com) /tilt-dev/tilt-extensions/)。有关更多详细信息，请参阅我们的 [tilt.dev 回顾](https://codefresh.io/kubernetes-tutorial/local-kubernetes-development-tilt-dev/)。
[Garden.io](https://garden.io/) is a tool designed for the full software lifecycle (and not just deployments). Its main appeal is  that it creates a model of your application dependencies [with great visualization](https://docs.garden.io/basics/how-garden-works) giving you great insights into what needs to be updated every time a change is made. Garden attempts to model the full software workflow  including [testing](https://docs.garden.io/using-garden/tests). You also have the ability to create your own tasks to improve your  daily workflow. This means that you can set up Garden as a mini CI  workflow (that you can call from your real CI service or Garden  Enterprise) allowing you to unify the way a developer works with your CI pipelines. One of the strong points of the approach is that proper  tests and pipelines can be run **before** pushing to git, which makes integration/e2e testing and debugging CI/CD issues much faster.

[Garden.io](https://garden.io/) 是为整个软件生命周期（而不仅仅是部署）设计的工具。它的主要吸引力在于它创建了一个应用程序依赖关系模型[具有出色的可视化](https://docs.garden.io/basics/how-garden-works)，让您深入了解每次需要更新的内容做出改变。 Garden 尝试对包括 [testing](https://docs.garden.io/using-garden/tests) 在内的完整软件工作流程进行建模。您还可以创建自己的任务来改进日常工作流程。这意味着您可以将 Garden 设置为迷你 CI 工作流（您可以从真正的 CI 服务或 Garden Enterprise 调用），从而统一开发人员使用 CI 管道的方式。该方法的优点之一是，可以在**推送到 git 之前**运行适当的测试和管道，这使得集成/e2e 测试和调试 CI/CD 问题的速度更快。

Garden also gives a lot of importance to extensibility. Even  Kubernetes support is actually created as a plugin/extensions and there  is already a set of providers for other platforms. Like Tilt.dev you can choose where the packaging/deployment takes place (on a [local ](https://docs.garden.io/guides/local-kubernetes)or [remote kubernetes cluster](https://docs. garden.io/guides/in-cluster-building)) and of course supports [live-reload capabilities](https://docs.garden.io/guides/hot-reload). See our [review of garden.io](https://codefresh.io/howtos/local-k8s-draft-skaffold-garden/) for more details.

Garden 也非常重视可扩展性。甚至 Kubernetes 支持实际上也是作为插件/扩展创建的，并且已经有一组用于其他平台的提供程序。像 Tilt.dev 一样，您可以选择打包/部署发生的位置（在 [本地](https://docs.garden.io/guides/local-kubernetes) 或 [远程 kubernetes 集群](https://docs. Garden.io/guides/in-cluster-building))，当然也支持 [实时重新加载功能](https://docs.garden.io/guides/hot-reload)。有关更多详细信息，请参阅我们的 [garden.io 评论](https://codefresh.io/howtos/local-k8s-draft-skaffold-garden/)。

[Skaffold ](https://skaffold.dev/)is a tool for local  Kubernetes development. It contains an opinionated workflow (that can  work the same on your local workstation or within a CI pipeline) and has built-in integration with many popular and not so popular build tools  such as [Bazel](https://bazel.build/), [Jib](https://github.com/GoogleContainerTools/jib)and [buildpacks](https://buildpacks.io/). See our [review of skaffold](https://codefresh.io/howtos/local-k8s-draft-skaffold-garden/) for more details.

[Skaffold ](https://skaffold.dev/)是Kubernetes本地开发的工具。它包含一个自以为是的工作流程（可以在您的本地工作站或 CI 管道中工作），并且与许多流行和不那么流行的构建工具（例如 [Bazel]（https://bazel.build/））进行了内置集成)、[Jib ](https://github.com/GoogleContainerTools/jib) 和 [buildpacks](https://buildpacks.io/)。有关更多详细信息，请参阅我们的 [skaffold 评论](https://codefresh.io/howtos/local-k8s-draft-skaffold-garden/)。

## Adopting Telepresence in your team

## 在您的团队中采用Telepresence

So is Telepresence worth having in your tool chest?

那么您的工具箱中是否值得拥有 Telepresence？

First of all if you are a Kubernetes Administrator or system operator the answer is undeniably yes. All the other solutions are targeted  strictly at developers. But if all you want is to see what endpoints are available in the cluster and want to run an adhoc bash/python script to do something that touches multiple services, Telepresence is much more  powerful that kubectl port-forward.

首先，如果您是 Kubernetes 管理员或系统操作员，答案无疑是肯定的。所有其他解决方案都严格针对开发人员。但是，如果您只想查看集群中可用的端点，并希望运行临时 bash/python 脚本来执行涉及多个服务的操作，那么 Telepresence 比 kubectl port-forward 强大得多。

If you are a developer, Telepresence works great for local  development since it has the fastest code loop ever (just code). No file syncing, no docker rebuild, no live container update and no local  cluster is needed. You need to take into account the fact that your  application runs on your laptop OUTSIDE of a kubernetes cluster,  accepting the risk of the dreaded “works on my machine” effect.

如果您是开发人员，Telepresence 非常适合本地开发，因为它拥有最快的代码循环（仅代码）。无需文件同步，无需 docker 重建，无需实时容器更新，也无需本地集群。您需要考虑到您的应用程序在 kubernetes 集群之外的笔记本电脑上运行的事实，接受可怕的“在我的机器上工作”效应的风险。

The preview feature of Telepresence is a nice addition, especially  for really critical production hotfixes (where you want to develop a  hotfix while your live users are unaffected).

Telepresence 的预览功能是一个很好的补充，特别是对于真正关键的生产修补程序（您希望在不影响实时用户的情况下开发修补程序）。

To get the full benefits of local Kubernetes development you should  couple Telepresence with another tool that actually deploys the  application in the cluster and get the best of both worlds (fast local  development, verification that the application will run the same in a  Kubernetes cluster).


为了获得本地 Kubernetes 开发的全部好处，您应该将 Telepresence 与另一个实际在集群中部署应用程序的工具结合起来，并获得两全其美（快速本地开发，验证应用程序将在 Kubernetes 集群中运行相同）。


### Kostis Kapelonis

Kostis is a software engineer/technical-writer  dual class character. He lives and breathes automation, good testing  practices and stress-free deployments. 
Kostis 是一名软件工程师/技术作家双重角色。他生活和呼吸自动化、良好的测试实践和无压力的部署。
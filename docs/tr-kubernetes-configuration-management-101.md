# What Configuration Management is and Why You Should Implement it in Your infrastructure

# 什么是配置管理以及为什么要在基础设施中实施它

Wikipedia defines [Configuration Management](https://en.wikipedia.org/wiki/Configuration_management) as “a [systems engineering](https://en.wikipedia.org/wiki/Systems_engineering) process for establishing and maintaining consistency of a product's performance, functional, and physical attributes with its requirements, design, and operational information throughout its life.”

维基百科将[配置管理](https://en.wikipedia.org/wiki/Configuration_management)定义为“一个[系统工程](https://en.wikipedia.org/wiki/Systems_engineering)过程，用于建立和维护产品在整个生命周期内的性能、功能和物理属性及其要求、设计和操作信息。”

Configuration Management has two main aspects:

1. Bringing up environments quickly from one or more templates.
2. Being able to replicate those environments easily by changing a few parameters.

配置管理有两个主要方面：

1. 从一个或多个模板快速调出环境。
2. 能够通过更改一些参数轻松复制这些环境。

Many Configuration Management tools exist and address the above requirements. For example, Ansible. Let’s say that you need to deploy WordPress to a virtual (or even physical) machine. All that you need to do is search Ansible Galaxy for a role that deploys WordPress and change its parameters according to your environment needs. There are other tools as well, and each of them has its deployable unit and a repository — for example, Puppet and its Puppet Forge and Chef and its Supermarket.

存在许多配置管理工具并满足上述要求。例如，Ansible。假设您需要将 WordPress 部署到虚拟（甚至物理）机器上。您需要做的就是在 Ansible Galaxy 中搜索部署 WordPress 并根据您的环境需要更改其参数的角色。还有其他工具，每个工具都有可部署的单元和存储库——例如，Puppet 及其 Puppet Forge 和 Chef 及其超市。

![configrations 1](https://www.magalix.com/hs-fs/hubfs/configrations%201.png?width=720&name=configrations%201.png)

You might be thinking that Kubernetes does not require configuration management in that sense. After all, you add all the requirements to the Docker image and let Kubernetes do the rest. But, Kubrenetes is not about images and containers only, you have other components that are built around containers like [Pods](http://www.magalix.com/blog/kubernetes-pods-101-the-cluster-sailors),[Services](http://www.magalix.com/blog/kubernetes-services-101-the-pods-interfaces), [Ingresses](https://kubernetes.io/docs/concepts/services-networking/ingress/), [configMaps](http://www.magalix.com/blog/the-configmap-pattern),[Secrets](http://www.magalix.com/blog/kubernetes-secrets-101), [ Volumes](http://www.magalix.com/blog/tag/storage), etc. So, while a traditional Configuration Management tool is only concerned with what runs inside the instance, a Kubernetes Configuration Management tool is responsible for what gets deployed to the instance as well as the infrastructure that surrounds this deployment and making it scalable.

您可能会认为 Kubernetes 在这个意义上不需要配置管理。毕竟，您将所有要求添加到 Docker 映像中，剩下的交给 Kubernetes 来完成。但是，Kubrenetes 不仅仅与图像和容器有关，您还有其他围绕容器构建的组件，例如 [Pods](http://www.magalix.com/blog/kubernetes-pods-101-the-cluster-sailors)，[服务](http://www.magalix.com/blog/kubernetes-services-101-the-pods-interfaces), [Ingresses](https://kubernetes.io/docs/concepts/services-networking/ingress/), [configMaps](http://www.magalix.com/blog/the-configmap-pattern),[Secrets](http://www.magalix.com/blog/kubernetes-secrets-101), [ Volumes](http://www.magalix.com/blog/tag/storage) 等。因此，传统的配置管理工具只关心实例内部运行的内容，而 Kubernetes 配置管理工具负责获取的内容部署到实例以及围绕此部署并使其可扩展的基础设施。

In this article, we explore a popular [Kubernetes](http://www.magalix.com/blog/kubernetes-101-concepts-and-why-it-matters) tool that can be used for Configuration Management: Helm.

在本文中，我们探索了一个可用于配置管理的流行 [Kubernetes](http://www.magalix.com/blog/kubernetes-101-concepts-and-why-it-matters) 工具：Helm。

# Helm



Helm can be thought of as the Package Manager for [Kubernetes](http://www.magalix.com/blog/kubernetes-101-concepts-and-why-it-matters). If Ubuntu uses apt, Centos uses yum, and macOS uses brew, then, in the same sense, Kubernetes uses Helm. With a single command, you can have a complete Kubernetes infrastructure up and running.

Helm 可以被认为是 [Kubernetes](http://www.magalix.com/blog/kubernetes-101-concepts-and-why-it-matters) 的包管理器。如果 Ubuntu 使用 apt，Centos 使用 yum，而 macOS 使用 brew，那么在同样的意义上，Kubernetes 使用 Helm。只需一个命令，您就可以启动并运行完整的 Kubernetes 基础设施。

![configrations 2](https://www.magalix.com/hs-fs/hubfs/configrations%202.png?width=720&name=configrations%202.png)


## Helm Installation

Helm has many installation options depending on your operating system. You can refer to [this page](https://helm.sh/docs/using_helm/#installing-helm) for the installation instructions specific to your OS.

Helm 有许多安装选项，具体取决于您的操作系统。您可以参考 [此页面](https://helm.sh/docs/using_helm/#installing-helm) 了解特定于您的操作系统的安装说明。

## Giving Helm a Test Drive

## 给 Helm 试驾

By default, Helm applies your commands to the default cluster context that your kubectl command uses. You can view the current context by issuing:

```yaml
kubectl config current-context
```


Helm operates in a very simple way:

- The helm tool works on the client computer.
- The commands issued by helm go to the Tiller tool, which runs on the server side (inside Kubernetes).
- Tiller apples the Chart to the cluster.

Helm 以非常简单的方式运行：

- helm 工具适用于客户端计算机。
- helm 发出的命令转到 Tiller 工具，该工具运行在服务器端（在 Kubernetes 内部）。
- Tiller 将 Chart 应用到集群中。

Let’s give the above steps a test drive by installing the MySQL database on our cluster.

让我们通过在集群上安装 MySQL 数据库来测试上述步骤。

The first step we need to do is install Tiller on the cluster:

```yaml
helm init --history-max 200
```


The above command will initialize helm locally (you can see several files get created in the command output) and also deploys Tiller to the cluster. Notice that we set the --history-max to an arbitrary number. This is recommended because otherwise Helm and Tiller will keep an indefinite amount of objects like configMaps, which may add unnecessary overhead.

上述命令将在本地初始化 helm（您可以在命令输出中看到创建了多个文件）并将 Tiller 部署到集群。请注意，我们将 --history-max 设置为任意数字。建议这样做，否则 Helm 和 Tiller 将保留无限量的对象，如 configMaps，这可能会增加不必要的开销。

Like it’s recommended to run apt update before running apt install on Ubuntu systems, we follow a similar practice with Kubernetes:

```yaml
helm install stable/mysql
```




The word stable here refers to the name of the Kubernetes repository that Helm will use when installing the Chart. You should immediately see an output similar to the following:

```yaml
NAME:   undercooked-saola
LAST DEPLOYED: Wed Jul 24 22:17:04 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                          DATA  AGE
undercooked-saola-mysql-test  1     0s

==> v1/PersistentVolumeClaim
NAME                     STATUS   VOLUME    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
undercooked-saola-mysql  Pending  hostpath  0s

==> v1/Pod(related)
NAME                                    READY  STATUS   RESTARTS  AGE
undercooked-saola-mysql-fc896fb6-zwfn6  0/1    Pending  0         0s

==> v1/Secret
NAME                     TYPE    DATA  AGE
undercooked-saola-mysql  Opaque  2     0s

==> v1/Service
NAME                     TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)   AGE
undercooked-saola-mysql  ClusterIP  10.101.209.21         3306/TCP  0s

==> v1beta1/Deployment
NAME                     READY  UP-TO-DATE  AVAILABLE  AGE
undercooked-saola-mysql  0/1    1           0          0s

NOTES:
MySQL can be accessed via port 3306 on the following DNS name from within your cluster:
undercooked-saola-mysql.default.svc.cluster.local

To get your root password run:

    MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default undercooked-saola-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)

To connect to your database:

1. Run an Ubuntu pod that you can use as a client:

    kubectl run -i --tty ubuntu --image=ubuntu:16.04 --restart=Never -- bash -il

2. Install the mysql client:

    $ apt-get update && apt-get install mysql-client -y

3. Connect using the mysql cli, then provide your password:
    $ mysql -h undercooked-saola-mysql -p

To connect to your database directly from outside the K8s cluster:
    MYSQL_HOST=127.0.0.1
    MYSQL_PORT=3306

    # Execute the following command to route the connection:
    kubectl port-forward svc/undercooked-saola-mysql 3306

    mysql -h ${MYSQL_HOST} -P${MYSQL_PORT} -u root -p${MYSQL_ROOT_PASSWORD}
```


Helm has already done a lot for us. It created configMap, a Persistent Volume, a Persistent Volume Claim, a Service, a Secret (for the root password) and a Deployment that manages the Pod.

Helm 已经为我们做了很多。它创建了 configMap、一个 Persistent Volume、一个 Persistent Volume Claim、一个 Service、一个 Secret（用于 root 密码）和一个管理 Pod 的 Deployment。

It also gives you some instructions about how to retrieve the root password and connect to the MySQL database both using another Pod or directly access the service from your own machine.

它还为您提供了一些有关如何使用另一个 Pod 或直接从您自己的机器访问服务来检索 root 密码和连接到 MySQL 数据库的说明。

## Getting Information About Your Deployed Charts

## 获取有关已部署图表的信息

You may have noticed that the moment you hit the ENTER key, Helm immediately printed its output. This behavior ensures that you don’t have a frozen terminal waiting for the Helm Chart to get deployed entirely as it may take a few minutes. However, at any time, you can see the current status of the Chart deployment by running:

```yaml
helm status
```


The output of this command gives you the current status of the chart deployment. In our specific example, it is also displaying instructions about how to connect to your database once the deployment is done. Let’s follow those instructions:

First, let’s grab the root password:

```yaml
$ MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default undercooked-saola-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
$ echo $MYSQL_ROOT_PASSWORD
QHG3e3TCHC

```


The command looks complex, but it’s not:

- We’re using the kubectl get secret to obtain a secret named undercooked-saola-mysql.
- We use jsonpath to extract just the string containing the password (the output of kubectl get is in JSON by default. jsonpath is a way of filtering data).
- Kubernetes Secrets are base64-encoded so we need to decode them using the base64 --decode command.
- The entire output of this chained command becomes the content of a variable called MYSQL\_ROOT\_PASSWORD. Displaying the contents of this variable reveals the root password that got generated for us: QHG3e3TCHC.

该命令看起来很复杂，但事实并非如此：

- 我们使用 kubectl get secret 来获取一个名为 undercooked-saola-mysql 的秘密。
- 我们使用 jsonpath 只提取包含密码的字符串（kubectl get 的输出默认为 JSON。jsonpath 是一种过滤数据的方式）。
- Kubernetes Secret 是 base64 编码的，因此我们需要使用 base64 --decode 命令对其进行解码。
- 这个链式命令的整个输出变成了一个名为 MYSQL\_ROOT\_PASSWORD 的变量的内容。显示此变量的内容会显示为我们生成的 root 密码：QHG3e3TCHC。

Now that we have the password, we can connect to the database instance in one of two ways: through another Pod or from the local machine. Let’s choose the first option as it’s closer to the real-world scenario:

```yaml
kubectl run -i --tty ubuntu --image=ubuntu:16.04 --restart=Never -- bash -il
```




The above command is a quick way to start and login to a container running Ubuntu instead of having to write and apply a definition file. Once you execute this command, you may need to hit ENTER to get the command prompt of the Ubuntu container:

```yaml
If you don't see a command prompt, try pressing enter
root@ubuntu:/#
```


Once inside, we need to install the mysql command line client to actually connect to our database:

```yaml
apt update && apt install -y mysql-client
```


The Helm Chart deployed a service for us so that we can reach our pod, its named undercooked-saola-mysql. Let’s use this service:

```yaml
root@ubuntu:/# mysql -h undercooked-saola-mysql -pQHG3e3TCHC
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.Commands end with ;or \g.
Your MySQL connection id is 11432
Server version: 5.7.14 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates.All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates.Other names may be trademarks of their respective
owners.

Type 'help;'or '\h' for help.Type '\c' to clear the current input statement.

mysql>

```


You can run the following commands to see the different resources that this chart created for you automatically:

```yaml
kubectl get deployments # The deployments
kubectl get pods # The pods
kubectl get secrets # The secrets
kubectl get pvc # The Persistent Volume Claims
kubectl get pv # The Persistent Volume
kubectl get svc # The service
```



# Customizing The Chart

# 自定义图表

The Helm Chart that we’ve just deployed is a great way to automate resource deployment. However, Kubernetes has taken a lot of decisions on our behalf, for example the volume size, the root password, whether to create a database schema in the process and what its name should be. Fortunately, Helm Charts can be highly customized to fit your exact needs. The first step we need to make is to examine what configuration options this chart supports:

```yaml
helm inspect stable/mysql

```


The output is too verbose to be displayed here, but you can have a look at the parameters part. This part contains many configuration options that you can set while executing helm against this chart. Let’s create another chart, set the root user password and create a new user with its own credentials.

输出过于冗长，无法在此处显示，但您可以查看参数部分。这部分包含许多配置选项，您可以在针对此图表执行 helm 时设置这些选项。让我们创建另一个图表，设置 root 用户密码并使用自己的凭据创建一个新用户。

Similar to Kubernetes resources, Helm Charts can be configured in one of two ways: declaratively through a definition file or imperatively through the command line options.

与 Kubernetes 资源类似，Helm Charts 可以通过以下两种方式之一进行配置：通过定义文件以声明方式或通过命令行选项以命令方式进行配置。

## Declarative Customization

## 声明式自定义

Create a new YAML file called config.yaml and add the following to it:

```yaml line-numbers
mysqlRootPassword: myrootpassword
mysqlUser: dbuser
mysqlPassword: mydbuserpassword

```


Create the Chart passing in the configuration file as follows:

```yaml
helm install -f config.yaml stable/mysql

```


Wait till the Pods are in the running state and let’s ensure that we have the passwords we set and that we can correctly connect to the new instance using the settings we configured:

```yaml
$ MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default whopping-grizzly-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
$ echo $MYSQL_ROOT_PASSWORD
myrootpassword

```


Notice that Helm named our new [deployment](http://www.magalix.com/blog/kubernetes-deployments-101) whopping-grizzly-mysql (this name is also referred to as the release object). Helm automatically chooses this name for you, but you can override this behavior by using the --name parameter. For example:

```yaml
helm install -f config.yaml --name my-awesome-chart stable/mysql

```


We can try to login with this password and also test logging-in with the user we created:

```yaml
$ kubectl delete pods ubuntu
$ kubectl run -i --tty ubuntu --image=ubuntu:16.04 --restart=Never -- bash -il
If you don't see a command prompt, try pressing enter.
root@ubuntu:/#
root@ubuntu:/# apt-get update && apt-get install mysql-client -y
# Output suppressed
root@ubuntu:/# mysql -h whopping-grizzly-mysql -pmyrootpassword
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.Commands end with ;or \g.
Your MySQL connection id is 184
Server version: 5.7.14 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates.All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates.Other names may be trademarks of their respective
owners.

Type 'help;'or '\h' for help.Type '\c' to clear the current input statement.

mysql> ^DBye
root@ubuntu:/# mysql -h whopping-grizzly-mysql -u mydbuser -pmydbpassword
mysql: [Warning] Using a password on the command line interface can be insecure.
ERROR 1045 (28000): Access denied for user 'mydbuser'@'10.1.0.129' (using password: YES)
root@ubuntu:/# mysql -h whopping-grizzly-mysql -u dbuser -pmydbuserpassword
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.Commands end with ;or \g.
Your MySQL connection id is 197
Server version: 5.7.14 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates.All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates.Other names may be trademarks of their respective
owners.

Type 'help;'or '\h' for help.Type '\c' to clear the current input statement.

mysql> ^DBye
root@ubuntu:/# logout

```




[![Learn How to Continuously Optimize your K8s Cluster](https://no-cache.hubspot.com/cta/default/3487587/963b2ba4-df32-4a3b-8781-f8fa4ff38750.png)](https://cta-redirect.hubspot.com/cta/redirect/3487587/963b2ba4-df32-4a3b-8781-f8fa4ff38750)

-redirect.hubspot.com/cta/redirect/3487587/963b2ba4-df32-4a3b-8781-f8fa4ff38750)

So, by replicating the steps we did before in our previous chart, we were able to connect to the database using the password that we set in the configuration file. We could also login with the user credentials we created for the non-root user.

因此，通过复制我们之前在之前图表中执行的步骤，我们能够使用我们在配置文件中设置的密码连接到数据库。我们还可以使用我们为非 root 用户创建的用户凭据登录。

## Imperative Customization

## 命令式定制

We would have achieved the same result if we used the imperative option of passing configuration options to Helm charts using a command like the following:

```yaml
helm install --set mysqlRootPassword=myrootpassword,mysqlUser=dbuser,mysqlPassword=mydbuserpassword stable/mysql

```


Notice that we are allowed to use --set in addition to -f _file_. However, in that case the --set parameters take precedence over the ones supplied in the configuration file. Common use cases of this method is when you define some default values in the configuration file and override them using the --set parameters.

请注意，除了 -f _file_ 之外，我们还可以使用 --set。但是，在这种情况下，--set 参数优先于配置文件中提供的参数。此方法的常见用例是当您在配置文件中定义一些默认值并使用 --set 参数覆盖它们时。

# Upgrading and Downgrading (Rolling Back) Charts

# 升级和降级（回滚）图表

You may have noticed that when we made configuration changes to our Chart, Helm created a new one for us with the new configuration. Creating new Kubernetes objects with every configuration change might not be the best option for everyone. Sometimes you may want to apply the changes to the existing Chart. You can do this using the helm upgrade command. Let’s - once again - change the root password for our database. This time we want Helm to apply the changes to the existing configuration without creating a new release.

您可能已经注意到，当我们对 Chart 进行配置更改时，Helm 使用新配置为我们创建了一个新的 Chart。在每次配置更改时创建新的 Kubernetes 对象可能不是每个人的最佳选择。有时您可能希望将更改应用于现有图表。您可以使用 helm upgrade 命令执行此操作。让我们再次更改数据库的 root 密码。这次我们希望 Helm 在不创建新版本的情况下将更改应用于现有配置。

Let’s check the image tag that is currently used with this Chart:

```yaml
helm inspect stable/mysql |grep imageTag
imageTag: "5.7.14"
imageTag: v0.10.0
|`imageTag`                                   |`mysql` image tag.|`5.7.14`                                             |
|`metrics.imageTag`                           |Exporter image                                                                               |`v0.10.0`                                            |

```


By default, this Chart uses MySQL version 5.7.14. Let’s assume that we are interested in running a newer version of the database engine, say 8. We need to change the imageTag parameter that the Chart uses. Our config.yaml should like this:

```yaml line-numbers
mysqlRootPassword: myrootpassword
mysqlUser: dbuser
mysqlPassword: myebuserpassword
imageTag: 8

```


Now, in order to apply this configuration to the same deployment, we use the upgrade subcommand:

```yaml
helm upgrade -f config.yaml whopping-grizzly stable/mysql

```


The upgrade subcommands takes two arguments: the release name and the Chart.

upgrade 子命令有两个参数：版本名称和图表。

If you run kubectl get pods you should see that we have one pod terminating and another one initializing.

如果您运行 kubectl get pods，您应该会看到我们有一个 Pod 正在终止，另一个正在初始化。

If you want to double check that the engine is using the correct version, first get the deployment name:

```yaml
$ kubectl get deployments
NAME                      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
undercooked-saola-mysql   1         1         1            1           2d
whopping-grizzly-mysql    1         1         1            1           21h

```


Our deployment name starts with the release name (whopping-grizly). Let’s investigate the controller properties by running:

```yaml
kubectl edit deployment whopping-grizzly-mysql

```


Search for image and ensure that it is using the correct tag. The line containing the image name should look as follows:

```yaml
- name: MYSQL_DATABASE
        image: mysql:8

```


## Rolling Back

##  滚回来

At the start of this article, we mentioned that we should limit the number of history items Helm stores. We set this number to 200, which means that Helm will keep up to 200 releases in history so that you can roll back to the required version.

在本文开头，我们提到应该限制 Helm 存储的历史项目数量。我们将此数字设置为 200，这意味着 Helm 将在历史记录中最多保留 200 个版本，以便您可以回滚到所需的版本。

To view the currently-stored releases in Helm run the following command:

```yaml
$ helm history whopping-grizzly
REVISION    UPDATED                     STATUS        CHART          DESCRIPTION
1           Fri Jul 26 15:11:09 2019    SUPERSEDED    mysql-1.3.0    Install complete
2           Sat Jul 27 12:06:49 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
3           Sat Jul 27 12:41:49 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
4           Sat Jul 27 12:42:36 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
5           Sat Jul 27 12:44:55 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
6           Sat Jul 27 12:48:19 2019    DEPLOYED      mysql-1.3.0    Upgrade complete

```




If you want to roll back to the first release ever, issue the following command:

```yaml
$ helm rollback whopping-grizzly 1
Rollback was a success.

```


The pod will get immediately replaced with a new one containing the configuration used in release 1.

该 pod 将立即替换为包含版本 1 中使用的配置的新 pod。

# Deleting Release

# 删除发布

If you want to remove a release you use the following command:

```yaml
$ helm delete whopping-grizzly
release "whopping-grizzly" deleted

```


Notice that this command removes all the resources defined by the release (pods, deployments, secrets,etc.)

请注意，此命令会删除发布定义的所有资源（pod、部署、秘密等）。

# Helm Repositories

# Helm 存储库

Like all other similar tools, Helm supports repositories. A Helm repository contains several ready-made Charts that fit most of your deployment needs. The stable repo is used by default. You can search for a Chart that provisions Redis for example by using a command like the following:

```yaml
helm search redis

```


You can also add more Repositories that Helm can use. For example, the incubator repo:

```yaml
helm repo add incubator https://kubernetes-charts-incubator.storage.googleapis.com/

```


Now, you can search this repository for all the Charts it contains:

```yaml
helm search incubator

```


Or search for a specific Chart by preceding the search term with the repo name:

```yaml
helm search incubator/prom

```


# TL;DR



- Configuration management is a practice used environment deployment and replication. They use templates to automatically deploy components with different configurations depending on parameterized values.

  配置管理是一种使用环境部署和复制的实践。他们使用模板根据参数化值自动部署具有不同配置的组件。

- There are configuration management tools for standalone infrastructures like Ansible, Chef, and Puppet. There are also configuration management tools for clustered applications like Kubernetes. In this article, we discussed Helm.

  有适用于 Ansible、Chef 和 Puppet 等独立基础架构的配置管理工具。还有用于集群应用程序（如 Kubernetes）的配置管理工具。在本文中，我们讨论了 Helm。

- Helm is made up of two parts: the helm command (client) that runs on your laptop (or client machine) and Tiller which is the server part of the application, it runs on Kubernetes (server side).

- In this article, we created a small lab in which we used Helm to automatically deploy the necessary Kubernetes components to run a MySQL database instance.

- There are several Chart parameters that can be customized to configure the environment. For example, the container image and its version.

- Helm supports upgrading the existing infrastructure with the new parameters. You can apply the new parameters through a configuration file or directly on the command line.

- You can rollback a release or more up to a defined limit that you can set when you initialize Helm.

- Helm uses the “stable” repository by default. However, you can add more repositories to the client and search for a particular Chart. 

- Helm 由两部分组成：在您的笔记本电脑（或客户端机器）上运行的 helm 命令（客户端）和作为应用程序服务器部分的 Tiller，它在 Kubernetes（服务器端）上运行。

- 在本文中，我们创建了一个小型实验室，在其中使用 Helm 自动部署运行 MySQL 数据库实例所需的 Kubernetes 组件。

- 有几个图表参数可以自定义以配置环境。例如，容器映像及其版本。

- Helm 支持使用新参数升级现有基础设施。您可以通过配置文件或直接在命令行上应用新参数。

- 您可以将一个或多个版本回滚到您在初始化 Helm 时可以设置的定义限制。

- Helm 默认使用“稳定”存储库。但是，您可以向客户端添加更多存储库并搜索特定图表。


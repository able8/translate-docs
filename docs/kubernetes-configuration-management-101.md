# What Configuration Management is and Why You Should Implement it in Your infrastructure

Wikipedia defines [Configuration Management](https://en.wikipedia.org/wiki/Configuration_management) as “a [systems engineering](https://en.wikipedia.org/wiki/Systems_engineering) process for establishing and maintaining consistency of a product's performance, functional, and physical attributes with its requirements, design, and operational information throughout its life.”

Configuration Management has two main aspects:

1. Bringing up environments quickly from one or more templates.
2. Being able to replicate those environments easily by changing a few parameters.

Many Configuration Management tools exist and address the above requirements. For example, Ansible. Let’s say that you need to deploy WordPress to a virtual (or even physical) machine. All that you need to do is search Ansible Galaxy for a role that deploys WordPress and change its parameters according to your environment needs. There are other tools as well, and each of them has its deployable unit and a repository — for example, Puppet and its Puppet Forge and Chef and its Supermarket.

![configrations 1](https://www.magalix.com/hs-fs/hubfs/configrations%201.png?width=720&name=configrations%201.png)

You might be thinking that Kubernetes does not require configuration management in that sense. After all, you add all the requirements to the Docker image and let Kubernetes do the rest. But, Kubrenetes is not about images and containers only, you have other components that are built around containers like [Pods](http://www.magalix.com/blog/kubernetes-pods-101-the-cluster-sailors), [Services](http://www.magalix.com/blog/kubernetes-services-101-the-pods-interfaces), [Ingresses](https://kubernetes.io/docs/concepts/services-networking/ingress/), [configMaps](http://www.magalix.com/blog/the-configmap-pattern), [Secrets](http://www.magalix.com/blog/kubernetes-secrets-101), [Volumes](http://www.magalix.com/blog/tag/storage), etc. So, while a traditional Configuration Management tool is only concerned with what runs inside the instance, a Kubernetes Configuration Management tool is responsible for what gets deployed to the instance as well as the infrastructure that surrounds this deployment and making it scalable.

In this article, we explore a popular [Kubernetes](http://www.magalix.com/blog/kubernetes-101-concepts-and-why-it-matters) tool that can be used for Configuration Management: Helm.

# Helm

Helm can be thought of as the Package Manager for [Kubernetes](http://www.magalix.com/blog/kubernetes-101-concepts-and-why-it-matters). If Ubuntu uses apt, Centos uses yum, and macOS uses brew, then, in the same sense, Kubernetes uses Helm. With a single command, you can have a complete Kubernetes infrastructure up and running.

![configrations 2](https://www.magalix.com/hs-fs/hubfs/configrations%202.png?width=720&name=configrations%202.png)


## Helm Installation

Helm has many installation options depending on your operating system. You can refer to [this page](https://helm.sh/docs/using_helm/#installing-helm) for the installation instructions specific to your OS.

## Giving Helm a Test Drive

By default, Helm applies your commands to the default cluster context that your kubectl command uses. You can view the current context by issuing:

```yaml
kubectl config current-context

```

Helm operates in a very simple way:

- The helm tool works on the client computer.
- The commands issued by helm go to the Tiller tool, which runs on the server side (inside Kubernetes).
- Tiller apples the Chart to the cluster.

Let’s give the above steps a test drive by installing the MySQL database on our cluster.

The first step we need to do is install Tiller on the cluster:

```yaml
helm init --history-max 200

```

The above command will initialize helm locally (you can see several files get created in the command output) and also deploys Tiller to the cluster. Notice that we set the --history-max to an arbitrary number. This is recommended because otherwise Helm and Tiller will keep an indefinite amount of objects like configMaps, which may add unnecessary overhead.

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
NAME                      	DATA  AGE
undercooked-saola-mysql-test  1 	0s

==> v1/PersistentVolumeClaim
NAME                 	STATUS   VOLUME	CAPACITY  ACCESS MODES  STORAGECLASS  AGE
undercooked-saola-mysql  Pending  hostpath  0s

==> v1/Pod(related)
NAME                                	READY  STATUS   RESTARTS  AGE
undercooked-saola-mysql-fc896fb6-zwfn6  0/1	Pending  0     	0s

==> v1/Secret
NAME                 	TYPE	DATA  AGE
undercooked-saola-mysql  Opaque  2 	0s

==> v1/Service
NAME                 	TYPE   	CLUSTER-IP 	EXTERNAL-IP  PORT(S)   AGE
undercooked-saola-mysql  ClusterIP  10.101.209.21     	3306/TCP  0s

==> v1beta1/Deployment
NAME                 	READY  UP-TO-DATE  AVAILABLE  AGE
undercooked-saola-mysql  0/1	1       	0      	0s

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

It also gives you some instructions about how to retrieve the root password and connect to the MySQL database both using another Pod or directly access the service from your own machine.

## Getting Information About Your Deployed Charts

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
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 11432
Server version: 5.7.14 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

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

The Helm Chart that we’ve just deployed is a great way to automate resource deployment. However, Kubernetes has taken a lot of decisions on our behalf, for example the volume size, the root password, whether to create a database schema in the process and what its name should be. Fortunately, Helm Charts can be highly customized to fit your exact needs. The first step we need to make is to examine what configuration options this chart supports:

```yaml
helm inspect stable/mysql

```

The output is too verbose to be displayed here, but you can have a look at the parameters part. This part contains many configuration options that you can set while executing helm against this chart. Let’s create another chart, set the root user password and create a new user with its own credentials.

Similar to Kubernetes resources, Helm Charts can be configured in one of two ways: declaratively through a definition file or imperatively through the command line options.

## Declarative Customization

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
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 184
Server version: 5.7.14 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> ^DBye
root@ubuntu:/# mysql -h whopping-grizzly-mysql -u mydbuser -pmydbpassword
mysql: [Warning] Using a password on the command line interface can be insecure.
ERROR 1045 (28000): Access denied for user 'mydbuser'@'10.1.0.129' (using password: YES)
root@ubuntu:/# mysql -h whopping-grizzly-mysql -u dbuser -pmydbuserpassword
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 197
Server version: 5.7.14 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> ^DBye
root@ubuntu:/# logout

```

[![Learn How to Continuously Optimize your K8s Cluster](https://no-cache.hubspot.com/cta/default/3487587/963b2ba4-df32-4a3b-8781-f8fa4ff38750.png)](https://cta-redirect.hubspot.com/cta/redirect/3487587/963b2ba4-df32-4a3b-8781-f8fa4ff38750)

So, by replicating the steps we did before in our previous chart, we were able to connect to the database using the password that we set in the configuration file. We could also login with the user credentials we created for the non-root user.

## Imperative Customization

We would have achieved the same result if we used the imperative option of passing configuration options to Helm charts using a command like the following:

```yaml
helm install --set mysqlRootPassword=myrootpassword,mysqlUser=dbuser,mysqlPassword=mydbuserpassword stable/mysql

```

Notice that we are allowed to use --set in addition to -f _file_. However, in that case the --set parameters take precedence over the ones supplied in the configuration file. Common use cases of this method is when you define some default values in the configuration file and override them using the --set parameters.

# Upgrading and Downgrading (Rolling Back) Charts

You may have noticed that when we made configuration changes to our Chart, Helm created a new one for us with the new configuration. Creating new Kubernetes objects with every configuration change might not be the best option for everyone. Sometimes you may want to apply the changes to the existing Chart. You can do this using the helm upgrade command. Let’s - once again - change the root password for our database. This time we want Helm to apply the changes to the existing configuration without creating a new release.

Let’s check the image tag that is currently used with this Chart:

```yaml
helm inspect stable/mysql | grep imageTag
imageTag: "5.7.14"
imageTag: v0.10.0
| `imageTag`                               	| `mysql` image tag.                                                                       	| `5.7.14`                                         	|
| `metrics.imageTag`                       	| Exporter image                                                                           	| `v0.10.0`                                        	|

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

If you run kubectl get pods you should see that we have one pod terminating and another one initializing.

If you want to double check that the engine is using the correct version, first get the deployment name:

```yaml
$ kubectl get deployments
NAME                  	DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
undercooked-saola-mysql   1     	1     	1        	1       	2d
whopping-grizzly-mysql	1     	1     	1        	1       	21h

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

At the start of this article, we mentioned that we should limit the number of history items Helm stores. We set this number to 200, which means that Helm will keep up to 200 releases in history so that you can roll back to the required version.

To view the currently-stored releases in Helm run the following command:

```yaml
$ helm history whopping-grizzly
REVISION    UPDATED            		 STATUS   	 CHART 		 DESCRIPTION
1  		 Fri Jul 26 15:11:09 2019    SUPERSEDED    mysql-1.3.0    Install complete
2  		 Sat Jul 27 12:06:49 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
3  		 Sat Jul 27 12:41:49 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
4  		 Sat Jul 27 12:42:36 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
5  		 Sat Jul 27 12:44:55 2019    SUPERSEDED    mysql-1.3.0    Upgrade complete
6  		 Sat Jul 27 12:48:19 2019    DEPLOYED 	 mysql-1.3.0    Upgrade complete

```

If you want to roll back to the first release ever, issue the following command:

```yaml
$ helm rollback whopping-grizzly 1
Rollback was a success.

```

The pod will get immediately replaced with a new one containing the configuration used in release 1.

# Deleting Release

If you want to remove a release you use the following command:

```yaml
$ helm delete whopping-grizzly
release "whopping-grizzly" deleted

```

Notice that this command removes all the resources defined by the release (pods, deployments, secrets,etc.)

# Helm Repositories

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
- There are configuration management tools for standalone infrastructures like Ansible, Chef, and Puppet. There are also configuration management tools for clustered applications like Kubernetes. In this article, we discussed Helm.
- Helm is made up of two parts: the helm command (client) that runs on your laptop (or client machine) and Tiller which is the server part of the application, it runs on Kubernetes (server side).
- In this article, we created a small lab in which we used Helm to automatically deploy the necessary Kubernetes components to run a MySQL database instance.
- There are several Chart parameters that can be customized to configure the environment. For example, the container image and its version.
- Helm supports upgrading the existing infrastructure with the new parameters. You can apply the new parameters through a configuration file or directly on the command line.
- You can rollback a release or more up to a defined limit that you can set when you initialize Helm.
- Helm uses the “stable” repository by default. However, you can add more repositories to the client and search for a particular Chart.


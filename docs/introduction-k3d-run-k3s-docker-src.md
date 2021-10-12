# Introduction to k3d: Run K3s in Docker

March 1, 2021
\|
By:
[Thorsten Klein](https://www.suse.com/c/author/thorsten-klein-95googlemail-com/ "View all posts by Thorsten Klein")

In this blog post, we’re going to talk about k3d, a tool that allows you to run throwaway Kubernetes clusters anywhere you have Docker installed. I’ve anticipated your questions — so let’s go!

## What is k3d?

k3d is a small program made for running a [K3s](https://k3s.io) cluster in Docker. K3s is a lightweight, CNCF-certified Kubernetes distribution and Sandbox project. Designed for low-resource environments, K3s is distributed as a single binary that uses under 512MB of RAM. To learn more about K3s, head over to [the documentation](https://rancher.com/docs/k3s/latest/en/) or check out this [blog post](https://rancher.com/blog/2019/2019-02-26-introducing-k3s-the-lightweight-kubernetes-distribution-built-for-the-edge/) or [video.](https://www.youtube.com/watch?v=hMr3prm9gDM)

k3d uses a [Docker image](https://hub.docker.com/r/rancher/k3s/tags) built from the [K3s repository](https://github.com/rancher/k3s) to spin up multiple K3s nodes in Docker containers on any machine with Docker installed. That way, a single physical (or virtual) machine (let’s call it Docker Host) can run multiple K3s clusters, with multiple server and agent nodes each, simultaneously.

## What Can k3d Do?

As of k3d version v4.0.0, released in January 2021, k3d’s abilities boil down to the following features:

- create/stop/start/delete/grow/shrink K3s clusters (and individual nodes)
  - via command line flags
  - via configuration file
- manage and interact with container registries that can be used with the cluster
- manage Kubeconfigs for the clusters
- import images from your local Docker daemon into the container runtime running in the cluster

Obviously, there’s way more to it and you can tweak everything in great detail.

## What is k3d Used for?

The main use case for k3d is local development on Kubernetes with little hassle and resource usage. The intention behind the initial development of k3d was to provide developers with an easy tool that allowed them to run a lightweight Kubernetes cluster on their development machine, giving them fast iteration times in a production-like environment (as opposed to running docker-compose locally vs. Kubernetes in production).

Over time, k3d also evolved into a tool used by operations to test some Kubernetes (or, specifically K3s) features in an isolated environment. For example, with k3d you can easily create multi-node clusters, deploy something on top of it, simply stop a node and see how Kubernetes reacts and possibly reschedules your app to other nodes.

Additionally, you can use k3d in your continuous integration system to quickly spin up a cluster, deploy your test stack on top of it and run integration tests. Once you’re finished, you can simply decommission the cluster as a whole. No need to worry about proper cleanups and possible leftovers.

We also provide a `k3d-dind` image (similar to dreams within dreams in the movie Inception, we’ve got containers within containers within containers.) With that, you can create a docker-in-docker environment where you run k3d, which spawns a K3s cluster in Docker. That means that you only have a single container (k3d-dind) running on your Docker host, which in turn runs a whole K3s/Kubernetes cluster inside.

## How Do I Use k3d?

1. [Install k3d](https://k3d.io/#installation) (and [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), if you want to use it)

**Note**: to follow along with this post, use at least k3d v4.1.1
2. Try one of the following examples or use the[documentation](https://k3d.io/usage/commands/) or the CLI help text to find your own way ( `k3d [command] --help`)

### The “Simple” Way

```
k3d cluster create
```

This single command spawns a K3s cluster with two containers: A Kubernetes control-plane node (server) and a load balancer (serverlb) in front of it. It puts both of them in a dedicated Docker network and exposes the Kubernetes API on a randomly chosen free port on the Docker host. It also creates a named Docker volume in the background as a preparation for image imports.

By default, if you don’t provide a name argument, the cluster will be named `k3s-default`

and the containers will show up as `k3d-<-role>-<#>`, so in this case `k3d-k3s-default-serverlb` and `k3d-k3s-default-server-0`.

k3d waits until everything is ready, pulls the Kubeconfig from the cluster and merges it with your default Kubeconfig (usually it’s in `$HOME/.kube/config` or whatever path your `KUBECONFIG` environment variable points to).

No worries, you can tweak that behavior as well.

Check out what you’ve just created using `kubectl` to show you the nodes: `kubectl get nodes`.

k3d also gives you some commands to list your creations: `k3d cluster|node|registry list`.

### The “Simple but Sophisticated” Way

`k3d cluster create mycluster --api-port 127.0.0.1:6445 --servers 3 --agents 2 --volume '/home/me/mycode:/code@agent[*]' --port '8080:80@loadbalancer'`

`` This single command spawns a K3s cluster with six containers:

-  1 load balancer
-  3 servers (control-plane nodes)
-  2 agents (formerly worker nodes)

With the `--api-port 127.0.0.1:6445`, you tell k3d to map the Kubernetes API Port ( `6443` internally) to `127.0.0.1`/localhost’s port `6445`. That means that you will have this connection string in your Kubeconfig: `server: https://127.0.0.1:6445` to connect to this cluster.

This port will be mapped from the load balancer to your host system. From there, requests will be proxied to your server nodes, effectively simulating a production setup, where server nodes also can go down and you would want to failover to another server.

The `--volume /home/me/mycode:/code@agent[*]` bind mounts your local directory `/home/me/mycode` to the path `/code` inside all ( `[*]` of your agent nodes). Replace `*` with an index (here: 0 or 1) to only mount it into one of them.

The specification telling k3d which nodes it should mount the volume to is called “node filter” and it’s also used for other flags, like the `--port` flag for port mappings.

That said, `--port '8080:80@loadbalancer'` maps your local host’s port `8080` to port `80` on the load balancer (serverlb), which can be used to forward HTTP ingress traffic to your cluster. For example, you can now deploy a web app into the cluster (Deployment), which is exposed (Service) externally via an Ingress such as `myapp.k3d.localhost`.

Then (provided that everything is set up to resolve that domain to your local host IP), you can point your browser to `http://myapp.k3d.localhost:8080` to access your app. Traffic then flows from your host through the Docker bridge interface to the load balancer. From there, it’s proxied to the cluster, where it passes via Ingress and Service to your application Pod.

**Note**: You have to have some mechanism set up to route to resolve `myapp.k3d.localhost` to your local host IP ( `127.0.0.1`). The most common way is using entries of the form `127.0.0.1 myapp.k3d.localhost` in your `/etc/hosts` file ( `C:\Windows\System32\drivers\etc\hosts` on Windows). However, this does not allow for wildcard entries ( `*.localhost`), so it may become a bit cumbersome after a while, so you may want to have a look at tools like [`dnsmasq` (MacOS/UNIX](https://en.wikipedia.org/wiki/Dnsmasq)) or [`Acrylic` (Windows)](https://stackoverflow.com/a/9695861/6450189) to ease the burden.

**Tip**: You can install the package `<a href="https://manpages.debian.org/unstable/libnss-myhostname/nss-myhostname.8.en.html" target="_blank" rel="noopener" data-index="40">libnss-myhostname</a>` on some systems (at least Linux operating systems including SUSE Linux and openSUSE), to auto-resolve `*.localhost` domains to `127.0.0.1`, which means you don’t have to fiddle around with e.g. `/etc/hosts`, if you prefer to test via Ingress, where you need to set a domain.

One interesting thing to note here: if you create more than one server node, K3s will be given the `--cluster-init` flag, which means that it swaps its internal datastore (by default that’s SQLite) for etcd.

### The “Configuration as Code” Way

As of k3d v4.0.0 (January 2021), we support config files to configure everything as code that you’d previously do via command line flags (and soon possibly even more than that).

As of this writing, the JSON-Schema used to validate the configuration file can be [found in the repository](https://github.com/rancher/k3d/blob/092f26a4e27eaf9d3a5bc32b249f897f448bc1ce/pkg/config/v1alpha2/schema.json).

Here’s an example config file:

```
# k3d configuration file, saved as e.g. /home/me/myk3dcluster.yaml
apiVersion: k3d.io/v1alpha2 # this will change in the future as we make everything more stable
kind: Simple # internally, we also have a Cluster config, which is not yet available externally
name: mycluster # name that you want to give to your cluster (will still be prefixed with `k3d-`)
servers: 1 # same as `--servers 1`
agents: 2 # same as `--agents 2`
kubeAPI: # same as `--api-port 127.0.0.1:6445`
hostIP: "127.0.0.1"
hostPort: "6445"
ports:
 - port: 8080:80 # same as `--port 8080:80@loadbalancer
nodeFilters:
 - loadbalancer
options:
k3d: # k3d runtime settings
wait: true # wait for cluster to be usable before returining; same as `--wait` (default: true)
timeout: "60s" # wait timeout before aborting; same as `--timeout 60s`
k3s: # options passed on to K3s itself
extraServerArgs: # additional arguments passed to the `k3s server` command
     - --tls-san=my.host.domain
extraAgentArgs: [] # addditional arguments passed to the `k3s agent` command
kubeconfig:
updateDefaultKubeconfig: true # add new cluster to your default Kubeconfig; same as `--kubeconfig-update-default` (default: true)
switchCurrentContext: true # also set current-context to the new cluster's context; same as `--kubeconfig-switch-context` (default: true)
```

Assuming that we saved this as `/home/me/myk3dcluster.yaml`, we can use it to configure a new cluster:

`k3d cluster create --config /home/me/myk3dcluster.yaml`

Note that you can still set additional arguments or flags, which will then take precedence (or will be merged) with whatever you have defined in the config file.

## What More Can I Do with k3d?

You can use k3d in even more ways, including:

- Create a cluster together with a k3d-managed container**registry**
- Use the cluster for fast development with**hot code reloading**
- Use k3d in combination with other development tools like`Tilt` or `Skaffold` ``
  - both can leverage the power of importing images via`k3d image import`
  - both can alternatively make use of a k3d-managed registry to speed up your development loop
- Use k3d in your CI system (we have a[PoC for that](https://github.com/iwilltry42/k3d-demo/blob/main/.drone.yml))
- Integrate it in your vscode workflow using the awesome new[community-maintained vscode extension](https://github.com/inercia/vscode-k3d)
- Use it to[set up K3s high availability](https://rancher.com/blog/2020/set-up-k3s-high-availability-using-k3d)

You can try all of these yourself by using prepared scripts in [this demo repository](https://github.com/iwilltry42/k3d-demo) or watch us showing them off in [one of our meetups](https://www.youtube.com/watch?v=d9JRb4fk5ag&feature=youtu.be).

Other than that, remember that k3d is a [community-driven](https://github.com/rancher/k3d/blob/main/CONTRIBUTING.md) project, so we’re always happy to hear from you on [Issues](https://github.com/rancher/k3d/issues), [Pull Requests](https://github.com/rancher/k3d/issues), [Discussions](https://github.com/rancher/k3d/discussions) and [Slack Chats](https://slack.rancher.io/)!

Ready to give k3d a try? [Start by downloading K3s](https://github.com/k3s-io/k3s/releases).

_Thorsten Klein is a DevOps Engineer at trivago and a freelance software engineer at SUSE. Thorsten is the maintener of k3d. Find him on [Twitter](https://twitter.com/iwilltry42) or visit his [website](https://iwilltry42.dev)._

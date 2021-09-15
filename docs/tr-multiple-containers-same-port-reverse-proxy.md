# How to Expose Multiple Containers On the Same Port

# 如何在同一个端口上暴露多个容器

August 28, 2021

2021 年 8 月 28 日

[Containers,](http://iximiuz.com/en/categories/?category=Containers)[Networking,](http://iximiuz.com/en/categories/?category=Networking) [Linux / Unix]( http://iximiuz.com/en/categories/?category=Linux / Unix)

[容器，](http://iximiuz.com/en/categories/?category=Containers)[网络，](http://iximiuz.com/en/categories/?category=Networking) [Linux / Unix]( http://iximiuz.com/en/categories/?category=Linux / Unix)

_**Disclaimer:** In 2021, there is still a place for simple setups with just one machine serving all traffic. So, no Kubernetes and no cloud load balancers in this post. Just good old Docker and Podman._

_**免责声明：** 在 2021 年，仍然有一个地方可以进行简单的设置，只需一台机器服务所有流量。因此，本文中没有 Kubernetes 和云负载均衡器。好老的 Docker 和 Podman。_

Even when you have just one physical or virtual server, it's often a good idea to run multiple instances of your application on it. Luckily, when the application is containerized, it's actually relatively simple. With multiple application containers, you get **horizontal scaling** and a much-needed **redundancy** for a very little price. Thus, if there is a sudden need for handling more requests, you can adjust the number of containers accordingly. And if one of the containers dies, there are others to handle its traffic share, so your app isn't a [SPOF](https://en.wikipedia.org/wiki/Single_point_of_failure) anymore.

即使您只有一台物理或虚拟服务器，在其上运行应用程序的多个实例通常也是一个好主意。幸运的是，当应用程序被容器化时，它实际上相对简单。借助多个应用程序容器，您可以以极低的价格获得**水平扩展** 和急需的**冗余**。因此，如果突然需要处理更多请求，您可以相应地调整容器数量。如果其中一个容器死亡，还有其他容器来处理其流量份额，因此您的应用不再是 [SPOF](https://en.wikipedia.org/wiki/Single_point_of_failure)。

The tricky part here is how to expose such a multi-container application to the clients. Multiple containers mean multiple listening sockets. But most of the time, clients just want to have a single point of entry.

这里的棘手部分是如何向客户端公开这样一个多容器应用程序。多个容器意味着多个侦听套接字。但大多数时候，客户只想有一个单一的入口点。

![Benefits of exposing multiple Docker containers on the same port](http://iximiuz.com/multiple-containers-same-port-reverse-proxy/multiple-containers-same-port-2000-opt.png)

Surprisingly or not, neither Docker nor Podman [support exposing multiple containers on the same host's port](https://github.com/docker/for-linux/issues/471) right out of the box.

不管是否令人惊讶，Docker 和 Podman [支持在同一主机的端口上公开多个容器](https://github.com/docker/for-linux/issues/471) 都不是开箱即用的。

Example: `docker-compose` failing scenario with `"Service specifies a port on the host. If multiple containers for this service are created on a single host, the port will clash."`

示例：`docker-compose` 失败场景，带有 `"Service 在主机上指定一个端口。如果在单个主机上为该服务创建多个容器，端口将发生冲突。"`

For instance, if you have the following docker-compose file:

例如，如果您有以下 docker-compose 文件：

```yaml
version: '3'
services:
app:
    image: kennethreitz/httpbin
    ports:
      - "80:80"  # httpbin server listens on the container's 0.0.0.0:80

```

And you want to scale up the `app` service with `docker-compose up --scale app=2`, you'll get the following error:

如果你想用 `docker-compose up --scale app=2` 来扩展 `app` 服务，你会得到以下错误：

```bash
WARNING: The "app" service specifies a port on the host.If multiple containers for this service are created on a single host, the port will clash.
Starting example_app_1 ... done
Creating example_app_2 ...
Creating example_app_2 ... error

ERROR: for example_app_2  Cannot start service app: driver failed programming external connectivity on endpoint example_app_2 (...): Bind for 0.0.0.0:80 failed: port is already allocated

ERROR: for app  Cannot start service app: driver failed programming external connectivity on endpoint example_app_2 (...): Bind for 0.0.0.0:80 failed: port is already allocated
ERROR: Encountered errors while bringing up the project.

```

The most widely suggested workaround is to use an extra container with a reverse proxy like Nginx, HAProxy, Envoy, or Traefik. Such a proxy should know the exact set of application containers and load balance the client traffic between them. The only port that needs to be exposed on the host in such a setup is the port of the proxy itself. Additionally, since modern reverse proxies usually come with advanced routing capabilities, you can get **canary and blue-green deployments** or even coalesce **diverse backend apps into a single fronted app** almost for free.

最广泛建议的解决方法是使用带有反向代理的额外容器，例如 Nginx、HAProxy、Envoy 或 Traefik。这样的代理应该知道应用程序容器的确切集合并负载平衡它们之间的客户端流量。在这种设置中，唯一需要在主机上公开的端口是代理本身的端口。此外，由于现代反向代理通常具有高级路由功能，您几乎可以免费获得**金丝雀和蓝绿部署**，甚至可以将**不同的后端应用程序合并为一个单一的前端应用程序**。

For any production setup, [I'd recommend going with a reverse proxy approach first](http://iximiuz.com#reverse-proxy). But in this post, I want to explore the alternatives. Are there other ways to expose multiple Docker or Podman containers on the same host's port? 

对于任何生产设置，[我建议首先使用反向代理方法](http://iximiuz.com#reverse-proxy)。但在这篇文章中，我想探索替代方案。还有其他方法可以在同一主机的端口上公开多个 Docker 或 Podman 容器吗？

The goal of the below exercise is manifold. First of all, to solidify the knowledge obtained while working on my [container networking](http://iximiuz.com/en/posts/container-networking-is-simple/) and [iptables](http://iximiuz.com/en/posts/laymans-iptables-101/) write-ups. But also to show that the proxy is not the only possible way to achieve a decent load balancing. So, if you are in a restricted setup where you can't use a proxy for some reason, the techniques from this post may come in handy.

以下练习的目标是多方面的。首先，巩固在我的[容器网络](http://iximiuz.com/en/posts/container-networking-is-simple/)和[iptables](http://iximiuz.com/en/posts/laymans-iptables-101/) 文章。但也表明代理不是实现体面负载平衡的唯一可能方法。因此，如果您处于由于某种原因无法使用代理的受限设置中，那么本文中的技术可能会派上用场。

## Multiple Containers / Same Port using SO\_REUSEPORT

## 使用 SO\_REUSEPORT 的多个容器/同一端口

Let's forget about containers for a second and talk about sockets in general.

让我们暂时忘掉容器，来谈谈一般的套接字。

To make a server socket [`listen()`](https://man7.org/linux/man-pages/man2/listen.2.html) on a certain address, you need to explicitly [`bind()` ](https://man7.org/linux/man-pages/man2/bind.2.html) it to an interface and port. For a long time, binding a socket to an _(interface, port)_ pair was an exclusive operation. If you bound a socket to a certain address from one process, no other processes on the same machine would be able to use the same address for their sockets until the original process closes its socket (hence, releases the port). And it's kind of reasonable behavior - an interface and port define a packet destination on a machine. Having ambiguous receivers would be bizarre.

要在某个地址上创建服务器套接字 [`listen()`](https://man7.org/linux/man-pages/man2/listen.2.html)，需要显式[`bind()` ](https://man7.org/linux/man-pages/man2/bind.2.html) 到一个接口和端口。长期以来，将套接字绑定到 _(interface, port)_ 对是一种独占操作。如果您将一个套接字绑定到一个进程的某个地址，那么在原始进程关闭其套接字（因此释放端口)之前，同一台机器上的任何其他进程都无法对其套接字使用相同的地址。这是一种合理的行为 - 接口和端口定义了机器上的数据包目的地。有模棱两可的接收者会很奇怪。

But... modern servers may need to handle tens of thousands of TCP connections per second. A single process [`accept()`](https://man7.org/linux/man-pages/man2/accept.2.html)-ing all the client connections quickly becomes a bottleneck with such a high connection rate. So, starting from Linux 3.9, [you can bind an arbitrary number of sockets to exactly the same _(interface, port)_ pair](https://lwn.net/Articles/542629/) as long as all of them use the `SO_REUSEPORT` socket option. The operating system then will make sure that TCP connections are evenly distributed between all the listening processes (or threads).

但是……现代服务器每秒可能需要处理数万个 TCP 连接。单个进程 [`accept()`](https://man7.org/linux/man-pages/man2/accept.2.html)-所有客户端连接很快成为具有如此高连接速率的瓶颈。所以，从 Linux 3.9 开始，[你可以将任意数量的套接字绑定到完全相同的 _(interface, port)_ 对](https://lwn.net/Articles/542629/) 只要它们都使用`SO_REUSEPORT` 套接字选项。然后操作系统将确保 TCP 连接在所有侦听进程（或线程)之间均匀分布。

Apparently, the same technique can be applied to containers. However, the `SO_REUSEPORT` option works only if all the sockets reside in the same network stack. And that's obviously not the case for the default Docker/Podman approach, where every container gets its own network namespace, hence an isolated network stack.

显然，同样的技术可以应用于容器。然而，`SO_REUSEPORT` 选项只有在所有套接字都驻留在同一个网络堆栈中时才有效。而默认的 Docker/Podman 方法显然不是这种情况，每个容器都有自己的网络命名空间，因此是一个隔离的网络堆栈。

The simplest way to overcome this is to sacrifice the isolation a bit and run all the containers in the host's network namespace with `docker run --network host`:

克服这个问题的最简单方法是牺牲一点隔离性，并使用 `docker run --network host` 运行主机网络命名空间中的所有容器：

![Multiple Docker containers listening on the same port with SO_REUSEPORT](http://iximiuz.com/multiple-containers-same-port-reverse-proxy/multiple-containers-same-port-so_reuseport-2000-opt.png)

But there is a more subtle way to share a single network namespace between multiple containers. Docker allows reusing a network namespace of an already existing container while launching a new one. So, we can start a _sandbox_ container that will do nothing but sleep. This container will originate a network namespace and also expose the target port to the host (other namespaces will also be created, but it doesn't really matter). All the application containers will then be attached to this network namespace using `docker run --network container:<sandbox_name>` syntax.

但是有一种更微妙的方法可以在多个容器之间共享单个网络命名空间。 Docker 允许在启动新容器时重用现有容器的网络命名空间。所以，我们可以启动一个 _sandbox_ 容器，它除了睡眠之外什么都不做。这个容器将创建一个网络命名空间，并将目标端口暴露给主机（也会创建其他命名空间，但这并不重要）。然后，所有应用程序容器都将使用 `docker run --network container:<sandbox_name>` 语法附加到这个网络命名空间。

_**NB:** We just reinvented Kubernetes pods here - check out the [Kubernetes CRI](https://github.com/kubernetes/cri-api/blob/d059f89d4bb00d7f29a89808f43e063ce35b50de/pkg/apis/services.go#L63-L79) spec._

_**注意：** 我们刚刚在这里重新发明了 Kubernetes pod - 查看 [Kubernetes CRI](https://github.com/kubernetes/cri-api/blob/d059f89d4bb00d7f29a89808f43e063ce35b50de/pkg/apis/services.go#L63-L79) 规格_

![Multiple Docker containers listening on the same port with SO_REUSEPORT and sandbox container network namespace](http://iximiuz.com/multiple-containers-same-port-reverse-proxy/multiple-containers-same-port-so_reuseport-netns-2000-opt.png)

-2000-opt.png)

Of course, all the instances of the application server need to set the `SO_REUSEPORT` option, so there won't be a port conflict, and the incoming requests will be evenly distributed between the containers listening on the same port.

当然，应用服务器的所有实例都需要设置`SO_REUSEPORT`选项，这样就不会有端口冲突，传入的请求会在监听同一个端口的容器之间均匀分布。

Example Go server using `SO_REUSEPORT` socket option.

示例 Go 服务器使用 SO_REUSEPORT 套接字选项。

Here is an example Go server that uses the `SO_REUSEPORT` option. [Setting socket options in Go](http://iximiuz.com/en/posts/go-net-http-setsockopt-example/) turned out to be slightly less trivial than I expected 🙈

这是一个使用 SO_REUSEPORT 选项的 Go 服务器示例。 [在 Go 中设置套接字选项](http://iximiuz.com/en/posts/go-net-http-setsockopt-example/) 结果比我预期的要简单一些🙈

```go
// http_server.go

package main

import (
    "context"
    "fmt"
    "net"
    "net/http"
    "os"
    "syscall"

    "golang.org/x/sys/unix"
)

func main() {
    lc := net.ListenConfig{
        Control: func(network, address string, conn syscall.RawConn) error {
            var operr error
            if err := conn.Control(func(fd uintptr) {
                operr = syscall.SetsockoptInt(
                    int(fd),
                    unix.SOL_SOCKET,
                    unix.SO_REUSEPORT,
                    1,
                )
            });err != nil {
                return err
            }
            return operr
        },
    }

    ln, err := lc.Listen(
        context.Background(),
        "tcp",
        os.Getenv("HOST")+":"+os.Getenv("PORT"),
    )
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, _req *http.Request) {
        w.Write([]byte(fmt.Sprintf("Hello from %s\n", os.Getenv("INSTANCE"))))
    })

    if err := http.Serve(ln, nil);err != nil {
        panic(err)
    }
}

```

Use the following _Dockerfile_ (not for production!) to containerize the above server:

使用以下 _Dockerfile_（不适用于生产！）来容器化上述服务器：

```dockerfile
FROM golang:1.16

COPY http_server.go .

ENV GO111MODULE=off

RUN go get golang.org/x/sys/unix

CMD ["go", "run", "http_server.go"]

```

Build it with `docker build -t http_server .`

使用`docker build -t http_server 构建它。

Here is a step by step instruction on how to launch the sandbox and the application containers:

以下是有关如何启动沙箱和应用程序容器的分步说明：

```bash
# Prepare the sandbox.
$ docker run -d --rm \
  --name app_sandbox \
  -p 80:8080 \
alpine sleep infinity

# Run first application container.
$ docker run -d --rm \
  --network container:app_sandbox \
  -e INSTANCE=foo \
  -e HOST=0.0.0.0 \
  -e PORT=8080 \
http_server

# Run second application container.
$ docker run -d --rm \
  --network container:app_sandbox \
  -e INSTANCE=bar \
  -e HOST=0.0.0.0 \
  -e PORT=8080 \
http_server

```

Testing it on a vagrant box with `curl` gave me the following results:

在带有 `curl` 的流浪盒子上测试它给了我以下结果：

```bash
# 192.168.37.99 is the external interface of the VM.

$ for i in {1..300};do curl -s 192.168.37.99 2>&1;done |sort |uniq -c
158 Hello from bar
142 Hello from foo

```

Pretty good distribution, huh?

很好的分布，对吧？

## Multiple Containers / Same Port using DNAT

## 使用 DNAT 的多个容器/同一端口

To understand the technique from this section, you need to know what happens when a container's port is published on the host. At first sight, it may look like there is indeed a process listening on the host and forwarding packets to the container (scroll to the right end):

要理解本节中的技术，您需要知道在主机上发布容器的端口时会发生什么。乍一看，看起来确实有一个进程在监听主机并将数据包转发到容器（滚动到右端）：

```bash
# Start a container with port 8080 published on host's port 80.
$ docker run -d --rm -p 80:8080 alpine sleep infinity

# Check listening TCP sockets.
$ sudo ss -lptn 'sport = :80'
State   Recv-Q  Send-Q  Local Address:Port  Peer Address:Port
LISTEN  0       128           0.0.0.0:80          0.0.0.0:*    users:(("docker-proxy",pid=4963,fd=4))
LISTEN  0       128              [::]:80             [::]:*    users:(("docker-proxy",pid=4970,fd=4))

```

[In actuality, the userland `docker-proxy` is rarely used.](https://windsock.io/the-docker-proxy/) Instead, a single iptables rule in the NAT table does all the heavy lifting. Whenever a packet destined to the _(host, published\_port)_ arrives, a destination address translation to _(container, target\_port)_ happens. So, the [port publishing boils down to adding this iptables rule upon the container startup](http://iximiuz.com/en/posts/container-networking-is-simple/#port-publishing).

[实际上，用户空间 `docker-proxy` 很少使用。](https://windsock.io/the-docker-proxy/) 相反，NAT 表中的单个 iptables 规则完成了所有繁重的工作。每当发往 _(host,published\_port)_ 的数据包到达时，就会发生到 _(container, target\_port)_ 的目的地地址转换。因此，[端口发布归结为在容器启动时添加此 iptables 规则](http://iximiuz.com/en/posts/container-networking-is-simple/#port-publishing)。

_**NB:** The iptables trick doesn't cover all the scenarios - for instance, traffic from `localhost` cannot be NAT-ed. So the `docker-proxy` is not fully useless_.

_**注意：** iptables 技巧并没有涵盖所有场景——例如，来自`localhost` 的流量不能被 NAT-ed。所以`docker-proxy` 并不是完全没用的_。

![Docker publishes container port on the host.](http://iximiuz.com/multiple-containers-same-port-reverse-proxy/docker-proxy-2000-opt.png)

The problem with the above DNAT is that it can do only the one-to-one translation. Thus, if we want to support multiple containers behind a single host's port, we need a more sophisticated solution. Luckily, [Kubernetes uses a similar trick in `kube-proxy` for the _Service ClusterIP to Pod IPs_ translation](https://arthurchiao.art/blog/cracking-k8s-node-proxy/) while implementing a [built-in service discovery](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/).

上述DNAT的问题在于它只能进行一对一的翻译。因此，如果我们想在单个主机的端口后面支持多个容器，我们需要一个更复杂的解决方案。幸运的是，[Kubernetes 在`kube-proxy` 中使用了类似的技巧来实现 _Service ClusterIP 到 Pod IPs_ 的转换](https://arthurchio.art/blog/cracking-k8s-node-proxy/) 同时实现了一个 [内置服务发现](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/)。

Long story short, iptables rules can be applied with some probability. So, if you have ten potential destinations for a packet, try applying the destination address translation to the first nine one of them with just a 10% chance. And if none of them worked out, apply a fallback for the very last destination with a 100% chance. As a result, you'll get ten equally loaded destinations.

长话短说，iptables 规则可以以一定的概率应用。因此，如果您有 10 个数据包的潜在目的地，请尝试将目的地地址转换应用于其中的前九个，只有 10% 的机会。如果它们都没有解决，则以 100% 的机会对最后一个目的地应用后备。因此，您将获得 10 个同等负载的目的地。

_**NB:** Of course, iptables are smart enough to apply the DNAT only to the new connections. For an already established connection, an existing address mapping is looked up on the fly._

_**注意：** 当然，iptables 足够聪明，可以将 DNAT 仅应用于新连接。对于已建立的连接，会即时查找现有地址映射。_

![Multiple Docker containers exposed on the same port with iptables NAT rules](http://iximiuz.com/multiple-containers-same-port-reverse-proxy/multiple-containers-same-port-iptables-2000-opt.png)

png)

The huge advantage of this approach comparing to the `SO_REUSEPORT` option is that it's absolutely transparent to the application.

与`SO_REUSEPORT` 选项相比，这种方法的巨大优势在于它对应用程序绝对透明。

Here is how you can quickly try it out.

以下是您可以快速试用的方法。

Example Go server.

示例 Go 服务器。

We can use a much simpler version of the test server:

我们可以使用更简单的测试服务器版本：

```go
// http_server.go
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, _req *http.Request) {
        w.Write([]byte(fmt.Sprintf("Hello from %s\n", os.Getenv("INSTANCE"))))
    })

    if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil);err != nil {
        panic(err)
    }
}

```

And the simplified _Dockerfile_:

以及简化的 _Dockerfile_：

```dockerfile
FROM golang:1.16

COPY http_server.go .

CMD ["go", "run", "http_server.go"]

```

Build it with:

使用以下命令构建它：

```bash
$ docker build -t http_server .

```

Run two application containers:

运行两个应用程序容器：

```bash
$ CONT_PORT=9090

$ docker run -d --rm \
  --name http_server_foo \
  -e INSTANCE=foo \
  -e PORT=$CONT_PORT \
http_server

$ docker run -d --rm \
  --name http_server_bar \
  -e INSTANCE=bar \
  -e PORT=$CONT_PORT \
http_server

```

Figure out the IP addresses of the containers:

找出容器的IP地址：

```bash
$ CONT_FOO_IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http_server_foo)
$ echo $CONT_FOO_IP
172.17.0.2

$ CONT_BAR_IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http_server_bar)
$ echo $CONT_BAR_IP
172.17.0.3

```

Configure iptables DNAT rules - for local ( `OUTPUT`) and external ( `PREROUTING`) traffic:

配置 iptables DNAT 规则 - 用于本地（`OUTPUT`）和外部（`PREROUTING`）流量：

```bash
$ FRONT_PORT=80

# Not secure!Use global ACCEPT only for tests.
$ sudo iptables -I FORWARD 1 -j ACCEPT

# DNAT - local traffic
$ sudo iptables -t nat -I OUTPUT 1 -p tcp --dport $FRONT_PORT \
    -m statistic --mode random --probability 1.0  \
    -j DNAT --to-destination $CONT_FOO_IP:$CONT_PORT

$ sudo iptables -t nat -I OUTPUT 1 -p tcp --dport $FRONT_PORT \
    -m statistic --mode random --probability 0.5  \
    -j DNAT --to-destination $CONT_BAR_IP:$CONT_PORT

# DNAT - external traffic
$ sudo iptables -t nat -I PREROUTING 1 -p tcp --dport $FRONT_PORT \
    -m statistic --mode random --probability 1.0  \
    -j DNAT --to-destination $CONT_FOO_IP:$CONT_PORT

$ sudo iptables -t nat -I PREROUTING 1 -p tcp --dport $FRONT_PORT \
    -m statistic --mode random --probability 0.5  \
    -j DNAT --to-destination $CONT_BAR_IP:$CONT_PORT

```

Testing it on a vagrant box game with `curl` gave me the following request distribution:

使用 `curl` 在一个流浪盒游戏上测试它给了我以下请求分布：

```bash
$ for i in {1..300};do curl -s 192.168.37.99 2>&1;done |sort |uniq -c
143 Hello from bar
157 Hello from foo

```

DON'T FORGET to clean up the iptables rules.

不要忘记清理 iptables 规则。

```bash
# Clean up filter rule (traffic forwarding)
$ sudo iptables -t filter -L -n --line-numbers

# Make sure your rule is on the first line, and then drop it
$ sudo iptables -t filter -D FORWARD 1

# Clean up nat rules (destination address replacement)
sudo iptables -t nat -L -n --line-numbers

# Make sure your rules are the first ones in the chains, and then drop them:
sudo iptables -t nat -D OUTPUT 1
sudo iptables -t nat -D OUTPUT 1
sudo iptables -t nat -D PREROUTING 1
sudo iptables -t nat -D PREROUTING 1

```

## Multiple Containers / Same Port using Reverse Proxy

## 使用反向代理的多个容器/同一端口

When it comes to proxies, the most challenging part is how to deal with volatile sets of upstreams. Containers come and go, and their IP addresses are ephemeral. So, whenever a list of application containers changes, the proxy needs to update its list of upstream IPs. Docker has a built-in DNS server that maintains the up-to-date DNS records for the containers. However, [Nginx doesn't refresh the DNS resolution results unless some magic with variables is applied](https://www.ameyalokare.com/docker/2017/09/27/nginx-dynamic-upstreams-docker.html). Configuring [Envoy proxy with a STRICT\_DNS cluster may show better results](https://www.redhat.com/en/blog/configuring-envoy-auto-discover-pods-kubernetes). However, the best results can be achieved only by [dynamically updating the proxy configuration on the fly listening to Docker events](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/).

说到代理，最具挑战性的部分是如何处理不稳定的上游集。容器来来去去，它们的 IP 地址是短暂的。因此，每当应用程序容器列表发生变化时，代理都需要更新其上游 IP 列表。 Docker 有一个内置的 DNS 服务器，用于维护容器的最新 DNS 记录。但是，[Nginx 不会刷新 DNS 解析结果，除非应用了一些带有变量的魔法](https://www.ameyalokare.com/docker/2017/09/27/nginx-dynamic-upstreams-docker.html)。配置 [使用 STRICT\_DNS 集群的 Envoy 代理可能会显示更好的结果](https://www.redhat.com/en/blog/configuring-envoy-auto-discover-pods-kubernetes)。然而，只有通过[动态更新代理配置以监听 Docker 事件](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-码头工人/)。

![Multiple Docker containers behind a reverse proxy.](http://iximiuz.com/multiple-containers-same-port-reverse-proxy/multiple-containers-same-port-proxy-2000-opt.png)

Luckily, there is a more modern proxy called [Traefik](https://traefik.io/traefik/) with built-in support for many service discovery mechanisms, including labeled Docker containers. If you want to see Traefik in action, check out my write-up on [canary container deployments](http://iximiuz.com/en/posts/traefik-canary-deployments-with-weighted-load-balancing/).

幸运的是，有一个更现代的代理叫做 [Traefik](https://traefik.io/traefik/)，它内置了对许多服务发现机制的支持，包括标记的 Docker 容器。如果您想看到 Traefik 的实际应用，请查看我关于 [canary 容器部署](http://iximiuz.com/en/posts/traefik-canary-deployments-with-weighted-load-balancing/) 的文章。

Well, that's it for now. Make code, not war!

嗯，暂时就这些。编写代码，而不是战争！

#### Resources

####  资源

- [Historical overview of SO\_REUSEADDR and SO\_REUSEPORT options](https://stackoverflow.com/questions/14388706/how-do-so-reuseaddr-and-so-reuseport-differ/14388707#14388707)
- [The SO\_REUSEPORT socket option](https://lwn.net/Articles/542629/) \- LWN article by Michael Kerrisk
- [The docker-proxy](https://windsock.io/the-docker-proxy/) 

- [SO\_REUSEADDR 和 SO\_REUSEPORT 选项的历史概述](https://stackoverflow.com/questions/14388706/how-do-so-reuseaddr-and-so-reuseport-difer/14388707#14388707)
- [SO\_REUSEPORT 套接字选项](https://lwn.net/Articles/542629/) \- Michael Kerrisk 的 LWN 文章
- [码头代理](https://windsock.io/the-docker-proxy/)

- [Cracking Kubernetes node proxy (aka kube-proxy)](https://arthurchiao.art/blog/cracking-k8s-node-proxy/)
- [Dynamic Nginx configuration for Docker with Python](https://www.ameyalokare.com/docker/2017/09/27/nginx-dynamic-upstreams-docker.html)
- [Automated Nginx Reverse Proxy for Docker](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/)
- [nginx-proxy/nginx-proxy](https://github.com/nginx-proxy/nginx-proxy) and [nginx-proxy/docker-gen](https://github.com/nginx-proxy/docker-gen) GitHub projects.
- [Configuring Envoy to Auto-Discover Pods on Kubernetes](https://www.redhat.com/en/blog/configuring-envoy-auto-discover-pods-kubernetes)

- [破解 Kubernetes 节点代理（又名 kube-proxy）](https://arthurchio.art/blog/cracking-k8s-node-proxy/)
- [Docker 与 Python 的动态 Nginx 配置](https://www.ameyalokare.com/docker/2017/09/27/nginx-dynamic-upstreams-docker.html)
- [Docker 的自动化 Nginx 反向代理](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/)
- [nginx-proxy/nginx-proxy](https://github.com/nginx-proxy/nginx-proxy) 和 [nginx-proxy/docker-gen](https://github.com/nginx-proxy/docker-gen) GitHub 项目。
- [配置 Envoy 以在 Kubernetes 上自动发现 Pod](https://www.redhat.com/en/blog/configuring-envoy-auto-discover-pods-kubernetes)

#### Related posts

####  相关文章

- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [Illustrated introduction to Linux iptables](http://iximiuz.com/en/posts/laymans-iptables-101/)
- [Traefik: canary deployments with weighted load balancing](http://iximiuz.com/en/posts/traefik-canary-deployments-with-weighted-load-balancing/)
- [Service Discovery in Kubernetes - Combining the Best of Two Worlds](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/)
- [Writing Web Server in Python: sockets](http://iximiuz.com/en/posts/writing-web-server-in-python-sockets/)
- [Exploring Go net/http Package - On How Not To Set Socket Options](http://iximiuz.com/en/posts/go-net-http-setsockopt-example/)

- [容器网络很简单！](http://iximiuz.com/en/posts/container-networking-is-simple/)
-  [Linux iptables 图解介绍](http://iximiuz.com/en/posts/laymans-iptables-101/)
- [Traefik：具有加权负载平衡的金丝雀部署](http://iximiuz.com/en/posts/traefik-canary-deployments-with-weighted-load-balancing/)
- [Kubernetes 中的服务发现 - 结合两全其美](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/)
- [用 Python 编写 Web 服务器：sockets](http://iximiuz.com/en/posts/writing-web-server-in-python-sockets/)
- [探索 Go net/http 包 - 关于如何不设置套接字选项](http://iximiuz.com/en/posts/go-net-http-setsockopt-example/)

[golang,](javascript: void 0) [socket,](javascript: void 0) [docker,](javascript: void 0) [iptables](javascript: void 0)

[golang,](javascript: void 0) [socket,](javascript: void 0) [docker,](javascript: void 0) [iptables](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)


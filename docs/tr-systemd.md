# Avoiding complexity with systemd

# 使用 systemd 避免复杂性

Saturday, June 26, 2021

Unix machines, since the early days of the operating system, have  been designed for multiple users to use concurrently. Traditionally  there is a set of “unprivileged” users used by people and system  services, and the root account which can generally do anything. Because  of the concept that most things in Unix are represented by a file, users could be allowed to perform various operations by adding them to groups and using filesystem permissions. There were also other functions which could not be delegated in this way–notably, binding to certain IP  ports. Various operating systems developed over the years have blurred  these lines a little, particularly on Linux which has features like  capabilities and ACLs that allow more control than the standard Unix  permission model provides.

Unix 机器，从操作系统的早期开始，就被设计为供多个用户同时使用。传统上有一组供人和系统服务使用的“非特权”用户，以及通常可以做任何事情的 root 帐户。由于 Unix 中的大多数事物都由文件表示的概念，因此可以允许用户通过将它们添加到组和使用文件系统权限来执行各种操作。还有其他功能不能以这种方式委托——特别是绑定到某些 IP 端口。多年来开发的各种操作系统已经稍微模糊了这些界限，特别是在具有功能和 ACL 等特性的 Linux 上，允许比标准 Unix 权限模型提供的更多控制。

Linux goes much further than this. There are mandatory access control  systems like SELinux or AppArmor that let you apply restrictions at the  kernel level outside of the software you're running. Features like  `cgroups` and namespaces combine to provide what we now call containers. Other features like `seccomp` allow software to opt-in to limits on its own ability to use various system calls.

Linux 远不止于此。有强制访问控制系统，如 SELinux 或 AppArmor，可让您在运行的软件之外的内核级别应用限制。 `cgroups` 和命名空间等特性结合起来提供了我们现在所说的容器。 “seccomp”等其他功能允许软件选择限制其自身使用各种系统调用的能力。

BSD operating systems have similar features, notably `pledge` and `unveil` on OpenBSD.

BSD 操作系统具有类似的功能，特别是 OpenBSD 上的“承诺”和“揭幕”。

Much of the original permission model remains to this day. If you  want to run a service that listens on a “well-known” port, those  numbered less than 1024, you generally need to be root to bind to the  port. There are other ways to allow this on Linux, such as applying a  capability called `CAP_NET_BIND_SERVICE` to the program you  want to run, but most server software that is designed to be portable  among Unix systems implements a feature called privilege dropping. The  service initially starts as root, binds to the ports that it requires,  and then calls some functions to set its own user and group IDs to an  unprivileged user. Ideally it does this before doing any significant  work in order to minimise the potential for an exploit to occur while  running as root.

大部分原始权限模型一直保留到今天。如果你想运行一个侦听“众所周知”端口的服务，那些端口号小于 1024，你通常需要以 root 身份绑定到该端口。在 Linux 上还有其他方法可以实现这一点，例如将称为“CAP_NET_BIND_SERVICE”的功能应用于您要运行的程序，但是大多数旨在在 Unix 系统之间移植的服务器软件都实现了一种称为权限删除的功能。该服务最初以 root 身份启动，绑定到它需要的端口，然后调用一些函数将其自己的用户和组 ID 设置为非特权用户。理想情况下，它会在执行任何重要工作之前执行此操作，以最大限度地减少以 root 身份运行时发生漏洞的可能性。

Even if the service isn’t going to bind to a privileged port,  sometimes there are files only readable by root that are required for  the service to operate. A good example of these are the private keys for TLS certificates, which often live in a root-owned directory under `/etc`. It’s common to do a similar trick to access these–start the service as  root, open the files, and then drop the privileges once they’ve been  read. Ideally the service will drop the privileges before parsing  anything. Parser code needs to be written very carefully, especially in  languages with manual memory management, and it has historically been a  source of security vulnerabilities.

即使服务不会绑定到特权端口，有时也有服务运行所需的文件只能由 root 读取。其中一个很好的例子是 TLS 证书的私钥，它通常位于 `/etc` 下的根目录中。使用类似的技巧来访问这些文件是很常见的——以 root 身份启动服务，打开文件，然后在读取它们后删除权限。理想情况下，该服务将在解析任何内容之前放弃特权。解析器代码需要非常仔细地编写，尤其是在具有手动内存管理的语言中，而且它在历史上一直是安全漏洞的来源。

All of this adds a bit of complexity to the services we write, which  it would be nice to avoid. It also adds to the attack surface: privilege dropping code has been a source of vulnerabilities, notably on a couple of occasions in Bash. Avoiding writing it at all, or at least  delegating it to other software with more testing than our own, would be good.

所有这些都为我们编写的服务增加了一些复杂性，最好避免这种情况。它还增加了攻击面：特权丢弃代码一直是漏洞的来源，特别是在 Bash 中的几次。完全避免编写它，或者至少将它委托给比我们自己的测试更多的其他软件，会很好。

It’s fairly common these days to write a service that runs on an  unprivileged port and then run some other software in front of it as a  reverse proxy–often nginx or Apache are used for this purpose. Depending on the use case these may provide some advantages, but they do require  additional configuration and will use some resources on the machines  operating the service.

现在编写一个在非特权端口上运行的服务，然后在它前面运行一些其他软件作为反向代理，这是相当普遍的——通常 nginx 或 Apache 用于此目的。根据用例，这些可能会提供一些优势，但它们确实需要额外的配置，并且会使用运行服务的机器上的一些资源。

# An example

#  一个例子

I’m working on my new startup idea–Lunch as a Service, a new SaaS  application to solve the complex problem of deciding what to have for  lunch. While I’m working on getting the VC funding together, I’m not  going to be deploying it on a huge Kubernetes setup with a service mesh  and a GitOps CI/CD platform. While that would be nice one day, for now  I’m going to make do with a VM.

我正在研究我的新创业理念——午餐即服务，这是一个新的 SaaS 应用程序，用于解决决定午餐吃什么的复杂问题。虽然我正在努力获得 VC 资金，但我不会将其部署在具有服务网格和 GitOps CI/CD 平台的大型 Kubernetes 设置上。虽然有一天那会很好，但现在我将使用虚拟机。

The core of the new service is `lunchd`. It looks like this.

新服务的核心是`lunchd`。它看起来像这样。

```go
package main
import (
        "fmt"
        "math/rand"
        "net/http"
        "time"
)
var lunchOptions = []string{
        "Sandwich",
        "Soup",
        "Salad",
        "Burger",
        "Sushi",
}
func getRandomLunch() string {
        return lunchOptions[rand.Intn(len(lunchOptions))]
}
func main() {
        rand.Seed(time.Now().UnixNano())
        http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
                fmt.Fprintf(w, "<h1>%s</h1>", getRandomLunch())
        })
        fmt.Println("Starting web server on port 8080")
        http.ListenAndServe(":8080", nil)
}
```




I’m using a reasonably recent Linux distribution with `systemd`, so I’m going to use that to start the service when the machine boots. The basic unit file looks like this.

我正在使用带有“systemd”的相当新的 Linux 发行版，因此我将在机器启动时使用它来启动服务。基本单元文件看起来像这样。

```ini
# This goes in /etc/systemd/system/lunchd.service
[Unit]
Description=Lunch as a Service

[Service]
ExecStart=/usr/local/bin/lunchd
User=lunchd

[Install]
WantedBy=multi-user.target
```


I’ve added a user called `lunchd` for the the service. After running `sudo systemctl daemon-reload`, I can then do `sudo systemctl start lunchd`. The service is hard coded to run on port 8080 right now, which is not a privileged port so it’ll start quite happily and run as a standard  user. The final line makes sure it starts when the system boots  normally.

我为该服务添加了一个名为“lunchd”的用户。运行 `sudo systemctl daemon-reload` 后，我可以执行 `sudo systemctl start Lund`。该服务被硬编码为现在在端口 8080 上运行，这不是一个特权端口，所以它会很高兴地启动并作为标准用户运行。最后一行确保它在系统正常启动时启动。

## Getting it ready for production

## 为生产做好准备

I can boot the application and have it run as a user other than root, which is a good start. However, I want it to be able to run on port 443 and use a TLS certificate. I would also like to use some of the  filesystem protection features that systemd advertises. Can I use all of these to avoid writing code to drop privileges?

我可以启动应用程序并让它以 root 以外的用户身份运行，这是一个好的开始。但是，我希望它能够在端口 443 上运行并使用 TLS 证书。我还想使用 systemd 宣传的一些文件系统保护功能。我可以使用所有这些来避免编写代码来放弃特权吗？

### Filesystem sandboxing

### 文件系统沙箱

Our service doesn’t require much in the way of filesystem access. `systemd` provides ways to restrict the parts of the filesystem the service can  see. By doing so, we can limit some of the opportunities for an attacker if the service is compromised.

我们的服务对文件系统访问的要求并不高。 `systemd` 提供了限制服务可以看到的文件系统部分的方法。通过这样做，如果服务受到威胁，我们可以限制攻击者的一些机会。

- `ProtectSystem` can be set to `true` to make `/usr` and `/boot` or `/efi` read-only for this process. If set to `full`, `/etc` is read-only too. `strict` makes the entire filesystem hierarchy read-only. This is fine for this  service as it doesn’t read anything, so we’ll enable that.
- `ProtectHome` can be set to `true` to make `/home`, `/root` and `/run/user` empty and inaccessible from the point of view of the service.
- `PrivateTmp` makes sure that the process’s temp  directories are only visible to itself, and not another process. Additionally, they’ll be emptied once the process finishes.

- `ProtectSystem` 可以设置为 `true` 以使 `/usr` 和 `/boot` 或 `/efi` 在这个进程中只读。如果设置为 `full`，`/etc` 也是只读的。 `strict` 使整个文件系统层次结构只读。这对这个服务很好，因为它不读取任何东西，所以我们将启用它。
- `ProtectHome` 可以设置为 `true` 以使 `/home`、`/root` 和 `/run/user` 为空并且从服务的角度来看无法访问。
- `PrivateTmp` 确保进程的临时目录只对它自己可见，而不对另一个进程可见。此外，一旦过程完成，它们将被清空。

The complete list of options can be found [in the systemd documentation](https://www.freedesktop.org/software/systemd/man/systemd.exec.html).

完整的选项列表可以在 [systemd 文档](https://www.freedesktop.org/software/systemd/man/systemd.exec.html) 中找到。

Let’s add some of these features to our unit file.

让我们将其中一些功能添加到我们的单元文件中。

```ini
[Unit]
Description=Lunch as a Service

[Service]
ExecStart=/usr/local/bin/lunchd
ProtectSystem=strict
ProtectHome=true
PrivateUsers=true
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```


There are far more options available here than I really want to get  into describing. The documentation linked above is a good reference. To  get some ideas for a service of your own, you can try a command similar  to

这里提供的选项比我真正想描述的要多得多。上面链接的文档是一个很好的参考。要获得您自己的服务的一些想法，您可以尝试类似于

```bash
sudo systemd-analyze security lunchd.service --no-pager
```


This gives you an overview of what security features systemd has, and which are enabled. It is probably not a matter of just enabling  everything it mentions, as then your service may not be able to do very  much–instead, you should consider these as suggestions and see if they  will work for you.

这使您可以大致了解 systemd 具有哪些安全功能，以及哪些已启用。这可能不仅仅是启用它提到的所有内容，因为那样您的服务可能无法做很多事情 - 相反，您应该将这些视为建议，看看它们是否适合您。

### Dynamic users

### 动态用户

Traditionally, many services have expected to run as a defined user  and group, usually specified in their configuration file. Some will  start and run as that user if they don’t need to do anything privileged, but others will start as root and switch. Many common internet services such as web and mail servers operate in this way.

传统上，许多服务期望作为定义的用户和组运行，通常在其配置文件中指定。如果他们不需要做任何特权操作，有些人会以该用户身份启动和运行，但其他人将以 root 身份启动并切换。许多常见的 Internet 服务（例如 Web 和邮件服务器）都以这种方式运行。

With systemd we can have a brand new user and group allocated for us  when the service starts. These users are prevented from changing any  state in the system except in the directories detailed above. Once the  service exits, the user is removed–it never actually exists in `/etc/passwd`. This makes it harder to create a persistent route in to a system after an exploit.

使用 systemd，我们可以在服务启动时为我们分配一个全新的用户和组。除了上面详述的目录之外，这些用户无法更改系统中的任何状态。一旦服务退出，用户就会被删除——它实际上从未存在于 `/etc/passwd` 中。这使得在漏洞利用后创建进入系统的持久路由变得更加困难。

```ini
[Unit]
Description=Lunch as a Service

[Service]
ExecStart=/usr/local/bin/lunchd
ProtectSystem=strict
ProtectHome=true
PrivateUsers=true
PrivateTmp=true

DynamicUser=yes

[Install]
WantedBy=multi-user.target
```


For this I’ve removed the system user I added earlier. After another `daemon-reload` and restarting the service, we can see that the service is still running as a user called `lunchd`, but it has a very high user ID:

为此，我删除了之前添加的系统用户。再一次`daemon-reload`并重新启动服务后，我们可以看到该服务仍然以名为`lunchd`的用户身份运行，但它的用户ID非常高：

```
mgdm@lunchbox:~$ ps aux |grep lunchd
lunchd     33146  0.2  0.4 1002788 4312 ?Ssl  20:28   0:00 /usr/local/bin/lunchd
mgdm@lunchbox:~$ id lunchd
uid=62840(lunchd) gid=62840(lunchd) groups=62840(lunchd)
```




None of these changes have, so far, required much in the way of  modification to our source code. At the most, we just need to make sure  our service can accept some configuration for the directories it writes  to.

到目前为止，这些更改都不需要对我们的源代码进行太多修改。最多，我们只需要确保我们的服务可以接受它写入的目录的一些配置。

### Using port 80

### 使用端口 80

Although we’ve added some security features, our service is still  running on port 8080. That’s fine for development purposes, but we’ve  run into the first of the problems I described at the start. We want to  bind to a privileged port, but the service is now definitely running as  an unprivileged user.

虽然我们添加了一些安全功能，但我们的服务仍然在 8080 端口上运行。这对于开发目的来说很好，但是我们遇到了我在开始时描述的第一个问题。我们想绑定到特权端口，但该服务现在肯定是作为非特权用户运行的。

`systemd` has a feature called socket activation, which  allows it to bind to a port and then hand a file descriptor for that  port to a process it launches. It can either launch the process once per connection, in a similar fashion to the old [inetd daemon](https://en.wikipedia.org/wiki/Inetd), or it can bind to the port once and hand it to a persistent process. This latter option is what we’ll do. It does require some modification  to the code, but none of it will ever run as root which further reduces  the attack surface.

`systemd` 有一个称为套接字激活的功能，它允许它绑定到一个端口，然后将该端口的文件描述符传递给它启动的进程。它可以以类似于旧的 [inetd 守护进程](https://en.wikipedia.org/wiki/Inetd) 的方式每个连接启动一次进程，或者它可以绑定到端口一次并将其交给一个持续的过程。后一个选项就是我们要做的。它确实需要对代码进行一些修改，但它不会以 root 身份运行，这进一步减少了攻击面。

It’s certainly possible to hand-write this code but in this case I’m  going to use a package from CoreOS. If you'd like to see a minimal  hand-written version, check out [this one from Lennart Poettering](https://github.com/systemd/portable-walkthrough-go/blob/master/main.go#L15-L31).

手写这个代码当然是可能的，但在这种情况下，我将使用 CoreOS 的一个包。如果您想查看最小的手写版本，请查看 [来自 Lennart Poettering 的这个](https://github.com/systemd/portable-walkthrough-go/blob/master/main.go#L15-L31)。

Here’s the modified service.

这是修改后的服务。

```go
package main
import (
        "fmt"
        "log"
        "math/rand"
        "net"
        "net/http"
        "time"
        "github.com/coreos/go-systemd/activation"
)
var lunchOptions = []string{
        "Sandwich",
        "Soup",
        "Salad",
        "Burger",
        "Sushi",
}
func getRandomLunch() string {
        return lunchOptions[rand.Intn(len(lunchOptions))]
}
func getListener() (net.Listener, error) {
        listeners, err := activation.Listeners()
        if err != nil ||len(listeners) != 1 {
                log.Printf("Excpected one listener, got %d: %s", len(listeners), err)
                listener, err := net.Listen("tcp", ":8080")
                return listener, err
        }
        return listeners[0], err
}
func main() {
        rand.Seed(time.Now().UnixNano())
        listener, err := getListener()
        if err != nil {
                log.Panicf("Could not set up listener: %s", err)
        }
        http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
                fmt.Fprintf(w, "<h1>%s</h1>", getRandomLunch())
        })
        log.Printf("Starting web server on port %s", listener.Addr().String())
}
```


Once we’ve built the new version, we can test the freshly-built version before deploying it:

构建新版本后，我们可以在部署之前测试新构建的版本：

```
mgdm@lunchbox:~$ systemd-socket-activate -l 8080 ./main
Listening on [::]:8080 as 3.
Communication attempt on fd 3.
Execing ./main (./main)
Starting web server on port 8080
```


This runs fine. To make it work this way permanently, we need to  create a socket file for systemd in addition to the service file. It  describes the port you want the systemd to bind to and pass to the  service. This can go in `/etc/systemd/system/lunchd.socket`.

这运行良好。为了让它永久地以这种方式工作，除了服务文件之外，我们还需要为 systemd 创建一个套接字文件。它描述了您希望 systemd 绑定到并传递给服务的端口。这可以进入`/etc/systemd/system/lunchd.socket`。

```ini
[Socket]
ListenStream = 80
BindIPv6Only = both
Accept=no

[Install]
WantedBy = sockets.target
```


After another `systemctl daemon-reload`, you can then type `systemctl start lunchd.socket` which will make it listen to the port.

在另一个 `systemctl daemon-reload` 之后，你可以输入 `systemctl start Lunchd.socket` 让它监听端口。

The `Accept` option provides a choice between systemd starting the service once and handing it the listening sockets (`Accept=no`), or starting a new instance of the service on each request (`Accept=yes`). This latter mode is quite like how `inetd` operates.

`Accept` 选项提供了一种选择，即 systemd 启动服务一次并将其交给侦听套接字（`Accept=no`），或者在每个请求上启动服务的新实例（`Accept=yes`）。后一种模式很像 `inetd` 的运作方式。

Another advantage of having systemd listen on these ports is that if  the service crashes, requests can still come in on port 80 and systemd  will take care of starting a new one for you. Theoretically at least,  this means you shouldn’t drop those requests.

让 systemd 监听这些端口的另一个好处是，如果服务崩溃，请求仍然可以在端口 80 上传入，systemd 会负责为您启动一个新的端口。至少从理论上讲，这意味着您不应该放弃这些请求。

### Adding HTTPS 

### 添加 HTTPS

For this we’ll need a certificate, which most of the time we can get  from LetsEncrypt. However, most LetsEncrypt clients will place the  private key somewhere in `/etc` in a directory only  accessible by root. This is fine, as having the private key owned by  another user makes it harder to steal in the event our service is  compromised. (We’ll just have to hope we don’t have [some kind of memory leak attack](https://en.wikipedia.org/wiki/Cloudbleed)). We want to be able to use the certificate, but we don’t want to use privilege dropping here.

为此，我们需要一个证书，大多数情况下我们可以从 LetsEncrypt 获得。但是，大多数 LetsEncrypt 客户端会将私钥放在 `/etc` 中的某个目录中，该目录只能由 root 访问。这很好，因为如果我们的服务受到损害，拥有另一个用户拥有的私钥会更难窃取。 （我们只希望我们没有[某种内存泄漏攻击](https://en.wikipedia.org/wiki/Cloudbleed))。我们希望能够使用证书，但我们不想在这里使用权限下降。

Normally for this, we’d use code something like the following, [taken from the Go documentation](https://golang.org/src/crypto/tls/example_test.go#L114):

通常为此，我们会使用类似以下的代码，[取自 Go 文档](https://golang.org/src/crypto/tls/example_test.go#L114)：

```go
func ExampleLoadX509KeyPair() {
    cert, err := tls.LoadX509KeyPair("example-cert.pem", "example-key.pem")
    if err != nil {
        log.Fatal(err)
    }
    cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
    listener, err := tls.Listen("tcp", ":2000", cfg)
    if err != nil {
        log.Fatal(err)
    }
    _ = listener
}
```


We can pretty much use this, although the paths are hard-coded so  this assumes the process is going to be able to open the files from  whatever user it’s running as. I’ve used `certbot` to create a LetsEncrypt certificate, which is stored in `/etc/letsencrypt/live/lunch.mgdm.net` and is only readable by root.

我们几乎可以使用它，尽管路径是硬编码的，因此假设该进程将能够从它运行的任何用户打开文件。我使用 `certbot` 创建了一个 LetsEncrypt 证书，它存储在 `/etc/letsencrypt/live/lunch.mgdm.net` 中，并且只能被 root 读取。

We can work around this using systemd. It can open the files and  present them in a modified filesystem view to the service. We’ll need to make some more changes to our code, but we were going to have to do  this anyway to add the TLS. The only thing to remember is to make the  paths to those files configurable. In this case, I’ve added flags called `key` and `certificate` which provide the paths to the TLS private key and the certificate with its chain.

我们可以使用 systemd 解决这个问题。它可以打开文件并将它们以修改后的文件系统视图呈现给服务。我们需要对代码进行更多更改，但无论如何我们都必须这样做才能添加 TLS。唯一要记住的是使这些文件的路径可配置。在这种情况下，我添加了名为“key”和“certificate”的标志，它们提供了到 TLS 私钥和证书及其链的路径。

The key thing here is the `LoadCredential` setting. It takes both a key and a path, separated by `:`. systemd will load the contents of the file at the specified path, and  expose it in a directory as a file named after the key. This directory  is given to the process as an environment variable called `CREDENTIALS_DIRECTORY`, which can also be used in the `ExecStart` line of the unit file. You can see this in action in the modified unit  file below. We’ve also made a change to the socket file in order to make it listen on port 443 instead.

这里的关键是“LoadCredential”设置。它需要一个键和一个路径，用`:`分隔。 systemd 将在指定路径加载文件的内容，并将其作为以密钥命名的文件公开在目录中。该目录作为名为“CREDENTIALS_DIRECTORY”的环境变量提供给进程，它也可以在单元文件的“ExecStart”行中使用。您可以在下面修改后的单元文件中看到这一点。我们还对套接字文件进行了更改，使其改为侦听端口 443。

#### lunchd.service

#### 午餐服务

```ini
[Unit]
Description=Lunch as a Service

[Service]
ExecStart=/usr/local/bin/lunchd -key=${CREDENTIALS_DIRECTORY}/key.pem -certificate=${CREDENTIALS_DIRECTORY}/chain.pem

LoadCredential=key.pem:/etc/letsencrypt/live/lunch.mgdm.net/privkey.pem
LoadCredential=chain.pem:/etc/letsencrypt/live/lunch.mgdm.net/fullchain.pem

ProtectSystem=strict
ProtectHome=true
PrivateUsers=true
PrivateTmp=true
DynamicUser=yes

[Install]
WantedBy=multi-user.target
```


#### lunchd.socket

```ini
[Socket]
ListenStream = 443
BindIPv6Only = both

[Install]
WantedBy = sockets.target
```


#### main.go

#### main.go

```go
package main
import (
     "crypto/tls"
     "flag"
     "fmt"
     "log"
     "math/rand"
     "net"
     "net/http"
     "time"
     "github.com/coreos/go-systemd/activation"
)
var lunchOptions = []string{
     "Sandwich",
     "Soup",
     "Salad",
     "Burger",
     "Sushi",
}
func getRandomLunch() string {
     return lunchOptions[rand.Intn(len(lunchOptions))]
}
func getListener() (net.Listener, error) {
     listeners, err := activation.Listeners()
     if err != nil || len(listeners) != 1 {
         log.Printf("Excpected one listener, got %d: %s", len(listeners), err)
         listener, err := net.Listen("tcp", ":8080")
         return listener, err
     }
     return listeners[0], err
}
func getCertificates() (string, string, error) {
     keyPath := flag.String("key", "", "The path to the private key")
     certPath := flag.String("certificate", "", "The path to the certificate")
     flag.Parse()
     if *keyPath == "" || *certPath == "" {
         return "", "", fmt.Errorf("Either or both of -key or -cert not set")
     }
     return *keyPath, *certPath, nil
}
func tlsSetup(keyPath string, certPath string, listener net.Listener) (net.Listener, error) {
     config := &tls.Config{
         Certificates:             make([]tls.Certificate, 1),
         NextProtos:               []string{"h2", "http/1.1"},
         PreferServerCipherSuites: true,
     }
     var err error
     log.Printf("Loading certs from key: %s and cert: %s", keyPath, certPath)
     config.Certificates[0], err = tls.LoadX509KeyPair(
         certPath,
         keyPath,
     )
     if err != nil {
         log.Printf("Failed to configure TLS: %s", err)
         return nil, err
     }
     return tls.NewListener(listener, config), nil
}
func main() {
     rand.Seed(time.Now().UnixNano())
     listener, err := getListener()
     if err != nil {
         log.Fatalf("Could not set up listener: %s", err)
     }
     keyPath, certPath, err := getCertificates()
     if err != nil {
         log.Fatalf("Could not load certificates: %s", err)
     }
     tlsListener, err := tlsSetup(keyPath, certPath, listener)
     http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
         fmt.Fprintf(w, "<h1>%s</h1>", getRandomLunch())
     })
     fmt.Printf("Starting web server on port %s", tlsListener.Addr().String())
     http.Serve(tlsListener, nil)
}
```

Once we make these changes, compile it, and deploy it, we can see the following in the journal by running `sudo journalctl -xe`:

一旦我们进行了这些更改、编译和部署，我们可以通过运行 `sudo journalctl -xe` 在日志中看到以下内容：

```
Jun 26 12:05:35 lunchbox lunchd[943]: 2021/06/26 12:05:35 Loading certs from key: /run/credentials/lunchd.service/key.pem and cert: /run/credentials/lunchd.service/chain.pem
```


This shows that the `CREDENTIALS_DIRECTORY` has been set to `/run/credentials/lunchd.service` in this case, and the private key and chain are exposed there. To make  this work, no changes have been made to the code beyond that required to enable TLS, which would be common to any platform I ran this code on. The only thing that I’ve done is made sure that I can specify the paths  to the key and certificate on the command line. The systemd-specific `CREDENTIALS_DIRECTORY` is only referenced in the unit file.

这表明在这种情况下`CREDENTIALS_DIRECTORY`已设置为`/run/credentials/lunchd.service`，并且私钥和链在那里暴露。为了完成这项工作，除了启用 TLS 所需的代码之外，没有对代码进行任何更改，这对于我运行此代码的任何平台都是通用的。我所做的唯一一件事是确保我可以在命令行上指定密钥和证书的路径。 systemd 特定的 `CREDENTIALS_DIRECTORY` 仅在单元文件中引用。

I could modify the service further so that systemd sets up listeners  on both ports 80 and 443, so the service deals with both HTTP and HTTPS  itself, but I think I’ll leave that as an exercise for the reader.

我可以进一步修改该服务，以便 systemd 在端口 80 和 443 上设置侦听器，因此该服务本身处理 HTTP 和 HTTPS，但我想我将把它留给读者作为练习。

# References

#  参考

I found the following articles helpful.

我发现以下文章很有帮助。

1. [The documentation for systemd.exec](https://www.freedesktop.org/software/systemd/man/systemd.exec.html)
2. [Options for hardening systemd service units](https://gist.github.com/ageis/f5595e59b1cddb1513d1b425a323db04)
3. [Integration of a Go service with systemd: socket activation by Vincent Bernat](https://vincent.bernat.ch/en/blog/2018-systemd-golang-socket-activation)
4. [The Debian docs on Service Sandboxing](https://wiki.debian.org/ServiceSandboxing)
5. [Walkthrough for Portable Services in Go](http://0pointer.net/blog/walkthrough-for-portable-services-in-go.html)

1. [systemd.exec的文档](https://www.freedesktop.org/software/systemd/man/systemd.exec.html)
2. [加固systemd服务单元的选项](https://gist.github.com/ageis/f5595e59b1cddb1513d1b425a323db04)
3. [Go 服务与 systemd 的集成：Vincent Bernat 的套接字激活](https://vincent.bernat.ch/en/blog/2018-systemd-golang-socket-activation)
4. [关于服务沙盒的 Debian 文档](https://wiki.debian.org/ServiceSandboxing)
5.  [Go 中的可移植服务演练](http://0pointer.net/blog/walkthrough-for-portable-services-in-go.html)

[The code is on GitHub](https://github.com/mgdm/lunchd) in case it’s useful to anyone. 

[代码在 GitHub 上](https://github.com/mgdm/lunchd) 以防它对任何人有用。


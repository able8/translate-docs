# Avoiding complexity with systemd

Saturday, June 26, 2021

Unix machines, since the early days of the operating system, have  been designed for multiple users to use concurrently. Traditionally  there is a set of “unprivileged” users used by people and system  services, and the root account which can generally do anything. Because  of the concept that most things in Unix are represented by a file, users could be allowed to perform various operations by adding them to groups and using filesystem permissions. There were also other functions which could not be delegated in this way–notably, binding to certain IP  ports. Various operating systems developed over the years have blurred  these lines a little, particularly on Linux which has features like  capabilities and ACLs that allow more control than the standard Unix  permission model provides. 

Linux goes much further than this. There are mandatory access control  systems like SELinux or AppArmor that let you apply restrictions at the  kernel level outside of the software you're running. Features like  `cgroups` and namespaces combine to provide what we now call containers. Other features like `seccomp` allow software to opt-in to limits on its own ability to use various system calls. 

BSD operating systems have similar features, notably `pledge` and `unveil` on OpenBSD.

Much of the original permission model remains to this day. If you  want to run a service that listens on a “well-known” port, those  numbered less than 1024, you generally need to be root to bind to the  port. There are other ways to allow this on Linux, such as applying a  capability called `CAP_NET_BIND_SERVICE` to the program you  want to run, but most server software that is designed to be portable  among Unix systems implements a feature called privilege dropping. The  service initially starts as root, binds to the ports that it requires,  and then calls some functions to set its own user and group IDs to an  unprivileged user. Ideally it does this before doing any significant  work in order to minimise the potential for an exploit to occur while  running as root. 

Even if the service isn’t going to bind to a privileged port,  sometimes there are files only readable by root that are required for  the service to operate. A good example of these are the private keys for TLS certificates, which often live in a root-owned directory under `/etc`. It’s common to do a similar trick to access these–start the service as  root, open the files, and then drop the privileges once they’ve been  read. Ideally the service will drop the privileges before parsing  anything. Parser code needs to be written very carefully, especially in  languages with manual memory management, and it has historically been a  source of security vulnerabilities.

All of this adds a bit of complexity to the services we write, which  it would be nice to avoid. It also adds to the attack surface: privilege dropping code has been a source of vulnerabilities, notably on a couple of occasions in Bash. Avoiding writing it at all, or at least  delegating it to other software with more testing than our own, would be good.

It’s fairly common these days to write a service that runs on an  unprivileged port and then run some other software in front of it as a  reverse proxy–often nginx or Apache are used for this purpose. Depending on the use case these may provide some advantages, but they do require  additional configuration and will use some resources on the machines  operating the service.

# An example

I’m working on my new startup idea–Lunch as a Service, a new SaaS  application to solve the complex problem of deciding what to have for  lunch. While I’m working on getting the VC funding together, I’m not  going to be deploying it on a huge Kubernetes setup with a service mesh  and a GitOps CI/CD platform. While that would be nice one day, for now  I’m going to make do with a VM.

The core of the new service is `lunchd`. It looks like this.

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

## Getting it ready for production

I can boot the application and have it run as a user other than root, which is a good start. However, I want it to be able to run on port 443 and use a TLS certificate. I would also like to use some of the  filesystem protection features that systemd advertises. Can I use all of these to avoid writing code to drop privileges?

### Filesystem sandboxing

Our service doesn’t require much in the way of filesystem access. `systemd` provides ways to restrict the parts of the filesystem the service can  see. By doing so, we can limit some of the opportunities for an attacker if the service is compromised.

- `ProtectSystem` can be set to `true` to make `/usr` and `/boot` or `/efi` read-only for this process. If set to `full`, `/etc` is read-only too. `strict` makes the entire filesystem hierarchy read-only. This is fine for this  service as it doesn’t read anything, so we’ll enable that.
- `ProtectHome` can be set to `true` to make `/home`, `/root` and `/run/user` empty and inaccessible from the point of view of the service.
- `PrivateTmp` makes sure that the process’s temp  directories are only visible to itself, and not another process.  Additionally, they’ll be emptied once the process finishes.

The complete list of options can be found [in the systemd documentation](https://www.freedesktop.org/software/systemd/man/systemd.exec.html).

Let’s add some of these features to our unit file.

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

```bash
sudo systemd-analyze security lunchd.service --no-pager
```

This gives you an overview of what security features systemd has, and which are enabled. It is probably not a matter of just enabling  everything it mentions, as then your service may not be able to do very  much–instead, you should consider these as suggestions and see if they  will work for you.

### Dynamic users

Traditionally, many services have expected to run as a defined user  and group, usually specified in their configuration file. Some will  start and run as that user if they don’t need to do anything privileged, but others will start as root and switch. Many common internet services such as web and mail servers operate in this way. 

With systemd we can have a brand new user and group allocated for us  when the service starts. These users are prevented from changing any  state in the system except in the directories detailed above. Once the  service exits, the user is removed–it never actually exists in `/etc/passwd`. This makes it harder to create a persistent route in to a system after an exploit.

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

```
mgdm@lunchbox:~$ ps aux | grep lunchd
lunchd     33146  0.2  0.4 1002788 4312 ?        Ssl  20:28   0:00 /usr/local/bin/lunchd
mgdm@lunchbox:~$ id lunchd
uid=62840(lunchd) gid=62840(lunchd) groups=62840(lunchd)
```

None of these changes have, so far, required much in the way of  modification to our source code. At the most, we just need to make sure  our service can accept some configuration for the directories it writes  to.

### Using port 80

Although we’ve added some security features, our service is still  running on port 8080. That’s fine for development purposes, but we’ve  run into the first of the problems I described at the start. We want to  bind to a privileged port, but the service is now definitely running as  an unprivileged user. 

`systemd` has a feature called socket activation, which  allows it to bind to a port and then hand a file descriptor for that  port to a process it launches. It can either launch the process once per connection, in a similar fashion to the old [inetd daemon](https://en.wikipedia.org/wiki/Inetd), or it can bind to the port once and hand it to a persistent process.  This latter option is what we’ll do. It does require some modification  to the code, but none of it will ever run as root which further reduces  the attack surface.

It’s certainly possible to hand-write this code but in this case I’m  going to use a package from CoreOS. If you’d like to see a minimal  hand-written version, check out [this one from Lennart Poettering](https://github.com/systemd/portable-walkthrough-go/blob/master/main.go#L15-L31).

Here’s the modified service.

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
        if err != nil || len(listeners) != 1 {
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

```
mgdm@lunchbox:~$ systemd-socket-activate -l 8080 ./main
Listening on [::]:8080 as 3.
Communication attempt on fd 3.
Execing ./main (./main)
Starting web server on port 8080
```

This runs fine. To make it work this way permanently, we need to  create a socket file for systemd in addition to the service file. It  describes the port you want the systemd to bind to and pass to the  service. This can go in `/etc/systemd/system/lunchd.socket`.

```ini
[Socket]
ListenStream = 80
BindIPv6Only = both
Accept=no

[Install]
WantedBy = sockets.target
```

After another `systemctl daemon-reload`, you can then type `systemctl start lunchd.socket` which will make it listen to the port.

The `Accept` option provides a choice between systemd starting the service once and handing it the listening sockets (`Accept=no`), or starting a new instance of the service on each request (`Accept=yes`). This latter mode is quite like how `inetd` operates.

Another advantage of having systemd listen on these ports is that if  the service crashes, requests can still come in on port 80 and systemd  will take care of starting a new one for you. Theoretically at least,  this means you shouldn’t drop those requests.

### Adding HTTPS

For this we’ll need a certificate, which most of the time we can get  from LetsEncrypt. However, most LetsEncrypt clients will place the  private key somewhere in `/etc` in a directory only  accessible by root. This is fine, as having the private key owned by  another user makes it harder to steal in the event our service is  compromised. (We’ll just have to hope we don’t have [some kind of memory leak attack](https://en.wikipedia.org/wiki/Cloudbleed)). We want to be able to use the certificate, but we don’t want to use privilege dropping here.

Normally for this, we’d use code something like the following, [taken from the Go documentation](https://golang.org/src/crypto/tls/example_test.go#L114):

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

We can work around this using systemd. It can open the files and  present them in a modified filesystem view to the service. We’ll need to make some more changes to our code, but we were going to have to do  this anyway to add the TLS. The only thing to remember is to make the  paths to those files configurable. In this case, I’ve added flags called `key` and `certificate` which provide the paths to the TLS private key and the certificate with its chain.

The key thing here is the `LoadCredential` setting. It takes both a key and a path, separated by `:`. systemd will load the contents of the file at the specified path, and  expose it in a directory as a file named after the key. This directory  is given to the process as an environment variable called `CREDENTIALS_DIRECTORY`, which can also be used in the `ExecStart` line of the unit file. You can see this in action in the modified unit  file below. We’ve also made a change to the socket file in order to make it listen on port 443 instead.

#### lunchd.service

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

```
Jun 26 12:05:35 lunchbox lunchd[943]: 2021/06/26 12:05:35 Loading certs from key: /run/credentials/lunchd.service/key.pem and cert: /run/credentials/lunchd.service/chain.pem
```

This shows that the `CREDENTIALS_DIRECTORY` has been set to `/run/credentials/lunchd.service` in this case, and the private key and chain are exposed there. To make  this work, no changes have been made to the code beyond that required to enable TLS, which would be common to any platform I ran this code on.  The only thing that I’ve done is made sure that I can specify the paths  to the key and certificate on the command line. The systemd-specific `CREDENTIALS_DIRECTORY` is only referenced in the unit file. 

I could modify the service further so that systemd sets up listeners  on both ports 80 and 443, so the service deals with both HTTP and HTTPS  itself, but I think I’ll leave that as an exercise for the reader.

# References

I found the following articles helpful.

1. [The documentation for systemd.exec](https://www.freedesktop.org/software/systemd/man/systemd.exec.html)
2. [Options for hardening systemd service units](https://gist.github.com/ageis/f5595e59b1cddb1513d1b425a323db04)
3. [Integration of a Go service with systemd: socket activation by Vincent Bernat](https://vincent.bernat.ch/en/blog/2018-systemd-golang-socket-activation)
4. [The Debian docs on Service Sandboxing](https://wiki.debian.org/ServiceSandboxing)
5. [Walkthrough for Portable Services in Go](http://0pointer.net/blog/walkthrough-for-portable-services-in-go.html)

[The code is on GitHub](https://github.com/mgdm/lunchd) in case it’s useful to anyone.


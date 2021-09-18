# Building containers without Docker

25 January 2020

In this post I'll outline several ways to build containers without the need for Docker itself. I'll use [OpenFaaS](https://github.com/openfaas/) as the case-study, which uses OCI-format container images for its workloads. The easiest way to think about OpenFaaS is as a CaaS platform for [Kubernetes](https://kubernetes.io) which can run microservices, and add in FaaS and event-driven tooling for free.

See also [OpenFaaS.com](https://openfaas.com/)

The first option in the post will show how to use the built-in buildkit option for Docker's CLI, then [buildkit](https://github.com/moby/buildkit) stand-alone (on Linux only), followed by Google's container builder, [Kaniko](https://github.com/GoogleContainerTools/kaniko).

This post covers tooling which can build an image from a Dockerfile, and so anything which limits the user to only Java (jib) or Go (ko) for instance is out of scope. I'll then wrap things up and let you know how to get in touch with suggestions, feedback and your own stories around wants and needs in container tooling.

## So what's wrong with Docker?

Nothing as such, Docker runs well on armhf, arm64, and on `x86_64`. The main Docker CLI has become a lot more than build/ship/run, and also lugs around several years of baggage, it now comes bundled with Docker Swarm and EE features.

> Update for Nov 2020: anyone using Docker's set of official base-images should also read: [Preparing for the Docker Hub Rate Limits](https://inlets.dev/blog/2020/10/29/preparing-docker-hub-rate-limits.html)

### Alternatives to Docker

There are a few efforts that attempt to strip "docker" back to its component pieces, the original UX we all fell in love with:

- [Docker](https://github.com/docker/docker) \- docker itself now uses containerd to run containers, and has support for enabling buildkit to do highly efficient, caching builds.

- [Podman](https://podman.io/) and [buildah](https://github.com/containers/buildah) combination - RedHat / IBM's effort, which uses their own OSS toolchain to generate OCI images. Podman is marketed as being daemonless and rootless, but still ends up having to mount overlay filesystems and use a UNIX socket.

- [pouch](https://github.com/alibaba/pouch) \- from Alibaba, pouch is billed as "An Efficient Enterprise-class Container Engine". It uses containerd just like Docker, and supports both container-level isolation with [runc](https://github.com/opencontainers/runc) and "lightweight VMs" such as [runV](https://github.com/hyperhq/runv). There's also more of a [focus on image distribution and strong isolation](https://github.com/alibaba/pouch/blob/master/docs/architecture.md).

- Stand-alone buildkit - buildkit was started by [Tõnis Tiigi](https://twitter.com/tonistiigi?lang=en) from Docker Inc as a brand new container builder with caching and concurrency in mind. buildkit currently only runs as a daemon, but you will hear people claim otherwise. They are forking the daemon and then killing it after a build.

- [img](https://github.com/genuinetools/img) \- img was written by [Jess Frazelle](https://github.com/jessfraz) and is often quoted in these sorts of guides and is a wrapper for buildkit. That said, I haven't seen traction with it compared to the other options mentioned. The project was quite active [until late 2018 and has only received a few patches since](https://github.com/genuinetools/img/commits/master). img claims to be daemonless, but it uses buildkit so is probably doing some trickery there. I hear that `img` gives a better UX than buildkit's own CLI `buildctr`, but it should also be noted that img is only released for `x86_64` and there are no binaries for armhf / arm64.


> An alternative to `img` would be `k3c` which also includes a runtime component and plans to support ARM architectures.

- [k3c](https://github.com/ibuildthecloud/k3c) \- Rancher's latest experiment which uses containerd and buildkit to re-create the original, classic, vanilla, lite experience of the original Docker version.

Out of all the options, I think that I like k3c the most, but it is very nascient and bundles everything into one binary which is likely to conflict with other software, at present it runs its own embedded containerd and buildkit binaries.

> Note: If you're a RedHat customer and paying for support, then you really should use their entire toolchain to get the best value for your money. I checked out some of the examples and saw one that used my "classic" blog post on multi-stage builds. See for yourself which style you prefer [the buildah example](https://github.com/containers/buildah/blob/master/demos/buildah_multi_stage.sh) vs. [Dockerfile example](https://blog.alexellis.io/mutli-stage-docker-builds).

So since we are focusing on the "build" piece here and want to look at relativelt stable options, I'm going to look at:

- buildkit in Docker,
- buildkit stand-alone
- and kaniko.

All of the above and more are now possible since the OpenFaaS CLI can output a standard "build context" that any builder can work with.

## Build a test app

Let's start with a Golang HTTP middleware, this is a cross between a function and a microservice and shows off how versatile OpenFaaS can be.

```sh
faas-cli template store pull golang-middleware

faas-cli new --lang golang-middleware \
build-test --prefix=alexellis2

```

- `--lang` specifies the build template
- `build-test` is the name of the function
- `--prefix` is the Docker Hub username to use for pushing up our OCI image

We'll get the following created:

```
./
├── build-test
│   └── handler.go
└── build-test.yml

1 directory, 2 files

```

The handler looks like this, and is easy to modify. Additional dependencies can be added through vendoring or [Go modules](https://blog.golang.org/using-go-modules).

```golang
package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world, input was: %s", string(input))))
}

```

### Build the normal way

The normal way to build this app would be:

```sh
faas-cli build -f build-test.yml

```

A local cache of the template and Dockerfile is also available at `./template/golang-middleware/Dockerfile`

There are three images that are pulled in for this template:

```
FROM openfaas/of-watchdog:0.7.3 as watchdog
FROM golang:1.13-alpine3.11 as build
FROM alpine:3.12

```

With the traditional builder, each of the images will be pulled in sequentially.

The wait a few moments and you're done, we now have that image in our local library.

We can also push it up to a registry with `faas-cli push -f build-test.yml`.

![seq](http://blog.alexellis.io/content/images/2020/01/seq.png)

### Build with Buildkit and Docker

This is the easiest change of all to make, and gives a fast build too.

```sh
DOCKER_BUILDKIT=1 faas-cli build -f build-test.yml

```

We'll see that with this approach, the Docker daemon automatically switches out its builder for buildkit.

Buildkit offers a number of advantages:

- More sophisticated caching
- Running later instructions first, when possible - i.e. downloading the "runtime" image, before the build in the "sdk" layer is even completed
- Super fast when building a second time

With buildkit, all of the base images can be pulled in to our local library at once, since the FROM (download) commands are not executed sequentially.

```
FROM openfaas/of-watchdog:0.7.3 as watchdog
FROM golang:1.13-alpine3.11 as build
FROM alpine:3.11

```

This option works even on a Mac, since buildkit is proxied via the Docker daemon running in the VM.

![dkit](http://blog.alexellis.io/content/images/2020/01/dkit.png)

### Build with Buildkit standalone

To build with Buildkit in a stand-alone setup we need to run buildkit separately on a Linux host, so we can't use a Mac.

`faas-cli build` would normally execute or fork `docker`, because the command is just a wrapper. So to bypass this behaviour we should write out a build context, that's possible via the following command:

```sh
faas-cli build -f build-test.yml --shrinkwrap

[0] > Building build-test.
Clearing temporary build folder: ./build/build-test/
Preparing ./build-test/ ./build/build-test//function
Building: alexellis2/build-test:latest with golang-middleware template. Please wait..
build-test shrink-wrapped to ./build/build-test/
[0] < Building build-test done in 0.00s.
[0] Worker done.

Total build time: 0.00

```

Our context is now available in the `./build/build-test/` folder with our function code and the template with its entrypoint and Dockerfile.

```sh
./build/build-test/
├── Dockerfile
├── function
│   └── handler.go
├── go.mod
├── main.go
└── template.yml

1 directory, 5 files

```

Now we need to run buildkit, we can build from source, or grab upstream binaries.

```sh
curl -sSLf https://github.com/moby/buildkit/releases/download/v0.6.3/buildkit-v0.6.3.linux-amd64.tar.gz | sudo tar -xz -C /usr/local/bin/ --strip-components=1

```

If you checkout the releases page, you'll also find buildkit available for armhf and arm64, which is great for multi-arch.

Run the buildkit daemon in a new window:

```sh
sudo buildkitd
WARN[0000] using host network as the default
INFO[0000] found worker "l1ltft74h0ek1718gitwghjxy", labels=map[org.mobyproject.buildkit.worker.executor:oci org.mobyproject.buildkit.worker.hostname:nuc org.mobyproject.buildkit.worker.snapshotter:overlayfs], platforms=[linux/amd64 linux/386]
WARN[0000] skipping containerd worker, as "/run/containerd/containerd.sock" does not exist
INFO[0000] found 1 workers, default="l1ltft74h0ek1718gitwghjxy"
WARN[0000] currently, only the default worker can be used.
INFO[0000] running server on /run/buildkit/buildkitd.sock

```

Now let's start a build, passing in the shrink-wrapped location as the build-context. The command we want is `buildctl`, buildctl is a client for the daemon and will configure how to build the image and what to do when it's done, such as exporting a tar, ignoring the build or pushing it to a registry.

```sh
buildctl build --help
NAME:
buildctl build - build

USAGE:

To build and push an image using Dockerfile:
    $ buildctl build --frontend dockerfile.v0 --opt target=foo --opt build-arg:foo=bar --local context=. --local dockerfile=. --output type=image,name=docker.io/username/image,push=true


OPTIONS:
   --output value, -o value  Define exports for build result, e.g. --output type=image,name=docker.io/username/image,push=true
   --progress value          Set type of progress (auto, plain, tty). Use plain to show container output (default: "auto")
   --trace value             Path to trace file. Defaults to no tracing.
   --local value             Allow build access to the local directory
   --frontend value          Define frontend used for build
   --opt value               Define custom options for frontend, e.g. --opt target=foo --opt build-arg:foo=bar
   --no-cache                Disable cache for all the vertices
   --export-cache value      Export build cache, e.g. --export-cache type=registry,ref=example.com/foo/bar, or --export-cache type=local,dest=path/to/dir
   --import-cache value      Import build cache, e.g. --import-cache type=registry,ref=example.com/foo/bar, or --import-cache type=local,src=path/to/dir
   --secret value            Secret value exposed to the build. Format id=secretname,src=filepath
   --allow value             Allow extra privileged entitlement, e.g. network.host, security.insecure
   --ssh value               Allow forwarding SSH agent to the builder. Format default|<id>[=<socket>|<key>[,<key>]]

```

Here's what I ran to get the equivalent of the Docker command with the `DOCKER_BUILDKIT` override:

```
sudo -E buildctl build --frontend dockerfile.v0 \
 --local context=./build/build-test/ \
 --local dockerfile=./build/build-test/ \
 --output type=image,name=docker.io/alexellis2/build-test:latest,push=true

```

Before running this command, you'll need to run `docker login`, or to create $HOME/.docker/config.json\` with a valid set of unencrypted credentials.

You'll see a nice ASCII animation for this build.

![buildkit-stand-alone](http://blog.alexellis.io/content/images/2020/01/buildkit-stand-alone.png)

### Build with `img` and buildkit

Since I've never used `img` and haven't really heard of it being used a lot with teams vs the more common options I thought I'd give it a shot.

First impressions are that multi-arch is not a priority and given the age of the project, may be unlikely to land. There is no binary for armhf or ARM64.

For `x86_64` the latest version is `v0.5.7` from 7 May 2019, built with Go 1.11, with Go 1.13 being the current release:

```sh
sudo curl -fSL "https://github.com/genuinetools/img/releases/download/v0.5.7/img-linux-amd64" -o "/usr/local/bin/img" \
	&& sudo chmod a+x "/usr/local/bin/img"

```

The build options look like a subset of `buildctl`:

```sh
img build --help
Usage: img build [OPTIONS] PATH

Build an image from a Dockerfile.

Flags:

  -b, --backend  backend for snapshots ([auto native overlayfs]) (default: auto)
  --build-arg    Set build-time variables (default: [])
  -d, --debug    enable debug logging (default: false)
  -f, --file     Name of the Dockerfile (Default is 'PATH/Dockerfile') (default: <none>)
  --label        Set metadata for an image (default: [])
  --no-cache     Do not use cache when building the image (default: false)
  --no-console   Use non-console progress UI (default: false)
  --platform     Set platforms for which the image should be built (default: [])
  -s, --state    directory to hold the global state (default: /home/alex/.local/share/img)
  -t, --tag      Name and optionally a tag in the 'name:tag' format (default: [])
  --target       Set the target build stage to build (default: <none>)

```

Here's what we need to do a build:

```sh
sudo img build -f ./build/build-test/Dockerfile -t alexellis2/build-test:latest ./build/build-test/

```

Now for one reason or another, `img` actually failed to do a successful build. It may be due to some of the optimizations to attempt to run as non-root.

![fail-build](http://blog.alexellis.io/content/images/2020/01/fail-build.png)

```sh
fatal error: unexpected signal during runtime execution
[signal SIGSEGV: segmentation violation code=0x1 addr=0xe5 pc=0x7f84d067c420]

runtime stack:
runtime.throw(0xfa127f, 0x2a)
	/home/travis/.gimme/versions/go1.11.10.linux.amd64/src/runtime/panic.go:608 +0x72
runtime.sigpanic()
	/home/travis/.gimme/versions/go1.11.10.linux.amd64/src/runtime/signal_unix.go:374 +0x2f2

goroutine 529 [syscall]:
runtime.cgocall(0xc9d980, 0xc00072d7d8, 0x29)
	/home/travis/.gimme/versions/go1.11.10.linux.amd64/src/runtime/cgocall.go:128 +0x5e fp=0xc00072d7a0 sp=0xc00072d768 pc=0x4039ee
os/user._Cfunc_mygetgrgid_r(0x2a, 0xc000232260, 0x7f84a40008c0, 0x400, 0xc0004ba198, 0xc000000000)

```

There seemed to be [three similar issues](https://github.com/genuinetools/img/issues/272) open.

### Build with Kaniko

Kaniko is Google's container builder which aims to sandbox container builds. You can use it as a one-shot container, or as a stand-alone binary.

I took a look at the build [in this blog post](https://blog.alexellis.io/quick-look-at-google-kaniko/)

```sh
docker run -v $PWD/build/build-test:/workspace \
 -v ~/.docker/config.json:/kaniko/config.json \
 --env DOCKER_CONFIG=/kaniko \
gcr.io/kaniko-project/executor:latest \
 -d alexellis2/build-test:latest

```

- The flag`-d` specifies where the image should be pushed after a successful build.
- The`-v` flag is bind-mounting the current directory into the Kaniko container, it also adds your `config.json` file for pushing to a remote registry.

![kaniko](http://blog.alexellis.io/content/images/2020/01/kaniko.png)

There is some support for caching in Kaniko, but it needs manual management and preservation since Kaniko runs in a one-shot mode, rather than daemonized like Buildkit.

### Summing up the options

- Docker - traditional builder

Installing Docker can be heavy-weight and add more than expected to your system. The builder is the oldest and slowest, but gets the job done. Watch out for the networking bridge installed by Docker, it can conflict with other private networks using the same private IP range.

- Docker - with buildkit

This is the fastest option with the least amount of churn or change. It's simply enabled by prefixing the command `DOCKER_BUILDKIT=1`

- Stand-alone buildkit

This option is great for in-cluster builds, or a system that doesn't need Docker such as a CI box or runner. It does need a Linux host and there's no good experience for using it on MacOS, perhaps by running an additional VM or host and accessing over TCP?


I also wanted to include a presentation by [Akihiro Suda]( [https://twitter.com/@](https://twitter.com/@) _AkihiroSuda_ /), a buildkit maintainer from NTT, Japan. This information is around 2 years old but provides another high-level overview from the landscape in 2018 [Comparing Next-Generation\
\
Container Image Building Tools](https://events19.linuxfoundation.org/wp-content/uploads/2017/11/Comparing-Next-Generation-Container-Image-Building-Tools-OSS-Akihiro-Suda.pdf)

This is the best option for [faasd users](https://github.com/alexellis/faasd), where users rely only on containerd and CNI, rather than Docker or Kubernetes.

- Kaniko

The way we used Kaniko still required Docker to be installed, but provided another option.


## Wrapping up

You can either use your normal container builder with OpenFaaS, or `faas-cli build --shrinkwrap` and pass the build-context along to your preferred tooling.

Here's examples for the following tools for building OpenFaaS containers:

- [Google Cloud Build](https://www.openfaas.com/blog/openfaas-cloudrun/)
- [GitHub Actions](https://lucasroesler.com/2019/09/action-packed-functions/)
- [Jenkins](https://docs.openfaas.com/reference/cicd/jenkins/) and
- [GitLab CI](https://docs.openfaas.com/reference/cicd/gitlab/).

In [OpenFaaS Cloud](https://docs.openfaas.com/openfaas-cloud/architecture/). we provide a complete hands-off CI/CD experience using the shrinkwrap approach outlined in this post and the buildkit daemon. For all other users I would recommend using Docker, or Docker with buildkit.

You can [build your own self-hosted OpenFaaS Cloud](https://www.openfaas.com/blog/ofc-private-cloud/) environment with GitHub or GitLab integration.

For [faasd users](https://github.com/openfaas/faasd), you only have containerd installed on your host instead of `docker`, so the best option for you is to download buildkit.

If you are interested in what Serverless Functions are and what they can do for you, why not checkout my new eBook and video workshop on Gumroad?

- [Checkout Serverless For Everyone](https://gumroad.com/l/serverless-for-everyone-else)

We did miss out one of the important parts of the workflow in this post, the deployment. Any OCI container can be deployed to the OpenFaaS control-plane on top of Kubernetes as long as its [conforms to the serverless workload definition](https://docs.openfaas.com/reference/workloads/). If you'd like to see the full experience of build, push and deploy, check out the [OpenFaaS workshop](https://github.com/openfaas/workshop/).

## Wrapping up

### Get help with Cloud Native, Docker, Go, CI & CD, or Kubernetes

Could you use some help with a difficult problem, an external view on a new idea or project? Perhaps you would like to build a technology proof of concept before investing more? Get in touch via [alex@openfaas.com](mailto:alex@openfaas.com) or book a session with me on [calendly.com/alexellis](https://calendly.com/alexellis/).

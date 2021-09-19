# Why we’re writing machine learning infrastructure in Go, not Python

## Production machine learning is about more than just algorithms

[Jan 8, 2020·5 min read](http://medium.com/@calebkaiser/why-were-writing-machine-learning-infrastructure-in-go-not-python-38d6a37e2d76?source=post_page-----38d6a37e2d76--------------------------------)

At this point, it should be a surprise to no one that Python is the [most popular](https://github.blog/2019-01-24-the-state-of-the-octoverse-machine-learning/) language for machine learning projects. While languages like R, C++, and Julia have their proponents—and use cases—Python remains the most universally embraced language, being used in every major machine learning framework.

So, naturally, our codebase at [Cortex](https://github.com/cortexlabs/cortex)—an open source platform for deploying machine learning models as APIs—is 87.5% Go.

![](https://miro.medium.com/max/60/1*jfWsTOsdPWlxjS4GuQkJLg.png?q=20)

Source: [Cortex GitHub](https://github.com/cortexlabs/cortex)

Machine learning algorithms, where Python shines, are just one component of a production machine learning system. To actually run a production machine learning API at scale, you need infrastructure that implements features like:

- Autoscaling, so that traffic fluctuations don’t break your API
- API management, to handle simultaneous API deployments
- Rolling updates, so that you can update models while still serving users

[Cortex](http://cortex.dev) is built to automate all of this infrastructure, along with other concerns like logging and cost optimizations.

Go is ideal for building software with these considerations, for a few reasons:

# 1\. Concurrency is crucial for machine learning infrastructure

A user can have many different models deployed as distinct APIs, all managed in the same Cortex cluster. In order for the Cortex Operator to manage these different deployments, it needs to wrangle a few different APIs. To name a couple:

- Kubernetes APIs, which Cortex calls to deploy models on the cluster.
- Various AWS APIs—EC2 Auto Scaling, S3, CloudWatch, and others—which Cortex calls to manage deployments on AWS.

The user doesn’t interact with any of these APIs directly. Instead, Cortex programmatically calls these APIs to provision clusters, launch deployments, and monitor APIs.

Making all of these overlapping API calls in a performative, reliable way is a challenge. Handling them concurrently is the most efficient way to do things, but it also introduces complexity, as now we have to worry about things like race conditions.

Go has an elegant, out of the box solution for this problem: **Goroutines**.

Goroutines are otherwise normal functions that Go executes concurrently. We could write an entire article digging into how Goroutines work under the hood, but at a high level, Goroutines are lightweight threads managed automatically by the Go runtime. Many Goroutines can fit on a single OS thread, and if a Goroutine blocks an OS thread, the Go runtime automatically moves the rest of the Goroutines over to a new OS thread.

Goroutines also offer a feature called “channels,” which allow Goroutines to pass messages between themselves, allowing us to schedule requests and prevent race conditions.

Implementing all of this functionality in Python may be doable with recent tools like asyncio, but the fact that Go is designed with this use case in mind makes our lives much easier.

# 2\. Building a cross-platform CLI is easier in Go

The Cortex CLI is a cross-platform tool that allows users to deploy models and manage APIs directly from the command line. The below GIF shows the CLI in action:

![](https://miro.medium.com/max/60/0*T17nRCxfd6j9gbB5?q=20)

Source: [Cortex GitHub](https://github.com/cortexlabs/cortex)

Originally, we wrote the CLI in Python, but trying to distribute it across platforms proved to be too difficult. Because Go compiles down to a single binary—no dependency management required—it offered us a simple solution to distributing our CLI across platforms without requiring much extra engineering effort.

The performance benefits of a compiled Go binary versus an interpreted language are also significant. According to the computer benchmarks game, Go is dramatically [faster than Python](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-python3.html).

It’s perhaps not coincidental that many other infrastructure CLI tools are written in Go, which brings us to our next point.

# 3\. The Go ecosystem is great for infrastructure projects

One of the benefits of open source is that you can learn from the projects you admire. For example, Cortex exists within the Kubernetes (which itself is written in Go) ecosystem. We were fortunate to have a number of great open source projects within that ecosystem to learn from, including:

- [**kubectl**](https://github.com/kubernetes/kubectl) **:** Kubernetes’ CLI
- [**minikube**](https://github.com/kubernetes/minikube): A tool for running Kubernetes locally
- [**helm**](https://github.com/helm/helm): A Kubernetes package manager
- [**kops**](https://github.com/kubernetes/kops) **:** A tool for managing production Kubernetes
- [**eksctl**](https://github.com/weaveworks/eksctl): The official CLI for Amazon EKS

All of the above are written in Go—and it’s not just Kubernetes projects. Whether you’re looking at [CockroachDB](https://github.com/cockroachdb/cockroach) or [Hashicorp’s](https://github.com/hashicorp) infrastructure projects, including [Vault](https://github.com/hashicorp/vault), [Nomad](https://github.com/hashicorp/nomad), [Terraform](https://github.com/hashicorp/terraform), [Consul](https://github.com/hashicorp/consul), and [Packer](https://github.com/hashicorp/packer), all of them are written in Go.

The popularity of Go in the infrastructure world has another effect, which is that most engineers interested in working on infrastructure are familiar with Go. This makes it easier to attract engineers. In fact, one of the best engineers at Cortex Labs found us by searching for Go jobs on AngelList—and we are lucky he found us.

# 4\. Go is just a pleasure to work with

The final note I’ll make on why we ultimately built Cortex in Go is that Go is just _nice_.

Relative to Python, Go is a bit more painful to get started with. Go’s unforgiving nature, however, is what makes it such a joy for large projects. We still heavily test our software, but static typing and compilation — two things that make Go a bit less comfortable for beginners — act as sort of guard rails for us, helping us to write (relatively) bug-free code.

There may be other languages you could argue offer a particular advantage, but on balance, Go best satisfies our technical and aesthetic needs.

# Python for machine learning, Go for infrastructure

We still love Python, and it has its place within Cortex, specifically around inference processing.

Cortex serves TensorFlow, PyTorch, scikit-learn, and other Python models, which means that interfacing with the models—as well as pre and post inference processing—are done in Python. However, even that Python code is packaged up into Docker containers, which are orchestrated by code that is written in Go.

If you’re interested in becoming a machine learning engineer, knowing Python is more or less non-negotiable. If you’re interested in working on machine learning _infrastructure_, however, you should seriously consider using Go.

_Are you an engineer interested in Go and machine learning? If so, consider_ [_contributing to Cortex_](https://github.com/cortexlabs/cortex) _!_

[**Caleb Kaiser**](http://medium.com/@calebkaiser?source=post_sidebar--------------------------post_sidebar-----------)

ML infrastructure ( [https://github.com/cortexlabs/cortex](https://github.com/cortexlabs/cortex)) Formerly at AngelList. Originally a Cadillac.

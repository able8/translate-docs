# The Next Generation of Kubernetes Native Postgres

July 07, 2021 [Jonathan S. Katz](https://blog.crunchydata.com/blog/author/jonathan-s-katz)

We're excited to announce the release of [PGO](https://github.com/CrunchyData/postgres-operator) 5.0, the open source [Postgres Operator](https://github.com/CrunchyData/postgres-operator) from [Crunchy Data](https://www.crunchydata.com/). While I'm very excited for you to [try out PGO 5.0](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/) and provide feedback, I also want to provide some background on this release.

When I joined Crunchy Data back in 2018, I had heard of [Kubernetes](https://kubernetes.io/) through my various open source activities, but I did not know much about it. I learned that we had been running Postgres on [Kubernetes](https://blog.crunchydata.com/blog/creating-a-postgresql-cluster-using-helm-for-kubernetes) and [OpenShift](http://blog.crunchydata.com/blog/advanced-crunchy-containers-for-postgresql) in production environments for years. This included the release of [one of the first Kubernetes Operators](http://blog.crunchydata.com/blog/postgres-operator-for-kubernetes)! It was quite remarkable to see how Crunchy lead the way in cloud native Postgres. I still remember how excited I was when I [got my first Postgres cluster up and running on Kubernetes](http://blog.crunchydata.com/blog/get-started-runnning-postgresql-on-kubernetes)!

Many things have changed in the cloud native world over the past three years. When I first started giving talks on the topic, I was answering questions like, "Can I run a database in a container?" coupled with "Should I run a database in a container?" The conversation has now shifted. My colleague Paul Laurence wrote an [excellent article](http://blog.crunchydata.com/blog/using-kubernetes-chances-are-you-need-a-database) capturing the current discourse. The question is no longer, "Should I run a database in Kubernetes" but " [Which database should I run in Kubernetes](http://blog.crunchydata.com/blog/using-kubernetes-chances-are-you-need-a-database)?" (Postgres!).

Along with this shift in discussion is a shift in expectation for how databases should work on Kubernetes. To do that, we need to understand the difference between an imperative workflow and a declarative workflow.

## Cloud Native Declarative Postgres

[PGO](https://github.com/CrunchyData/postgres-operator), the open source [Postgres Operator from Crunchy Data](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/), was initially designed for running cloud native Postgres using an imperative workflow. Many operations required the use of a command-line utility called "pgo". For convenience, "pgo" follows  the conventions of the Kubernetes " [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/)" command line tool. "pgo" includes many conveniences for managing Postgres on Kubernetes, including:

- Creating databases:pgo create cluster hippo
- Maintenance operations:pgo update cluster hippo --memory=2Gi
- Backups:pgo backup hippo
- Cloning clusters:pgo create cluster rhino --restore-from=hippo

and more.

A declarative workflow is one where you describe what you want, and your application "makes it happen." An example of this is SQL: you write a query, e.g.:

```
<span>SELECT</span> <span>*</span> <span>FROM</span> animals <span>WHERE</span> animal_type <span>=</span> <span>'hippo'</span>;
```

You are asking your database "Find me all the animals that are hippos". You don't care how the database actually does this: it could search over an index or perform a sequential scan. You know that for your end result, you want all the hippos.

This is a very powerful concept: instead of building towards your end result, you describe what you want it to be. A well-designed declarative engine can optimize your experience using the software.

This was the founding principle on which we built the new version of our Kubernetes Operator for Postgres.

## Designing the Next Generation Postgres Operator

We had several goals while designing the next generation of [Postgres Operator](https://github.com/CrunchyData/postgres-operator). We wanted to make it both [easy to get started running Postgres on Kubernetes with PGO](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/) and easy to manage "Day 2" type operations like cluster resizing. We also wanted to ensure our Postgres databases could withstand the rigors of running in a Kubernetes environment.

While building the new PGO, we focused on creating a seamless, declarative user experience that included the production-ready Postgres features introduced in prior releases. We also leveraged modern Kubernetes features like [server-side apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/). We architected PGO so that it works behind the scenes maintaining the Postgres environment that you requested.

For example, let's say you accidentally delete your connection pooler Deployment. Before this release, [pgMonitor](https://github.com/CrunchyData/pgmonitor) would send you alerts indicating your database is disconnected. After investigating, you realize that you need to manually recreate your pgBouncer deployment. With the updated declarative approach, PGO now automatically heals your environment. When PGO detects the missing Deployment, it will instantaneously recreate your connection pooler and make it seem like nothing happened!

Think of it as the Postgres Operator automatically managing your operations for you.

## Deploying Postgres the GitOps Way

There is an ongoing push to use [infrastructure as code](https://en.wikipedia.org/wiki/Infrastructure_as_code) ( [IAC](https://en.wikipedia.org/wiki/Infrastructure_as_code)) methodologies such as [GitOps](https://www.gitops.tech/) with Kubernetes. Because of this, we wanted PGO to work well with tools such as [Kustomize](https://kustomize.io/), [Helm](https://helm.sh/), and [OLM](https://olm.operatorframework.io/) throughout the lifetime of a Postgres cluster.

We took Kelsey Hightower's " [Stop scripting and start shipping](https://twitter.com/kelseyhightower/status/953638870888849408)" to heart. We made a simple, one-step operation to create your Postgres database and connect your application in a safe and secure way. We also allowed for everything to be modifiable through a declarative workflow. We fine-tuned the experience by running lots of tests through [ArgoCD](https://argoproj.github.io/argo-cd/). We also created a series of [examples for deploying Postgres clusters in various Kubernetes scenarios](https://github.com/CrunchyData/postgres-operator-examples).

Part of managing your infrastructure using GitOps involves deploying applications with minimal disruptions. In Kubernetes, there are certain modifications that can cause downtime. For example, updating to a newer Postgres bugfix release causes Kubernetes to create new Pods. We designed PGO to make sure that when these necessary Day 2 actions occur, they are performed with minimal to zero downtime.

## The Next Generation Cloud Native Postgres is Now

I've been around open source for a long time, and through the years I've seen how open source innovation begets innovation. A decade ago, Postgres [added support for the JSON data type](http://blog.crunchydata.com/blog/better-json-in-postgres-with-postgresql-14). Adding JSON spurred many advances throughout the application development and relational database landscape!

I'm seeing something similar now occurring with Kubernetes. Our team built a powerful, declarative Postgres Operator thanks to advances in the core Kubernetes ecosystem. We can now run production cloud native Postgres by maintaining a few YAML files.

Even with all these advances in cloud native technology, what excites me is that we're only starting to scratch the surface of what we can do. The good news is that the next generation of cloud native Postgres is here, and it's [open source](https://github.com/CrunchyData/postgres-operator).

[Crunchy Postgres for Kubernetes](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/) 5.0 is built on the redesigned [PGO](https://github.com/CrunchyData/postgres-operator), the open source [Postgres Operator](https://github.com/CrunchyData/postgres-operator). PGO 5.0 combines many years experience running Postgres on Kubernetes with a robust operations feature set to deliver a modern cloud native Postgres distribution. You can check out our [full press release](http://blog.crunchydata.com/news/next-generation-crunchy-postgres-for-kubernetes-released) that also includes info on our PGO 5.0 webinar.

We cordially invite you to [try out PGO](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/), take it for a test drive, and let us know what you think.

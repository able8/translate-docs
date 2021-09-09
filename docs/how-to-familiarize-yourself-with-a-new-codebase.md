# How to familiarize yourself with a new codebase

A few weeks ago, a tweet made me take a second and think about something that I'd never consciously considered before; how can you approach an unfamiliar codebase and start to understand it?

[https://twitter.com/d\_feldman/status/1336407539928477697?s=21](https://twitter.com/d_feldman/status/1336407539928477697?s=21)

It got me thinking about how I would approach a new repo that I'd never seen before but needed to make a contribution against, like a bug fix. I remembered my early days of learning [Kubernetes](https://kubernetes.io/), and wanting to make requests to the its API (because using the command line wasn't good enough for me, apparently). I had been trying to work out how to automatically deploy a particular branch of a GitLab repo into a cluster every time someone pushed to it. I had big ideas about automating DNS, setting up automated certificates, and adding a Slackbot to notify you whenever a new deploy happened.

If I remember correctly, I got a proof of concept working, and then it never went much past that. Given how popular [GitOps](https://www.cloudbees.com/gitops/what-is-gitops) has become, maybe I should have stuck with it! When I started delving into the Kuberenetes side of the project, I was completely and utterly lost. The documentation didn't have much in the way of _how_ to use the API (I'm sure nowadays things are much better), and reading the Kubernetes source code was a complete non-starter because well, that thing is a monster. I remember thinking to myself that I just needed to replicate what kubectl was doing to create a new Deployment.

So I gave up trying to read Kubernetes' source, and moved over to the [source for kubectl](https://github.com/kubernetes/kubectl). This is where I started to make some headway! I was able to follow straight from the `main()` function to the `apply` command, down through the logic until it started making API requests. It felt so good to finally get an answer, and to just import some Go packages to make it all work in short order!\*

This is the background behind my answer to the tweet above:

[https://twitter.com/cohix/status/1336408360770531331?s=21](https://twitter.com/cohix/status/1336408360770531331?s=21)

Since that project years ago, I've sort of instinctively followed this strategy whenever I need to reason about a new codebase because well, it works! Only recently did this tweet make me think about it concretely, and I'm glad it did. I tried to replicate this purposefully to test my strategy. I went to a [large open-source repo](https://github.com/fluxcd/flux2) and tried to find the code where it installed itself into a cluster. Using this strategy, I started with the tool's `main()` and then was able to find my way to the `install` command, which led me down to where the installation happens (funnily enough, by calling `kubectl`).

I think it's important for any developer to understand not only how to reason about an unfamiliar codebase, but also to realize that an important way that we learn is by trial and error. When we try something and it works, it brings us joy and we'll continue to do it, even if we don't realize it. I think it's a good idea to take a second to think about these moments when they happen, take a mental note of it so that next time you come across a similar problem, you can consciously use your previous learning and expand upon the strategies you've developed over time.

The reason I wanted to turn this tweet into a full blog post is because it made me realize that one of the goals for [Atmo](https://github.com/suborbital/atmo) is to make it easier for developers to reason about applications. Since Atmo uses a declarative format for building backend applications (using WebAssembly modules), there is always one canonical entrypoint; the Directive. From there, you can easily reason about what the application is doing because it is [declarative instead of imperative](https://stackoverflow.com/a/1784702). This is one of the things that made Kubernetes so popular. Being able to describe your application in a simple format, and then have a system "make it happen" is a magical thing, and Atmo strives to do exactly that.

Atmo is gaining new functionality every week. If you want to learn more, check out [the Suborbital homepage](https://suborbital.dev)

- When I say "short order", I'm sure it still took several days to get everything working, but once you unblock yourself on a big problem, everything after that just seems to fly by.

Cover Photo by [Rafif Prawira](https://unsplash.com/@rafifatmaka) on [Unsplash](https://unsplash.com/s/photos/maze)

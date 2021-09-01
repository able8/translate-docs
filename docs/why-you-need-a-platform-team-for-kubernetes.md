# Why you need a platform team for Kubernetes

Setting up a Kubernetes cluster can be deceptively simple, as there are plenty of installers to create a basic cluster in minutes. However, that’s only the start of the actual work. Kubernetes moves fast; when it’s a critical part of your infrastructure, there’s a host of things you need to look out for to maintain a healthy cluster. More often than not, it’s wise to have a dedicated team to run Kubernetes.

## Integrations

For most use cases, Kubernetes is not enough on its own. It might be better to think of it as a framework for building platforms. You will need to integrate Kubernetes with other tools and systems to make it useful. The exact integrations depend on your use case — there’s no silver bullet. You’ll need to figure out your own optimal setup.

To give you an idea, here are a few common integrations you might need to set up and configure:

- ExternalDNS for interfacing with cloud DNS services, to have human-readable addresses for services

- Let’s Encrypt for automatic HTTPS

- A network plugin such as Calico or Flannel

- A CI/CD system such as Spinnaker, GitLab, or Jenkins to deploy to Kubernetes

- Authentication with LDAP, AD, or another system for using your company’s regular user accounts

- Prometheus for monitoring

- Alertmanager for alerting


Granted, you might not need all of these, and some of them may be provided by your Kubernetes distribution out of the box. Still, if you end up using them, you’ll need to understand how they work and be able to operate them. All of this integration work adds a lot of complexity, and you’ll need to spend a lot of time learning how it all works. It can be difficult for one person to grasp all the details of the entire system, so it’s better to share the effort with multiple people familiar with Kubernetes and its integrations.

## Keeping up with the times

Kubernetes is still a fast-moving project. There are four releases every year – old versions are maintained only for a year, after which there are no security updates. While running a managed service, such as GKE, EKS, or AKS, helps significantly with upgrading, it’s still not entirely effortless.

Keeping a Kubernetes cluster updated and secure is complex. Besides keeping Kubernetes itself up to date, there’s also work involved in upgrading all the integrations needed to create a fully functional platform. Running a managed Kubernetes service does not necessarily help with keeping them updated. Another continuous update cycle comes from the various container images involved in running the software that is the reason for running Kubernetes in the first place.

There are good reasons for keeping all the software involved up to date: Kubernetes clusters have a lot of potential for attacks due to their complex and distributed nature. On top of that, they make for attractive targets if looking for computing power. One example of attacks against Kubernetes is cryptojacking where attackers hijack clusters to mine cryptocurrency. [Tesla famously fell victim to such an attack a couple of years ago](https://redlock.io/blog/cryptojacking-tesla) and recently such attacks have become more sophisticated as evidenced by the [Hildegard malware](https://unit42.paloaltonetworks.com/hildegard-malware-teamtnt/).

## The human aspect

As with any IT system you operate, you’ll need someone who can keep it running at all times. As Kubernetes is a platform on top of which you run other services, this is even more important: outages have a broader impact and might impair your operations significantly.

As a rule of thumb, you’ll need at least three people who can independently solve Kubernetes-related issues — but you’ll likely need more than that. Consider these scenarios as a member of a Kubernetes maintenance team of three:

- One of your teammates leaves for another job. How long will it take to hire and/or train a replacement?

- Your team has a face-to-face meeting. It turns out one of your teammates was coming down with the flu and inadvertently spread it to your other colleague: now two out of your team of three have the flu and you’re alone.

- You have a teammate on vacation and another one sick. There’s an issue related to one of the integrations, but you don’t know it too well – your coworker who’s drinking mai tais on the beach is the expert on that one. Consider the lead time for contacting your teammate, getting them on a computer, and in the mindset to solve your problem.


It’s not unreasonable to assume that you’ll have someone available most of the time, but exceptions are always around the corner. There is always a significant risk of running into a situation where you don’t have enough people around to solve an incident promptly. You’ll need to adjust your availability target accordingly. If you need 24/7 operations, three people likely won’t cut it.

## No dedicated platform team?

By this point, I hope I’ve convinced you that operating and maintaining a Kubernetes cluster is not a one-person operation. Even if you try not having a dedicated team by spreading the knowledge across multiple development teams, you still need someone available to troubleshoot issues at all times. What would be the benefits of having a dedicated Kubernetes platform team instead of a distributed model?

One significant downside to distributing Kubernetes maintainers across multiple teams is the increased amount of context switching. This is a recipe for a constant tug of war between running Kubernetes and the main team’s duties (if they’re only tasked with Kubernetes, why not just put them in a dedicated platform team instead?).

Constant context switching gets stressful as the maintainers need to make tough decisions about what to prioritize. This is exacerbated by the fact that operations work is unevenly distributed and unpredictable: Often things are quiet and systems keep humming along without any issues, but an incident can suddenly and unexpectedly turn your week upside down.

All in all, it’s simply hard to keep up with Kubernetes while also juggling other work.

## Conclusion

If your organization is large enough and you can afford to have a dedicated team to maintain Kubernetes, it will save you a lot of time and effort compared to other options for managing computing resources. You need to spend time to save time: there are a certain initial investment and ongoing maintenance overhead to keep things running smoothly.

It’s worth noting that while having a platform team has its benefits, you should be careful not to build it into a silo. It’s best to think of the Kubernetes platform as an internal service with internal customers you need to keep happy. And for that, you need to collaborate closely with them to make sure you are building a good platform for their specific use cases.

If you have a smaller organization and can’t yet justify a dedicated team just for Kubernetes, it might mean sacrifices in the quality and reliability of the platform. Consider your options carefully: it might not be possible to get a production-ready platform out of Kubernetes (depending on what “production-ready” means for you).

* * *

Risto Laurikainen is a DevOps Consultant with a decade of experience in building cloud computing platforms.

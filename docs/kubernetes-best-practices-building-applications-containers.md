# 7 best practices: Building applications for containers and Kubernetes     

Let’s examine key considerations for  building new applications specifically for containers and Kubernetes,  according to cloud-native experts

March 2, 2020

Don’t let the growing popularity of [containers](https://enterprisersproject.com/tags/containers) and [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes?intcmp=701f2000000tjyaaaa&extidcarryover=true&sc_cid=70160000000h0axaaq) dupe you into thinking that you should use them to run any and every type of application. You need to distinguish between “can” and “should.”

One basic example of this distinction is the difference between  building an app specifically to be run in containers and operated with  Kubernetes (some would refer to this as [cloud-native](https://www.redhat.com/en/topics/cloud-native-apps?intcmp=701f2000000tjyaAAA) development) and using these containers and orchestration for existing monolithic apps.

Building  new applications specifically for containers and Kubernetes might be the best starting point for teams just beginning their container work.

We’ll cover the latter scenario in an upcoming post. Today, we’re  focused on some of the key considerations for building new applications  specifically for containers and Kubernetes – in part because the  so-called “greenfield” approach might be the better starting point for  teams just beginning with containers and orchestration.

“Containers [and orchestration] are a technical vehicle for building, deploying, and running cloud-native applications,” says Rani Osnat, VP  of strategy at [Aqua Security](https://www.aquasec.com/). “I typically recommend to those starting their journey with containers  to use a new, simple greenfield application as their test case.”

**[ Want to learn about building and deploying Kubernetes Operators? Get the free eBook: [O’Reilly: Kubernetes Operators: Automating the Container Orchestration Platform.](https://www.redhat.com/en/resources/oreilly-kubernetes-operators-automation-ebook?intcmp=701f2000000tjyaAAA) ]**

## **How to develop apps for containers and Kubernetes**

We asked Osnat and other cloud-native experts to share their top tips for developing apps specifically to be run in containers using  Kubernetes. Let’s dive into six of their best recommendations.

### 1. Think and build modern

If you’re building a new home today, you’ve got different styles and  approaches than you would have, say, 50 years ago. The same is true with software: You’ve got new tools and approaches at your disposal.

“If you’re building an app, build it in a modern way!” says Miles Ward, CTO at [SADA](https://sada.com/). Ward points to [microservices](https://enterprisersproject.com/article/2017/8/how-explain-microservices-plain-english) and the [12-factor methodology](https://12factor.net/) as chief examples of modern application development.

Ward notes that while microservices and containers can work well  together, the pairing is not actually a necessity, at least under the  right conditions. “Microservices are also mentioned often in conjunction with Kubernetes; however, it is absolutely not a hard requirement,”  Ward says. “A monolithic approach can also work, provided it can scale  horizontally either as a single horizontal deployment or as multiple  deployments with different endpoints to the same codebase.”

The same is true of twelve-factor: “Twelve-factor is a useful starting point, but its tenets aren’t necessarily law,” Ward says.



If you’re building an app from scratch, give strong consideration to the microservices approach.

But if you’re building an application from scratch – as Osnat advises teams do when they are getting started with containers and  orchestration – give strong consideration to the microservices approach.

“To maximize the benefits of using containers, architect your app as a microservices app, in a way that will allow it to function even when  individual containers are being refreshed,” Osnat advises. “It should  also be structured so container images represent units that can be  released independently, allowing for an efficient  [CI/CD](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up) implementation.”

“Modern” development can be defined in various ways. If you’re  building an application for containers and Kubernetes, it essentially  means making choices that suit these packaging and deployment  technologies. Here are two more examples:

- **Define container images as logical units that can scale independently:** “For instance, it usually makes sense to implement databases, logging,  monitoring, load-balancing, and user session components as their own  containers or groups of containers,” Osnat says.
- **Consider cloud-native APIs:** “Kubernetes has powerful API extension mechanisms,” says Vaibhav Kamra, VP of engineering and co-founder at [Kasten](https://www.kasten.io/). “By integrating with those, you can immediately take advantage of  existing tools in the ecosystem like command-line utilities and  authentication.”

“Modern” is a good thing from a software development perspective, too.

“The great thing about most modern languages and frameworks is that  they are overwhelmingly container-friendly,” says Ravi Lachhman, DevOps  advocate at [Harness](https://harness.io/). “Going back even a few years, runtimes like Java had a difficult time  respecting container boundaries, and the dreaded out of memory killer  would run seemingly arbitrarily to the operator. Today, because of the  popularity of containers and orchestrators, especially Kubernetes, the  languages and frameworks have evolved to live in the new paradigm.”

**[ As Kubernetes adoption grows, what can IT leaders expect? Read also: [5 Kubernetes trends to watch in 2020](https://enterprisersproject.com/article/2020/1/kubernetes-trends-watch-2020?sc_cid=70160000000h0axaaq). ]**

### 2. CI/CD and automation are your friends

Automation is a critical characteristic of container orchestration;  it should be a critical characteristic of virtually all aspects of  building an application to be run in containers on Kubernetes.  Otherwise, the operational burden can be overwhelming.

“Build applications and services with automation as minimum table stakes,” recommends Chander Damodaran, chief architect at [Brillio](https://www.brillio.com/). “With the proliferation of services and components, this can become an unmanageable issue.”



A well-conceived  CI/CD pipeline can bake automation into many phases of your development and deployment processes.

A well-conceived  [CI/CD pipeline](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up) is an increasingly popular approach to baking automation into as many  phases of your development and deployment processes as possible. Check  out our recent primer for IT leaders: [How to build a CI/CD pipeline.](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up)

Another way to think about the value of automation: It will make your mistakes – which are almost inevitable, especially early on – easier to bounce back from.

“Using any new platform requires a lot of trial and error, and the  ease of using Kubernetes does not excuse you from taking necessary  precautions,” says Lachhman from Harness. “Having a robust continuous  delivery pipeline in place can ensure that confidence-building standards like testing, security, and change management strategies are followed  to ensure your applications are running effectively.”

**[ Why does Kubernetes matter to IT leaders? Learn more about Red Hat's [point of view](https://www.redhat.com/en/topics/containers/kubernetes-approach?intcmp=7013a000002DSiEAAW). ]**

### 3. Keep container images as light as possible

Another key principle when developing an application for containers  and Kubernetes: Keep your container images as small as possible for  performance, security, and other reasons.



Make sure to remove all other packages – including shell utilities – that are not required by the application. 

“Only include what you absolutely need. Often images contain packages which are not needed to run the contained application,” says Ken  Mugrage, principal technologist in the office of the CTO at [ThoughtWorks](https://www.thoughtworks.com/). Make sure to remove all other packages – including shell utilities –  that are not required by the application. This not only makes the images smaller but reduces the attack surface for security issues, he says.

This is a good example of how building a containerized application  might require a shift in traditional practices for some development  teams.

“Developers need to rethink how they develop applications. For  example, create a smaller container and base image,” says Nilesh Deo,  director of product marketing at [CloudBolt](https://www.cloudbolt.io/). “The smaller the image, the faster it can load, and as a result, the application will be faster.”

Let’s examine four other best practices:

  

### 4. Don’t blindly trust images

As is common in software development, there’s a chance you can reuse  or repurpose existing components rather than build them from scratch.  The same principle can apply to containers. Just don’t make the mistake  of having blind faith in container images, especially not from a  security perspective.



“Far too many people choose an image from a repository with some sort of application stack already installed.” 

“Far too many people choose an image from a repository with some sort of application stack already installed,” Mugrage says. “Often these  images are poorly built, and the risk of security issues can’t be  ignored. Any images you use, even ones in your own repositories, should  be scanned for vulnerabilities and compliance on every run of your  deployment pipeline.”

### 5. Plan for observability, telemetry, and monitoring from the start



Kubernetes' self-healing capabilities are a piece of the platform’s appeal, but they also  underscore the need for proper visibility.

Failure is essentially a part of the plan with containers and microservices, but it’s more [a matter of managing failure](https://cloud.ibm.com/docs/cloud-native?topic=cloud-native-observability-cn) rather than trying to avoid it altogether. Kubernetes’ self-healing  capabilities are a piece of the platform’s appeal, but they also  underscore the need for proper visibility into your applications and  environments. This is where observability, telemetry, and monitoring  become key.

“Kubernetes has built-in mechanisms for resiliency, which create a  need for comprehensive monitoring as a best practice,” says Andrei  Zbikowski, software engineer at [Sentry.io](https://sentry.io/). “Its self-healing functions can restart failed containers or replace  and terminate others when certain health parameters are not met. While  this will keep applications up and running initially, it can actually  conceal growing problems.”

Zbikowski adds that a lack of visibility into your code might mean an app is throwing errors hourly, for example, even though health metrics  show everything up and running normally.

“It is important to monitor applications, as well as containers and  back-end systems,” Zbikowski says. “A comprehensive approach to  monitoring will provide greater visibility into issues and events so  that problems can be identified and remediated before there is any  significant impact to users.”

Trying to bolt monitoring on to your containerized applications later down the line might lead to unsatisfactory results, Mugrage says.  “Think about observability and monitoring from the beginning.  Troubleshooting distributed applications is hard and has to be included  in the application design. Adding a monitoring solution later will leave you disappointed.” (Mugrage points to this [piece on observability](https://martinfowler.com/articles/domain-oriented-observability.html) as an additional resource.)

“There is a big toolbox of cloud-native technologies available to  build sophisticated monitoring, tracing, service meshes, and dashboards  into your applications,” says [Red Hat](https://www.redhat.com/en?intcmp=701f2000000tjyaAAA) technologist evangelist [Gordon Haff](https://enterprisersproject.com/user/gordon-haff). He adds, “Prometheus, Jaeger, Kiali, and Istio are just a few of the  projects you may hear about, and new ones are popping up all the time.  However, the choice can be overwhelming and integrating all the tools  yourself can be a challenging distraction.” (This is where you can  consider instead an integrated enterprise open source product like Red  Hat OpenShift Container Platform, Haff notes. )

**[ Read also: [OpenShift and Kubernetes: What’s the difference?](https://www.redhat.com/en/blog/openshift-and-kubernetes-whats-difference?intcmp=701f2000000tjyaAAA&extIdCarryOver=true&sc_cid=70160000000cYRWAA2) ]**

### 6. Consider starting with stateless applications

One early line of thinking about containers and Kubernetes has been  that running stateless apps is a lot easier than running stateful apps  (such as databases). That’s changing with the growth of [Kubernetes Operators](https://enterprisersproject.com/article/2019/2/kubernetes-operators-plain-english), but teams new to Kubernetes might still be better served by beginning with stateless applications.

“The benefits of running applications in containers and Kubernetes  are numerous, and there are steps that developers can take to make sure  they are prospering from all those benefits. If I had to pick just one,  the most important step is developing applications with a stateless,  rather than stateful, backend,” says Chris Parmer, co-founder at [Plotly](https://plot.ly/). “Through a stateless backend, development teams can ensure there are no long-running connections or mutable states that make it harder to  scale. Developers will also be able to deploy applications more easily  with zero downtime and enable end-user requests to be delivered in  parallel to different containers.”

**[ Also read [Kubernetes Operators: 4 facts to know](https://enterprisersproject.com/article/2020/2/kubernetes-operators-4-things-know). ]**

Parmer notes that scalability is one of the major draws of running  containers on Kubernetes; that benefit will be easier to realize with a  stateless app.

“Stateless applications make it easy to migrate and scale as  necessary to meet the business needs of the organization, allowing teams to add or remove containers at will,” Parmer says. “By using web  application frameworks that are built upon stateless backends, you get  the most out of your Kubernetes cluster.”

### 7. Remember, this is hard

MORE ON KUBERNETES

- [How to explain Kubernetes Operators in plain English](https://enterprisersproject.com/article/2019/2/kubernetes-operators-plain-english)
- [Kubernetes in production vs. Kubernetes in development: 4 myths](https://enterprisersproject.com/article/2018/11/kubernetes-production-4-myths-debunked)
- [Kubernetes: 6 secrets of successful teams ](https://enterprisersproject.com/article/2020/2/kubernetes-6-secrets-success)

“None of the abstractions that exist in Kubernetes today make the  underlying systems any easier to understand. They only make them easier  to use,” says Chris Short, Red Hat OpenShift principal technical  marketing manager. “If this were easy, everyone would be doing it  already. The industry would be moving on from the Kubernetes hype to the next big thing. This stuff is hard. We’re doing container orchestration while abstracting away the need to manage much other than the state of  the cluster and the infrastructure underneath it. [Etcd](https://etcd.io/) is a huge Kubernetes dependency that a lot of people have had nicely  tucked away from them. There is networking, security, and everything  else wrapped up in Kubernetes too. If your teams aren’t expecting  failure and ready to learn from mistakes, then I need to figure out how  you built the perfect Kubernetes environment.”

**[ Kubernetes terminology, demystified: Get our [Kubernetes glossary](https://enterprisersproject.com/kubernetes-glossary?sc_cid=70160000000h0axaaq) cheat sheet for IT and business leaders. ]**
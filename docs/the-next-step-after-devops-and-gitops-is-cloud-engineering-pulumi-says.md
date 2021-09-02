# The Next Step after DevOps and GitOps Is Cloud Engineering, Pulumi Says

#### 3 May 2021 7:55am,   by [Mary Branscombe](https://thenewstack.io/author/marybranscombe/ "Posts by Mary Branscombe")


If we going to treat infrastructure as code, shouldn’t infrastructure engineers have access to the same tools that make software engineers productive and even the same languages? That’s the theory behind [Pulumi](https://thenewstack.io/pulumi-uses-real-programming-languages-to-enforce-cloud-best-practices/), which has just released version 3 of its open source platform.

What founder and CEO [Joe Duffy](https://www.linkedin.com/in/joejduffy/), calls a “cloud engineering platform” is an attempt to “distill a lot of the lessons learned from helping developers build modern cloud applications, helping infrastructure teams increasingly apply engineering disciplines to the way they’re doing infrastructure and help the entire team really ship faster with confidence.”

“People are realizing the only way to keep up with the pace of the modern cloud and the level of innovation is to empower developers to be more self serve but also learn from the past decades in software engineering and apply that to the way we’re doing infrastructure.”

“Not everybody can spend two years going on the journey of figuring out how to do cloud engineering: we want to make sure that everybody has access to this on day one.”

## Cloud Engineering Culture

“Cloud engineering uses standard software engineering and tools across infrastructure, app dev and compliance to simplify and homogenize the complexity of modern cloud environments,” IDC DevOps research director [Jim Mercer](https://www.idc.com/getdoc.jsp?containerId=PRF005085) explained to the New Stack.

To Justin Fitzhugh, whose job title at Pulumi customer Snowflake is vice president of cloud engineering, the term represents the continued evolution from the separate cultures of development and operations.

“With DevOps, we saw developers and operations teams working together on a product or a specific deliverable. We saw applications and systems being developed jointly where operational concerns were being taken into account in the development phase and inversely the developers were concerned about how they scale and how do you manage as you go into production.”

Cloud engineering takes that to the next level. “It’s more of a proactive engineering culture where we’re building tooling, software and components to build, manage, instantiate and take care of the infrastructure. Any change to our infrastructure is a commit to a codebase which runs through a CI/CD pipeline where you can diff and you can test and has appropriate reviews just as any other software change would and then is pushed to production through automated pipelines. We look at it as a software engineering function, just focused on infrastructure as opposed to UI or back end engineering.”

Infrastructure still requires a specific discipline, Fitzhugh noted, because their concerns are integration, scale, abstraction and how to interact with and operate cloud services, but much of that can be automated.

![](https://cdn.thenewstack.io/media/2021/04/40f21f70-pulumi-01-1024x791.png)

Some DevOps teams keep their infrastructure configuration in git repos, but too often operational changes are made by running through the steps on a to-do list, even if that’s kept in a wiki or a configuration management system. “As opposed to how do I codify that, how do I commit that to a codebase, how do I make that part of a reproducible test run and have it be part of a pipeline.”

Snowflake uses Go extensively, along with everything from Ansible to Python, and it runs on AKS, EKS and GKE so needs to deal with different APIs, account metrics and cloud primitives. “We use Pulumi to wrap that into a single workflow orchestration.”

Duffy called Pulumi “connective tissue” and noted that, increasingly, security engineering teams are part of that same workflow because Pulumi users are trying to get away from silos where they have to deal with YAML and multiple domain-specific languages, with different delivery mechanisms for infrastructure — many of them manual — even if they’re doing automated delivery of containers for some of their workloads.

And increasingly, cloud native requires orchestration across multiple services; doing that with bash scripts can be painful, especially when you think about security.

“We find the line starts to become blurry with these modern architectures. Is a serverless function infrastructure or is that an application? Is a queue or a pub sub-topic part of the application or is that infrastructure? For a lot of folks leaning into the modern cloud, it’s a little bit of both. That’s why it’s important to have one way to build, deploy and manage both the applications and the infrastructure and having one substrate to manage all those things, especially when they have interdependencies is really powerful.”

![](https://cdn.thenewstack.io/media/2021/04/e007cab2-pulumi-02-1024x791.png)

## Developer Terms for Infrastructure Features

The new features in Pulumi 3.0 are familiar concepts for developers that are useful for infrastructure teams too, with a cloud twist.

Teams could already share and reuse code in [Pulumi Resources](https://thenewstack.io/infrastructure-is-code-and-with-pulumi-2-0-so-is-architecture-and-policy/), but thanks to the underlying language-neutral model of the cloud objects you’re working with, Pulumi Packages now make that work across multiple languages and for higher-level Components rather than low-level Resources.

“Previously, if you wrote your package in Node.js. you were tied to the Node ecosystem or if you wrote it in Python, it was only available in Python. Customers want to be able to write packages in Python if the infrastructure team is using Python and defining a Kubernetes cluster component but they want to enable their developers to spin those things up from Node.js or Go or their favorite language,” Duffy explained.

Packages can be written in any language Pulumi supports and consumed from any other supported language. Pulumi is also supplying some multilanguage components; the first handles provisioning and managing AWS EKS clusters, which is frequently complex. “We’re bootstrapping an ecosystem here,” Duffy suggested and noted that a numbers of customers will publish their own once the feature is generally available.

Native providers are a new type of Pulumi Package that automatically generates interfaces from the Azure, Google, and Amazon API specifications. Since September 2020, Azure has added 166 new services or new features in existing services; “They’re shipping services left and right all the time and if you can’t access them, that’s a problem.”

Because Microsoft documents the entire Azure surface in OpenAPI, Pulumi can support the new features immediately. The Azure native provider includes twice as many services as the provider Pulumi had built manually, like Azure Static Websites which was implemented within an hour of release. Azure native support is GA now, with Google Cloud in public preview and AWS coming in the second half of this year.

Fitzhugh noted that the fast support of new features differentiates Pulumi. “Cloud providers are moving their APIs ahead, sometimes at a breakneck pace; having support for that in real-time is critical. Especially when we’re looking at Kubernetes-based offerings from the cloud providers, they’re iterating and moving forward so fast that a lot of the frameworks would have trouble keeping up.”

Because they use all three major cloud providers, Snowflake is keen to give its developers a uniform platform for deploying containers and uses Pulumi to get that abstraction. “We’re not trying to build a complete other cloud platform on top of it, because where I see people get in trouble is when they try to almost reinvent what the cloud offer, and build it on top of it, because you’re never going to iterate as fast as the cloud,” Fizhugh explained.

“We’re focusing a lot on what the onboarding and user experience looks like for the end users: how do we streamline that, how do we make it as efficient as possible. But this is all enabled by the fact that we can describe what that environment looks like via code and we can iterate on it quickly at scale across, our many, many deployments across multiple cloud providers.”

The new Automation API is another feature that will help by letting customers build infrastructure-as-code into their own software. “What if infrastructure as code wasn’t a CLI-based experience; if it was just a library you could use within other programs,” Duffy explained.

This is particularly useful for SaaS providers like Cockroach Labs whose database as a service manages Kubernetes clusters behind the scenes on behalf of their customers but organizations are also using it to build self-service portals like Snowflake’s internal platform. “You can link with Pulumi and manage your infrastructure, and still get all the power of infrastructure as code, but not have to have this clunky CLI-based interface that you drive programmatically. You can build your own self-service portals: we’ve seen people building their own custom tools to do multiregion rollouts and lots of complex scenarios.”

![](https://cdn.thenewstack.io/media/2021/04/15a93a1a-pulumi-03-1024x791.png)

## Next for Pulumi

For Fitzhugh, the Automation API, and the new RBAC support that allows single sign-on with SAML, fine-grained permissions and identity workflows that synchronize with a central identity provider, move them a step closer to cloud engineering. “Touching the infrastructure in any way, either creating, managing or tearing it down is ideally all done through appropriate means; you really need a CI/CD pipeline to test and then deploy in an automated way. Our customers are asking us how can we have fewer humans interacting with the production environment and more pushing it through a pipeline. Identity management and the compliance and regulation pieces with SAML SSO and more granular access control piece is really useful to us.”

Identity is important for organizations using Pulumi for hybrid cloud; while 80% of Pulumi customers rely heavily on public cloud, hybrid and private cloud are also important. Pulumi supports Azure Arc and AWS Outposts; it’s already being used with the EKS distro and will support EKS Anywhere. Other customers are using Pulumi with vSphere, Duffy noted; “One customer is using Pulumi to spin up bare metal worldwide for new data centers.”

The Pulumi service doesn’t ever see your identity or credentials for the cloud provider, leaving you to do credential management the way you choose. “But we’ve almost become counselors on how to properly manage cloud credentials because unfortunately, we see it done incorrectly so often,” Duffy told us. “We try to push some of the best CI/CD integrations, we help folks do temporary credentials so there are no long-lived credentials — we see that all the time as an antipattern, but you look at the CI/CD providers themselves and they tell you to do that, which is really, really bad. We try to help people secure the CI/CD pipeline; when all deployments go through that, you can lock that down and secure it.”

He noted the complexity of dealing with so many systems that have their own identity models and suggested that in the future, Pulumi would be able to help with more of that. “Part of the whole cloud engineering platform vision is that we’re going to bring more into the fold of Pulumi,. If you want me to do delivery for you, great. If you just want to click a button to do some deployments, great. So there’s going to be a lot more of that and that will get us more into the identity credential management business.”

Duffy also talked about managing more of the dynamic state of infrastructure. “Pulumi captures the static topology of your infrastructure and applications which is beneficial and allows you to validate a lot of things, but it doesn’t validate everything I think the next frontier is capturing dynamic dependencies and understanding the semantic connections between them.” That could build on tools like the AWS security group analyzer for tracking down problems like the firewall rule that’s stopping two machines connecting.

Pulumi will also continue to improve and extent language support. Before adding new languages, the focus is on improving existing language support: Python support now includes static type checkers and the Go libraries are smaller and so faster to load. Pulumi 4.0 will include new languages: PowerShell, JVM and Ruby. “You can use Pulumi with PowerShell, it’s just the API’s aren’t designed to be PowerShell friendly, or idiomatically PowerShell.”

Customers coming from Chef and [Puppet](https://puppet.com/?utm_content=inline-mention) ecosystem are driving the interest in Ruby, Duffy said. “It’s actually the number one upvoted issue out of all of our open source issues.”

![](https://cdn.thenewstack.io/media/2021/04/46acfd5d-pulumi-04-1024x791.png)

Puppet is a sponsor of The New Stack.


Feature image par [stokpic](https://pixabay.com/fr/users/stokpic-692575/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=600468) de [Pixabay](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=600468).


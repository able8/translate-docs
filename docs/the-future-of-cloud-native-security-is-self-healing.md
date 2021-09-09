# The Future of Cloud Native Security Is Self-Healing

#### 26 Oct 2020 9:29am,   by [Jon Jarboe](https://thenewstack.io/author/jonjarboe/ "Posts by Jon Jarboe")

![](https://cdn.thenewstack.io/media/2020/10/98663b08-woman-3918661_1280-1024x646.jpg)

Jon has been helping software development organizations improve processes and tools for over 20 years, in contexts ranging from embedded systems to complex distributed applications and roles including support, development, customer success, management, pre- and post-sales.](https://www.linkedin.com/in/jon-jarboe/)

Security has always been a concern in the cloud. Individuals and organizations alike have a vested interest in securing their private information. Given the constant barrage of breach announcements, organizations clearly struggle to meet their data security obligations. Some argue that it’s an issue of priorities or regulation, but I think they are missing the point: it is simply too easy to fail when it comes to securing data in the cloud.

Breaches represent a real risk to organizations in regulated industries such as financial services and healthcare, with consequences to the business and even personal liability for senior leadership. Businesses in unregulated industries have faced fewer legal consequences, but legislators are showing an increasing appetite for consumer protections — which will affect all organizations. GDPR and CCPA are two recent examples, but they are certainly not the last. Eventually, every organization will likely have significant legal obligations that affect their approach to information security.

Whether you’re motivated by your obligation to users or your legal obligations, the goal is the same: to prevent breaches. More specifically, to reduce the opportunity for failure while delivering secure products more consistently.

## We’ve Faced Similar Challenges Before

That goal may sound familiar because we’ve faced similar problems before. Not so long ago, organizations were under increasing pressure from a different direction; in fact, they still are. Competitive pressures demand that organizations deliver innovation more quickly, and economic pressures require that they do so more predictably and at a lower cost. Delivered products need to be reliable and available because the businesses of end users depend on them. These pressures led to the DevOps revolution and the rise of the cloud, which have been indisputable successes.

**Sponsor Note**

![sponsor logo](https://cdn.thenewstack.io/media/2020/04/fe1154fd-accurics@2x.png)

Accurics enables compliance, governance, and security across the full cloud native stack in hybrid and multi-cloud environments. It seamlessly scans cloud automation code for risks before the stack is provisioned and monitors production cloud stacks for changes that introduce risk.

In 2019, the Cloud Native Computing Foundation (CNCF) [surveyed the industry](https://www.cncf.io/wp-content/uploads/2020/08/CNCF_Survey_Report.pdf) and found that cloud native technologies are widespread. In production environments, 84% of organizations are using containers; 78% use Kubernetes, and at least 41%, serverless. Industry analysts such as [IDC expect](https://www.idc.com/getdoc.jsp?containerId=prUS45613519) these trends to continue, with more than 90% of apps being cloud native by 2025 and two-thirds of enterprises deploying daily. Clearly, the cloud is the future of software.

While this is fantastic for the delivery of innovation, it complicates the problem of securing it. The good news is that many of the challenges that complicate effective security processes — such as manual processes, siloed expertise, and poor communication — have been solved before. These are the same challenges affecting development, quality and operations teams, that DevOps is successfully overcoming.

## A Way Forward

In many ways, securing systems can be seen as another iteration of DevOps — breaking up the security silo and integrating that expertise into the DevOps process. And much like earlier iterations of DevOps, it will require new approaches, tools and adjustments. As with everything DevOps, there are no silver bullets or “plug and play” solutions. Every team is different, and the best solution for any team will necessarily be built upon their unique needs and capabilities.

That said, we’re not lost in the wilderness. Many organizations have at least started their DevOps journey and have a basic understanding of what works for them. They probably have some experience with the impact that these changes will have on teams and schedules. They hopefully recognize the direction in which success lies. Those with experience in these areas — namely, the developers and DevOps teams that have already navigated these waters — can help lead the way to more secure releases.

We know that DevOps strategies — such as optimizing manual processes and automating them, establishing continuous feedback loops, and improving collaboration — have measurable results. Let’s start there.

One of the most disruptive aspects of traditional security processes is the late-stage security review. It’s hard to start early because you need to review the actual behavior, so many security teams assess security after deployment — in the staging or production environment (occasionally even in the live customer environment!) When issues are found, there’s very little context and it requires a lot of research to figure out what to do about them. There is a lot of room for improvement here: in the manual assessment process, the manual reporting and tracking process, and the manual remediation process. Truth be told, most of that effort is wasted because the issues are rarely fixed — they end up in a large backlog of issues that are never scheduled.

## The Infrastructure Can Heal Itself

It doesn’t have to be that way. Infrastructure-as-code (IaC) — the definition of resources and relationships in a codified format — is foundational for modern software teams, as a way to consistently and reliably build or rebuild the infrastructure as needed.

In addition to enabling continuous delivery and continuous deployment, IaC can also improve the security review process, because it defines what the runtime environment will look like. Security teams can start threat modeling and security assessments long before deployment actually happens, which means DevOps teams receive feedback in contexts where they can act on it.

That’s an improvement, but it’s still fundamentally a disruptive process. After a security problem is identified and reported, somebody needs to fix it with urgency. The level of effort required to manually fix that problem will necessarily displace other planned work, and automation will only make things worse. What’s needed is a way to not only identify issues early but also eliminate the manual work required to investigate and fix the problem. That is, to have a self-healing infrastructure that can not only identify problems but fix them autonomously so that planned work continues apace.

## Self-Healing Is the Future

As long as security relies on expert tools and processes, it cannot possibly fit into modern DevOps workflows. Modern software teams rely on infrastructure-as-code to improve consistency and scalability, and we can leverage those same workflows to improve security. By embedding appropriate security controls into the development process, that infrastructure can self-diagnose and self-heal — eliminating the all-too-common bottlenecks that limit the effectiveness of today’s security programs.

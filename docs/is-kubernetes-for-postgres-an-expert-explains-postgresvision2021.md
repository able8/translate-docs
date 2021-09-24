### Is Kubernetes for Postgres? An expert explains

by [Patrick Nelson](https://siliconangle.com/author/patricknelson/)

A database’s level of simplicity is the most important consideration for assessing when it’s worth shifting it to an advanced, sophisticated infrastructure platform, according to a Kubernetes application development specialist who works closely with PostgreSQL, an open-source relational database management system.

“The more commodifiable a particular database instance is, the better candidate it is to move,” said [Josh Berkus](https://www.linkedin.com/in/josh-berkus-1792412/)) (pictured), Kubernetes community manager at Red Hat Inc.

Berkus, however, is not simply talking about levels of complication that might hinder the migration, but that there’s less gain to be made by doing the move under certain circumstances. “The real advantage of moving stuff to Kubernetes is your ability to automate things,” he explained.

Berkus spoke with [Dave Vellante](https://twitter.com/dvellante), host of theCUBE, SiliconANGLE Media’s livestreaming studio, during [Postgres Vision 2021](https://www.thecube.net/postgres-vision-2021). They discussed just how appropriate Kubernetes is for Postgres data. _(\\* Disclosure below.)_

### What databases to migrate

Massive company-wide databases, where the entire business operation is run using one elaborate database would be an example of a less-suitable candidate, according to Berkus.

“To the extent that you can describe [a] particular database, what it does, who needs to use it, what’s in it in a simple one pager, then that’s probably a really good candidate for hosting on Kubernetes,” Berkus said.

Less appropriate are databases that are so complicated it gets hard, or even impossible, to explain the inputs and outputs.

“I’ve worked with people who have one big database, where the database is three terabytes in size. It powers their reporting system and their customer’s system and the web portal and everything else in one database,” Berkus said. “That’s the one that’s really going to be a hard call and that you might, in fact, never physically migrate to Kubernetes.”

Equally, databases that aren’t going to be taking advantage of automation are less optimal candidates. Berkus advises that one should assess whether application workflow and team organization can handle the new setup. If that’s in place, and particularly if development is unified, along with an infra team that owns everything, “then those people are going to be a really good candidate for moving that stack to Kubernetes.”

But even if all the considerations don’t fall into place, all is not lost. Berkus believes users should take a look at Service Catalog in Kubernetes if they’re dead set on the migration. That’s a way of exposing an external service in Kubernetes and making it appear a Kubernetes service.

“That’s what I tend to do with those kinds of [complicated] databases,” he said.

Exceptions to the aforementioned do exist, however. Big data infrastructure on Postgres is a contender for Kubernetes, according to Berkus.

“A lot of modern data analysis and data-mining platforms are built on top of Postgres,” he said. “Part of how they do their work is they actually run a bunch of little Postgres instances that they federate together.” Kubernetes then can be the tool that lets you to manage “all of those little Postgres instances.”

Don’t expect substantial performance differences by moving to Kubernetes if you’re already on cloud or network storage and have databases that can share hardware systems, according to the Postgres Kubernetes expert.

“Your whole point of moving to Kubernetes in general is going to be: take advantage of the automation,” Berkus concluded.

Watch the complete video interview below, and be sure to check out more of SiliconANGLE’s and theCUBE’s coverage of [Postgres Vision 2021](https://www.thecube.net/postgres-vision-2021). _(\\* Disclosure: TheCUBE is a paid media partner for the Postgres Vision event. Neither EnterpriseDB Corp., the sponsor for theCUBE’s event coverage, nor other sponsors have editorial control over content on theCUBE or SiliconANGLE.)_

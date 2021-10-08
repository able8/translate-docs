# Facebook outage More details about the October 4 outage

POSTED ON OCTOBER 5, 2021 TO [Networking & Traffic](https://engineering.fb.com/category/networking-traffic/)

![More details about the Oct. 4 Facebook outage](https://engineering.fb.com/wp-content/uploads/2021/10/datahall.jpg)

By [Santosh Janardhan](https://engineering.fb.com/author/santosh-janardhan/ "Posts by Santosh Janardhan")

Now that our platforms are up and running as usual after yesterday’s outage, I thought it would be worth sharing a little more detail on what happened and why — and most importantly, how we’re learning from it.

This outage was triggered by the system that manages our global backbone network capacity. The backbone is the network Facebook has built to connect all our computing facilities together, which consists of tens of thousands of miles of fiber-optic cables crossing the globe and linking all our data centers.

Those data centers come in different forms. Some are massive buildings that house millions of machines that store data and run the heavy computational loads that keep our platforms running, and others are smaller facilities that connect our backbone network to the broader internet and the people using our platforms.

When you open one of our apps and load up your feed or messages, the app’s request for data travels from your device to the nearest facility, which then communicates directly over our backbone network to a larger data center. That’s where the information needed by your app gets retrieved and processed, and sent back over the network to your phone.

The data traffic between all these computing facilities is managed by routers, which figure out where to send all the incoming and outgoing data. And in the extensive day-to-day work of maintaining this infrastructure, our engineers often need to take part of the backbone offline for maintenance — perhaps repairing a fiber line, adding more capacity, or updating the software on the router itself.

This was the source of yesterday’s outage. During one of these routine maintenance jobs, a command was issued with the intention to assess the availability of global backbone capacity, which unintentionally took down all the connections in our backbone network, effectively disconnecting Facebook data centers globally. Our systems are designed to audit commands like these to prevent mistakes like this, but a bug in that audit tool prevented it from properly stopping the command.

This change caused a complete disconnection of our server connections between our data centers and the internet. And that total loss of connection caused a second issue that made things worse.

One of the jobs performed by our smaller facilities is to respond to DNS queries. DNS is the address book of the internet, enabling the simple web names we type into browsers to be translated into specific server IP addresses. Those translation queries are answered by our authoritative name servers that occupy well known IP addresses themselves, which in turn are advertised to the rest of the internet via another protocol called the border gateway protocol (BGP).

To ensure reliable operation, our DNS servers disable those BGP advertisements if they themselves can not speak to our data centers, since this is an indication of an unhealthy network connection. In the recent outage the entire backbone was removed from operation,  making these locations declare themselves unhealthy and withdraw those BGP advertisements. The end result was that our DNS servers became unreachable even though they were still operational. This made it impossible for the rest of the internet to find our servers.

All of this happened very fast. And as our engineers worked to figure out what was happening and why, they faced two large obstacles: first, it was not possible to access our data centers through our normal means because their networks were down, and second, the total loss of DNS broke many of the internal tools we’d normally use to investigate and resolve outages like this.

Our primary and out-of-band network access was down, so we sent engineers onsite to the data centers to have them debug the issue and restart the systems. But this took time, because these facilities are designed with high levels of physical and system security in mind. They’re hard to get into, and once you’re inside, the hardware and routers are designed to be difficult to modify even when you have physical access to them. So it took extra time to activate the secure access protocols needed to get people onsite and able to work on the servers. Only then could we confirm the issue and bring our backbone back online.

Once our backbone network connectivity was restored across our data center regions, everything came back up with it. But the problem was not over — we knew that flipping our services back on all at once could potentially cause a new round of crashes due to a surge in traffic. Individual data centers were reporting dips in power usage in the range of tens of megawatts, and suddenly reversing such a dip in power consumption could put everything from electrical systems to caches at risk.

Helpfully, this is an event we’re well prepared for thanks to the “storm” drills we’ve been running for a long time now. In a storm exercise, we simulate a major system failure by taking a service, data center, or entire region offline, stress testing all the infrastructure and software involved. Experience from these drills gave us the confidence and experience to bring things back online and carefully manage the increasing loads. In the end, our services came back up relatively quickly without any further systemwide failures. And while we’ve never previously run a storm that simulated our global backbone being taken offline, we’ll certainly be looking for ways to simulate events like this moving forward.

Every failure like this is an opportunity to learn and get better, and there’s plenty for us to learn from this one. After every issue, small and large, we do an extensive review process to understand how we can make our systems more resilient. That process is already underway.

We’ve done extensive work hardening our systems to prevent unauthorized access, and it was interesting to see how that hardening slowed us down as we tried to recover from an outage caused not by malicious activity, but an error of our own making. I believe a tradeoff like this is worth it — greatly increased day-to-day security vs. a slower recovery from a hopefully rare event like this. From here on out, our job is to strengthen our testing, drills, and overall resilience to make sure events like this happen as rarely as possible.

### Read More in Networking & Traffic

[View All](https://engineering.fb.com/category/networking-traffic/)

AUG 9, 2021
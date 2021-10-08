##       A Tale of DNS & BGP: The Facebook Outage, October 2021    

![img](https://riskledger-website-media-uploads.s3-eu-west-1.amazonaws.com/chrome-dns-error-facebook.png)

This is a short tale. One of despair and intrigue, as we realise the  fragility of the global internet that we all so love and adore. On  Monday 4th October 2021, people stopped scrolling through Facebook, they gave up posting selfies to Instagram, they ceased texting on WhatsApp,  and Facebook employees abandoned doing any work. For all Facebook-owned  websites were down, thanks to a couple of three letter acronyms: DNS and BGP.

While Facebook have not released any key details nor a postmortem,  the internet is built on a series of open source standards and protocols that allow us to remotely inspect some of the fallout.

When you type `facebook.com`, `riskledger.com`  or any other name into your web browser of choice, an invisible  background process begins. This starts with the Domain Name System, or  DNS for short. It is often said to be the phone book of the internet. It is responsible for converting long series of arcane numbers, useful  only to machines, to something memorable.

The DNS system is contacted by your web browser asking simply, "what  is the internet address for facebook.com?". Now if one device was  answering every DNS question, it would quickly become overwhelmed, so  over the years a distributed hierarchy formed. This starts with a DNS  implementation built right into your device's operating system, to  devices called "recursers" operated by your ISP and sometimes large  organizations such as [Google](https://dns.google/) or [CloudFlare](https://1.1.1.1/), and even scaling up to central "root" devices that know the whereabouts of every .com, .co.uk and more.

Facebook in this case, operates a set of intermediary DNS servers  that are responsible for everything between your ISP's recursers and the roots. These are responsible for facebook.com, instagram.com,  whatsapp.com and everything else they operate. These servers are not  responding. This is what we see above from our web browser when it hints to us the error is `DNS_PROBE_FINISHED_NXDOMAIN`. Our web browser tried it's best to work out the internet address for Facebook, but it didn't get a reply.

While DNS may allow our computers to rapidly translate from names  useful to humans to numbers useful for computers, how does your device,  with it's own internet address, traverse the global internet to reach  Facebook?

No two devices on the internet are directly connected. At home, you  will have your residential ISP. In a datacentre, there will be multiple  commercial ISPs. Given two internet addresses, how do the two  communicate? This is where routing comes in.

![multiple internet routers](https://riskledger-website-media-uploads.s3-eu-west-1.amazonaws.com/bgp-routing.png)

Routing is the system by which a path between two devices is  calculated and established, potentially traversing dozens of ISP  networks in the process. Given there must be tens-of-thousands of ISPs,  how do we ensure they all can speak to each other?

We establish a standard! In 1989 no less, the Border Gateway  Protocol, or BGP for short, was accepted by the internet community as [RFC 1105](https://datatracker.ietf.org/doc/html/rfc1105). The document laid out a protocol by which routers operated by different companies and ISPs could exchange routing information which each other.

So what does this have to do with Facebook? We discussed earlier how  our web browsers were not receiving a response to their DNS questions,  this was however not a result of DNS itself, it was a result of  Facebook's routers ceasing to speak BGP with the rest of the internet,  and ultimately all of our ISP's routers stopped knowing where to send  our DNS requests, or any traffic to Facebook for that matter.

While we can speculate all day as to exactly what issue occurred on  the 4th of October in Facebook's infrastructure, it is however fun to  arm-chair investigate using open source information, and send [#hugops](https://www.pagerduty.com/blog/hugops-in-practice/) to the Facebook's network operations team!
# Scaling to 100k Users

# 扩展到 10 万用户

Feb 3, 2020  [Alex Pareto](https://alexpareto.com/)

2020 年 2 月 3 日 [亚历克斯帕累托](https://alexpareto.com/)

Many startups have been there - what feels like legions of  new users are signing up for accounts every day and the engineering team is scrambling to keep things running.

许多初创公司都在那里 - 感觉就像每天都有大量新用户注册帐户，而工程团队正在争先恐后地保持运行。

It’s a good a problem to have, but information on how to take a web  app from 0 to hundreds of thousands of users can be scarce. Usually  solutions come from either massive fires popping up or by identifying  bottlenecks (and often times both).

这是一个很好的问题，但有关如何将 Web 应用程序从 0 用户带到数十万用户的信息可能很少。通常解决方案来自于突然出现的大规模火灾或通过识别瓶颈（通常两者兼而有之）。

With that said, I’ve noticed that many of the main patterns for  taking a side project to something highly scalable are relatively  formulaic.

话虽如此，我注意到许多将副项目带到高度可扩展的主要模式是相对公式化的。

This is an attempt to distill the basics around that formula into  writing. We’re going to take our new photo sharing website, Graminsta,  from 1 to 100k users.

这是尝试将围绕该公式的基础知识提炼成写作。我们将把我们的新照片共享网站 Graminsta 的用户从 1 增加到 10 万。

## 1 User: 1 Machine

## 1 用户：1 台机器

Nearly every application, be it a website or a mobile app - has three key components: an API, a database, and a client (usually an app or a  website). The database stores persistent data. The API serves requests  for and around that data. The client renders that data to the user.

几乎每个应用程序，无论是网站还是移动应用程序 - 都具有三个关键组件：API、数据库和客户端（通常是应用程序或网站）。数据库存储持久数据。 API 为该数据及其周围的请求提供服务。客户端将该数据呈现给用户。

I’ve found in modern application development that thinking of the  client as a completely separate entity from the API makes it much easier to reason about scaling the application.

我发现在现代应用程序开发中，将客户端视为与 API 完全独立的实体，可以更容易地推理扩展应用程序。

When we first start building the application, it’s alright for all  three of these things to run on one server. In a way, this resembles our development environment: one engineer runs the database, API, and  client all on the same computer.

当我们第一次开始构建应用程序时，这三个东西都可以在一台服务器上运行。在某种程度上，这类似于我们的开发环境：一名工程师在同一台计算机上运行数据库、API 和客户端。

In theory, we could deploy this to the cloud on a single DigitalOcean Droplet or AWS EC2 instance like below:

理论上，我们可以将其部署到单个 DigitalOcean Droplet 或 AWS EC2 实例上的云中，如下所示：

![/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.11.27_AM.png](https://alexpareto.com/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.11.27_AM.png)


With that said, if we expect Graminsta to be used by more than 1  person, it almost always makes sense to split out the database layer.

话虽如此，如果我们希望 Graminsta 被超过 1 人使用，那么拆分数据库层几乎总是有意义的。

## 10 Users: Split out the Database Layer

## 10 个用户：拆分数据库层

Splitting out the database into a managed service like Amazon’s RDS  or Digital Ocean’s Managed Database will serve us well for a long time. It's slightly more expensive than self-hosting on a single machine or  EC2 instance - but with these services you get a lot of easy add ons out of the box that will come in handy down the line: multi-region  redundancy, read replicas, automated backups, and more.

将数据库拆分为托管服务，例如 Amazon 的 RDS 或 Digital Ocean 的托管数据库，将在很长一段时间内为我们提供良好的服务。它比单台机器或 EC2 实例上的自托管稍贵 - 但是通过这些服务，您可以获得许多开箱即用的简单附加组件：多区域冗余、只读副本、自动化备份等。

Here’s what the Graminsta system looks like now:

这是 Graminsta 系统现在的样子：

![/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.13.17_AM.png](https://alexpareto.com/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.13.17_AM.png)


## 100 Users: Split Out the Clients

## 100 个用户：拆分客户端

Lucky for us, our first few users love Graminsta. Now that traffic is starting to get more steady, it’s time to split out the client. One  thing to note is that **splitting out** entities is a key  aspect of building a scalable application. As one part of the system  gets more traffic, we can split it out so that we can handle scaling the service based on its own specific traffic patterns.

幸运的是，我们的前几个用户喜欢 Graminsta。现在流量开始变得更加稳定，是时候拆分客户端了。需要注意的一件事是**拆分**实体是构建可扩展应用程序的一个关键方面。当系统的某一部分获得更多流量时，我们可以将其拆分，以便我们可以根据其自身特定的流量模式来处理扩展服务。

This is why I like to think of the client as separate from the API. It makes it very easy to reason about building for multiple platforms:  web, mobile web, iOS, Android, desktop apps, third party services etc.  They’re all just clients consuming the same API.

这就是我喜欢将客户端与 API 分开的原因。这使得为多个平台构建的推理变得非常容易：Web、移动 Web、iOS、Android、桌面应用程序、第三方服务等。它们都只是使用相同 API 的客户端。

In the same vein, the biggest feedback we’re getting from users is  that they want Graminsta on their phones. So we’re going to launch a  mobile app while we’re at it.

同样，我们从用户那里得到的最大反馈是他们希望在他们的手机上安装 Graminsta。因此，我们将在此期间推出一个移动应用程序。

Here’s what the system looks like now:

这是系统现在的样子：

![/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-02-03_at_9.56.51_PM.png](https://alexpareto.com/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-02-03_at_9.56.51_PM.png)

_PM.png)

## 1,000 Users: Add a Load Balancer.

## 1,000 个用户：添加负载均衡器。

Things are picking up at Graminsta. Users are uploading photos left  and right. We’re starting to get more sign ups. Our lonely API instance  is having trouble keeping up with all the traffic. We need more compute  power! 

Graminsta 的情况正在好转。用户正在左右上传照片。我们开始获得更多注册。我们孤独的 API 实例无法跟上所有流量。我们需要更多的计算能力！

Load balancers are very powerful. The key idea is that we place a  load balancer in front of the API and it will route traffic to an  instance of that service. This allows for horizontal scaling (increasing the amount of requests we can handle by adding more servers running the same code).

负载均衡器非常强大。关键思想是我们在 API 前面放置一个负载均衡器，它将流量路由到该服务的一个实例。这允许水平扩展（通过添加更多运行相同代码的服务器来增加我们可以处理的请求量）。

We’re going to place a separate load balancer in front of our web  client and our API. This means we can have multiple instances running  our API and web client code. The load balancer will route requests to  whichever instance has the least traffic.

我们将在 Web 客户端和 API 前面放置一个单独的负载均衡器。这意味着我们可以有多个实例运行我们的 API 和 Web 客户端代码。负载均衡器会将请求路由到流量最少的实例。

What we also get out of this is redundancy. When one instance goes  down (maybe it gets overloaded or crashes), then we have other instances still up to respond to incoming requests - instead of the whole system  going down.

我们还从中得到的是冗余。当一个实例出现故障（可能它会过载或崩溃）时，我们还有其他实例仍在响应传入的请求——而不是整个系统出现故障。

A load balancer also enables autoscaling. We can set up our load  balancer to increase the number of instances during the Superbowl when  everyone is online and decrease the number of instances when all of our  users are asleep.

负载均衡器还支持自动缩放。我们可以设置我们的负载均衡器，以在每个人都在线时增加 Superbowl 期间的实例数量，并在我们所有用户都睡着时减少实例数量。

With a load balancer, our API layer can scale to practically  infinity, we will just keep adding instances as we get more requests.

使用负载均衡器，我们的 API 层可以扩展到几乎无限，随着我们收到更多请求，我们将继续添加实例。

![/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.25.50_AM.png](https://alexpareto.com/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.25.50_AM.png)


> Side note: At this point what we have so far is very similar to  what PaaS companies like Heroku or AWS’s Elastic Beanstalk provide out  of the box (and why they’re so popular). Heroku puts the database on a  separate host, manages a load balancer with autoscaling, and lets you  host your web client separately from your API. This is a great reason to use a service like Heroku for projects or early stage startups - all of the necessary basics come out of the box.

> 旁注：到目前为止，我们所拥有的与 Heroku 或 AWS 的 Elastic Beanstalk 等 PaaS 公司提供的开箱即用非常相似（以及它们为何如此受欢迎）。 Heroku 将数据库放在单独的主机上，管理具有自动缩放功能的负载均衡器，并允许您将 Web 客户端与 API 分开托管。这是将 Heroku 之类的服务用于项目或早期初创公司的重要原因——所有必要的基础知识都是开箱即用的。

## 10,000 Users: CDN

## 10,000 名用户：CDN

We probably should have done this from the beginning, but we’re  moving fast here at Graminsta. Serving and uploading all these images is starting to put way too much load on our servers.

我们可能应该从一开始就这样做，但我们在 Graminsta 的发展速度很快。提供和上传所有这些图像开始给我们的服务器带来太多负载。

We should be using a cloud storage service to host static content at  this point: think images, videos, and more (AWS’s S3 or Digital Ocean’s  Spaces). In general, our API should avoid handling things like serving  images and image uploads.

此时我们应该使用云存储服务来托管静态内容：想想图像、视频等（AWS 的 S3 或 Digital Ocean 的 Spaces）。一般来说，我们的 API 应该避免处理诸如提供图像和图像上传之类的事情。

The other thing we get out of a cloud storage service is a CDN (in  AWS this is an add-on called Cloudfront, but many cloud storage services offer it out of the box). A CDN will automatically cache our images at  different data centers throughout the world.

我们从云存储服务中获得的另一件事是 CDN（在 AWS 中，这是一个名为 Cloudfront 的附加组件，但许多云存储服务都提供开箱即用的功能）。 CDN 将自动将我们的图像缓存在世界各地的不同数据中心。

While our main data center may be hosted in Ohio, if someone requests an image from Japan, our cloud provider will make a copy and store it  at their data center in Japan. The next person to request that image in  Japan will then recieve it much faster. This is important when we need  to serve larger file sizes like images or videos that take a long time  to load + send across the world.

虽然我们的主要数据中心可能位于俄亥俄州，但如果有人从日本请求图像，我们的云提供商将制作副本并将其存储在他们位于日本的数据中心。下一个在日本请求该图像的人将更快地收到它。当我们需要提供更大的文件大小时，这很重要，例如需要很长时间才能加载和发送到世界各地的图像或视频。

![/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.30.06_AM.png](https://alexpareto.com/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.30.06_AM.png)


## 100,000 Users: Scaling the Data Layer

## 100,000 个用户：扩展数据层

The CDN helped us out a lot - things are booming at Graminsta. A  YouTube celebrity, Mavid Mobrick, just signed up and posted us on their  story. The API CPU and memory usage is low across the board - thanks to  our load balancer adding 10 API instances to the environment - but we’re starting to get a lot of timeouts on requests…why is everything taking  so long?

CDN 帮了我们很多忙——Graminsta 的事情正在蓬勃发展。 YouTube 名人 Mavid Mobrick 刚刚注册并发布了他们的故事。 API CPU 和内存使用率普遍较低——这要归功于我们的负载均衡器向环境中添加了 10 个 API 实例——但我们开始遇到很多请求超时......为什么一切都需要这么长时间？

After some digging we see it: the Database CPU is hovering at 80-90%. We’re maxed out.

经过一番挖掘，我们看到：数据库 CPU 徘徊在 80-90%。我们用完了。

Scaling the data layer is probably the trickiest part of the  equation. While for API servers serving stateless requests, we can  merely add more instances, the same is not true with *most* database systems. In this case, we’re going to explore the popular relational database systems (PostgreSQL, MySQL, etc.).

缩放数据层可能是等式中最棘手的部分。而对于服务无状态请求的 API 服务器，我们只能添加更多实例，但对于*大多数*数据库系统而言，情况并非如此。在这种情况下，我们将探索流行的关系数据库系统（PostgreSQL、MySQL 等）。

### Caching 

### 缓存

One of the easiest ways to get more out of our database is by  introducing a new component to the system: the cache layer. The most  common way to implement a cache is by using an in-memory key value store like Redis or Memcached. Most clouds have a managed version of these  services: Elasticache on AWS and Memorystore on Google Cloud.

从我们的数据库中获取更多信息的最简单方法之一是向系统引入一个新组件：缓存层。实现缓存的最常见方法是使用内存中的键值存储，如 Redis 或 Memcached。大多数云都有这些服务的托管版本：AWS 上的 Elasticache 和 Google Cloud 上的 Memorystore。

The cache comes in handy when the service is making lots of repeated  calls to the database for the same information. Essentially we hit the  database once, save the information in the cache, and never have to  touch the database again.

当服务对数据库进行大量重复调用以获取相同信息时，缓存就派上用场了。本质上，我们访问数据库一次，将信息保存在缓存中，再也不必接触数据库。

For example, in Graminsta every time someone goes to Mavid Mobrick’s  profile page, the API layer requests Mavid Mobrick’s profile information from the database. This is happening over and over again. Since Mavid  Mobrick’s profile information isn’t changing on every request, that info is a great candidate to cache.

例如，在 Graminsta 中，每次有人访问 Mavid Mobrick 的个人资料页面时，API 层都会从数据库中请求 Mavid Mobrick 的个人资料信息。这一次又一次地发生。由于 Mavid Mobrick 的个人资料信息不会在每次请求时都发生变化，因此该信息非常适合缓存。

We’ll cache the result from the database in Redis under the key `user:id` with an expiration time of 30 seconds. Now when someone goes to Mavid  Mobrick’s profile, we check Redis first and just serve the data straight out of Redis if it exists. Despite Mavid Mobrick being the most popular on the site, requesting the profile puts hardly any load on our  database.

我们将在 Redis 中将数据库的结果缓存在键“user:id”下，过期时间为 30 秒。现在，当有人访问 Mavid Mobrick 的个人资料时，我们首先检查 Redis，如果存在，则直接从 Redis 中提供数据。尽管 Mavid Mobrick 是网站上最受欢迎的，但请求个人资料几乎不会给我们的数据库带来任何负载。

The other plus of most cache services, is that we can scale them out  easier than a database. Redis has a built in Redis Cluster mode that, in a similar way to a load balancer[1](https://alexpareto.com/scalability/systems/2020/02/03/scaling-100k.html#fn:1) , lets us distribute our Redis cache across multiple machines (thousands if one so pleases).

大多数缓存服务的另一个优点是我们可以比数据库更容易扩展它们。 Redis 有一个内置的 Redis 集群模式，类似于负载均衡器[1](https://alexpareto.com/scalability/systems/2020/02/03/scaling-100k.html#fn:1) ，让我们在多台机器上分布我们的 Redis 缓存（如果有的话，数千个)。

Nearly all highly scaled applications take ample advantage of  caching, it’s an absolutely integral part of making a fast API. Better  queries and more performant code are all a part of the equation, but  without a cache none of these will be sufficient to scale to millions of users.

几乎所有高度扩展的应用程序都充分利用了缓存，这是制作快速 API 绝对不可或缺的一部分。更好的查询和更高性能的代码都是等式的一部分，但如果没有缓存，这些都不足以扩展到数百万用户。

### Read Replicas

### 只读副本

The other thing we can do now that our database has started to get  hit quite a bit, is to add read replicas using our database management  system. With the managed services above, this can be done in one-click. A read replica will stay up to date with your master DB, and will be able to be used for SELECT statements.

我们现在可以做的另一件事是我们的数据库已经开始受到相当多的攻击，那就是使用我们的数据库管理系统添加只读副本。使用上述托管服务，可以一键完成。只读副本将与您的主数据库保持同步，并可用于 SELECT 语句。

Here’s our system now:

现在是我们的系统：

![/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.35.01_AM.png](https://alexpareto.com/assets/Scaling%20to%20100k%20Users/Screen_Shot_2020-01-21_at_8.35.01_AM.png)


## Beyond

##  超过

As the app continues to scale, we’re going to want to focus on  splitting out services that can be scaled independently. For example, if we start to make use of websockets, it would make sense to pull out our websocket handling code. We can put this on new instaces behind their  own load balancer that can scale up and down based on how many websocket connections have been opened or closed, independently of how many HTTP  requests we have coming in.

随着应用程序的不断扩展，我们将专注于拆分可以独立扩展的服务。例如，如果我们开始使用 websockets，那么拉出我们的 websocket 处理代码是有意义的。我们可以把它放在他们自己的负载均衡器后面的新实例上，它可以根据打开或关闭的 websocket 连接的数量进行扩展和缩减，而与我们传入的 HTTP 请求的数量无关。

We’re also going to continue to bump up against limitations on the  data layer. This is when we are going to want to start looking into  partitioning and sharding the database. These both require more  overhead, but effectively allow the data layer to scale infinitely.

我们还将继续挑战数据层的限制。这是我们要开始研究对数据库进行分区和分片的时候。这些都需要更多的开销，但有效地允许数据层无限扩展。

We will want to make sure that we have monitoring installed using a  service like New Relic or Datadog. This will ensure we understand what  requests are slow and where improvement needs to be made. As we scale we will want to be focused on finding bottlenecks and fixing them - often  by taking advantage of some of the ideas in the previous sections.

我们将要确保我们使用 New Relic 或 Datadog 等服务安装了监控。这将确保我们了解哪些请求是缓慢的以及需要改进的地方。随着我们的扩展，我们希望专注于发现瓶颈并修复它们——通常是利用前面部分中的一些想法。

Hopefully at this point, we have some other people on the team to help out as well!

希望在这一点上，我们团队中还有其他人可以提供帮助！

#### References:

####  参考：

This post was inspired by one of [my favorite posts on High Scalability](http://highscalability.com/blog/2016/1/11/a-beginners-guide-to-scaling-to-11-million-users-on-amazons.html). I wanted to flesh the article out a bit more for the early stages and  make it a bit more cloud agnostic. Definitely check it out if you’re  interested in these kind of things.

这篇文章的灵感来自[我最喜欢的关于高可扩展性的文章](http://highscalability.com/blog/2016/1/11/a-beginners-guide-to-scaling-to-11-million-users-on-amazons.html)。我想在早期阶段更充实这篇文章，让它更加与云无关。如果您对这些事情感兴趣，请务必检查一下。

#### Translations

#### 翻译

Readers have kindly translated this post into [Chinese](https://www.infoq.cn/article/Tyx5HwaD9OKNX4xzFaFo),[Korean](https://leonkim.dev/systems/scaling-100k/), and [Spanish] (https://www.ibidemgroup.com/edu/salto-a-100000-usuarios/).

读者已将这篇文章翻译成[中文](https://www.infoq.cn/article/Tyx5HwaD9OKNX4xzFaFo)、[韩文](https://leonkim.dev/systems/scaling-100k/)和[西班牙文]（https://www.ibidemgroup.com/edu/salto-a-100000-usuarios/)。

#### Footnotes 

#### 脚注

1. While similar in terms of allowing us to spread load across  multiple instances, the underlying implementation of Redis Cluster is  much different than a load balancer. [↩](https://alexpareto.com/scalability/systems/2020/02/03/scaling-100k.html#fnref:1)

1. 虽然在允许我们将负载分散到多个实例方面类似，但 Redis Cluster 的底层实现与负载均衡器有很大不同。 [↩](https://alexpareto.com/scalability/systems/2020/02/03/scaling-100k.html#fnref:1)





Hey! My name's Alex. I work on engineering at    [Brex](https://brex.com). I'm based in SF.
I write here mostly about software and startups.

嘿！我叫亚历克斯。我在 [Brex](https://brex.com) 从事工程工作。我在SF。
我在这里写的主要是关于软件和初创公司。

Previously I led engineering at    [NTWRK](https://thentwrk.com) and founded    [demeanor.co](https://www.demeanor.co) (Y    Combinator S18, acquired by NTWRK). Before that, I worked on the video    platform at [Facebook](http://facebook.com). I    studied at [USC](https://www.usc.edu) and    [Andover](http://www.andover.edu/).

之前我在[NTWRK](https://thentwrk.com)领导工程并创立了[demeanor.co](https://www.demeanor.co)（Y Combinator S18，被NTWRK收购)。在此之前，我在 [Facebook](http://facebook.com) 的视频平台工作。我在 [USC](https://www.usc.edu) 和 [Andover](http://www.andover.edu/) 学习。

I love to chat. Find me at    [a@alexpareto.com](mailto:a@alexpareto.com) or on Twitter    [@alexpareto](https://twitter.com/alexpareto).



我喜欢聊天。在 [a@alexpareto.com](mailto:a@alexpareto.com) 或 Twitter [@alexpareto](https://twitter.com/alexpareto) 上找到我。



 




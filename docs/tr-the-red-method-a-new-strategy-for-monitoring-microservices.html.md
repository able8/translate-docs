# The RED method: A new strategy for monitoring microservices

# RED 方法：监控微服务的新策略

https://www.infoworld.com/article/3638693/the-red-method-a-new-strategy-for-monitoring-microservices.html

### By using the RED metrics—rate, error, and duration—you can get a solid understanding of how your services are performing for end-users.

### 通过使用 RED 指标（速率、错误和持续时间），您可以深入了解您的服务对最终用户的执行情况。

Monitoring an application is crucial for providing a quality product and experience for users. But simply collecting a ton of application metrics doesn’t solve the true problem. What software companies need is a way to get actionable insights from their metrics so they can quickly fix any issues their users are experiencing.

监控应用程序对于为用户提供优质产品和体验至关重要。但是仅仅收集大量的应用程序指标并不能解决真正的问题。软件公司需要的是一种从他们的指标中获得可行见解的方法，以便他们能够快速解决用户遇到的任何问题。

Enter the RED method.

输入 RED 方法。

## **RED method origins**

## **RED 方法起源**

The RED method is a monitoring methodology coined by Tom Wilkie based on what he learned while working at Google. RED is derived from some best practices established at Google known as the “Four Golden Signals,” developed by Google’s SRE team.

RED 方法是 Tom Wilkie 根据他在 Google 工作时学到的知识创造的一种监控方法。 RED 源自 Google 建立的一些最佳实践，称为“四个黄金信号”，由 Google 的 SRE 团队开发。

The primary rationale behind RED is that previous monitoring philosophies and methodologies such as [the USE method](https://www.infoworld.com/article/3638772/making-the-use-method-of-monitoring-useful.html) didn't fully align with the objectives of software companies and modern software architectures. USE applies more to hardware and infrastructure, while the RED method intends to focus on what users of an application are actually experiencing.

RED 背后的主要理由是以前的监控理念和方法，例如 [USE 方法](https://www.infoworld.com/article/3638772/making-the-use-method-of-monitoring-useful.html)没有完全符合软件公司和现代软件架构的目标。 USE 更多地应用于硬件和基础设施，而 RED 方法旨在关注应用程序用户的实际体验。

The goal of the RED method is to ensure that the software application functions properly for the end-users above all else. In the modern era of microservice architectures, containers, and cloud infrastructure, metrics related to hardware aren’t nearly as important as long as your service level objectives (SLOs) are being met.

RED 方法的目标是确保软件应用程序为最终用户正常运行，高于一切。在微服务架构、容器和云基础设施的现代时代，只要您的服务水平目标 (SLO) 得到满足，与硬件相关的指标就不再那么重要了。

## RED method explained

## RED 方法解释

RED stands for rate, errors, and duration. These represent the three key metrics you want to monitor for each service in your architecture:

RED 代表速率、错误和持续时间。这些代表了您要为架构中的每个服务监控的三个关键指标：

- Rate - The number of requests the service is handling per second.
- Error - The number of failed requests per second.
- Duration - The amount of time each request takes.

- 速率 - 服务每秒处理的请求数。
- 错误 - 每秒失败的请求数。
- 持续时间 - 每个请求所花费的时间。

Using these three metrics, you can get a solid understanding of how your services are performing. The number of requests gives you a baseline for how much traffic is going to your service. The portion of those requests that are errors lets you know if a service is functioning within your SLO. Finally, the duration of time it takes for each request to be handled by your service gives you insight into the overall user experience of your application.

使用这三个指标，您可以深入了解服务的执行情况。请求数为您提供了一个基准，了解有多少流量流向您的服务。这些请求中的错误部分可让您了解服务是否在您的 SLO 中运行。最后，您的服务处理每个请求所需的时间可以让您深入了解应用程序的整体用户体验。

## Benefits of the RED method

## RED 方法的好处

The first benefit of the RED method is helping to reduce the cognitive load required for engineers to determine why a service is having issues. RED abstracts away the internal details of each service into something that can be understood across the entire architecture. This not only means problems can be solved faster, but also that it is easier to scale an operations team because members can now be on-call for services they didn’t write themselves.

RED 方法的第一个好处是有助于减少工程师确定服务出现问题的原因所需的认知负担。 RED 将每个服务的内部细节抽象为可以在整个架构中理解的东西。这不仅意味着可以更快地解决问题，而且更容易扩展运营团队，因为成员现在可以随叫随到，以获取不是他们自己编写的服务。

The RED abstraction makes it easy to understand what is going wrong and to determine how to fix it. Even if the service they are trying to fix is effectively a black box that they don’t understand internally, the engineer can look at telemetry data and determine the best action to improve the user experience. Because the same metrics are used for every service the amount of training time or service-specific knowledge is reduced as well.

RED 抽象可以很容易地理解出了什么问题并确定如何修复它。即使他们试图修复的服务实际上是一个他们内部无法理解的黑匣子，工程师也可以查看遥测数据并确定改善用户体验的最佳措施。因为每项服务都使用相同的指标，所以培训时间或服务特定知识的数量也会减少。

Another benefit of the RED method is that it more closely aligns with users and the company’s overall objectives. Users don’t care about your infrastructure. They don’t care about your CPU utilization, your memory usage, or any other hardware metrics. They care if they start seeing error messages when they use your app. They care if pages on your website take a long time to load. The RED method makes it very clear when a service isn’t living up to your SLO and your users are having a poor experience. 

RED 方法的另一个好处是它更接近用户和公司的整体目标。用户不关心您的基础设施。他们不关心你的 CPU 使用率、内存使用率或任何其他硬件指标。他们关心他们在使用您的应用程序时是否开始看到错误消息。他们关心您网站上的页面是否需要很长时间才能加载。当服务不符合您的 SLO 并且您的用户体验不佳时，RED 方法非常清楚。

A final benefit of the RED method is that automating tasks and alerts across your services becomes easier. Automating repetitive tasks is simpler and safer because all services are treated the same. You can also standardize things like dashboard layouts across services because the same three metrics are being used.

RED 方法的最后一个好处是跨服务自动化任务和警报变得更加容易。自动化重复性任务更简单、更安全，因为所有服务都被同等对待。您还可以跨服务标准化仪表板布局等内容，因为使用了相同的三个指标。

## Limitations of the RED method

## RED 方法的局限性

All of those benefits don’t mean the RED method is perfect. The RED method is primarily designed for request-driven applications, so for use cases that involve batch processing or streaming, it may not provide the insight you need.

所有这些好处并不意味着 RED 方法是完美的。 RED 方法主要是为请求驱动的应用程序设计的，因此对于涉及批处理或流式处理的用例，它可能无法提供您需要的洞察力。

A second downside is that the “external” view of RED means that you could have a hard time knowing how close a service is to failing. A slight increase in traffic may cause your response duration to increase and you may not have internal application metrics to determine why. Using the RED method means your metrics can be interpreted differently depending on multiple factors, so it does require deliberate implementation.

第二个缺点是 RED 的“外部”视图意味着您可能很难知道服务有多接近失败。流量的轻微增加可能会导致您的响应持续时间增加，并且您可能没有内部应用程序指标来确定原因。使用 RED 方法意味着您的指标可以根据多个因素进行不同的解释，因此它确实需要经过深思熟虑的实施。

The good news is that the RED method was never intended as a way to cover all aspects of monitoring. Tom Wilkie recommends that the RED monitoring methodology be used in combination with other monitoring methods like USE to give teams full monitoring coverage of their application. 

好消息是 RED 方法从未打算用作涵盖监控所有方面的方法。 Tom Wilkie 建议将 RED 监控方法与 USE 等其他监控方法结合使用，以使团队能够全面监控其应用程序。


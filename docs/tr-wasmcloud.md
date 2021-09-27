# What’s next, after Kubernetes?

# Kubernetes 之后的下一步是什么？

*[Stuart Harris](https://awesome.red-badger.com/stuartharris/) — 14th January 2021*

*[斯图尔特·哈里斯](https://awesome.red-badger.com/stuartharris/)——2021 年 1 月 14 日*

Some *good* things came out of 2020! An exciting one, for me,  was the progress that the global collective of open source software  engineers has been making towards the future of services in the Cloud.

2020 年发生了一些*好的*事情！对我来说，令人兴奋的是，全球开源软件工程师集体朝着云服务的未来所取得的进展。

Microservices are continuing to gain traction for Cloud applications and [Kubernetes](https://kubernetes.io) has, without question, become their de facto hosting environment. But I think that could all be about to change.

微服务继续受到云应用程序的青睐，[Kubernetes](https://kubernetes.io) 毫无疑问已成为它们事实上的托管环境。但我认为这一切都可能会改变。

Kubernetes is really good. But it does nothing to address what I  think is one of the biggest problems we have with microservices — the  ratio of functional code (e.g. our core business logic) to  non-functional code (e.g. talking to a database) is way too low. If you  open up the code of any microservice, or ask any cross-functional team,  you’ll see what I mean. You can’t see the functional wood for the  non-functional trees. As an industry, we’re spending way too much time  and effort on things that really don’t matter (but are still needed to  get the job done). This means that we can’t move fast enough (or build  enough features) to provide enough *real* value.

Kubernetes 真的很好。但它没有解决我认为微服务最大的问题之一——功能性代码（例如我们的核心业务逻辑）与非功能性代码（例如与数据库对话）的比率太低了。如果你打开任何微服务的代码，或者询问任何跨职能团队，你就会明白我的意思。对于非功能性树木，您看不到功能性木材。作为一个行业，我们在真正无关紧要的事情上花费了太多时间和精力（但仍然需要完成工作）。这意味着我们不能足够快地移动（或构建足够的功能）来提供足够的*真实*价值。

In this post, I first of all want to explore the Onion architecture,  how it applies to microservices and how we might peel off the outer,  non-functional layers of the onion, so that we can focus on the  functional core.

在这篇文章中，我首先想探索 Onion 架构，它如何应用于微服务以及我们如何剥离洋葱的外部非功能层，以便我们可以专注于功能核心。

We'll also see how Kubernetes can be augmented to support this idea (with a service mesh like [Istio](https://istio.io/) or [Linkerd](https://linkerd.io/), and a distributed application runtime like [Dapr](https://dapr.io/)).

我们还将看到如何增强 Kubernetes 以支持这个想法（使用像 [Istio](https://istio.io/) 或 [Linkerd](https://linkerd.io/) 这样的服务网格，以及分布式应用程序运行时，如 [Dapr](https://dapr.io/))。

Finally, and most importantly we'll ask what comes after Kubernetes  (spoiler: a WebAssembly actor runtime, possibly something like [WasmCloud](https://wascc.dev/)) that can support core business logic more natively, allowing us to write that logic in any language, run it literally anywhere, and  securely connect it to capability providers that we don't have to write  ourselves (but could if we needed to).

最后，也是最重要的一点，我们会问 Kubernetes（剧透：WebAssembly actor 运行时，可能类似于 [WasmCloud](https://wascc.dev/)）之后会发生什么，它可以更原生地支持核心业务逻辑，让我们能够用任何语言编写该逻辑，在任何地方运行它，并将其安全地连接到我们不必自己编写的能力提供者（但如果我们需要，可以这样做)。

## 1. The Onion Architecture

## 1. 洋葱架构

Similar to [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) (also known as “Ports and Adapters”) and [Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), the [Onion Architecture](https://jeffreypalermo.com/2008/07/the-onion-architecture-part-1/) advocates a structure for our application that allows us to segregate core business logic.

类似于 [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))（也称为“端口和适配器”）和 [Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)，[洋葱架构](https://jeffreypalermo.com/2008/07/the-onion-architecture-part-1/)为我们的应用程序提倡一种允许我们分离核心业务逻辑的结构。

Imagine the concentric layers of an onion where you can only call  inwards (i.e. from an outer layer to an inner layer). Let’s see how this might work by starting at its core.

想象一下洋葱的同心层，您只能向内调用（即从外层到内层）。让我们从它的核心开始，看看它是如何工作的。

I’ve augmented each layer’s description with a simple code example. I’ve used [Rust](https://www.rust-lang.org/) for this, because it’s awesome! [Fight me](https://blog.red-badger.com/now-is-a-really-good-time-to-make-friends-with-rust). Even if you don’t know Rust, it should be easy to understand this  example, but I’ve added a commentary that may help, just in case. You  can try out the example from this [Github repository](https://github.com/StuartHarris/onion).

我用一个简单的代码示例扩充了每一层的描述。我为此使用了 [Rust](https://www.rust-lang.org/)，因为它太棒了！[与我战斗](https://blog.red-badger.com/now-is-a-really-good-time-to-make-friends-with-rust)。即使你不了解 Rust，也应该很容易理解这个例子，但我添加了一个评论，可能会有所帮助，以防万一。您可以从 [Github 存储库](https://github.com/StuartHarris/onion) 中尝试示例。

![Onion Architecture](https://awesome.red-badger.com/stuartharris/wasmcloud/onion.svg)

The *core* is pure in the functional sense, i.e. it has no  side-effects. This is where our business logic resides. It is  exceptionally easy to test because its pure functions only take and  return values. In our example, our *core* is just a single function that takes 2 integers and adds them together. In the *core*, we don’t think about IO at all.

*core* 在功能意义上是纯粹的，即它没有副作用。这是我们的业务逻辑所在的地方。它非常容易测试，因为它的纯函数只接受和返回值。在我们的示例中，我们的 *core* 只是一个函数，它接受 2 个整数并将它们加在一起。在 *core* 中，我们根本不考虑 IO。

```
/// 1. Pure.Don't think about IO at all
mod core {
    pub fn add(x: i32, y: i32) -> i32 {
        x + y
    }
}
```

Surrounding the *core* is the *domain*, where we do  think about IO, but not its implementation. This layer orchestrates our  logic, providing hooks to the outside world, whilst having no knowledge  of that world (databases etc.). 

围绕 *core* 的是 *domain*，我们确实在其中考虑了 IO，但并未考虑其实现。这一层编排我们的逻辑，提供到外部世界的钩子，同时不了解那个世界（数据库等）。

In our code example, we have to use an asynchronous function. Calling out to a database (or something else, we actually don’t care yet) will  take some milliseconds, so it’s not something we want to stop for. The `async` keyword tells the compiler to return a `Future` which may complete at some point. The `Result` is implicitly wrapped in this `Future`.

在我们的代码示例中，我们必须使用异步函数。调用数据库（或其他东西，我们实际上并不关心）将花费几毫秒，所以我们不想停下来。 `async` 关键字告诉编译器返回一个可能在某个时间完成的 `Future`。 `Result` 隐含在这个 `Future` 中。

Importantly, our function takes another function as an argument. It’s this latter function that will actually do the work of going to the  database, so it will also need to return a `Future` and we will need to `await` for it to be completed. Incidentally, the question mark after the `await` allows the function to exit early with an error if something went wrong.

重要的是，我们的函数接受另一个函数作为参数。后一个函数将实际执行访问数据库的工作，因此它还需要返回一个 `Future`，我们需要 `await` 以使其完成。顺便说一下，如果出现问题，`await` 后面的问号允许函数提前退出并显示错误。

```
/// 2. think about IO but not its implementation
mod domain {
    use super::core;
    use anyhow::Result;
    use std::future::Future;

    pub async fn add<Fut>(get_x: impl Fn() -> Fut, y: i32) -> Result<i32>
    where
        Fut: Future<Output = Result<i32>>,
    {
        let x = get_x().await?;
        Ok(core::add(x, y))
    }
}
```

Those 2 inner layers are where all our application logic resides. Ideally we wouldn’t write any other code. However, in real life, we have to talk to databases, an event bus, or another service, for example. So the outer 2 layers of the onion are, sadly, necessary.

这两个内层是我们所有应用程序逻辑所在的地方。理想情况下，我们不会编写任何其他代码。然而，在现实生活中，我们必须与数据库、事件总线或其他服务进行对话。因此，遗憾的是，洋葱的外两层是必要的。

The *infra* layer is where our IO code goes. This is the code that knows how to do things like calling a database.

*infra* 层是我们的 IO 代码所在。这是知道如何执行诸如调用数据库之类的操作的代码。

```
/// 3. IO implementation
mod infra {
    use anyhow::Result;

    pub async fn get_x() -> Result<i32> {
        // call DB, which returns 7
        Ok(7)
    }
}
```

And, finally, the *api* layer is where we interact with our users. We present an API and wire up dependencies (in this example, by passing our *infra* function into our *domain* function):

最后，*api* 层是我们与用户交互的地方。我们提供了一个 API 并连接了依赖项（在这个例子中，通过将我们的 *infra* 函数传递到我们的 *domain* 函数中）：

```
/// 4. inject dependencies
mod api {
    use super::{domain, infra};
    use anyhow::Result;

    pub async fn add(y: i32) -> Result<i32> {
        let result = domain::add(infra::get_x, y).await?;
        Ok(result)
    }
}
```

We’ll need an entry-point for our service:

我们的服务需要一个入口点：

```
fn main() {
    async_std::task::block_on(async {
        println!(
            "When we add 3 to the DB value (7), we get {:?}",
            api::add(3).await
        );
    })
}
```

Then, when we run it we see it works!

然后，当我们运行它时，我们看到它起作用了！

```
cargo run
    Finished dev [unoptimized + debuginfo] target(s) in 0.06s
     Running `target/debug/onion`
When we add 3 to the DB value (7), we get Ok(10)
```

Ok, now we have that out of the way, let's see how we can shed the  outer 2 layers, so that when we write a service, we only need to worry  about our *domain* and our *core* (ie what really matters ).

好的，现在我们已经解决了这个问题，让我们看看如何去掉外面的 2 层，这样当我们编写服务时，我们只需要担心我们的 *domain* 和我们的 *core*（即什么才是真正重要的） ）。

## 2. Microservices in Kubernetes

## 2. Kubernetes 中的微服务

So today, we typically host our microservices in Kubernetes, something like this:

所以今天，我们通常在 Kubernetes 中托管我们的微服务，如下所示：

![Microservices in Kubernetes](https://awesome.red-badger.com/stuartharris/wasmcloud/microservices.svg)

If each microservice talks to its own database, say in a cloud hosted service such as Azure CosmosDB, then each would include the same  libraries and similar glue code in order to talk to the DB. Even worse,  if each service is written in a different language, then we would be  including (and maintaining) different libraries and glue code for each  language.

如果每个微服务都与自己的数据库通信，比如在 Azure CosmosDB 等云托管服务中，那么每个微服务都将包含相同的库和类似的胶水代码，以便与数据库通信。更糟糕的是，如果每个服务都是用不同的语言编写的，那么我们将为每种语言包含（并维护）不同的库和粘合代码。

This problem is addressed today, for networking-related concerns, by a Service Mesh such as [Istio](https://istio.io/) or [Linkerd](https://linkerd.io/). These products abstract away traffic, security, policy and  instrumentation into a sidecar container in each pod. This helps a lot  because we now no longer need to implement this functionality in each  service (and in each service’s language).

今天，对于与网络相关的问题，这个问题已通过服务网格解决，例如 [Istio](https://istio.io/) 或 [Linkerd](https://linkerd.io/)。这些产品将流量、安全性、策略和仪器抽象到每个 Pod 中的 sidecar 容器中。这很有帮助，因为我们现在不再需要在每个服务（以及每个服务的语言)中实现此功能。

![Microservices with Service Mesh](https://awesome.red-badger.com/stuartharris/wasmcloud/servicemesh.svg)

But, and this is where the fun starts, we can also apply the same  logic to abstracting away other application concerns such as those in  the outer 2 layers of our onion. 

但是，这就是有趣的开始，我们也可以应用相同的逻辑来抽象出其他应用程序关注点，例如洋葱的外 2 层中的应用程序关注点。

Amazingly, there is an open source product available today that does just this! It’s called [Dapr](https://dapr.io/) (Distributed Application Runtime). It’s from the Microsoft stable and  is currently approaching its 1.0 release (v1.0.0-rc.2). Although it’s  only a year old, it has a very active community and has come a long way  already with many community-built components that interface with a wide  variety of popular cloud products.

令人惊讶的是，今天有一种开源产品可以做到这一点！它被称为 [Dapr](https://dapr.io/)（分布式应用程序运行时）。它来自 Microsoft 稳定版，目前正在接近其 1.0 版本（v1.0.0-rc.2)。虽然它只有一年的历史，但它拥有一个非常活跃的社区，并且已经有了许多社区构建的组件，这些组件与各种流行的云产品接口。

Dapr abstracts away IO-related concerns (i.e. those in our *infra* and *api* layers) and adds distributed application capabilities. If you use Dapr in Kubernetes, it is also implemented as a sidecar:

Dapr 抽象了与 IO 相关的问题（即我们的 *infra* 和 *api* 层中的问题）并添加了分布式应用程序功能。如果你在 Kubernetes 中使用 Dapr，它也是作为 sidecar 实现的：

![Microservices with Dapr](https://awesome.red-badger.com/stuartharris/wasmcloud/dapr.svg)

In fact, we can use Dapr and a Service Mesh together, ending up with 2 sidecars and our service (with no networking or IO concerns) in each  pod:

事实上，我们可以同时使用 Dapr 和 Service Mesh，最终在每个 pod 中得到 2 个 sidecar 和我们的服务（没有网络或 IO 问题）：

![Microservices with Service Mesh and Dapr](https://awesome.red-badger.com/stuartharris/wasmcloud/servicemesh_and_dapr.svg)

Now we’re getting somewhere! Our service becomes business logic and  nothing else! This is incredibly important! Now, when we look at the  source code for our service, we can see the wood — because all the  non-functional, non-core, non-business-logic, dull, repetitive,  boilerplate code is no longer there.

现在我们正在到达某个地方！我们的服务变成了业务逻辑，没有别的！这非常重要！现在，当我们查看服务的源代码时，我们可以看到木头——因为所有非功能性、非核心、非业务逻辑、枯燥、重复的样板代码都不再存在。

What’s more, our service is now much more portable. It can literally  run anywhere, because how it connects to the world around it is the  responsibility of Dapr, and is configured declaratively (in Yaml files,  just like a service mesh). If you wanted to move the service from Azure  to AWS, or even to the edge, you could, without *any* code changes.

更重要的是，我们的服务现在更加便携。它实际上可以在任何地方运行，因为它如何连接到周围的世界是 Dapr 的责任，并且是声明式配置的（在 Yaml 文件中，就像服务网格一样）。如果您想将服务从 Azure 迁移到 AWS，甚至迁移到边缘，则无需*任何*代码更改即可。

## 3. The Actor model

## 3. Actor 模型

Once the outer layers have been shed, we’re left with a much smaller  service, that concerns itself only with doing a job. It’s beginning to  look a bit like an Actor. Incidentally, Dapr has support for Virtual  Actors as well as the services that we already described. This gives us a flexible deployment model for our core logic.

一旦外层脱落，我们就剩下一个小得多的服务，它只关心做一份工作。它开始看起来有点像演员。顺便说一下，Dapr 支持 Virtual Actors 以及我们已经描述过的服务。这为我们的核心逻辑提供了灵活的部署模型。

So what is the [Actor Model](https://en.wikipedia.org/wiki/Actor_model)? Briefly, it’s an architectural pattern that allows small pieces of  business logic to run (and maintain own state) by receiving and sending  messages. Actors are inherently concurrent because they process messages in series. They can only process messages, send messages to other  actors, create other actors, and determine how to handle the next  message (e.g. by keeping state). The canonical example is an actor that  represents your bank account. When you send it a withdrawal message, it  deducts from your balance. The way that the next message is handled will depend on the new balance.

那么什么是 [Actor 模型](https://en.wikipedia.org/wiki/Actor_model)？简而言之，它是一种架构模式，它允许通过接收和发送消息来运行（并维护自己的状态）小块业务逻辑。 Actor 本质上是并发的，因为它们按顺序处理消息。它们只能处理消息，向其他参与者发送消息，创建其他参与者，并确定如何处理下一条消息（例如，通过保持状态)。规范示例是代表您的银行帐户的参与者。当您向其发送提款消息时，它会从您的余额中扣除。处理下一条消息的方式将取决于新的余额。

[Erlang OTP](https://en.wikipedia.org/wiki/Erlang_(programming_language)) is probably the most famous example of an actor (or “process”) runtime  organised in supervisor trees. Processes are allowed to crash and errors propagate up the tree. It turns out that this pattern is reliable,  safe, and massively concurrent. Which is why it’s been so good, for so  long, in telecoms applications.

[Erlang OTP](https://en.wikipedia.org/wiki/Erlang_(programming_language))可能是在主管树中组织的actor（或“进程”)运行时最著名的例子。允许进程崩溃并且错误沿树向上传播。事实证明，这种模式是可靠、安全和大规模并发的。这就是为什么它在电信应用中如此出色、如此之久的原因。

The Dapr Virtual Actor building block can place actors on suitable  nodes, hydrating and dehydrating them (with their state) as required. There may be thousands of actors running (or memoised) at any one time.

Dapr 虚拟 Actor 构建块可以将 Actor 放置在合适的节点上，根据需要对它们（及其状态）进行补水和脱水。任何时候都可能有成千上万的演员在运行（或记住）。

Depending on what type of application we are building, we can run our logic as a service behind the Dapr sidecar, or as actors supervised by  the runtime. Or both. Either way, we have written the logic in the  language that is most appropriate for the job, and we haven’t had to  pollute our code with concerns about how it talks with the outside  world.

根据我们正在构建的应用程序类型，我们可以将我们的逻辑作为 Dapr sidecar 背后的服务运行，或者作为由运行时监督的角色运行。或两者。无论哪种方式，我们都用最适合工作的语言编写了逻辑，而且我们不必担心代码与外部世界的对话方式来污染我们的代码。

## 4. WebAssembly

## 4. WebAssembly

There’s one thing that can make our code even more portable: [WebAssembly](https://webassembly.org/)(Wasm). In December 2019 WebAssembly [became a W3C recommendation](https://www.w3.org/2019/12/pressrelease-wasm-rec.html.en) and is now the fourth language of the Web (alongside HTML, CSS and JS). 

有一件事可以使我们的代码更具可移植性：[WebAssembly](https://webassembly.org/)(Wasm)。 2019 年 12 月，WebAssembly [成为 W3C 推荐](https://www.w3.org/2019/12/pressrelease-wasm-rec.html.en)，现在是 Web 的第四种语言（与 HTML、CSS 和JS)。

Wasm is great in the browser; all modern browsers support it. But, in my opinion, it becomes really useful on the server, where there are  already [tens of different Wasm runtimes](https://github.com/appcypher/awesome-wasm-runtimes) that we can choose from, each with different characteristics (such as  just-in-time vs ahead-of-time compilation, etc). Arguably, the most  popular runtime is [Wasmtime](https://github.com/bytecodealliance/wasmtime), which implements a specification called [WebAssembly System Interface (WASI)](https://wasi.dev/) from the [ Bytecode Alliance](https://bytecodealliance.org/) that is specifically designed to run untrusted code safely in a Wasm sandbox on the server.

Wasm 在浏览器中很棒；所有现代浏览器都支持它。但是，在我看来，它在服务器上变得非常有用，那里已经有 [数十种不同的 Wasm 运行时](https://github.com/appcypher/awesome-wasm-runtimes)可供我们选择，每个都有不同的特性（例如即时编译与提前编译等)。可以说，最流行的运行时是 [Wasmtime](https://github.com/bytecodealliance/wasmtime)，它实现了一个名为 [WebAssembly System Interface (WASI)](https://wasi.dev/) 的规范，来自 [ Bytecode Alliance](https://bytecodealliance.org/) 专门设计用于在服务器上的 Wasm 沙箱中安全地运行不受信任的代码。

This safety is important — modern microservices are assembled from  multiple open source libraries and frameworks and it seems irresponsible to run this code as though we trust it. Yet that’s what we do, all the  time. Docker containers everywhere could be hosting malicious code that  is just hanging around, hiding, and waiting for someone with a black hat to trigger its exploit.

这种安全性很重要——现代微服务是由多个开源库和框架组装而成的，像我们信任它一样运行这些代码似乎是不负责任的。然而，这就是我们一直在做的事情。无处不在的 Docker 容器可能托管着恶意代码，这些代码只是徘徊、隐藏并等待戴着黑帽子的人触发其漏洞利用。

We should be building security *in*, rather than building it *on*. Today, we wrap our containerised microservices with a CyberSecurity  industry, instead of making it impossible, in the first instance, for  code to do anything that we haven’t specifically said it can do. Shifting security left, like this, is called [DevSecOps](https://www.redhat.com/en/topics/devops/what-is-devsecops), which is more than just [DevOps](https://en.wikipedia.org/wiki/DevOps) — it's about designing security into our applications from the ground up.

我们应该*in* 构建安全性，而不是*on* 构建它。今天，我们用网络安全行业包装我们的容器化微服务，而不是首先让代码做任何我们没有明确表示可以做的事情。像这样向左移动安全被称为 [DevSecOps](https://www.redhat.com/en/topics/devops/what-is-devsecops)，这不仅仅是 [DevOps](https://en.wikipedia.org/wiki/DevOps)——这是关于从头开始为我们的应用程序设计安全性。

## 5. WasmCloud

## 5. WasmCloud

So to recap, we want to reduce our code to just the functional business logic — the *core* — without having to worry about how it talks with the outside world. We want to run this, safely, in a sandbox — e.g. Wasmtime. We need a  reliable, secure orchestration framework to supervise and manage our  small code components (or Actors). We want location transparency and  full portability, so we don’t have to worry about whether we’re  deploying to the cloud, on-prem, or at the edge. And we want to design  all this security in, from the ground up.

因此，回顾一下，我们希望将我们的代码简化为功能性业务逻辑——*核心*——而不必担心它如何与外部世界对话。我们希望在沙箱中安全地运行它——例如是时候了。我们需要一个可靠、安全的编排框架来监督和管理我们的小代码组件（或 Actor）。我们想要位置透明性和完全可移植性，因此我们不必担心我们是部署到云、本地还是边缘。我们希望从头开始设计所有这些安全性。

I think [WasmCloud](https://wascc.dev/) (in the process of being renamed from [waSCC](https://wascc.dev/) – the WebAssembly Secure Capabilities Connector) is heading in this  direction. Something like this will be what comes next, after  Kubernetes.

我认为 [WasmCloud](https://wascc.dev/)（正在从[waSCC](https://wascc.dev/) – WebAssembly 安全功能连接器重命名的过程中)正朝着这个方向发展。在 Kubernetes 之后，接下来会出现这样的事情。

![A waSCC host](https://awesome.red-badger.com/stuartharris/wasmcloud/wascc-host.svg)

A WasmCloud (or waSCC) host securely connects cryptographically  signed Wasm actors to the declared capabilities of well-known providers. The actors are placed and managed by the host nodes, which can  self-form a Lattice when connected as [NATs](https://nats.io/) leaf nodes. Actors can be placed near to suitable providers or  distributed across heterogeneous networks that span on-prem, Cloud,  edge, IoT, embedded devices, etc.

WasmCloud（或 waSCC）主机将加密签名的 Wasm 参与者安全地连接到知名提供商的声明功能。 Actor 由主机节点放置和管理，当作为 [NATs](https://nats.io/) 叶节点连接时，主机节点可以自行形成一个 Lattice。参与者可以放置在合适的供应商附近，也可以分布在跨本地、云、边缘、物联网、嵌入式设备等的异构网络中。

![A waSCC Lattice](https://awesome.red-badger.com/stuartharris/wasmcloud/wascc-lattice.svg)

WasmCloud is not the only thing out there that is following this path. [Lunatic](https://github.com/lunatic-lang/lunatic) is also interesting (and also written in Rust). Go check it out.

WasmCloud 并不是唯一遵循这条道路的东西。 [Lunatic](https://github.com/lunatic-lang/lunatic) 也很有趣（也是用 Rust 编写的)。去看看吧。

It may be a while before the Wasm actor model becomes viable for  production applications, but it’s definitely one to watch. Personally, I can’t wait for the time when we can literally write distributed  applications by just concentrating on the real work we need to do. In  the meantime, we can get going now by using [Dapr](https://dapr.io/), which is good for production workloads today.

Wasm actor 模型在生产应用程序中变得可行可能还需要一段时间，但它绝对值得一看。就我个人而言，我迫不及待地等待我们可以通过专注于我们需要做的实际工作来真正地编写分布式应用程序。同时，我们现在可以开始使用 [Dapr](https://dapr.io/)，这对当今的生产工作负载很有用。

By the way, [I made a small side-project](https://github.com/redbadger/rpi-wascc-demo) that creates a WasmCloud Lattice across my MacBook and a Raspberry Pi,  and wrote a Wasm actor that controls an OLED display via a native  provider running on the Pi host. WasmCloud is moving fast and I will try to keep this up to date, but if you fancy it, have a play, and raise an issue if you want to chat. 

顺便说一句，[我做了一个小的副项目](https://github.com/redbadger/rpi-wascc-demo)，它在我的 MacBook 和 Raspberry Pi 上创建了一个 WasmCloud Lattice，并编写了一个控制通过在 Pi 主机上运行的本机提供程序的 OLED 显示器。 WasmCloud 发展很快，我会尽量保持最新状态，但如果你喜欢它，可以玩一玩，如果你想聊天，可以提出问题。


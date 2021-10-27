# What is Developer Experience and why should we care?

[#devrel](http://dev.to/t/devrel) [#developerexperience](http://dev.to/t/developerexperience) [#engineering](http://dev.to/t/engineering) [#discuss](http://dev.to/t/discuss)

Information Technology (IT) is undoubtedly one of the most important industries today and one that is ever-growing. Every company is becoming an IT company these days. From Taxis to food delivery to banking, every industry is dominated by companies that are IT companies first and domain second. With that growth, the demand for software and tools used by other developers also grows.

As an industry it took us some time to realize the importance of user experience (UX), you will understand what I'm talking about if you have tried using the internet or any software before the 2000s, but fortunately, we took notice, and today there are entire departments dedicated to user research and UX in software development.

User experience matters for any software, but if your primary consumers are developers, then there is something else that matters more than UX. Its developer experience (DX) as developers make IT possible.

## What is Developer Experience?

It is the overall feeling that a developer gets when using a technical product in her/his development workflow. It is akin to UX but from a developer's perspective. Let's take an example for the sake of non-developer folks out there.

[![What is Developer experience](https://res.cloudinary.com/practicaldev/image/fetch/s--13WaG65m--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://i.imgur.com/EB23Cv3.png)](https://res.cloudinary.com/practicaldev/image/fetch/s--13WaG65m--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://i.imgur.com/EB23Cv3.png)

Let's say you are building a cool and fancy product that lets developers add an image gallery to their applications, as part of the product you provide:

- an API, to get random optimized images from your service;
- a JS SDK to add this easily on WordPress sites;
- a web application to manage the images.

Now for a developer who would use this product the DX is going to be the sum of:

- The experience using the API, like:

  - how easy it was to onboard
  - how easy it is to integrate into their app?
  - simplicity of the API and the resulting code
  - learning curve of the API
  - how informative are error messages?
  - does it follow known standards and structure?
  - how easy it is to debug.
- Performance of the API.
- Documentation of the API and product.
- The ease of use of the JS SDK if they are using it.
- The user experience of the web app when managing images.

And this experience determines if the developer is going to consider using your product for the next project or not.

We could also loosely measure DX as the inverse of the amount of frustration a developer has when using a product. Sometimes these frustrations could be outside of your control but regardless it is going to affect the DX as after all we are human and our emotions and feelings bound us.

> Of course other factors like features, pricing, sales, marketing, and so on will get you through the door, but good DX is what will keep you in the room.

Of course, there are more things that can be done in the above example to make DX even better, like providing example applications, video tutorials, blog posts showing various use cases, CLI tools for debugging, and so on.

## Why should you care about Developer Experience?

You should care about the DX of your product for the same reasons you care about the UX and some more. If you are a developer, just think of what kind of experience you would want when using a similar product. A good DX also shows empathy on your part for your primary users.

Developers are a very opinionated bunch. We love our opinions and we love to defend our favorite language, technology, and tools. Heck, we are even ready to go to war over something as trivial as [tabs vs spaces](https://www.reddit.com/r/programming/comments/p1j1c/tabs_vs_spaces_vs_both/).

So if your product has great DX the developers using it will love it and will evangelize and defend your product to the death. You might even gain a community of ardent supporters for your product that no amount of marketing can get you.

But if the DX is bad, they are going to badmouth your product. If you are a developer, I think you know what I'm talking about and I'm pretty sure you have done this a lot.

Another reason for focusing on DX is that it will make marketing and sales much easier as you have less friction with your end-users and fewer things that you need to convince them about.

The recent surge in the interest for developer advocacy has also helped to bring to the limelight the importance of DX and without good DX there is not much you can do about developer advocacy and evangelism.

## How can we have a great Developer Experience?

A product with great DX helps a developer to get up and running quickly and reach her/his goal with minimal frustrations.

So let's see what are some of the common things that could help make great DX. Please note that this is not an exhaustive list and there are many more things that could help, depending on the specific product/use case.

### If you are building APIs

- Stick to known and highly adopted standards and conventions. Do not reinvent the wheel or try to come up with new fancy conventions as it will increase the learning curve and make it harder for developers to onboard.
- Provide good error handling. Adopt something like[RFC 7807](https://datatracker.ietf.org/doc/html/rfc7807) to provide a consistent and easy-to-use error handling for your API. Errors are unavoidable but making developers scratch their heads when it happens is avoidable to an extend.
- Provide consistent and easy-to-use documentation. Provide an[OpenAPI](https://www.openapis.org/) or [RAML](http://raml.org/) file that describes your API and the endpoints. Also, provide easy-to-use interactive online API documentation like [Swagger](https://swagger.io/) if possible.
- Provide SDKs and libraries for most popular languages and frameworks.

### If you are building development tools/products

- Good UX, which is tailored for developers. Like providing high contrast options. Good keyboard navigation, consistent with an industry standard.
- Customizability, did I mention developers are an opinionated bunch?
- Easy to install on different platforms. Support different OSes. Provide multiple installation methods, especially, support the most popular installation method/package manager used by the ecosystem relevant to your product.
- Easy-to-use and well documented.

### If you are building SDKs/Libraries/Frameworks

- Consistency and following industry standards and conventions. Avoid reinventing the wheel without a solid reason.
- Play well with other SDKs/Frameworks in the ecosystem.
- Provide inline code documentation wherever possible. Developers appreciate not requiring to leave their IDE.
- Provide great documentation, examples, tutorials, interactive learning, and so on.
- Easy to use error reporting system.
- Open source when possible.
- Avoid gatekeeping as much as possible. No registration-only webinars and no white papers that ask for my mother's favorite color.

### General

> In general, always ask this question, how can my product make a developer's day better?

- Make sure the product does exactly what it claims to do.
- Make sure the product is reliable and performant otherwise good DX will not cut it.
- Focus on easy-to-use self-service rather than traditional support channels. Developers hate bureaucratic processes. Provide transparent support channels for those edge cases but most developers would prefer self-service if available rather than talking to a support person.
- Simple onboarding. Avoid having to go through sales or other channels just to get started with development or for trying out your product.
- Provide modern tooling or make your stuff compatible with modern tooling. Nobody wants to download and set up something from the 80s to run your software.
- Make trying out your product as easy as possible.
- Make documentation easy to find and navigate. Avoid having to jump through hoops to find documentation.
- Make developer resources easy to find. Make a developer-focused section on your company website as a landing point for developer resources.
- Avoid jargon-rich fancy marketing as most of us can see right through it. Keep it simple and to the point.
- Build developer advocacy rather than technical evangelism.
- Build a community around your product.

## Conclusion

We are slowly transitioning to an era where the importance of developers is being recognized and their influence on decision making is no longer something companies can take for granted.

This is very clear from the fact that more and more companies are building developer relationship teams and hiring developer advocates rather than marketing evangelists.

In this crowded space, being developer-focused used to be a differentiator but things are going towards the same situation that happened with UX where it became a must have rather than a good to have.

The same will happen for DX as well and if you are building a product that is going to be used by developers, then you should start caring. Building developer experience and developer advocacy don't happen overnight.

* * *

If you like this article, please leave a like or a comment.

# OpenTelemetry Steps up to Manage the Mayhem of Microservices

17 Apr 2020

![](https://cdn.thenewstack.io/media/2020/04/89962ad8-open-telemetry.jpg)

Work continues to make [OpenTelemetry](https://opentelemetry.io/about/) the standard set of vendor-neutral specifications and associated tools for capturing cloud native operational data.

Bringing the work close to production-readiness the project has released a collector, as well as software development kits (SDKs) to support a number of different languages.

For end-users, the new release candidate offers end-users “confidence that OpenTelemetry is safe to evaluate because it is now at a place where it’s nearing feature completion, runtime stability, and API stability,” wrote [Liz Fong-Jones](https://twitter.com/lizthegrey), principal developer advocate for observability software provider and project contributor [Honeycomb](https://www.honeycomb.io/), in an e-mail.

She notes that the release signals to monitoring software and services providers that it is time to write their own exporters for the project, and it provides “a smoother and less friction path for users to get critical telemetry data into their app.”

The work sets the stage for making “high-quality telemetry a built-in feature for any cloud native software,” said [Ben Sigelman](https://www.linkedin.com/in/bensigelman/), CEO and co-founder of cloud native monitoring provider [Lightstep](https://lightstep.com/), one of the contributors to the project.

The work is particularly pertinent as concerns grow that running microservices-based architectures in large-scale environments [may be too unwieldy](https://thenewstack.io/kelsey-hightower-and-ben-sigelman-debate-microservices-vs-monoliths/), given the difficulties in debugging across multiple components. “The challenge with telemetry in cloud native is that there’s a lot of it, especially for tracing data,” Sigelman said. A [sandbox project](https://landscape.cncf.io/selected=open-telemetry) of the Cloud Native Computing Foundation, OpenTelemetry would be a key enabler in making microservices manageable (OpenTelemetry itself was a 2019 merger between two overlapping projects, [OpenTracing and OpenCensus](https://thenewstack.io/opentracing-opencensus-merge-into-a-single-new-project-opentelemetry/)).

An open-source collaborative project across 82 different companies, OpenTelemetry is building a set of libraries, agents, and other components to aid in the observation, management and debugging of microservices and distributed applications. Telemetry is built on the three pillars of metrics, logs, and tracing, providing the context needed to trace an individual transaction as it flows through multiple components

The goal with OpenTelemetry is not to provide a platform for observability, but rather to provide a standard substrate to collect and convey operational data so it can be used in monitoring and observational platforms, either of the open source or commercial variety.

Historically, when an enterprise would purchase a package for systems monitoring, all the agents that would be attached to the resources would be specific to that provider’s implementation. If a customer wanted to change out, the applications and infrastructure would have to be entirely re-instrumented, Sigelman explained. By using the OpenTelemetry, users could instrument their systems once and pick the best and visualization and analysis products for their workloads, and not worry about lock-in.

In addition to Honeycomb and Lightstep, some of the largest vendors in the monitoring field, as well as the largest end-users are participating, including Google, Microsoft, Splunk, Postmates, and Uber.

### New Collectors and SDK

The new collector is crucial, explained Honeycomb’s Fong-Jones, in that it narrows the minimum scope of what vendors must support in order to ingest telemetry. “It enables vendors to write one exporter in Golang and have _any_ [Jaeger](https://www.jaegertracing.io/), [Zipkin](https://zipkin.io/), or [OTLP](https://github.com/open-telemetry/oteps/blob/master/text/0035-opentelemetry-protocol.md) (OpenTelemetry’s new telemetry protocol) producer transmit to that vendor via the collector.”

![](https://cdn.thenewstack.io/media/2020/04/0d6104d3-opentelemetry.jpg)

OpenTelemetry Collector is the ‘swiss army knife’ for critical telemetry data that can be used downstream in a number of tools including Honeycomb for production system Observability and improved resilience.

The Honeycomb software itself, which originally supported both OpenCensus and OpenTracing, can now be streamlined by using only OpenTelemetry to format and deliver into Honeycomb.

The SDKs are also vital in that they allow developers to write instrumentation in their own favored languages, while still adhering to the OpenTelemetry API specifications, Fong-Jones explained. Languages currently supported include Erlang, GoLang, Java, JavaScript,  and Python.

[Open Keynote: (Open)Telemetry Makes Observability Simple – Sarah Novotny & Liz Fong-Jones on YouTube.](https://www.youtube.com/watch?v=W_8MHdtrgZE&list=WL&index=2&t=0s)

[Open OpenTelemetry: Overview & Backwards Compatibility of OpenTracing + OpenCensus – Steve Flanders on YouTube.](https://www.youtube.com/watch?v=UdRqts403G4&list=WL&index=3&t=12s)

[Open Beyond Getting Started: Using OpenTelemetry to Its Full Potential – Sergey Kanzhelev & Morgan McLean on YouTube.](https://www.youtube.com/watch?v=FlghuHDlQdM&list=WL&index=4&t=51s)

Lightstep is a sponsor of The New Stack.

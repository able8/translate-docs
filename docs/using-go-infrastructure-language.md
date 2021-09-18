# Using Go as the Infrastructure Language at NodeSource

in [NodeSource](http://nodesource.com/blog/category/nodesource)
on Mar 08 2019

- [Cloud](http://nodesource.com/blog/tag/cloud "Cloud")
- [AWS](http://nodesource.com/blog/tag/aws "AWS")
- [Azure](http://nodesource.com/blog/tag/azure "Azure")
- [Golang](http://nodesource.com/blog/tag/golang "Golang")

NodeSource is a cloud-first company, with resources on AWS, Google Cloud Platform, and Azure. Traditionally Python was a common choice of programming language as it could be used by infrastructure teams. But as Google developed Go to solve problems that Python could not, Go quickly gained new users over the course of just a few years. Go has become the more popular option, and many tools that were once a staple for infrastructure in Python now have a Go version. Some of these tools, like graphite, were originally created in Python but have been recreated for use with Go due to their popularity and growing demand.

Every organization has its own set of needs and constraints, so while I’m a big fan of Go, I recognize that it isn’t the right choice for everyone. For NodeSource, some features in Go were of particular interest and use to us, and made Go the better choice for us as we built infrastructure, including:

- [Static linking](https://en.wikipedia.org/wiki/Static_library): Static linking makes it easy to deploy what I build. All Go binaries are statically linked, and copying them to the target platform allows for easy execution without managing dependencies. In our case, I don’t even need to install the runtime when creating a new instance, since that too is included in the binary. This is a much simpler task than it would be if I used Python, which requires that the runtime as well as dependencies (requirements.txt) be installed before code can run. With Go, if you can build the binary; you can deploy the binary.
- [Goroutines](https://tour.golang.org/concurrency/1): We live in a world with multiple CPUs (or vCPUs if you prefer). To maximize performance and computing power, we need threading or multiprocessing. However, threading can be painful (insert your favorite pthread or Java threading jokes here). So with Go, threading is as simple as `go myfunc()`; it has never been easier to thread with a language such as Go.

The above two benefits are more or less "core" to the language itself; but, with any language comes the need for third-party modules.

A few additional modules which I consider to be very important for writing Go programs include:

- [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) make it very easy to create software that supports loading configuration files and having command-line switches. Any moderately-important software should have at minimum CLI switches, and any really important software should probably load its configuration from a JSON or YAML file. Both of these modules allow you to do that.
- [logrus](https://github.com/Sirupsen/logrus) is another important module for logging. It's not required, but makes software feel more polished and ready for production. This is definitely preferable to just using `fmt.Printf` statements.
- [jsonq](https://github.com/jmoiron/jsonq) this library makes it easier to read JSON data, even though it is already easy to marshal/unmarshal JSON-encoded data in Go.

Finally, because our team uses Go as our preferred cloud infrastructure language, I would also suggest the following critical modules that are used to interact with our cloud environments:

- [aws-sdk-go](https://github.com/aws/aws-sdk-go) is the official AWS SDK for Go. I have used it extensively for interacting with AWS services including (but not limited to) S3, ECS, and EC2.
- The[Slack API in Go](https://github.com/nlopes/slack) allows me to post messages to slack from anything written in Go—as part of a Slack-friendly team, this has been critical for me.
- The[Docker SDK for Go](https://docs.docker.com/develop/sdk/#go-sdk) is the module to write Go code that can interact with a Docker host; critical if you’re planning to write any of your own tooling to pull data from your containerized environment.
- The official[Go client library for Prometheus](https://prometheus.io/docs/guides/go-application/) allows me to define and expose Prometheus metrics from any Go binary I’ve written.

### One More Thing....

If you’re an infrastructure professional and interested in working with NodeSource, we’re currently looking for an awesome [Site Reliability Engineer](https://nodesource.bamboohr.com/jobs/view.php?id=31) to join our team. Please check out our [careers page](https://nodesource.com/careers) for other open positions and to apply.

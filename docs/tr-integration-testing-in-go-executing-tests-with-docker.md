# Integration Testing in Go: Part I - Executing Tests with Docker

# Go 中的集成测试：第一部分 - 使用 Docker 执行测试

George ShawMarch 18, 2019

乔治·肖 2019 年 3 月 18 日

### Introduction

###  介绍

> *“Testing leads to failure, and failure leads to understanding.” - \*Burt Rutan**

> *“测试导致失败，失败导致理解。” - \*伯特·鲁坦**

Burt Rutan is an aerospace engineer who designed Voyager, the first plane to fly around the world without stopping or refueling. Although Rutan was  not a software engineer, his words speak volumes to the importance of  testing, even testing software. Testing software in all forms is  extremely important, whether it be unit, integration, system, or  acceptance testing. However, depending on the project, one form of  testing can be more valuable than the others. In other words, sometimes  one form of testing can lead to better understanding about the health  and integrity of the software than the other forms.

伯特·鲁坦 (Burt Rutan) 是一位航空航天工程师，他设计了航海者号 (Voyager)，这是第一架无需停车或加油即可环球飞行的飞机。虽然 Rutan 不是软件工程师，但他的话充分说明了测试，甚至是测试软件的重要性。各种形式的软件测试都非常重要，无论是单元测试、集成测试、系统测试还是验收测试。但是，根据项目的不同，一种形式的测试可能比其他形式更有价值。换句话说，有时一种形式的测试可以比其他形式更好地理解软件的健康和完整性。

When  developing a web service, I believe a strong set of integration tests  can provide a better understanding of the service than other types of  tests. Integration tests are a form of software testing that tests the  interaction of your code against the dependencies your application is  leveraging, such as databases and messaging systems. Without integration tests, it’s difficult to trust the end-to-end operation of a web  service. I believe this is true because individual units of code being  tested in a web service rarely provide the same level of insight as  integration tests.

在开发 Web 服务时，我相信一组强大的集成测试可以比其他类型的测试更好地理解服务。集成测试是一种软件测试形式，用于测试您的代码与应用程序所利用的依赖项（例如数据库和消息传递系统）之间的交互。如果没有集成测试，就很难信任 Web 服务的端到端操作。我相信这是真的，因为在 Web 服务中测试的单个代码单元很少提供与集成测试相同级别的洞察力。

This is the first entry of a two part series  about integration testing in Go. The ideas, code and processes shared in this series aim to be easily extendable to the web service projects you are working on. In this post, I will show you how to setup your web  service projects to use Docker and Docker Compose to run your Go tests  and dependencies in a restrictive computing environment that doesn’t  have Go pre-installed.

这是关于 Go 集成测试的两部分系列的第一篇。本系列中共享的想法、代码和流程旨在轻松扩展到您正在处理的 Web 服务项目。在这篇文章中，我将向您展示如何设置您的 Web 服务项目以使用 Docker 和 Docker Compose 在没有预安装 Go 的限制性计算环境中运行您的 Go 测试和依赖项。

### Why use Docker and Docker Compose

### 为什么使用 Docker 和 Docker Compose

What attracts many developers to Docker is how you’re able to load  applications on your host machine without the burden of having to  install and manage them manually. This means you can load complex  software including, but not limited to, databases (e.g. Postgres),  messaging systems (e.g. Kafka) and monitoring systems (e.g. Prometheus). All of this is done by downloading a set of images that represent the  application and all of its dependencies.

Docker 吸引许多开发人员的原因是您能够在主机上加载应用程序，而无需手动安装和管理它们。这意味着您可以加载复杂的软件，包括但不限于数据库（例如 Postgres）、消息系统（例如 Kafka）和监控系统（例如 Prometheus）。所有这一切都是通过下载一组代表应用程序及其所有依赖项的图像来完成的。

*Note: For more  information on containers, Docker has a webpage devoted to the  definition of a container and highlights the differences and  similarities between a container and a virtual machine that can be found [here](https://www.docker.com/resources/what-container).*

*注意：有关容器的更多信息，Docker 有一个专门介绍容器定义的网页，并强调了容器和虚拟机之间的异同，可以在 [此处](https://www.docker.com)找到/resources/what-container).*

Docker Compose is an orchestration tool that aids in building, running, and  networking a group of containers together inside of a single sandbox. With a single command, `docker-compose up`, you can make your Docker Compose file come to life. All the services defined in the  compose file will become containers running as a group within their own  networked sandbox and run as configured. This is in contrast to manually building, running, and networking each of your containers in order to  allow them to run together, communicate with each other, and persist  data.

Docker Compose 是一种编排工具，可帮助在单个沙箱内构建、运行和联网一组容器。使用一个命令 `docker-compose up`，你可以让你的 Docker Compose 文件变得生动起来。撰写文件中定义的所有服务将成为在其自己的网络沙箱中作为一个组运行的容器，并按配置运行。这与手动构建、运行和联网每个容器以允许它们一起运行、相互通信和保留数据形成对比。

Since Docker Compose allows you to group different  applications together and run them within a single networked sandbox,  you can start and stop an entire suite of applications with a single  command. You can even handpick certain applications to run from the  group. This group of applications can be deployed as a single unit and  be built and tested by a CI (continuous integration) environment. Docker Compose ultimately helps ensure that your application is consistent  across any environment it’s tested and deployed on.

由于 Docker Compose 允许您将不同的应用程序组合在一起并在单个网络沙箱中运行它们，因此您可以使用单个命令启动和停止整个应用程序套件。您甚至可以从组中挑选要运行的某些应用程序。这组应用程序可以作为一个单元部署，并由 CI（持续集成）环境构建和测试。 Docker Compose 最终有助于确保您的应用程序在其测试和部署的任何环境中保持一致。

*Note: For  more information on Docker Compose, visit the webpage for the overview  of Docker Compose on the official Docker website [here](https://docs.docker.com/compose/overview/).* 

*注意：有关 Docker Compose 的更多信息，请访问 Docker 官方网站 [此处](https://docs.docker.com/compose/overview/) 上的 Docker Compose 概述网页。*

Another big benefit of Docker and Docker Compose is that they help facilitate  an easier transition when bringing new developers into a project. Instead of having complex documentation regarding how a development  environment is installed and managed, new developers just need to  execute a few Docker and Docker Compose commands to get started. The  Docker CLI takes care of downloading required images if they don’t  currently exist on the host machine when an application is started.

Docker 和 Docker Compose 的另一大好处是，它们有助于在将新开发人员带入项目时更轻松地过渡。新开发人员无需编写关于如何安装和管理开发环境的复杂文档，只需执行一些 Docker 和 Docker Compose 命令即可开始。如果应用程序启动时主机上当前不存在所需的图像，Docker CLI 会负责下载所需的图像。

### Using Docker and Docker Compose to Run Tests

### 使用 Docker 和 Docker Compose 运行测试

The web service application referenced throughout this series exposes a  simple CRUD based REST API with a Postgres database. The project uses  Docker to run a Postgres database for both production and testing. The  tests for this application need to be able to run in a local development environment that already has Go installed and a restricted environment  where Go does not exist.

本系列中引用的 Web 服务应用程序公开了一个简单的基于 CRUD 的 REST API 和一个 Postgres 数据库。该项目使用 Docker 来运行用于生产和测试的 Postgres 数据库。此应用程序的测试需要能够在已安装 Go 的本地开发环境和不存在 Go 的受限环境中运行。

The following Docker Compose file  supports the ability to run integration tests for the project in both  environments I mentioned above. In this section, I will break down the  configuration options that I chose and why I chose them.

以下 Docker Compose 文件支持在我上面提到的两种环境中为项目运行集成测试的能力。在本节中，我将分解我选择的配置选项以及我选择它们的原因。

*Listing 1*

*清单 1*

```yaml
version: '3'

networks:
  integration-tests-example-test:
    driver: bridge

services:
  listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
    depends_on:
      - db
    networks:
      - integration-tests-example-test
  db:
    image: postgres:11.1
    ports:
      - "5432:5432"
    expose:
      - "5432"
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: root
      POSTGRES_DB: testdb
    restart: on-failure
    networks:
      - integration-tests-example-test
```

In listing 1, you see the Docker Compose file that  defines the services for the project that are required to run the tests. This file has three main keys: `version`, `networks`, and `services`. The `version` key defines the version of Docker Compose you’re using. The `networks` key defines one or more network configurations that can be available to a given service. The `services` key defines the containers to be started and their configuration.

在清单 1 中，您会看到 Docker Compose 文件，该文件定义了运行测试所需的项目服务。该文件具有三个主要键：`version`、`networks` 和`services`。 `version` 键定义了您正在使用的 Docker Compose 版本。 `networks` 键定义了一个或多个可用于给定服务的网络配置。 `services` 键定义了要启动的容器及其配置。

*Listing 2*

*清单 2*

```yaml
networks:
  integration-tests-example:
    driver: bridge
```

By having your service definitions in one compose file  they are automatically placed within the same network by default and  therefore can communicate with each other. However, it’s a best practice to create a network for your services as opposed to using the default  network. The top-level `networks` configuration defines the name of the network and the driver it uses, the bridge driver in this case.

通过将您的服务定义放在一个组合文件中，默认情况下它们会自动放置在同一网络中，因此可以相互通信。但是，最佳做法是为您的服务创建网络而不是使用默认网络。顶级“networks”配置定义了网络的名称和它使用的驱动程序，在这种情况下是桥接驱动程序。

The bridge driver is the default driver provided by Docker which creates a  private internal network for containers to communicate within. The  services are told to use the created network within their service  definition configuration in the compose file.

桥接驱动程序是 Docker 提供的默认驱动程序，它创建了一个私有内部网络，供容器在内部进行通信。服务被告知在撰写文件中的服务定义配置中使用创建的网络。

*Listing 3*

*清单 3*

```yaml
services:
  listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
// ... omitted code…
  db:
// ... omitted code…
```

The `services` key has two immediate child keys, `listd_tests` and `db`. The `listd_tests` container defines its image by specifying a Dockerfile. The `context` key denotes that all host paths should be relative to the current working directory, as denoted by a `.`.

`services` 键有两个直接子键，`listd_tests` 和 `db`。 `listd_tests` 容器通过指定一个 Dockerfile 来定义它的镜像。 `context` 键表示所有主机路径都应该相对于当前工作目录，用 `.` 表示。

*Listing 4*

*清单 4*

```yaml
listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
    depends_on:
      - db
    volumes:
      - $PWD:/go/src/github.com/george-e-shaw-iv/integration-tests-example
```

The `depends_on` key tells the `listd_tests` service to wait to start until the `db` service has already started. In addition to asserting start order of services, this key will disallow the `listd_tests` service from being run independently of the `db` service. The `volumes` key tells compose to mount the current working directory, denoted by `$PWD` (**P**rint **W**orking **D**irectory), to `/go/src/github. com/george-e-shaw-iv/integration-tests-example` within the container, which is where the code will be located and tested.

`depends_on` 键告诉 `listd_tests` 服务等待启动，直到 `db` 服务已经启动。除了断言服务的启动顺序之外，该键还将禁止 `listd_tests` 服务独立于 `db` 服务运行。 `volumes` 键告诉 compose 将当前工作目录（由 `$PWD`（**P**rint **W**orking **D**irectory）表示）挂载到 `/go/src/github。 com/george-e-shaw-iv/integration-tests-example` 在容器内，这是代码将被定位和测试的地方。

*Listing 5*

*清单 5*

```yaml
listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
    depends_on:
      - db
    networks:
      - integration-tests-example-test
```

Finally, the service is given a network to communicate  on when running inside the sandbox. This was originally defined in the  top-level `networks` configuration key in listing 2.

最后，当在沙箱内运行时，服务被赋予一个网络来进行通信。这最初是在清单 2 中的顶级 `networks` 配置键中定义的。

*Listing 6*

*清单 6*

```yaml
db:
    image: postgres:11.1
```

The container in the next service definition, `db`, defines its image by [using a image hosted at Docker Hub](https://hub.docker.com/_/postgres), the `postgres:1.11` image. The Docker CLI is smart enough to know to look in the Docker Hub image repository if it can’t find the image on your local machine.

下一个服务定义中的容器 `db`，通过[使用托管在 Docker Hub 的镜像](https://hub.docker.com/_/postgres) 定义其镜像，即 `postgres:1.11` 镜像。 Docker CLI 足够聪明，它知道如果在您的本地机器上找不到镜像，就可以查看 Docker Hub 镜像存储库。

*Listing 7*

*清单 7*

```yaml
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
```

For security reasons, by default none of the container  ports are accessible from your host machine. This proves to be a problem when running integration tests locally, as the integrated service is  pretty worthless if it isn’t accessible. The `ports` key defines port mappings from your host machine to your container in the following format: `"HOST_PORT:CONTAINER_PORT"`. The preceding definition in listing 7 ensures that port 5432 on your machine is mapped to port 5432 on the `db` container, as that is the port that Postgres is running on within the container by default.

出于安全原因，默认情况下，您的主机无法访问任何容器端口。在本地运行集成测试时，这被证明是一个问题，因为如果无法访问集成服务，它就毫无价值。 `ports` 键定义了从主机到容器的端口映射，格式如下：`"HOST_PORT:CONTAINER_PORT"`。清单 7 中的上述定义确保您的机器上的端口 5432 映射到 `db` 容器上的端口 5432，因为默认情况下 Postgres 在容器内运行的端口。

*Listing 8*

*清单 8*

```yaml
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
    expose:
      - "5432"
```

In the same manner that container ports aren’t exposed  to the host machine by default, container ports are also not exposed to  containers running within the networked sandbox by default. This is true even if they are on the same network. In order to expose a port to  other containers running within the networked sandbox, the `expose` configuration key needs to be set.

与默认情况下容器端口不向主机公开的方式相同，默认情况下容器端口也不向在网络沙箱中运行的容器公开。即使它们在同一网络上也是如此。为了向网络沙箱中运行的其他容器公开端口，需要设置 `expose` 配置键。

*Note: In the case of the `postgres:1.11` image, port 5432 has already been exposed thanks to the person who  created the image. Since you don’t know if the image was created with  the port already exposed unless you look at an image’s Dockerfile, it is best to define the `expose` key, even if it’s redundant.*

*注意：在`postgres:1.11`镜像的情况下，由于创建镜像的人，端口5432已经暴露。由于除非查看镜像的 Dockerfile，否则您不知道镜像是否是使用已经公开的端口创建的，因此最好定义 `expose` 键，即使它是多余的。*

*Listing 9*

*清单 9*

```yaml
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
    expose:
      - "5432"
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: root
      POSTGRES_DB: testdb
    restart: on-failure
    networks:
      - integration-tests-example-test
```

The final configuration options that `db` needs are `environment`, `restart`, and `networks`. The `networks` key is given the name of the already defined network, not unlike the previous service definition. The `restart` key is given the value `on-failure` to ensure that the service will automatically restart if it fails at any point during its execution. The `environment` option can receive a list of environment variables which are then set  in the container’s shell. Most hosted images for popular applications,  such as postgres, have environment variables that can be specified to  configure the application that the image provides.

`db` 需要的最终配置选项是 `environment`、`restart` 和 `networks`。 `networks` 键被赋予已定义网络的名称，与之前的服务定义不同。 `restart` 键的值为 `on-failure` 以确保服务在执行期间的任何时候失败时都会自动重启。 `environment` 选项可以接收环境变量列表，然后在容器的 shell 中设置这些变量。大多数流行应用程序的托管图像（例如 postgres）都有环境变量，可以指定这些环境变量来配置图像提供的应用程序。

### Running The Tests

### 运行测试

With the Docker Compose file ready to go, the next step is to build the image based on a dockerfile that is referenced in the `listd_tests` service. This dockerfile defines an image that is capable of running  the integration tests for the entire service. Once the image is created, then tests can be run.

准备好 Docker Compose 文件后，下一步是基于在 `listd_tests` 服务中引用的 dockerfile 构建映像。这个 dockerfile 定义了一个能够为整个服务运行集成测试的图像。创建映像后，即可运行测试。

#### Building an Image Capable of Running Tests

#### 构建能够运行测试的图像

In order to build an image capable of running tests, four things have to be defined inside the dockerfile:

为了构建能够运行测试的镜像，必须在 dockerfile 中定义四件事：

Grab a base image that has the latest stable version of Go installed on it. Install `git` for Go modules. Copy the testable code into the container. Run the tests.

获取安装了最新稳定版 Go 的基础镜像。为 Go 模块安装 `git`。将可测试的代码复制到容器中。运行测试。

Let’s break these steps down and analyze the instructions that the dockerfile needs to carry them out.

让我们分解这些步骤并分析 dockerfile 执行它们所需的指令。

*Listing 10*

*清单 10*

```dockerfile
FROM golang:1.12-alpine
```

Listing 10 shows step 1 of 4. The image I’ve chosen as the base operating system image is `golang:1.11-alpine`. This image comes pre-installed with the latest stable version of Go at the time of writing this blog post.

清单 10 显示了第 1 步（共 4 步）。我选择作为基本操作系统映像的映像是 `golang:1.11-alpine`。在撰写本博文时，此映像预装了最新的稳定版 Go。

*Listing 11*

*清单 11*

```dockerfile
FROM golang:1.11-alpine

RUN set -ex;\
    apk update;\
    apk add --no-cache git
```

Because the Alpine OS is very lightweight, you must manually install the `git` dependency on top of the base Alpine image. Listing 11 shows step 2 where `git` is added to the image in order to use Go modules. The `apk update` command is ran before adding `git` to ensure the latest version of `git` is installed. If your project happens to use `cgo`, then you must also manually install `gcc` and its required libraries as well.

由于 Alpine 操作系统非常轻量级，您必须在基本 Alpine 映像之上手动安装 `git` 依赖项。清单 11 显示了第 2 步，其中将 `git` 添加到图像中以使用 Go 模块。在添加 `git` 之前运行 `apk update` 命令以确保安装最新版本的 `git`。如果您的项目碰巧使用了 `cgo`，那么您还必须手动安装 `gcc` 及其所需的库。

*Listing 12*

*清单 12*

```dockerfile
FROM golang:1.12-alpine

RUN set -ex;\
    apk update;\
    apk add --no-cache git

WORKDIR /go/src/github.com/george-e-shaw-iv/integration-tests-example/
```

For ease of use, in listing 12 the `WORKDIR` instruction is set to `/go/src/github.com/george-e-shaw-iv/integration-tests-example/` so that the rest of the instructions will be relative to that directory, which is within the container's `$GOPATH`. Step 3 of the process, copying the testable code into the container, is already taken care of due to the fact that a volume with the testable  code was mounted in listing 4.

为了便于使用，在清单 12 中，`WORKDIR` 指令设置为 `/go/src/github.com/george-e-shaw-iv/integration-tests-example/`，以便其余的指令相对于该目录，该目录位于容器的 `$GOPATH` 中。该过程的第 3 步，将可测试代码复制到容器中，由于在清单 4 中挂载了一个带有可测试代码的卷，因此已经完成。

*Listing 13*

*清单 13*

```dockerfile
FROM golang:1.12-alpine

RUN set -ex;\
    apk update;\
    apk add --no-cache git

WORKDIR /go/src/github.com/george-e-shaw-iv/integration-tests-example/

CMD CGO_ENABLED=0 go test ./...
```

Finally, listing 13 shows step 4, running the tests. This is accomplished using `go test ./...` with the `CMD` instruction.

最后，清单 13 显示了第 4 步，运行测试。这是通过使用 `go test ./...` 和 `CMD` 指令来完成的。

The tests are run with `CGO_ENABLED=0` as an inline environment variable because the tests in the sample project don’t use `cgo` and the alpine base image does not ship with a C compiler. Disabling `cgo` in this manner is necessary even if your project has no `cgo` code within it since Go will still attempt to use standard C libraries for certain networking tasks if `cgo` is enabled.

测试使用 `CGO_ENABLED=0` 作为内联环境变量运行，因为示例项目中的测试不使用 `cgo`，并且 alpine 基础映像不附带 C 编译器。以这种方式禁用 `cgo` 是必要的，即使你的项目中没有 `cgo` 代码，因为如果启用 `cgo`，Go 仍然会尝试使用标准 C 库来执行某些网络任务。

*Note: The code for the entire Dockerfile defining the custom image capable of running Go tests from within it can be found [here](https://github.com/george-e-shaw-iv/integration-tests-example/blob/master/cmd/listd/deploy/Dockerfile.test).*

*注意：整个 Dockerfile 的代码定义了能够从其中运行 Go 测试的自定义映像，可以在 [here](https://github.com/george-e-shaw-iv/integration-tests-example/blob/master/cmd/listd/deploy/Dockerfile.test).*

Now that the dockerfile that defines the image is written, the following Docker Compose command can bring up the `listd_test` and `db` services which will run all integration tests and report the outcome.

现在定义镜像的 dockerfile 已经写好了，下面的 Docker Compose 命令可以启动 `listd_test` 和 `db` 服务，它们将运行所有集成测试并报告结果。

*Listing 14*

*清单 14*

```
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```

The `--abort-on-container-exit` flag is  necessary as the other containers that contain the integrated services  will hang after the tests have finished running if the flag is omitted.

`--abort-on-container-exit` 标志是必要的，因为如果省略该标志，其他包含集成服务的容器将在测试运行完成后挂起。

### Clean-up

###  清理

*Listing 15*

*清单 15*

```makefile
test:
    docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
    docker-compose -f docker-compose.test.yml down --volumes
```

Stopping and removing containers, volumes, and networks  is a really important step that often gets neglected after running  tests. Figuring out why your tests are broken due to data that has  persisted from the last test run is a less than trivial bug that can be  easily avoided. To prevent this from happening, I created a simple `makefile` rule, `test`, showcased in listing 14, to build, run, and teardown containers without any human intervention.

停止和删除容器、卷和网络是一个非常重要的步骤，在运行测试后经常被忽略。弄清楚为什么您的测试由于上次测试运行中保留的数据而被破坏是一个可以轻松避免的不重要的错误。为了防止这种情况发生，我创建了一个简单的 `makefile` 规则 `test`，如清单 14 所示，用于在没有任何人工干预的情况下构建、运行和拆除容器。

*Listing 16*

*清单 16*

```makefile
test-db-up:
    docker-compose -f docker-compose.test.yml up --build db

test-db-down:
    docker-compose -f docker-compose.test.yml down --volumes db
```

The rules in listing 15 work best in a restricted  environment since they start both services in the Compose file. In order to achieve the same effect for local testing the `test-db-up` rule in listing 16 can be used before running any integration test, and `testdb-down` after all tests have been ran.

清单 15 中的规则在受限环境中效果最好，因为它们在 Compose 文件中启动了这两个服务。为了在本地测试中获得相同的效果，可以在运行任何集成测试之前使用清单 16 中的 `test-db-up` 规则，并在运行所有测试之后使用 `testdb-down`。

### Conclusion 

###  结论

In this post, I showed you how to setup your web service projects to use  Docker and Docker Compose. The files I reviewed allow you to run your Go tests and dependencies in a restrictive computing environment that  didn’t have Go pre-installed. In the next part of the series, I will  showcase the Go code required to set-up a test suite for the web  service, which will be the basis for writing insightful integration  tests, as well as writing actual integration tests.

在这篇文章中，我向您展示了如何设置 Web 服务项目以使用 Docker 和 Docker Compose。我查看的文件允许您在没有预安装 Go 的限制性计算环境中运行 Go 测试和依赖项。在本系列的下一部分中，我将展示为 Web 服务设置测试套件所需的 Go 代码，这将是编写有见地的集成测试以及编写实际集成测试的基础。

*Note: This entire series of posts draws its examples from [this repository](https://github.com/george-e-shaw-iv/integration-tests-example).*

*注意：这整个系列的帖子都从 [这个存储库](https://github.com/george-e-shaw-iv/integration-tests-example) 中提取示例。*

# Go Training

# 去训练

We have taught Go to thousands of developers all around the world since  2014. There is no other company that has been doing it longer and our  material has proven to help jump start developers 6 to 12 months ahead  of their knowledge of Go. We know what knowledge developers need in  order to be productive and efficient when writing software in Go.

自 2014 年以来，我们已向世界各地的数千名开发人员教授 Go。没有其他公司这样做的时间更长，而且我们的材料已被证明可以帮助开发人员提前 6 到 12 个月快速入门。我们知道开发人员在使用 Go 编写软件时需要什么知识才能提高生产力和效率。

Our classes are perfect for both experienced and beginning engineers. We  start every class from the beginning and get very detailed about the  internals, mechanics, specification, guidelines, best practices and  design philosophies. We cover a lot about "if performance matters" with a focus on mechanical sympathy, data oriented design, decoupling and  writing production software.

我们的课程非常适合经验丰富的工程师和初级工程师。我们从头开始每堂课，并非常详细地了解内部结构、机制、规范、指南、最佳实践和设计理念。我们涵盖了很多关于“如果性能很重要”的内容，重点是机械共鸣、面向数据的设计、解耦和编写生产软件。

![Capital One](https://www.ardanlabs.com/images/client-logos/white/training-client01.png)

![Cisco](https://www.ardanlabs.com/images/client-logos/white/training-client02.png)

![Visa](https://www.ardanlabs.com/images/client-logos/white/training-client03.png)

![Teradata](https://www.ardanlabs.com/images/client-logos/white/training-client04.png)

![Red Ventures](https://www.ardanlabs.com/images/client-logos/white/training-client05.png)

Interested in Ultimate Go Corporate Training and special pricing?

对 Ultimate Go 企业培训和特价感兴趣？

[Let’s Talk Corporate Training!](mailto:hello@ardanlabs.com?Subject=Let’s Talk Ultimate Go Corporate Training and special pricing!)

[让我们谈谈企业培训！](mailto:hello@ardanlabs.com?Subject=让我们谈谈 Ultimate Go 企业培训和特价！)

# Join Our Online Education Program

# 加入我们的在线教育计划

Our courses have been designed from training over 4,000 engineers since  2013 and they go beyond just being a language course. Our goal is to  challenge every student to think about what they are doing and why. 

自 2013 年以来，我们的课程是通过培训 4,000 多名工程师而设计的，它们不仅仅是一门语言课程。我们的目标是挑战每个学生思考他们在做什么以及为什么这样做。


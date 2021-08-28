# [Automating Go Integration Tests With Docker](https://www.dudley.codes/posts/2020.10.02-golang-docker-integration-tests/)

# [使用 Docker 自动化 Go 集成测试](https://www.dudley.codes/posts/2020.10.02-golang-docker-integration-tests/)

2020-10-02

2020-10-02

Go developers have everything needed to start writing automated unit tests using the `go test` command baked directly into the compiler toolchain. By hooking into the `testing` package’s lifecycle and importing Docker’s client libraries, we can  automate integration tests that manage their own Docker containers.

Go 开发人员拥有使用直接嵌入编译器工具链的 go test 命令开始编写自动化单元测试所需的一切。通过连接到 `testing` 包的生命周期并导入 Docker 的客户端库，我们可以自动化管理他们自己的 Docker 容器的集成测试。

> See also: accompanying example code [on GitHub](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang/06-dockerized).

> 另请参阅：随附的示例代码 [在 GitHub 上](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang/06-dockerized)。

## Unit Tests vs. Integration Tests

## 单元测试与集成测试

Unit tests are automated processes that execute small, isolated pieces of  code (e.g., functions) to test that they are performing correctly. If we had a function to calculate the area of a square:

单元测试是执行小的、独立的代码段（例如，函数）以测试它们是否正确执行的自动化过程。如果我们有一个函数来计算正方形的面积：

```go
func SquareArea(length int) int {
    return length * length
}
```


A unit test would feed it a known value and verify the result matches our expected output:

单元测试会为其提供一个已知值并验证结果是否与我们的预期输出匹配：

```go
func TestAverage(t *testing.T) {
    if actual: = SquareArea(12);actual != 144 {
            t.Errorf("Expected `%d` but got `%d`", 144, actual)
    }
}
```


A single test isn’t all that beneficial. But a  slew of tests checking results from a wide range of length inputs gives  us confidence that no matter what happens to the rest of our codebase,  the `SquareArea()` function is working as expected.

一次测试并不是那么有益。但是，从各种长度输入中检查结果的大量测试让我们相信，无论我们的代码库的其余部分发生什么，“SquareArea()”函数都按预期工作。

> See also: [Go Unit Test Exhibits](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang)

> 另见：[Go 单元测试展品](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang)

Most of our time, as developers, is spent integrating all the isolated  pieces of code; tests that check the result of these compositions are  referred to as *integration tests*. Unlike unit tests that are  generally straightforward to write, integration tests are more  challenging to craft, especially when the code under examination  integrates with an external dependency such as a database.

作为开发人员，我们的大部分时间都花在集成所有孤立的代码段上；检查这些组合物结果的测试被称为*整合测试*。与通常易于编写的单元测试不同，集成测试的编写更具挑战性，尤其是当被检查的代码与外部依赖项（例如数据库）集成时。

## Overview

##  概述

Fortunately, we can run the external dependency in a Docker container; we can  automate our integration tests! In this example, we’re going to hook  into the `testing` package’s lifecycle to download a Docker image for `PostgreSQL,` spin it up as a container with some basic settings, execute our integration tests, and finally terminate the container.

幸运的是，我们可以在 Docker 容器中运行外部依赖；我们可以自动化我们的集成测试！在这个例子中，我们将挂钩 `testing` 包的生命周期以下载 `PostgreSQL` 的 Docker 镜像，将它作为具有一些基本设置的容器启动，执行我们的集成测试，最后终止容器。

- Docker will need to be installed on the host for these examples to work.
- While the example code is specific for PostgreSQL, this flow can target almost any Dockerized dependency.

- 需要在主机上安装 Docker 才能使这些示例工作。
- 虽然示例代码特定于 PostgreSQL，但此流程几乎可以针对任何 Dockerized 依赖项。

## Entry Point

##  入口点

I like to keep these integration tests apart from our regular unit tests by creating an `integration_test.go` file that will serve as their entry point, and constraining it with a [build tag](https://golang.org/cmd/go/#hdr-Build_constraints).

我喜欢通过创建一个 `integration_test.go` 文件作为它们的入口点，并用 [build 标签](https://golang.org/cmd/去/#hdr-Build_constraints）。

```go
// +build integration

package postgresql
```


This way, the fast-running unit tests can continue to be part of our regular development cycle using `go test ./...`, and we can opt into the slower-running integration tests as needed (such as before a commit) with `go test ./... --tags integration --count=1`.

这样，快速运行的单元测试可以继续成为我们使用 go test ./... 的常规开发周期的一部分，并且我们可以根据需要选择运行较慢的集成测试（例如在提交之前）使用 `go test ./... --tags 集成 --count=1`。

## Automating the Service Container

## 自动化服务容器

First, we’ll define the setting need to start and connect to the Docker container:

首先，我们将定义启动和连接到 Docker 容器所需的设置：

```go
const  (
    // The canonical name of the image on Docker Hub
    _dockerImage = "docker.io/library/postgres:13"
    // The name of the database to be used by PostgreSQL
    _dbName   = "some_db"
    // The password to be used by PostgreSQL
    _password = "some_password"
)
```


Next, we’ll create a function to manage the  lifecycle of the container. It’ll pull the image from Docker Hub, spin  it up into a container, and return both an `error` value  along with a lambda function that will stop the running container. The  function requires the TCP port, mapping it from the host operating  system to the container’s internal service port

接下来，我们将创建一个函数来管理容器的生命周期。它将从 Docker Hub 拉取镜像，将其旋转到一个容器中，并返回一个 `error` 值以及一个将停止运行容器的 lambda 函数。该函数需要TCP端口，将其从宿主操作系统映射到容器内部服务端口

```go
// Create a PostgreSQL Docker container mapped to the specified TCP port.
func startPostgreContainer(port int) (closer func(), err error) {
    // Create a client for interacting with the local Docker client.
    cli, err := client.NewEnvClient()
    if err != nil {
        return nil, fmt.Errorf("Failed to get docker client, %w", err)
    }

    // Download the PostgreSQL image from Docker Hub.
    r, err := cli.ImagePull(context.TODO(), _dockerImage, types.ImagePullOptions{})
    if err != nil {
        return nil, fmt.Errorf("Failed to pull docker image, %w", err)
    }

    // Copy the download status from reader r to standard out.
    if _, err := io.Copy(os.Stdout, r);err != nil {
        return nil,
            fmt.Errorf("couldn't fetch docker image %q: %w", _dockerImage, err)
    }

    // Define a configuration for spinning up the container, injecting
    // the needed environment variables.
    containerCfg := container.Config{
        Env: []string{
            fmt.Sprintf("POSTGRES_DB=%s", _dbName),
            fmt.Sprintf("POSTGRES_PASSWORD=%s", _password),
        },
        Image: _dockerImage,
    }

    // Define a configuration for the container's host settings.
    hostCfg := container.HostConfig{
        AutoRemove: true,
        PortBindings: nat.PortMap{
            "5432/tcp": []nat.PortBinding{
                {HostPort: fmt.Sprintf("%d/tcp", port)},
            },
        },
    }

    // Define a unique name for the container.
    containerName := fmt.Sprintf("test_postgresql_%d", time.Now().Unix())

    // Spin up the Docker image into a running container.
    cont, err := dockerClient.ContainerCreate(context.Background(), &containerCfg, &hostCfg, nil, containerName)
    if err != nil {
        return nil, fmt.Errorf("Failed to create container, %w", err)
    }

    // Create an anonymous function that will shut down the running container.
    closeContainer := func() {
        if err := dockerClient.ContainerRemove(context.Background(), cont.ID, types.ContainerRemoveOptions{
            RemoveVolumes: true,
            RemoveLinks:   true,
            Force:         true,
        });err != nil {
            fmt.Printf("failed to remove container: %s", err.Error())
        }
    }

    // Spin up the container using the defined configurations.
    if err := dockerClient.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{});err != nil {
        closeContainer()
        return nil, fmt.Errorf("Failed to start container, %w", err)
    }

    // Return a lambda function to stop the container, along with a nil error.
    return func() {
        closeContainer()

        timeout := 10 * time.Second

        if err := dockerClient.ContainerStop(context.Background(), cont.ID, &timeout);err != nil {
            fmt.Printf("failed to stop container: %s", err.Error())
        }
    }, nil
```




## System Under Test

## 被测系统

In this example, the integration tests are for a PostgreSQL package, which uses an exported struct named `Broker` to expose its functionality. We’ll set up a *System Under Test* variable at the package level used by all of the integration tests.

在此示例中，集成测试针对 PostgreSQL 包，该包使用名为“Broker”的导出结构来公开其功能。我们将在所有集成测试使用的包级别设置一个 *System Under Test* 变量。

```go
var (
    // The System Under Test shared by all PostgreSQL integration tests.
    _sut *Broker = nil
)
```


## Hooking Into the Testing Flow

## 进入测试流程

Since our integration tests depended on a running Docker container, we’ll  define a testing entry point that will spin up (and verify) the  container, execute the tests, and then shut down the container. Since  this entry function exists above the scope of the `*testing.T` reference, it must print to the host's [standard out stream](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_(stdout))and terminate to the operating system with an [exit status](https://en.wikipedia.org/wiki/Exit_status) when needed.

由于我们的集成测试依赖于正在运行的 Docker 容器，我们将定义一个测试入口点，该入口点将启动（并验证）容器，执行测试，然后关闭容器。由于这个入口函数存在于 `*testing.T` 引用的范围之上，它必须打印到主机的 [标准输出流](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_(stdout))和需要时以 [退出状态](https://en.wikipedia.org/wiki/Exit_status) 终止到操作系统。

```go
// TestMain serves as the entry point for Go's test command.
func TestMain(m *testing.M) {
    // Get a free TCP port from the operating system.
    port := 0
    if addr, err := net.ResolveTCPAddr("tcp", "localhost:0");err == nil {
        if ln, err := net.ListenTCP("tcp", addr);err == nil {
            ln.Close()
            port = ln.Addr().(*net.TCPAddr).Port
        } else {
            fmt.Printf("Failed to connect to local TCP stack: %s", err.Error())
            os.Exit(1)
        }
    } else {
        fmt.Printf("Failed to resolve address for localhost: %s", err.Error())
        os.Exit(1)
    }

    // Start up the PostgreSQL container, capturing the lambda function
    // to close the running container.
    closer, err := startPostgreContainer(port)
    if err != nil {
        fmt.Printf("Failed to create Docker container for %s: %s", _dockerImage, err.Error())
        os.Exit(1)
    }

    // Define the configuration used by the SUT.
    cfg := Config{
        Host:     "localhost",
        Port:     port,
        User:     "postgres",
        Password: _password,
        DBName:   _dbName,
    }

    // Create the SUT and link to the testing-package level pointer.
    _sut, err = New(cfg)
    if err != nil {
        fmt.Printf("Failed to create new PostgreSQL broker: %s", err)
        closer()
        os.Exit(1)
    }

    // Ensure the container starts by waiting for a ping response,
    // with a 12 seconds timeout.
    for i := 0;;i++ {
        time.Sleep(2 * time.Second)

        if err := _sut.database.Ping();err == nil {
            break
        } else if i > 6 {
            fmt.Printf("PostgreSQL container did not respond to ping in %d attempts: %s",
                i, err.Error())
            closer()
            os.Exit(1)
        }
    }

    // Handoff test execution to all Test<NAME> functions utilizing the
    // *testing.T reference and capture the results of the tests.
    exitCode := m.Run()

    // Shut down the SUT.
    if err := _sut.Close();err != nil {
        fmt.Printf("Failed to properly close PostgreSQL agent: %s", err.Error())
    }

    // Shut down the Docker container.
    closer()

    // Exit to the operating system, returning the test results.
    os.Exit(exitCode)
}
```


## An Example Integration Test

## 一个示例集成测试

With the container setup and teardown automated, we can write our integration tests like our other tests, utilizing `*testing.T`. Since the PostgreSQL service is shared by all tests using `_sut`, each test should take care of its unique prerequisites. In this  example, the test will create a database table that won’t be used by any other tests. If the setup fails, we abort out using `t.Fatal()`.

通过容器设置和拆卸自动化，我们可以像其他测试一样使用 `*testing.T` 编写集成测试。由于 PostgreSQL 服务由所有使用 `_sut` 的测试共享，所以每个测试都应该注意其独特的先决条件。在这个例子中，测试将创建一个不会被任何其他测试使用的数据库表。如果设置失败，我们使用 `t.Fatal()` 中止。

```go
func TestRowCount(t *testing.T) {
    // Create an empty table used only for these checks.
    tableName := fmt.Sprintf("rowcount_%d", time.Now().Unix())
    createQuery := fmt.Sprintf("CREATE TABLE %s (k VARCHAR (8))", tableName)

    if _, err := _sut.database.Exec(createQuery);err != nil {
        t.Fatalf("Failed to create table %q for row count testing: %s", tableName, err)
    }

    t.Run("empty table should report 0 rows", func(t *testing.T) {
        count, err := _sut.RowCount(tableName)

        if err != nil {
            t.Errorf("Failed while executing RowCount(): %s", err.Error())
        }

        if count != 0 {
            t.Errorf("Counting the rows of an empty table should result in `0` not `%d`", count)
        }
    })

    // Additional checks (with their own setup) targeting _sut.RowCount()
    // would go here.
}
```


## Conclusion 

##  结论

The example code in this post is purposely minimal. For a fuller example, see [the example repo on GitHub](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang/06-dockerized).If you play with this flow, I’d love to hear about your experience [via Twitter](https://twitter.com/dudleycodes).

这篇文章中的示例代码特意精简。有关更完整的示例，请参阅 [GitHub 上的示例存储库](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang/06-dockerized)。如果你玩这个流程，我很想听听你的体验[通过 Twitter](https://twitter.com/dudleycodes)。

Happy testing!

测试愉快！

## Troubleshooting Tips

## 故障排除提示

### Unexpected Container Behavior

### 意外的容器行为

The process of getting the container correctly configured can take multiple iterations of trial and error. Some tips for troubleshooting during  this stage:

正确配置容器的过程可能需要多次反复试验。在此阶段进行故障排除的一些提示：

- Verify `TestMain()` gives the containerized application enough time to start up - some applications start-up much slower than others.
- Remove the setting `AutoRemove: true` from `container.HostConfig`, so that the container isn’t removed after it’s stopped, and then inspect its output using `docker logs <container-names>` from the command line.
- Verify you can interact with the running container using expected settings and off the shelve tools (such as the `pq` command-line tool for PostgreSQL).

- 验证`TestMain()` 为容器化应用程序提供了足够的启动时间 - 一些应用程序的启动速度比其他应用程序慢得多。
- 从 `container.HostConfig` 中删除设置 `AutoRemove: true`，以便容器在停止后不会被删除，然后从命令行使用 `docker logs <container-names>` 检查其输出。
- 验证您可以使用预期的设置和现成的工具（例如 PostgreSQL 的 `pq` 命令行工具）与正在运行的容器进行交互。

### Caching Issues

### 缓存问题

Go’s test command will use previous results where it can, by skipping all  checks against code that hasn’t changed. Unfortunately, with these  integration tests, external dependency could have changed, and the Go  toolchain has no way to know. To avoid this pitfall, force all tests to  run one time by adding `--count=1`.

Go 的 test 命令将尽可能使用以前的结果，跳过对未更改代码的所有检查。不幸的是，通过这些集成测试，外部依赖可能已经改变，Go 工具链无法知道。为了避免这个陷阱，通过添加 `--count=1` 强制所有测试运行一次。

`go test ./... --tags integration --count=1`.

`去测试 ./... --tags 集成 --count=1`。

> See also: [issue #3799](https://github.com/golang/go/issues/23799), a discussion around implementing a feature so that test can identify themselves as un-uncacheable.

> 另见：[issue #3799](https://github.com/golang/go/issues/23799)，关于实现一个特性的讨论，以便测试可以将自己标识为不可缓存的。

### Container Not Shutting Down

### 容器没有关闭

When calling `os.Exit()` defer statements will not execute. Make sure the `closer` lambda function is called explicitly before calling `os.Exit()`.

调用`os.Exit()` defer 语句时不会执行。确保在调用 `os.Exit()` 之前显式调用了 `closer` lambda 函数。

### Failed to pull docker image, repository name must be canonical

### 无法拉取 docker 镜像，存储库名称必须是规范的

When using the Docker SDK, images must be pulled using their full, canonical names, so `postgres:13` becomes `docker.io/library/postgres:13`

使用 Docker SDK 时，必须使用完整的规范名称提取图像，因此 `postgres:13` 变为 `docker.io/library/postgres:13`

### Integration Tests From Inside a Docker Container

### Docker 容器内部的集成测试

Executing these integration tests from within a Docker container is outside this post’s scope. Read up on *Docker In Docker* and bind mounting `/var/run/docker.sock`.

在 Docker 容器内执行这些集成测试超出了本文的范围。阅读 *Docker In Docker* 并绑定挂载 `/var/run/docker.sock`。

## Additional Resources
- [Develop with Docker Engine SDKs](https://docs.docker.com/engine/api/sdk/) 

## 其他资源
- [使用 Docker Engine SDK 开发](https://docs.docker.com/engine/api/sdk/)


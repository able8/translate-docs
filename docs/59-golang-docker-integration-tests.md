# [Automating Go Integration Tests With Docker](https://www.dudley.codes/posts/2020.10.02-golang-docker-integration-tests/)

2020-10-02

Go developers have everything needed to start writing automated unit tests using the `go test` command baked directly into the compiler toolchain. By hooking into the `testing` package’s lifecycle and importing Docker’s client libraries, we can  automate integration tests that manage their own Docker containers.

> See also: accompanying example code [on GitHub](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang/06-dockerized).

## Unit Tests vs. Integration Tests

Unit tests are automated processes that execute small, isolated pieces of  code (e.g., functions) to test that they are performing correctly. If we had a function to calculate the area of a square:

```go
func SquareArea(length int) int {
    return length * length
}
```

A unit test would feed it a known value and verify the result matches our expected output:

```go
func TestAverage(t *testing.T) {
    if actual: = SquareArea(12); actual != 144 {
            t.Errorf("Expected `%d` but got `%d`", 144, actual)
    }
}
```

A single test isn’t all that beneficial. But a  slew of tests checking results from a wide range of length inputs gives  us confidence that no matter what happens to the rest of our codebase,  the `SquareArea()` function is working as expected.

> See also: [Go Unit Test Exhibits](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang)

Most of our time, as developers, is spent integrating all the isolated  pieces of code; tests that check the result of these compositions are  referred to as *integration tests*. Unlike unit tests that are  generally straightforward to write, integration tests are more  challenging to craft, especially when the code under examination  integrates with an external dependency such as a database.

## Overview

Fortunately, we can run the external dependency in a Docker container; we can  automate our integration tests! In this example, we’re going to hook  into the `testing` package’s lifecycle to download a Docker image for `PostgreSQL,` spin it up as a container with some basic settings, execute our integration tests, and finally terminate the container.

- Docker will need to be installed on the host for these examples to work.
- While the example code is specific for PostgreSQL, this flow can target almost any Dockerized dependency.

## Entry Point

I like to keep these integration tests apart from our regular unit tests by creating an `integration_test.go` file that will serve as their entry point, and constraining it with a [build tag](https://golang.org/cmd/go/#hdr-Build_constraints).

```go
// +build integration

package postgresql
```

This way, the fast-running unit tests can continue to be part of our regular development cycle using `go test ./...`, and we can opt into the slower-running integration tests as needed (such as before a commit) with `go test ./... --tags integration --count=1`.

## Automating the Service Container

First, we’ll define the setting need to start and connect to the Docker container:

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
    if _, err := io.Copy(os.Stdout, r); err != nil {
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
        }); err != nil {
            fmt.Printf("failed to remove container: %s", err.Error())
        }
    }

    // Spin up the container using the defined configurations.
    if err := dockerClient.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{}); err != nil {
        closeContainer()
        return nil, fmt.Errorf("Failed to start container, %w", err)
    }

    // Return a lambda function to stop the container, along with a nil error.
    return func() {
        closeContainer()

        timeout := 10 * time.Second

        if err := dockerClient.ContainerStop(context.Background(), cont.ID, &timeout); err != nil {
            fmt.Printf("failed to stop container: %s", err.Error())
        }
    }, nil
```

## System Under Test

In this example, the integration tests are for a PostgreSQL package, which uses an exported struct named `Broker` to expose its functionality. We’ll set up a *System Under Test* variable at the package level used by all of the integration tests.

```go
var (
    // The System Under Test shared by all PostgreSQL integration tests.
    _sut *Broker = nil
)
```

## Hooking Into the Testing Flow

Since our integration tests depended on a running Docker container, we’ll  define a testing entry point that will spin up (and verify) the  container, execute the tests, and then shut down the container. Since  this entry function exists above the scope of the `*testing.T` reference, it must print to the host’s [standard out stream](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_(stdout)) and terminate to the operating system with an [exit status](https://en.wikipedia.org/wiki/Exit_status) when needed.

```go
// TestMain serves as the entry point for Go's test command.
func TestMain(m *testing.M) {
    // Get a free TCP port from the operating system.
    port := 0
    if addr, err := net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
        if ln, err := net.ListenTCP("tcp", addr); err == nil {
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
    for i := 0; ; i++ {
        time.Sleep(2 * time.Second)

        if err := _sut.database.Ping(); err == nil {
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
    if err := _sut.Close(); err != nil {
        fmt.Printf("Failed to properly close PostgreSQL agent: %s", err.Error())
    }

    // Shut down the Docker container.
    closer()

    // Exit to the operating system, returning the test results.
    os.Exit(exitCode)
}
```

## An Example Integration Test

With the container setup and teardown automated, we can write our integration tests like our other tests, utilizing `*testing.T`. Since the PostgreSQL service is shared by all tests using `_sut`, each test should take care of its unique prerequisites. In this  example, the test will create a database table that won’t be used by any other tests. If the setup fails, we abort out using `t.Fatal()`.

```go
func TestRowCount(t *testing.T) {
    // Create an empty table used only for these checks.
    tableName := fmt.Sprintf("rowcount_%d", time.Now().Unix())
    createQuery := fmt.Sprintf("CREATE TABLE %s (k VARCHAR (8))", tableName)

    if _, err := _sut.database.Exec(createQuery); err != nil {
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

The example code in this post is purposely minimal. For a fuller example, see [the example repo on GitHub](https://github.com/dudleycodes/UnitTestExhibits/tree/master/golang/06-dockerized). If you play with this flow, I’d love to hear about your experience [via Twitter](https://twitter.com/dudleycodes).

Happy testing!

## Troubleshooting Tips

### Unexpected Container Behavior

The process of getting the container correctly configured can take multiple iterations of trial and error. Some tips for troubleshooting during  this stage:

- Verify `TestMain()` gives the containerized application enough time to start up - some applications start-up much slower than others.
- Remove the setting `AutoRemove: true` from `container.HostConfig`, so that the container isn’t removed after it’s stopped, and then inspect its output using `docker logs <container-names>` from the command line.
- Verify you can interact with the running container using expected settings and off the shelve tools (such as the `pq` command-line tool for PostgreSQL).

### Caching Issues

Go’s test command will use previous results where it can, by skipping all  checks against code that hasn’t changed. Unfortunately, with these  integration tests, external dependency could have changed, and the Go toolchain has no way to know. To avoid this pitfall, force all tests to  run one time by adding `--count=1`.

`go test ./... --tags integration --count=1`.

> See also: [issue #3799](https://github.com/golang/go/issues/23799), a discussion around implementing a feature so that test can identify themselves as un-uncacheable.

### Container Not Shutting Down

When calling `os.Exit()` defer statements will not execute. Make sure the `closer` lambda function is called explicitly before calling `os.Exit()`.

### Failed to pull docker image, repository name must be canonical

When using the Docker SDK, images must be pulled using their full, canonical names, so `postgres:13` becomes `docker.io/library/postgres:13`

### Integration Tests From Inside a Docker Container

Executing these integration tests from within a Docker container is outside this post’s scope. Read up on *Docker In Docker* and bind mounting `/var/run/docker.sock`.

## Additional Resources
- [Develop with Docker Engine SDKs](https://docs.docker.com/engine/api/sdk/)
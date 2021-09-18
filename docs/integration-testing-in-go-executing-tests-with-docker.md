# Integration Testing in Go: Part I - Executing Tests with Docker

George ShawMarch 18, 2019

### Introduction

> *“Testing leads to failure, and failure leads to understanding.” - \*Burt Rutan**

Burt Rutan is an aerospace engineer who designed Voyager, the first plane to fly around the world without stopping or refueling. Although Rutan was  not a software engineer, his words speak volumes to the importance of  testing, even testing software. Testing software in all forms is  extremely important, whether it be unit, integration, system, or  acceptance testing. However, depending on the project, one form of  testing can be more valuable than the others. In other words, sometimes  one form of testing can lead to better understanding about the health  and integrity of the software than the other forms.

When  developing a web service, I believe a strong set of integration tests  can provide a better understanding of the service than other types of  tests. Integration tests are a form of software testing that tests the  interaction of your code against the dependencies your application is  leveraging, such as databases and messaging systems. Without integration tests, it’s difficult to trust the end-to-end operation of a web  service. I believe this is true because individual units of code being  tested in a web service rarely provide the same level of insight as  integration tests.

This is the first entry of a two part series  about integration testing in Go. The ideas, code and processes shared in this series aim to be easily extendable to the web service projects you are working on. In this post, I will show you how to setup your web  service projects to use Docker and Docker Compose to run your Go tests  and dependencies in a restrictive computing environment that doesn’t  have Go pre-installed.

### Why use Docker and Docker Compose

What attracts many developers to Docker is how you’re able to load  applications on your host machine without the burden of having to  install and manage them manually. This means you can load complex  software including, but not limited to, databases (e.g. Postgres),  messaging systems (e.g. Kafka) and monitoring systems (e.g. Prometheus). All of this is done by downloading a set of images that represent the  application and all of its dependencies.

*Note: For more  information on containers, Docker has a webpage devoted to the  definition of a container and highlights the differences and  similarities between a container and a virtual machine that can be found [here](https://www.docker.com/resources/what-container).*

Docker Compose is an orchestration tool that aids in building, running, and  networking a group of containers together inside of a single sandbox.  With a single command, `docker-compose up`, you can make your Docker Compose file come to life. All the services defined in the  compose file will become containers running as a group within their own  networked sandbox and run as configured. This is in contrast to manually building, running, and networking each of your containers in order to  allow them to run together, communicate with each other, and persist  data.

Since Docker Compose allows you to group different  applications together and run them within a single networked sandbox,  you can start and stop an entire suite of applications with a single  command. You can even handpick certain applications to run from the  group. This group of applications can be deployed as a single unit and  be built and tested by a CI (continuous integration) environment. Docker Compose ultimately helps ensure that your application is consistent  across any environment it’s tested and deployed on.

*Note: For  more information on Docker Compose, visit the webpage for the overview  of Docker Compose on the official Docker website [here](https://docs.docker.com/compose/overview/).*

Another big benefit of Docker and Docker Compose is that they help facilitate  an easier transition when bringing new developers into a project.  Instead of having complex documentation regarding how a development  environment is installed and managed, new developers just need to  execute a few Docker and Docker Compose commands to get started. The  Docker CLI takes care of downloading required images if they don’t  currently exist on the host machine when an application is started.

### Using Docker and Docker Compose to Run Tests

The web service application referenced throughout this series exposes a  simple CRUD based REST API with a Postgres database. The project uses  Docker to run a Postgres database for both production and testing. The  tests for this application need to be able to run in a local development environment that already has Go installed and a restricted environment  where Go does not exist.

The following Docker Compose file  supports the ability to run integration tests for the project in both  environments I mentioned above. In this section, I will break down the  configuration options that I chose and why I chose them.

*Listing 1*

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

*Listing 2*

```yaml
networks:
  integration-tests-example:
    driver: bridge
```

By having your service definitions in one compose file  they are automatically placed within the same network by default and  therefore can communicate with each other. However, it’s a best practice to create a network for your services as opposed to using the default  network. The top-level `networks` configuration defines the name of the network and the driver it uses, the bridge driver in this case.

The bridge driver is the default driver provided by Docker which creates a  private internal network for containers to communicate within. The  services are told to use the created network within their service  definition configuration in the compose file.

*Listing 3*

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

*Listing 4*

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

The `depends_on` key tells the `listd_tests` service to wait to start until the `db` service has already started. In addition to asserting start order of services, this key will disallow the `listd_tests` service from being run independently of the `db` service. The `volumes` key tells compose to mount the current working directory, denoted by `$PWD` (**P**rint **W**orking **D**irectory), to `/go/src/github.com/george-e-shaw-iv/integration-tests-example` within the container, which is where the code will be located and tested.

*Listing 5*

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

*Listing 6*

```yaml
db:
    image: postgres:11.1
```

The container in the next service definition, `db`, defines its image by [using a image hosted at Docker Hub](https://hub.docker.com/_/postgres), the `postgres:1.11` image. The Docker CLI is smart enough to know to look in the Docker Hub image repository if it can’t find the image on your local machine.

*Listing 7*

```yaml
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
```

For security reasons, by default none of the container  ports are accessible from your host machine. This proves to be a problem when running integration tests locally, as the integrated service is  pretty worthless if it isn’t accessible. The `ports` key defines port mappings from your host machine to your container in the following format: `"HOST_PORT:CONTAINER_PORT"`. The preceding definition in listing 7 ensures that port 5432 on your machine is mapped to port 5432 on the `db` container, as that is the port that Postgres is running on within the container by default.

*Listing 8*

```yaml
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
    expose:
      - "5432"
```

In the same manner that container ports aren’t exposed  to the host machine by default, container ports are also not exposed to  containers running within the networked sandbox by default. This is true even if they are on the same network. In order to expose a port to  other containers running within the networked sandbox, the `expose` configuration key needs to be set.

*Note: In the case of the `postgres:1.11` image, port 5432 has already been exposed thanks to the person who  created the image. Since you don’t know if the image was created with  the port already exposed unless you look at an image’s Dockerfile, it is best to define the `expose` key, even if it’s redundant.*

*Listing 9*

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

### Running The Tests

With the Docker Compose file ready to go, the next step is to build the image based on a dockerfile that is referenced in the `listd_tests` service. This dockerfile defines an image that is capable of running  the integration tests for the entire service. Once the image is created, then tests can be run.

#### Building an Image Capable of Running Tests

In order to build an image capable of running tests, four things have to be defined inside the dockerfile:

Grab a base image that has the latest stable version of Go installed on it. Install `git` for Go modules. Copy the testable code into the container. Run the tests.

Let’s break these steps down and analyze the instructions that the dockerfile needs to carry them out.

*Listing 10*

```dockerfile
FROM golang:1.12-alpine
```

Listing 10 shows step 1 of 4. The image I’ve chosen as the base operating system image is `golang:1.11-alpine`. This image comes pre-installed with the latest stable version of Go at the time of writing this blog post.

*Listing 11*

```dockerfile
FROM golang:1.11-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git
```

Because the Alpine OS is very lightweight, you must manually install the `git` dependency on top of the base Alpine image. Listing 11 shows step 2 where `git` is added to the image in order to use Go modules. The `apk update` command is ran before adding `git` to ensure the latest version of `git` is installed. If your project happens to use `cgo`, then you must also manually install `gcc` and its required libraries as well.

*Listing 12*

```dockerfile
FROM golang:1.12-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /go/src/github.com/george-e-shaw-iv/integration-tests-example/
```

For ease of use, in listing 12 the `WORKDIR` instruction is set to `/go/src/github.com/george-e-shaw-iv/integration-tests-example/` so that the rest of the instructions will be relative to that directory, which is within the container’s `$GOPATH`. Step 3 of the process, copying the testable code into the container, is already taken care of due to the fact that a volume with the testable  code was mounted in listing 4.

*Listing 13*

```dockerfile
FROM golang:1.12-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /go/src/github.com/george-e-shaw-iv/integration-tests-example/

CMD CGO_ENABLED=0 go test ./...
```

Finally, listing 13 shows step 4, running the tests. This is accomplished using `go test ./...` with the `CMD` instruction.

The tests are run with `CGO_ENABLED=0` as an inline environment variable because the tests in the sample project don’t use `cgo` and the alpine base image does not ship with a C compiler. Disabling `cgo` in this manner is necessary even if your project has no `cgo` code within it since Go will still attempt to use standard C libraries for certain networking tasks if `cgo` is enabled.

*Note: The code for the entire Dockerfile defining the custom image capable of running Go tests from within it can be found [here](https://github.com/george-e-shaw-iv/integration-tests-example/blob/master/cmd/listd/deploy/Dockerfile.test).*

Now that the dockerfile that defines the image is written, the following Docker Compose command can bring up the `listd_test` and `db` services which will run all integration tests and report the outcome.

*Listing 14*

```
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```

The `--abort-on-container-exit` flag is  necessary as the other containers that contain the integrated services  will hang after the tests have finished running if the flag is omitted.

### Clean-up

*Listing 15*

```makefile
test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes
```

Stopping and removing containers, volumes, and networks  is a really important step that often gets neglected after running  tests. Figuring out why your tests are broken due to data that has  persisted from the last test run is a less than trivial bug that can be  easily avoided. To prevent this from happening, I created a simple `makefile` rule, `test`, showcased in listing 14, to build, run, and teardown containers without any human intervention.

*Listing 16*

```makefile
test-db-up:
	docker-compose -f docker-compose.test.yml up --build db

test-db-down:
	docker-compose -f docker-compose.test.yml down --volumes db
```

The rules in listing 15 work best in a restricted  environment since they start both services in the Compose file. In order to achieve the same effect for local testing the `test-db-up` rule in listing 16 can be used before running any integration test, and `testdb-down` after all tests have been ran.

### Conclusion

In this post, I showed you how to setup your web service projects to use  Docker and Docker Compose. The files I reviewed allow you to run your Go tests and dependencies in a restrictive computing environment that  didn’t have Go pre-installed. In the next part of the series, I will  showcase the Go code required to set-up a test suite for the web  service, which will be the basis for writing insightful integration  tests, as well as writing actual integration tests.

*Note: This entire series of posts draws its examples from [this repository](https://github.com/george-e-shaw-iv/integration-tests-example).*

# Go Training

We have taught Go to thousands of developers all around the world since  2014. There is no other company that has been doing it longer and our  material has proven to help jump start developers 6 to 12 months ahead  of their knowledge of Go. We know what knowledge developers need in  order to be productive and efficient when writing software in Go.

Our classes are perfect for both experienced and beginning engineers. We  start every class from the beginning and get very detailed about the  internals, mechanics, specification, guidelines, best practices and  design philosophies. We cover a lot about "if performance matters" with a focus on mechanical sympathy, data oriented design, decoupling and  writing production software.

![Capital One](https://www.ardanlabs.com/images/client-logos/white/training-client01.png)

![Cisco](https://www.ardanlabs.com/images/client-logos/white/training-client02.png)

![Visa](https://www.ardanlabs.com/images/client-logos/white/training-client03.png)

![Teradata](https://www.ardanlabs.com/images/client-logos/white/training-client04.png)

![Red Ventures](https://www.ardanlabs.com/images/client-logos/white/training-client05.png)

Interested in Ultimate Go Corporate Training and special pricing?

[Let’s Talk Corporate Training!](mailto:hello@ardanlabs.com?Subject=Let’s Talk Ultimate Go Corporate Training and special pricing!)

# Join Our Online Education Program

Our courses have been designed from training over 4,000 engineers since  2013 and they go beyond just being a language course. Our goal is to  challenge every student to think about what they are doing and why.
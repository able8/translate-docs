# Developing A RESTful API With Golang And A MongoDB NoSQL Database

# 使用 Golang 和 MongoDB NoSQL 数据库开发 RESTful API

**This tutorial was updated on June 27, 2019 to reflect the latest versions of the technologies mentioned.**

**本教程于 2019 年 6 月 27 日更新，以反映所提及技术的最新版本。**

If you’ve been following along, you’re probably familiar with my love of  Node.js and the Go programming language. Over the past few weeks I've  been writing a lot about API development with MongoDB and Node.js, but  did you know that [MongoDB](https://www.mongodb.com/) also has an official SDK for [Golang ](https://golang.org/)? As of now the SDK is in beta, but at least it exists and is progressing.

如果你一直在关注，你可能熟悉我对 Node.js 和 Go 编程语言的热爱。在过去的几周里，我写了很多关于使用 MongoDB 和 Node.js 进行 API 开发的文章，但是您知道吗 [MongoDB](https://www.mongodb.com/) 也有 [Golang] 的官方 SDK ](https://golang.org/)？截至目前，SDK 处于测试阶段，但至少它存在并且正在取得进展。

The good news is that it isn’t difficult to develop with the Go SDK for MongoDB and you can accomplish quite a bit with it.

好消息是，使用 Go SDK for MongoDB 进行开发并不困难，您可以使用它完成很多工作。

In this tutorial we’re going to take a look at building a simple REST API  that leverages the Go SDK for creating data and querying in a MongoDB  NoSQL database.

在本教程中，我们将看看构建一个简单的 REST API，它利用 Go SDK 在 MongoDB NoSQL 数据库中创建数据和查询。

Before going forward, we’re going to assume that  you already have an instance of MongoDB configured and Golang is  installed and configured as well. If you need help, I wrote a simple  tutorial for [deploying MongoDB with Docker](https://www.thepolyglotdeveloper.com/2019/01/getting-started-mongodb-docker-container-deployment/) that you can check out . It is a great way to get up and running quickly.

在继续之前，我们将假设您已经配置了一个 MongoDB 实例，并且还安装并配置了 Golang。如果你需要帮助，我写了一个简单的[使用 Docker 部署 MongoDB教程，你可以查看.这是快速启动和运行的好方法。

## Creating a New Go Project with the MongoDB Dependencies

## 使用 MongoDB 依赖项创建一个新的 Go 项目

There are many ways to create a REST API with Golang. We're going to be  making use of a popular multiplexer that I wrote about in a previous  tutorial titled, [Create a Simple RESTful API with Golang](https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/).

有很多方法可以使用 Golang 创建 REST API。我们将使用一个流行的多路复用器，我在之前的标题为 [使用 Golang 创建一个简单的 RESTful API](https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/)。

Create a new directory within your **$GOPATH** and add a **main.go** file. Before we start adding code, we need to obtain our dependencies. From the command line, execute the following:

在 **$GOPATH** 中创建一个新目录并添加一个 **main.go** 文件。在我们开始添加代码之前，我们需要获取我们的依赖项。从命令行，执行以下命令：

```bash
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
```


The above commands will get our multiplexer as well as the MongoDB SDK for Golang. With the dependencies available, open the **main.go** file and include the following boilerplate code:

上述命令将获得我们的多路复用器以及 Golang 的 MongoDB SDK。有了可用的依赖项，打开 **main.go** 文件并包含以下样板代码：

```golang
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Person struct {
    ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
    Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {}
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) { }
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) { }

func main() {
    fmt.Println("Starting the application...")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, _ = mongo.Connect(ctx, clientOptions)
    router := mux.NewRouter()
    router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
    http.ListenAndServe(":12345", router)
}
```


So what is happening in the above code? Well,  we’re not actually doing any MongoDB logic, but we are configuring our  API and connecting to the database.

那么上面的代码中发生了什么？好吧，我们实际上并没有做任何 MongoDB 逻辑，但是我们正在配置我们的 API 并连接到数据库。

In our example, the `Person` data structure will be the data that we wish to work with. We have both JSON and BSON annotations so that we can work with MongoDB BSON data  and receive or respond with JSON data. In the `main` function we are connecting to an instance of MongoDB, which in my scenario is on my local computer, and configuring our API routes.

在我们的示例中，`Person` 数据结构将是我们希望使用的数据。我们有 JSON 和 BSON 注释，以便我们可以使用 MongoDB BSON 数据并使用 JSON 数据接收或响应。在 main 函数中，我们连接到 MongoDB 的一个实例，在我的场景中，它在我的本地计算机上，并配置我们的 API 路由。

While we could create a full CRUD API, we’re just going to work with three endpoints. You could easily expand upon this to do updates and deletes.

虽然我们可以创建一个完整的 CRUD API，但我们只会使用三个端点。您可以轻松地扩展它以进行更新和删除。

With the boilerplate code out of the way, we can focus on each of our endpoint functions.

有了样板代码，我们就可以专注于我们的每个端点功能。

## Designing API Endpoints for HTTP Interaction 

## 为 HTTP 交互设计 API 端点

Assuming that we’re working with a fresh instance of MongoDB, the first thing to do might be to create data. For this reason we’re going to work on the `CreatePersonEndpoint` to receive client data.

假设我们正在使用一个新的 MongoDB 实例，首先要做的可能是创建数据。出于这个原因，我们将使用 `CreatePersonEndpoint` 来接收客户端数据。

```golang
func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var person Person
    _ = json.NewDecoder(request.Body).Decode(&person)
    collection := client.Database("thepolyglotdeveloper").Collection("people")
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    result, _ := collection.InsertOne(ctx, person)
    json.NewEncoder(response).Encode(result)
}
```


In the above code we are setting the response type for JSON because most web applications can easily work with JSON data. In the request there will be a JSON payload which we are decoding into a native data structure based on the annotations.

在上面的代码中，我们为 JSON 设置响应类型，因为大多数 Web 应用程序可以轻松地处理 JSON 数据。在请求中将有一个 JSON 有效负载，我们根据注释将其解码为原生数据结构。

Now that we have  data to work with, we connect to a particular database within our  instance and open a particular collection. With a connection to that  collection we can insert our native data structure data and return the  result which would be an object id.

现在我们有了要处理的数据，我们连接到实例中的特定数据库并打开特定集合。通过与该集合的连接，我们可以插入我们的本机数据结构数据并返回将是对象 ID 的结果。

Since we know the id, we can work towards obtaining that particular document from the database.

由于我们知道 id，我们可以努力从数据库中获取该特定文档。

```golang
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    params := mux.Vars(request)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var person Person
    collection := client.Database("thepolyglotdeveloper").Collection("people")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(person)
}
```


With the `GetPersonEndpoint` we are  passing an id as the route parameter and converting it to an object id. After getting a collection to work with we can make use of the `FindOne` function and a filter based on our id. This single result can then be decoded to a `Person` object.

使用`GetPersonEndpoint`，我们将一个id 作为路由参数传递并将其转换为一个对象id。在获得一个集合后，我们可以使用`FindOne` 函数和一个基于我们 id 的过滤器。然后可以将这个单个结果解码为“Person”对象。

As long as there were no errors, we can return the person which should include an `id`, a `firstname`, and a `lastname` property.

只要没有错误，我们就可以返回应包含“id”、“firstname”和“lastname”属性的人。

This leads us to the most complicated of our three endpoints.

这将我们引向三个端点中最复杂的一个。

```golang
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var people []Person
    collection := client.Database("thepolyglotdeveloper").Collection("people")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var person Person
        cursor.Decode(&person)
        people = append(people, person)
    }
    if err := cursor.Err();err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(people)
}
```


The `GetPeopleEndpoint` should return  all documents within our collection. This means we have to work with  cursors in MongoDB similar to how you might work with cursors in an  RDBMS. Using a `Find` with no filter criteria we can loop through our cursor, decoding each iteration and adding it to a slice of the `Person` data type.

`GetPeopleEndpoint` 应该返回我们集合中的所有文档。这意味着我们必须在 MongoDB 中使用游标，类似于在 RDBMS 中使用游标的方式。使用没有过滤条件的“Find”，我们可以遍历游标，解码每次迭代并将其添加到“Person”数据类型的切片中。

Provided there were no errors along the way, we can encode the slice and return  it as JSON to the client. If we wanted to, we could add filter criteria  so that only specific documents are returned rather than everything. The filter criteria would include document properties and the anticipated  values.

如果一路上没有错误，我们可以对切片进行编码并将其作为 JSON 返回给客户端。如果我们愿意，我们可以添加过滤条件，以便仅返回特定文档而不是所有文档。过滤条件将包括文档属性和预期值。

## Conclusion

##  结论

You just saw how to create a simple REST API with [Golang](https://golang.org/) and [MongoDB](https://www.mongodb.com/). This is a step up from my basic tutorial titled, [Create a Simple RESTful API with Golang](https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/) . 

您刚刚看到了如何使用 [Golang](https://golang.org/) 和 [MongoDB](https://www.mongodb.com/) 创建一个简单的 REST API。这是我的基础教程 [使用 Golang 创建一个简单的 RESTful API](https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/) 的升级版.

Being that the MongoDB SDK for Go is in beta, the [official documentation](https://docs.mongodb.com/drivers/go) is pretty terrible. I found that much of it didn’t work out of the box  and it advised things that might not be the best approach. While I did  my best to fill in the gap, there may be better ways to accomplish  things. I’m open to hearing these better ways in the comments.

由于 MongoDB SDK for Go 处于测试阶段，[官方文档](https://docs.mongodb.com/drivers/go) 非常糟糕。我发现其中的大部分内容都不是开箱即用的，它建议了可能不是最佳方法的事情。虽然我已尽力填补空白，但可能有更好的方法来完成任务。我愿意在评论中听到这些更好的方法。

A video version of this tutorial can be found below. 

本教程的视频版本可以在下面找到。


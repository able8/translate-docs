# Mocking gRPC in Go

# 在 Go 中模拟 gRPC

One of the huge benefits of using gRPC is the ability to  autogenerate our client and server stubs from the protocol buffer  definitions.

使用 gRPC 的巨大好处之一是能够从协议缓冲区定义自动生成我们的客户端和服务器存根。

In the same way, we can we generate our own code by building a plugin for the protocol buffer compiler (protoc).

同样，我们可以通过为协议缓冲区编译器（protoc）构建一个插件来生成我们自己的代码。

The plugin we're creating is going to auto-generate gRPC response  messages so that we can build a mock gRPC server in Go (Golang).

我们正在创建的插件将自动生成 gRPC 响应消息，以便我们可以在 Go (Golang) 中构建模拟 gRPC 服务器。

## Goal

##  目标

Given the following `proto` file:

给定以下 `proto` 文件：

```proto
service MyService {
  rpc Get(Request) returns (Response) {}
}
message Request {}
message Response {
  string id = 1;
}
```

We want to be able to generate something like:

我们希望能够生成如下内容：

```go
func (m *MyServiceMock) Get(context.Context, *Request) (*Response, error) {
    return Response{Id: "f267332d-0d9a-4220-b055-63b661b600db"}, nil
}
```

## Creating a protoc Go plugin

## 创建一个 protoc Go 插件

All we have to do is implement the `Plugin` interface defined in `protobuf/protoc-gen-go/generator`:

我们所要做的就是实现`protobuf/protoc-gen-go/generator`中定义的`Plugin`接口：

```go
type grpcmock struct {
    *generator.Generator
}
func New() generator.Plugin {
    return &grpcmock{}
}
func (g *grpcmock) Name() string {
    return "grpcmock"
}
func (g *grpcmock) Init(gen *generator.Generator) {
    g.Generator = gen
}
func (g *grpcmock) Generate(file *generator.FileDescriptor) {
    for _, service := range file.FileDescriptorProto.Service {
        g.mockService(file, service)
    }
}
func (g *grpcmock) GenerateImports(file *generator.FileDescriptor) {
    imports := generator.NewPluginImports(g.Generator)
    imports.GenerateImports(file)
}
```

## Mocking a Service

## 模拟服务

Now that we have satisfied our Plugin interface we can implement our mock services:

现在我们已经满足了我们的插件接口，我们可以实现我们的模拟服务：

```go
func (g *grpcmock) mockService(file *generator.FileDescriptor, service *descriptor.ServiceDescriptorProto) {
    origServName := service.GetName()
    servName := generator.CamelCase(origServName)
    servTypeName := fmt.Sprintf("%sMock", servName)

    g.P(`type `, servTypeName, ` struct {}`)
    g.P()
    for _, method := range service.Method {
        g.mockMethod(servTypeName, method)
    }
}
```

This gives us our mock struct that we can create our methods from to implement the gRPC interface for our service:

这为我们提供了我们的模拟结构，我们可以从中创建我们的方法来为我们的服务实现 gRPC 接口：

```go
func (g *grpcmock) mockMethod(servTypeName string, method *descriptor.MethodDescriptorProto) {
    methName := generator.CamelCase(method.GetName())
    inType := g.typeName(method.GetInputType())
    outType := g.typeName(method.GetOutputType())

    g.P(`func (m *`, servTypeName, `) `, methName, `(context.Context, *`, inType, `) (*`, outType, `, error){`)
    g.In()
    
    msg := g.objectNamed(method.GetOutputType()).(*generator.Descriptor)
    g.P(`res := `)
    g.generateMockMessage(msg)
    g.P(`return res, nil`)

    g.Out()
    g.P(`}`)
}
```

The last part is to generate the response message and it's fields:

最后一部分是生成响应消息及其字段：

```go
func (g *grpcmock) generateMockMessage(msg *generator.Descriptor) {
    msgName := g.TypeName(msg)
    g.P(msgName, `{`)
    g.In()
    for _, field := range msg.Field {
        fieldName := g.GetFieldName(msg, field)
        if field.IsString() {
            g.P(fieldName, `: "f267332d-0d9a-4220-b055-63b661b600db",`)
        }
    }
    g.Out()
    g.P(`}`)
}
```

For the above `string` field we are returning a hard coded value, but we can easily replace this for auto-generated values.

对于上面的 `string` 字段，我们将返回一个硬编码值，但我们可以轻松地将其替换为自动生成的值。

## Building our plugin

## 构建我们的插件

All that is left to do is wire up our plugin (errors ignored for brevity):

剩下要做的就是连接我们的插件（为简洁起见，忽略错误）：

```go
func main() {
    gen := generator.New()
    data, _ := ioutil.ReadAll(os.Stdin)
    proto.Unmarshal(data, gen.Request)
    gen.CommandLineParameters(gen.Request.GetParameter())
    gen.WrapTypes()
    gen.SetPackageNames()
    gen.BuildTypeNameMap()

    gen.GeneratePlugin(plugin.New())
    for i := 0;i < len(gen.Response.File);i++ {
        gen.Response.File[i].Name = proto.String(strings.Replace(*gen.Response.File[i].Name, ".pb.go", ".mock.go", -1))
    }
    
    data, _ = proto.Marshal(gen.Response)
    os.Stdout.Write(data)
}
```

By convention of protoc plugins our binary must be prefixed with `protoc-gen-` so in our case we are going to build our binary as `protoc-gen-gogrpcmock`.

按照 protoc 插件的约定，我们的二进制文件必须以 `protoc-gen-` 为前缀，所以在我们的例子中，我们将把我们的二进制文件构建为 `protoc-gen-gogrpcmock`。

Now we can use our plugin to generate the mock implementation of our service:

现在我们可以使用我们的插件来生成我们服务的模拟实现：

```bash
$ protoc -I.--go_out=plugin=grpc=:/out --gogrpcmock_out=:/out src/*.proto
```

## Usage

##  用法

Now that we have our generated output we can use it (along with the gRPC output stub) to build a mock gRPC Server:

现在我们有了生成的输出，我们可以使用它（连同 gRPC 输出存根）来构建一个模拟 gRPC 服务器：

```go
grpcServer := grpc.NewServer()
myservice.RegisterMyServiceServer(s, &myservice.MyServiceMock{})
lis, _ := net.Listen("tcp", 50501)
grpcServer.Serve(lis)
```

## Batteries included 

## 包括电池

If you are looking for a complete solution to generate mock data for your gRPC service(s) then you can take a look at my [gRPC Mock Generator](https://github.com/SafetyCulture/s12-proto/tree/master/protobuf/protoc-gen-gogrpcmock); just install via `go get` and start using the `protoc-gen-gogrpcmock` plugin (FYI: I use the Gogo version of protobuf for this plugin)

如果您正在寻找为您的 gRPC 服务生成模拟数据的完整解决方案，那么您可以查看我的 [gRPC Mock Generator](https://github.com/SafetyCulture/s12-proto/tree/master/protobuf/protoc-gen-gogrpcmock);只需通过 `go get` 安装并开始使用 `protoc-gen-gogrpcmock` 插件（仅供参考：我使用 Gogo 版本的 protobuf 这个插件)

```bash
$ go get -u github.com/SafetyCulture/s12-proto/protobuf/protoc-gen-gogrpcmock
$ protoc -I.--gogo_out=plugin=grpc=:/out --gogrpcmock_out=:/out src/*.proto
```


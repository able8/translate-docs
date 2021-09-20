# Mocking gRPC in Go

One of the huge benefits of using gRPC is the ability to  autogenerate our client and server stubs from the protocol buffer  definitions.

In the same way, we can we generate our own code by building a plugin for the protocol buffer compiler (protoc).

The plugin we're creating is going to auto-generate gRPC response  messages so that we can build a mock gRPC server in Go (Golang).

## Goal

Given the following `proto` file:

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

```go
func (m *MyServiceMock) Get(context.Context, *Request) (*Response, error) {
    return Response{Id: "f267332d-0d9a-4220-b055-63b661b600db"}, nil
}
```

## Creating a protoc Go plugin

All we have to do is implement the `Plugin` interface defined in `protobuf/protoc-gen-go/generator`:

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

Now that we have satisfied our Plugin interface we can implement our mock services:

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

## Building our plugin

All that is left to do is wire up our plugin (errors ignored for brevity):

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
    for i := 0; i < len(gen.Response.File); i++ {
		gen.Response.File[i].Name = proto.String(strings.Replace(*gen.Response.File[i].Name, ".pb.go", ".mock.go", -1))
    }
    
    data, _ = proto.Marshal(gen.Response)
    os.Stdout.Write(data)
}
```

By convention of protoc plugins our binary must be prefixed with `protoc-gen-` so in our case we are going to build our binary as `protoc-gen-gogrpcmock`.

Now we can use our plugin to generate the mock implementation of our service:

```bash
$ protoc -I. --go_out=plugin=grpc=:/out --gogrpcmock_out=:/out src/*.proto
```

## Usage

Now that we have our generated output we can use it (along with the gRPC output stub) to build a mock gRPC Server:

```go
grpcServer := grpc.NewServer()
myservice.RegisterMyServiceServer(s, &myservice.MyServiceMock{})
lis, _ := net.Listen("tcp", 50501)
grpcServer.Serve(lis)
```

## Batteries included

If you are looking for a complete solution to generate mock data for your gRPC service(s) then you can take a look at my [gRPC Mock Generator](https://github.com/SafetyCulture/s12-proto/tree/master/protobuf/protoc-gen-gogrpcmock); just install via `go get` and start using the `protoc-gen-gogrpcmock` plugin (FYI: I use the Gogo version of protobuf for this plugin)

```bash
$ go get -u github.com/SafetyCulture/s12-proto/protobuf/protoc-gen-gogrpcmock
$ protoc -I. --gogo_out=plugin=grpc=:/out --gogrpcmock_out=:/out src/*.proto
```

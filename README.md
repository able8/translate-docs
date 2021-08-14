# translate-docs

使用 [Google 翻译 API](https://translate.google.cn/)，将输入文档按段落生成 英文 - 中文 对照翻译。

有时阅读英文文档，理解不够准确，使用中英文对照理解更准确。在改进翻译中，加深理解。

### 如何使用

```
./translate_darwin_amd64 -f example-input.md

```

Output:
```
2021/08/14 23:31:06 Input file is "example-input.md"
2021/08/14 23:31:06 Translating ...
2021/08/14 23:31:06 Done. Generated output file:  tr-example-input.md
```

### go build 生成 Mac、linux、Windows 平台可执行文件

```sh
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o translate_darwin_amd64 translate.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o translate_linux_amd64 translate.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o translate_windows_amd64.exe translate.go
```

## docs

1. [Handling Client Requests Properly with Kubernetes 使用 Kubernetes 正确处理客户端请求](docs/tr-1-Handling-Client-Requests-Properly-with-Kubernetes.md)
1. [Zero Downtime Server Updates For Your Kubernetes Cluster Kubernetes 集群的零停机服务器更新](docs/tr-2-zero-downtime-server-updates-for-your-kubernetes-cluster.md)
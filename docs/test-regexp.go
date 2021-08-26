package main

import (
	"fmt"
	"regexp"
)

func main() {
	// str := "[负载平衡](https://docs.cilium.io/en/stable/intro/#load-balancing）用于容器之间和外部服务的流量"
	str := "[这个代码片段](https://github.com/eliben/code-for-blog/blob/master/2020/go-fake-stdio/snippets/redirect -cgo-stdout.go）演 "
	str = regexp.MustCompile(`]\(http(.*?) (.*?) `).ReplaceAllString(str, "]($1$2")
	fmt.Println(str)
}

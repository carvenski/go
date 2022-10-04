package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
)

// go mod vendor 拷贝vendor到当前项目目录
// go build -mod=vendor 使用当前项目的vendor来编译
// 可以修改以下文件来增加自定义功能
// vendor/github.com/elazarl/goproxy/proxy.go  112行
func main() {
    proxy := goproxy.NewProxyHttpServer()
    //proxy.Verbose = true
    port := ":8002"
    log.Println("proxy start at:", port)
    log.Fatal(http.ListenAndServe(port, proxy))
}

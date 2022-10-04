package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"

    "cpic_http_proxy/https"
)

func main() {
    // start https server at localhost:18002
    go https.StartServer()

    // start http proxy at :8002
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    port := ":8002"
    log.Println("proxy start at:", port)
    log.Fatal(http.ListenAndServe(port, proxy))
}

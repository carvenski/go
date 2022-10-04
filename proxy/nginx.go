package main

import (
       "log"
        "net/http"
        "net/http/httputil"
        "net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
        backend, _ := url.Parse("http://localhost:8080")
        proxy := httputil.NewSingleHostReverseProxy(backend)
        proxy.ServeHTTP(w, r)
}

// golang实现nginx的负载均衡和静态文件的功能,可简单代替nginx使用.
func main() {
        // 静态资源
        http.Handle("/bimp-ui/", http.StripPrefix("/bimp-ui/",
                http.FileServer(http.Dir("dist"))))
        // 动态请求
        http.HandleFunc("/bimp/", Handler)
        log.Println("nginx started.")
        err := http.ListenAndServe(":37798", nil)
        if err != nil {
                log.Fatal("error")
        }
}


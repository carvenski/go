package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Handler struct {
	backends []string
}

func (this *Handler) getBackend() string {
    // random select
	return this.backends[rand.Intn(len(this.backends))]
}

// 定制reverse proxy功能可修改文件 $GOROOT/src/net/http/httputil/reverseproxy.go
func (this *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend, err := url.Parse(this.getBackend())
    //log.Println("proxy to backend=", backend)
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(backend)
	r.Host = backend.Host
	proxy.ServeHTTP(w, r)
	//log.Println(r.backendAddr + r.Method + r.URL.String() + r.Proto + r.UserAgent())
}

func main() {
	bind := ":37799"
	backends := []string{
		"http://21.50.131.33:37798",
		"http://21.50.131.34:37798",
		"http://21.50.131.35:37798",
	}
	log.Printf("Proxy Listen on %s", bind)
	log.Printf("Proxy Forward to %s", backends)
	handler := &Handler{backends: backends}
	err := http.ListenAndServe(bind, handler)
	if err != nil {
		log.Fatalln("err: ", err)
	}
}

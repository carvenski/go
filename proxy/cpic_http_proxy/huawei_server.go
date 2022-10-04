package main

import (
        "log"
        "net/http"
        "io/ioutil"
)

var client *http.Client = &http.Client{}

type Handler struct {
}

func (this *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

    // do original req
    origin_url := req.Header.Get("origin_url")
    proxyReq, err := http.NewRequest(req.Method, origin_url, req.Body)
    proxyReq.Header = req.Header
    resp, err := client.Do(proxyReq)
    if err != nil {
        log.Println("origin resp err=", err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    //log.Println("originla resp=", string(body), err)
    for k, _ := range resp.Header {
        if (k == "Proxy-Connection"){
            continue;
        }
        w.Header().Add(k, resp.Header.Get(k))
    }
    w.Header().Add("zzz", "xyz")
    w.Write(body)
    log.Println("| origin URL=", origin_url)
    return
}

func main() {
        bind := ":58888"
        log.Printf("Proxy Listen on %s", bind)
        handler := &Handler{}
        err := http.ListenAndServe(bind, handler)
        if err != nil {
                log.Fatalln("err: ", err)
        }
}

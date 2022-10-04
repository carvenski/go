package https

import (
	"io"
	"log"
	"net/http"
	"io/ioutil"
)

var client *http.Client = &http.Client{}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	// https CONNECT req
	if req.Method == "CONNECT" {
		io.WriteString(w, "HTTP/1.1 200 Connection established\r\n\r\n")
		return
	}

	// forward http req to HuaWeiCloud Http Server
	huawei_url := "http://123.60.89.227:58888"
	origin_url := "https://"+ req.Host + req.URL.String()
    proxyReq, err := http.NewRequest(req.Method, huawei_url, req.Body)
    proxyReq.Header = req.Header
    proxyReq.Header.Set("origin_url", origin_url)
	resp, err := client.Do(proxyReq)
    if err != nil {
        log.Println("HuaWeiCloud resp err=", err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    // log.Println("Resp=\n", string(body), err)
    // log.Println("Header=", resp.Header)    
    for k, _ := range resp.Header {
    	w.Header().Add(k, resp.Header.Get(k))
    }
    w.Write(body)
    return
}

func StartServer() {
	port := ":18002"
	http.HandleFunc("/", HelloServer)
	log.Println("https server listen at", port)
	err := http.ListenAndServeTLS(port, "https/server.crt", "https/server.key", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

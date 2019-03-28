package main

import (
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is a https server.\n"))
	log.Println("there is 1 user been here.")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Println("https server listen at :50000...")
	//生成证书和私钥: openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -days 3650
	//(注意:公钥和证书是2个东西,证书中包含公钥和其他组织等信息,然后发给用户浏览器)
	err := http.ListenAndServeTLS(":50000", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

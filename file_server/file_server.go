package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/bimp-ui/", http.StripPrefix("/bimp-ui/", 
		http.FileServer(http.Dir("/tmp"))))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"io"
	"log"
	"fmt"
	"net/http"

	// use pprof+graphviz: (apt install graphviz)
	// go tool pprof http://ip:port/debug/pprof/profile           cpu
	// go tool pprof http://ip:port/debug/pprof/heap              mem
	// go tool pprof http://ip:port/debug/pprof/block             block
	// pprof command: top/web
	_"net/http/pprof" 
)

var PORT = ":50000"

func hello(w http.ResponseWriter, r *http.Request) {
	block1()
	block2()	
	io.WriteString(w, "hello, world!\n")
}


func block1() {
	fib1(40)
}

func block2() {
	fib2(40)
}

func fib1(n int) int {
	if (n <= 2) {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	if (n <= 2) {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

func main() {
	http.HandleFunc("/", hello)
	log.Println(fmt.Sprintf("HTTP Server Started %s", PORT))
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("HTTP Server Error:", err)
	}
}




package main

import (
	"time"
	"log"
	"net/http"

	// use pprof+graphviz: (apt install graphviz)
	// go tool pprof http://ip:port/debug/pprof/profile           cpu
	// go tool pprof http://ip:port/debug/pprof/heap              mem
	// go tool pprof http://ip:port/debug/pprof/block             block
	// pprof command: top/web
	_ "net/http/pprof"
)

func block1() {
	for {
		time.Sleep(time.Second*1)
		fib1(1000)
	}
}

func block2() {
	for {
		time.Sleep(time.Second*1)
		fib2(1000)
	}
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
	log.Println("main started")
	go block1()
	go block2()
    
    // start a http server to get pprof file
    go log.Println(http.ListenAndServe(":50000", nil))
    select{}
}




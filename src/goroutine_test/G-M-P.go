package main

import (
	"time"
	"fmt"
	"runtime"
)

//need use slice here, because slice is reference but array not.
var x = []int{0}  

func main() {
	// *******************************************************************************
	// G-M-P: G/goroutines => M/threads => P/processors
	// *******************************************************************************
	// if =1, use only 1 processor for threads, no real multi threads race here !
	// if >1, use multi processors for threads, so multi threads race now !
	var num = 1
	// var num = 4
	runtime.GOMAXPROCS(num) 
	// *******************************************************************************

	fmt.Println("before:")
	fmt.Println(x)
	for i := 0; i < 1000; i++ {
		go testGorountine(x)		
	}
	time.Sleep(5*time.Second)
	fmt.Println("after:")
	fmt.Println(x)  // 1000 or random <= 1000 !!
}

func testGorountine(a []int) {
	a[0] += 1
}



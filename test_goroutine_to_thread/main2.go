package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("main start")

	for i := 0; i < 100; i++ {
		go io_func()
	}

	time.Sleep(time.Duration(120) * time.Second)
	fmt.Println("main end")

}


func io_func() {
	// 看书上说:系统调用等操作,go的一个协程会占用一个线程.这里测试并没有测到这个结果? 是测试用例写的不对还是因为top看不到内核态线程?
	fmt.Println("1 io func start")

	//系统调用阻塞型
    blo := make([]byte, 100)
    os.Stdin.Read(blo)

	fmt.Println("1 io func end")
}


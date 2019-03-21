package main

import (
	"fmt"
    "os"
    "time"
)

func main() {
	fmt.Println("main start...")

	for i := 0; i < 200; i++ {
		go io_func()
	}

	time.Sleep(time.Duration(60) * time.Second)
	fmt.Println("main end...")
}


func io_func() {
	// 看书上说:系统调用等3种操作,go的协程会占用一个线程.这里测试发现线程数的确和协程数成线性增长...好在只是系统调用等3种情形下.
	fmt.Println("1 io func start")

    // syscall操作,读取控制台输入,          这个好像不是syscall操作?那文件io是吗?测试的结果是8个线程. 到底什么才是系统调用操作??
    var buffer [512]byte
    n, err := os.Stdin.Read(buffer[:])
    if err != nil {
         fmt.Println("read error:", err)
         return
    }
    fmt.Println("count:", n, ", msg:", string(buffer[:]))

	fmt.Println("1 io func end")
}


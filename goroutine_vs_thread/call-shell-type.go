package main

import (
	"fmt"
    "os/exec"
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

    //调用外部shell执行命令
    command := "curl --connect-timeout 10 http://google.com"
    cmd := exec.Command("bash", "-c", command)
    out, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out)

	fmt.Println("1 io func end")
}


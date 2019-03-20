package main

import (
	"fmt"
    "os/exec"
	"time"
    "runtime/debug"
)

func main() {
	fmt.Println("main start")

    //设置go程序最大能够使用的线程数量
    debug.SetMaxThreads(100)

	for i := 0; i < 100; i++ {
		go io_func()
	}

	time.Sleep(time.Duration(120) * time.Second)
	fmt.Println("main end")

}


func io_func() {
	// 看书上说:系统调用等操作,go的一个协程会占用一个线程.这里测试发现线程数的确和协程数成线性增长...好在只是系统调用等3种情形下.
	fmt.Println("1 io func start")


    //调用shell
    c := "curl --connect-timeout 60 http://google.com"
    cmd := exec.Command("bash", "-c", c)
    out, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out)

	fmt.Println("1 io func end")
}


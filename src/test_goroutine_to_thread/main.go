package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	f := func() {}
	fmt.Println("main args:", os.Args)
	if os.Args[1] == "cpu" {
		f = cpu_func
	} else if os.Args[1] == "io" {
		f = io_func
	} else {
		fmt.Println("please choose from cpu/io")
		os.Exit(1)
	}

	for i := 0; i < 100; i++ {
		go f()
	}

	time.Sleep(time.Duration(120) * time.Second)
	fmt.Println("end")

}

func cpu_func() {
	// cpu型的是一个线程可以切换运行N个协程,有个屁用啊... M:N
	fmt.Println("1 cpu func start")
	time.Sleep(time.Duration(60) * time.Second)
	fmt.Println("1 cpu func over")
}

func io_func() {
	// 看书上说:网络请求/系统调用等io操作,go会使用一个线程对应一个协程.和python一样. 1:1 但实测并非如此...??
	fmt.Println("1 io func start")

	c := "curl --connect-timeout 60 http://google.com"
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	fmt.Println("1 io func end")
}

// 测试结果不是1:1,仍然是M:N的?? go一直是起了10个左右线程?? io型的go也并没有线性增加线程数量??

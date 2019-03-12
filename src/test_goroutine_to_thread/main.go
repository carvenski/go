package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"os/exec"
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

	for i := 0; i < 1000; i++ {
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

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)[50:100])

	/*
		c := "curl --connect-timeout 60 http://google.com"
		cmd := exec.Command("bash", "-c", c)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
	*/

	fmt.Println("1 io func end")
}

// 测试结果不是1:1,仍然是M:N的?? go一直是起了10个左右线程?? 网络请求的io型的go也并没有线性增加线程数量??
// 这么看的话,go协程是真的M:N到线程了!? 即使是本例中的1000个网络请求这样的操作,也只使用了10个线程...擦,真的比python协程好??

// 或许,无论怎样,go的协程之于python的协程,最大的好处至少是: go团队替你打包了所有底层调度映射细节操作,你只需要关心产品的并发业务开发即可！
// 至少在并发程序编写这块,go的确大大的提高了开发效率,提高了生产效率,这就是它的优点.

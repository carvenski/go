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
	// 看书上说:系统调用等操作,go的一个协程会占用一个线程.这里测试并没有测到这个结果? 是测试用例写的不对还是因为top看不到内核态线程?
	fmt.Println("1 io func start")

	/* 发起http网络请求
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
	*/

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

// 无论go还是python,http网络请求类的io操作都是可以做到M:1的!! 而http网络请求类操作就是最常见的协程应用场景.

// 测试结果不是1:1,仍然是M:N的.go一直是起了10个左右线程.http网络请求的io操作go并没有线性增加线程数量.这点python也做到了,tornado/gevent的效果
// 的确,go协程是真的M:N到线程了.本例中的1000个http网络请求这样的操作,也只使用了10个线程...

// 无论怎样,go的协程之于python的协程,最大的好处至少是: go团队替你打包了所有底层调度映射细节操作,你只需要关心产品的并发业务开发即可！
// 至少在并发程序编写这块,go的确大大的提高了开发效率,提高了生产效率,这就是它的优点.在并发需求的程序中应当优先采用.

// 实际上,python的io型也不是1:1,也可以是M:1 => 在tornado和gevent中,N个发起http网络请求的协程也是只使用了一个单线程的.只是有些比如访问mysql的
// 操作,本质上是http访问,但是却没有人使用AsyncHttpClient来为mysql写这个异步client(需要按照mysql协议写http数据交互解析这些事),
// 需要自己写.这个就不方便了...效率低,造轮子.而go里面这些都有了,拿来就用,所以说生产效率高.

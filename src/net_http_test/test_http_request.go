package main

import (
	"fmt"
	"io/ioutil"
	"net/http" // how it is build by using go's Gevent-like non-block socket ?
	"time"
)

func main() {
	fmt.Println("start-----------")
	for i := 0; i < 10; i++ {
		// httpGet(i)  //if you don't add go, it will just execute by order, no goroutines no concurrency !
		go httpGet(i) // 所以说同一个go函数,既可以写普通顺序阻塞代码,也可以写协程异步非阻塞代码...就看前面加不加go
	}
	time.Sleep(time.Second * 10)
	fmt.Println("end-----------")
}

func httpGet(i int) {
	fmt.Println("in goroutine---------", i)
	//when http request(io block), this goroutine yield cpu to let other goroutine run.
	//actually go has a non-block Gevent-like socket essencially. http request use the socket essencially.
	//==============================================================================
	//    mysql query => http request => go socket(Gevent-like non-block socket)
	//==============================================================================
	resp, err := http.Get("http://www.github.com/yxzoro/go")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body)[50:100])
	fmt.Println("out goroutine---------", i)
}

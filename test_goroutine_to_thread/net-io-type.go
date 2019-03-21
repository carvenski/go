package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
    "time"
)

func main() {
	fmt.Println("start...")

	for i := 0; i < 100; i++ {
		go io_func()
	}

	time.Sleep(time.Duration(60) * time.Second)
	fmt.Println("end...")

}

func io_func() {
	fmt.Println("1 io func start")

	// 发起http网络请求, 网络请求属于被go特殊处理的netpoll操作,它不同于syscall那种操作,是不会导致线程阻塞的.
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
	fmt.Println(string(body)[0:20])

	fmt.Println("1 io func end")
}


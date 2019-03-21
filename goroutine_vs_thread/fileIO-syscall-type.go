package main

import (
	"fmt"
    "time"
    "io/ioutil"
)

func main() {
	fmt.Println("main start...")

	for i := 0; i < 100; i++ {
		go io_func()
	}

	time.Sleep(time.Duration(60) * time.Second)
	fmt.Println("main end...")
}


func io_func() {
	fmt.Println("1 io func start")

    // 文件读写操作,属于系统调用,所以这里看到线程数量在上涨,
    // 当然这里的文件没那么大,所以前面处理完的线程可以给后面的协程继续使用,所以线程数不完全等于协程数,线程不够了才自动创建.
    // 但足以证明,go里面在系统调用等3种操作的时候,协程是1:1占用线程的.
    dat, err := ioutil.ReadFile("./go")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Print(string(dat)[0:50])

	fmt.Println("1 io func end")
}


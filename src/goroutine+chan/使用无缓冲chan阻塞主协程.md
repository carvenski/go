
#### 只需要使用一个chan(无缓冲或缓冲chan都可以)即可在main函数中实现等待所有协程完成的join功能,好方法!
```go
package main

import (
	"fmt"
)

var ch chan int = make(chan int) 
//或者使用缓冲chan也可以: var ch chan int = make(chan int 10)
//区别只在于: 无缓冲的信道是数据一个一个流进流出, 而有缓冲信道可以是数据一批流进,一批流出的...

func t(i int) {
	// do your work here
	fmt.Println("------goroutine woring in----", i)
	ch <- i
}

func main() {
	fmt.Println("----main start-----")
	goroutine_num := 10
	for i := 0; i < goroutine_num; i++ {
		go t(i)
	}

	for i := 0; i < goroutine_num; i++ {
		<-ch //只需要使用一个chan即可在main函数中等待所有协程完成,实现类似进程/线程的join函数功能,是个实用的好方法!
	}
	fmt.Println("----main end-----")
}
```

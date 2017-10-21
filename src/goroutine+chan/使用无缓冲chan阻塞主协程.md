
#### 只需要使用一个无缓冲chan即可在main函数中实现等待所有协程完成的join功能,好方法!
```go
package main

import (
	"fmt"
)

var ch chan int = make(chan int)

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
		<-ch //只需要使用一个chan即可在main函数中等待所有协程完成,好方法!
	}
	fmt.Println("----main end-----")
}
```

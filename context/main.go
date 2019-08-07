package main

import (
	"log"
	"context"
	"time"
)

/*
Context包:
是专门用来简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用
比如有一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine
这样的话， 我们就可以通过Context，来跟踪这些goroutine，并且通过Context来控制他们的目的
这就是Go语言为我们提供的Context，中文可以称之为“上下文”
另外一个实际例子是，在Go服务器程序中，每个请求都会有一个goroutine去处理
然而，处理程序往往还需要创建额外的goroutine去访问后端资源，比如数据库、RPC服务等
由于这些goroutine都是在处理同一个请求，所以它们往往需要访问一些共享的资源，比如用户身份信息、认证token、请求截止时间等
而且如果请求超时或者被取消后,所有的goroutine都应该马上退出并且释放相关的资源,这种情况也需要用Context来为我们取消掉所有goroutine
*/
func main() {
	log.Println("main start")
	// pass sub context to each goroutine, then call cancel() in main to cancel all goroutines
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go test(ctx, i)
	}
	time.Sleep(3*time.Second)
	cancel()
	log.Println("\n\n====> cancel() already called in main !\n")
	time.Sleep(2*time.Second)
	log.Println("main stop")
}

func test(ctx context.Context, n int) {
	log.Printf("goroutine %d start...", n)
	loop:
		for i := 0; i < 10; i++ {
			select {
				// <-ctx.Done() will success after cancel() called
			    case <-ctx.Done():
			    	log.Printf("goroutine %d stop by cancel()...", n)
			    	break loop
			    default:			
					time.Sleep(1*time.Second)
					log.Printf("goroutine %d print %d...", n, i)
			}
		}
}










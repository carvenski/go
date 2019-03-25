package main

import (
	"fmt"
	"runtime"
	"time"
)

func haha() {
	fmt.Print("----sleeping")
	time.Sleep(5 * time.Second)             //这里的time.sleep()显然不是同步阻塞版的,go内部有处理了.
	fmt.Print("----sleeping end")

}


//-----------------------------------------------------------------------------------------------------------------
//可以看到, 5个goroutine 的却使用了 1个thread, 但是运行效果的却是并发的.
//在python,一个同步阻塞的coroutine要用一个thread来跑其同步阻塞的任务的(本质上是背后的多线程)...
//***********************************************************************
// 无论python和go:
//   cpu型的和网络io型的操作,都是可以做到N:1的!
//   但在少数3种情况下的操作(如系统调用等)仍然是要1:1的,底层都是要占用1个线程的.
//***********************************************************************
//------------------------------------------------------------------------------------------------------------------


func main() {
	// we know that:
	//   1 CPU can be used by multi process/thread,
	//   actually multi process/thread each get little time span of CPU to use.

	// in same:
	//   1 thread can be used by multi goroutine,
	//   actually multi goroutine each get little time span of thread(CPU) to use.

	//在go里面,可以把 thread 理解成 CPU ,把 goroutine 理解成 process/thread，
	//多个goroutine被调度轮询使用1个thread的CPU,就像操作系统里多个process/thread被调度轮询使用1个CPU一样. 实现并发效果.
	//（都是采用 <分时间片> 的方法来实现的并发效果...）

	runtime.GOMAXPROCS(1)

	fmt.Print("----main start")

	for i := 0; i < 5; i++ {
		go haha()

	}

	time.Sleep(6 * time.Second)
	fmt.Print("----main end")

}

//output:
// ----main start
// ----sleeping----sleeping----sleeping----sleeping----sleeping
// ----sleeping end----sleeping end----sleeping end----sleeping end----sleeping end
// ----main end




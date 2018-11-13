package main

import (
	"fmt"
	"runtime"
	"time"
)

func haha() {
	fmt.Print("----sleeping")
	time.Sleep(5 * time.Second)
	fmt.Print("----sleeping end")

}


//-----------------------------------------------------------------------------------------------------------------
//可以看到, 5个goroutine 的却只映射/使用了 1个thread, 但是运行效果的却是并发的!!这个就真牛逼了(go实现了真的多协程单线程)...  
//而在python中的话,是每一个coroutine都要使用一个thread来跑其同步阻塞的任务的(本质上还是背后的多线程)...                       
//*****************************************
//  在go里面  goroutine : thread == N ：1 (或者可选 M ：N)
//  在python里面 coroutine : thread == 1 ：1
//*****************************************
//------------------------------------------------------------------------------------------------------------------


func main() {
	// ---------------------------------------------------------------------------------------------------------
	// there is only 1 thread in backfround here, so,
	// how can 5 goroutines run in same time with only 1 thread ??

	//my understanding:
	// we know that:
	//   1 CPU can be used by multi process/thread,
	//   actually multi process/thread each get little time span of CPU to use.

	// in same:
	//   1 thread can be used by multi goroutine,
	//   actually multi goroutine each get little time span of thread(CPU) to use.

	//在go里面,可以把 thread 理解成 CPU ,把 goroutine 理解成 process/thread，
	//多个goroutine被调度轮询使用1个thread的CPU,就像操作系统里多个process/thread被调度轮询使用1个CPU一样. 实现并发效果!!
	//（都是采用 <分时间片> 的方法来实现的并发效果...）
	// is it right ??

	runtime.GOMAXPROCS(1)
	// --------------------------------------------------------------------------------------------------------

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




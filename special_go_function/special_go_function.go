package main

import (
	"fmt"
	"time"
)

/* **************************************************************************
在python中:
协程中的函数必须是异步非阻塞的,而不能是同步阻塞的 !
同步阻塞的函数改造成异步非阻塞的需要使用多线程等异步操作包装下才行.
而这2种函数显然不是同一个.

但在go里面:
一个函数即是同步阻塞的又是异步非阻塞的!同一个函数既能当同步的又能当异步的使用.
go底层自动处理了 => go里面func前加了go关键字之后就会自动转成异步的使用了!!
go底层把这些python里的异步函数/协程/loop等技术细节都封装好了,你只管方便的使用...
这个真TM太方便了!!绝对是go语言使用协程的一大简洁方便之处!!

=============================================================================
(同步转异步一般都是采用线程池的方法来实现的异步操作效果,
go底层采用的是linux内核的线程池做异步操作,很高效的.)

其实可以这么理解golang中的go关键字:
  1.go中的函数本身还是正常的同步阻塞函数,
    所以正常写同步代码时它的效果也就是正常的同步阻塞效果,
  2.但当你在函数前面加上go关键字后,go会自动的把该函数扔到底层的线程池中去跑,
    所以自然的有了异步非阻塞的并发运行效果.
    如果你不需要等待该函数的结果,那么主协程和其他协程都并发运行而已,
    如果你需要等待该函数的结果,就使用channel在这等待结果即可.

而且可以这么理解go语言的goroutine+channel并发编程模型:
  直接想象成你在使用java/python写代码,有一些并发需求,
  然后你采用'多线程+生产-消费者模式'来写代码实现并发效果而已!
  只是go的底层使用linux内核的线程池,比一般java/python的多线程更高效而已.

本质上:
在其他语言里的[使用多线程+P-C模型来实现异步/并发效果],
在go里,变成了[使用多协程(底层还是linux内核多线程)+P-C模型来实现异步/并发效果]
只是所谓的goroutine协程(M:N映射到linux内核的线程)比一般的线程更轻量级更高效.
************************************************************************** */

//例如:
//同一个x函数,既可以作为同步阻塞的函数正常使用,
//又可以作为异步非阻塞的函数在协程中使用.
func x(a string) {
	fmt.Print("====" + a + " start " + " ====\n")
	//即相当于python的time.sleep又相当于gevent.sleep或tornado.gen.sleep !!
	time.Sleep(1 * time.Second)
	fmt.Print("====" + a + " end" + " ====\n")
}

func sync_test() {
	fmt.Print("==== sync_test start ====\n")

	// 同步版本:
	time.Sleep(1 * time.Second)
	x("A")
	x("B")
	x("C")

	fmt.Print("==== sync_test end ====\n")
	time.Sleep(4 * time.Second)
}

func async_test() {
	fmt.Print("==== async_test start ====\n")

	// 异步版本:
	//(函数前加了go之后,会被扔到linux内核里的线程池里去异步并发运行)
	//(如果需要等待结果,就使用channel即可)
	go time.Sleep(1 * time.Second)
	go x("A")
	go x("B")
	go x("C")

	fmt.Print("==== async_test end ====\n")
	time.Sleep(4 * time.Second)
}

func main() {
	fmt.Print("TEST sync:\n")
	sync_test()
	fmt.Print("\n\n")
	fmt.Print("TEST async:\n")
	async_test()
}

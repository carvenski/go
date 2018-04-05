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
(go底层自动处理了 => go里面func前加了go关键字之后就会自动转成异步的使用了!!)

        **** 这个真TM太方便了!!绝对是go语言使用协程的一大简洁方便之处!! ****
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

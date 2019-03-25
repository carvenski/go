package main

import (
	"fmt"
)

type A interface {
	Haha()
}

type B struct {
}

type C struct {
}

func (b B) Haha(x int) {
	fmt.Println(x)
}

func (c *C) Haha(x int) { //结构体相关的/绑定的方法 => 统一使用 func (t *T) f() 指针写法
	fmt.Println(x)
}

func main() {
	a1 := B{}
	a1.Haha(1)
	a2 := &B{}
	a2.Haha(2)

	a3 := C{} //
	a3.Haha(3)
	a4 := &C{}
	a4.Haha(4)
}

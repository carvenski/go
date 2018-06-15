package main

import (
	"fmt"
)

type X struct {
	Id   int
	Name string
}

func (x X) test1(a int) {
	fmt.Println(x.Id)
}

func (x *X) test2(a int) {
	fmt.Println(x.Id)
}

func main() {
	x1 := X{Id: 10, Name: "haha1"}
	x2 := &X{Id: 11, Name: "haha2"}
	x1.test1(99) // 结构体变量会自动识别是否是指针并自动转换
	x1.test2(99)
	x2.test1(99)
	x2.test2(100)
}

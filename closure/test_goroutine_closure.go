package main

import (
	"fmt"
	"time"
)

var x = "X"

func main(){
	fmt.Println("== main start")
	a := 1
	b := []int{1,2,3}
	c := map[int]int{1:1, 2:2}
	// 闭包可以引用到函数外部的变量 不会报错 后面等待gc来回收变量即可。
	// 而且 => 闭包引用的还是指针 !!! 而不是值拷贝 !!!
        // 这里外部的 abcx 的值都会被修改。。
	go func(){
		for {
			time.Sleep(time.Second*1)
			// 这里的abcx都是指针引用
			fmt.Printf("a=%v b=%v c=%v x=%v \n", a, b, c, x)
			a = a + 1
			b = append(b, 4)
			c[3] = 3
			x = x + "X"
		}
	}()
	for {
		time.Sleep(time.Second*3)
		fmt.Println("main")
	}
}

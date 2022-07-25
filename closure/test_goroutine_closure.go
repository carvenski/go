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
  
	// 闭包可以引用到函数外部的变量 不会报错 后面等待gc来回收变量即可
	go func(){
		for {
			time.Sleep(time.Second*1)
			fmt.Printf("a=%v b=%v c=%v x=%v \n", a, b, c, x)
		}
	}()
  
	for {
		time.Sleep(time.Second*3)
		fmt.Println("main")
	}
}

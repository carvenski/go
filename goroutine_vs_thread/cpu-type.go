package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start...")

	for i := 0; i < 100; i++ {
		go cpu_func()
	}

	time.Sleep(time.Duration(60) * time.Second)
	fmt.Println("end...")
}

func cpu_func() {
	// cpu型的一个线程可以切换运行N个协程,可有个屁用...还不如顺序调用...
	fmt.Println("1 cpu func start")
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println("1 cpu func over")
}


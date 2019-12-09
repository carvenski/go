package main

import (
	"log"
	"runtime/debug" // 使用debug.Stack()打印异常调用栈
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("error: %v", err)
			log.Println("---------error stack----------")
			log.Println(string(debug.Stack()))
			log.Println("-----------------------------")
		}
	}()

	log.Println("start...")
	foo()
	log.Println("end...")
}

func foo() {
	bar()
}

func bar() {
	panic("Panic Here")
}

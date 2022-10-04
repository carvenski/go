/*
go就是模仿了C的异常处理：
做法就是靠函数的返回值来判断是否出错了
C是判断返回值是否是-1, go判断返回的err是否是空
写起来就是一堆if判断, 真的不如写try方便。。
所以后来的java等语言才加入了try啊写起来更简洁方便
go还倒回去了这个设计是真不咋的。
可能是意识到了这一点,又加了个defer recover来模拟try。
*/
package main

import (
	"log"
)

func main() {
	// defer用法1: defer => 实现函数整体try的效果:
	// 开头第一个defer用来捕获函数中的任何位置的异常,相当于在函数整体上套个try
	// 捕获未知异常,报错并返回
	defer func() {
		if err := recover(); err != nil {
			log.Println("catch exception")
			log.Println("function stop and return because of unexpected exception !")
		}
	}()

	log.Println("function main start......")
	log.Println("step 1")

	func() {
		// defer用法2: defer + 匿名函数 => 实现局部try的效果:
		// 局部的defer用来捕获函数中的某特定几句可能发生的的异常,相当于在函数局部几句代码上套个try
		// 捕获局部异常,处理后继续执行代码
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception ")
				log.Println("continue run step 3 after handle exception !")
			}
		}()

		log.Println("step 2")
		println("exception happen at step 3")
		panic("")
		log.Println("step 3")
	}()

	log.Println("step 4")
	println("exception happen at step 5")
	panic("")
	log.Println("step 5")
	log.Println("step 6")
	log.Println("function main finished......")

}

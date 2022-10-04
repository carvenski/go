package main

import (
	"init_test/lib" // 因为这里先导入了lib包,所以先执行了lib包里的init函数
	"log"
)

func init() { // 然后才执行main包里面的init函数
	log.Println("main init")
}

func main() { // 最后才是main函数的执行
	log.Println("main() func")
	lib.Test()
}

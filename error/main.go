package main

import (
	"fmt"
)

func main() {
	fmt.Println("test error")
	err := test(0)
	fmt.Println(err)
	fmt.Println("end")
}

func test(x int) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("ERROR: %v", p)  // 最外层的异常捕捉
		}
	}()
	_ = 1/x
	return
}




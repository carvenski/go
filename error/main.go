package main

import (
	"fmt"
    "errors"
)

func main() {
	fmt.Println("test error")
	err1, err2 := test(0)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println("end")
}

func test(x int) (err1,err2 error) {
	defer func() {
		if p := recover(); p != nil {
			err1 = fmt.Errorf("ERROR: %v", p)  // 最外层的异常捕捉
			err2 = errors.New("thorws an error")
		}
	}()
	_ = 1/x
	return
}




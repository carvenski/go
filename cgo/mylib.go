package main

import (
	"C"
)

//export foo
func foo(a int64, b int64) int64 {
	var sum int64 = 0
	for i := int64(0); i < 1000000; i++ {
		sum += i
	}
	return sum
}

func main() {}

// 编译.so库文件: go build -buildmode=c-shared -o mylib.so mylib.go

package main

import (
	"fmt"
	"os"
)

func main() {
	// never ignore the err !! always check the err and handle the err situation !!
	f, err := os.Open("haha.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("file: ", f)

	/*  if you ignore err, there will be unexpected exception happens in code...
	f, _ := os.Open("haha.txt")
	fmt.Println("file: ", f)
	*/
}

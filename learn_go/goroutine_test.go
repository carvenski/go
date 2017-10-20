
package main

import (
    "fmt"
)

var quit chan bool = make(chan bool)

func main() {
    go testGorountine()
    <-quit
}

func testGorountine() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello world!")
    }
    quit <- true
}



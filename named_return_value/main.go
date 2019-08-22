package main

import (
    "log"
    "errors"
)

func test(i int) (v int, e error) {
    defer func(){
        if p := recover(); p != nil {
            v, e = 0, p.(error)  // return value in defer: use "named return value" grammar...shit
        }
    }()

    if i == 0 {
        log.Println("panic")
        panic( errors.New("panic 0 error") )
    } else {
        log.Println("no panic")
        return 1, nil
    }
}

func main() {
    log.Println( test(0) )
    log.Println( "" )
    log.Println( test(1) )
}





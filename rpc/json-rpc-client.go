package main

import (
    "fmt"
    "net/rpc/jsonrpc"
    "os"
)

type Args struct {
    A, B int
}

type quo struct {
    Quo, Rem int
}

func main() {
    service := "127.0.0.1:35000"
    client, err := jsonrpc.Dial("tcp", service)
    if err != nil {
        fmt.Println("dial error:", err)
        os.Exit(1)
    }

    args := Args{10, 2}
    var reply int
    err = client.Call("Arith.Muliply", args, &reply)
    if err != nil {
        fmt.Println("Arith.Muliply call error:", err)
        os.Exit(1)
    }
    fmt.Println("the arith.mutiply json rpc call result is :", reply)

    var quto quo
    err = client.Call("Arith.Divide", args, &quto)
    if err != nil {
        fmt.Println("Arith.Divide call error:", err)
        os.Exit(1)
    }
    fmt.Println("the arith.devide json rpc call result is :", quto.Quo, quto.Rem)

}





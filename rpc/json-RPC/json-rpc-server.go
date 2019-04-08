package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "os"
)

type Args struct {
    A, B int
}

func checkError(err error) {
    if err != nil {
        fmt.Fprint(os.Stderr, "Usage: %s", err.Error())
        os.Exit(1)
    }
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (t *Arith) Muliply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A * args.B
    quo.Rem = args.A / args.B
    return nil
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)

    fmt.Println("json rpc server listen at :35000")
    tcpAddr, err := net.ResolveTCPAddr("tcp", ":35000")
    checkError(err)

    Listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        conn, err := Listener.Accept()
        fmt.Println("there is one connection here.")
        if err != nil {
            fmt.Fprint(os.Stderr, "accept err: %s", err.Error())
            continue
        }
        jsonrpc.ServeConn(conn)
    }
}





package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

type Arith int

// 这里rpc方法的参数都要全用指针 ?
func (t *Arith) Multiply(args *Args, reply *int) error {
    log.Println("finish one Multiply rpc call.")
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith) // 在rpc server中注册提供方法的对象
	rpc.HandleHTTP()    // 支持通过tcp或http方式的rpc实现
	l, e := net.Listen("tcp", ":35000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Println("rpc server listen at :35000")
	http.Serve(l, nil)
}

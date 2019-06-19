package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

/*
Gob简介
Gob 是 Go 的一个序列化数据结构的编码解码工具，在 Go 标准库中内置encoding/gob包
以供使用。一个数据结构使用 Gob 进行序列化之后，能够用于网络传输。与 JSON 或 XML 这种
基于文本描述的数据交换语言不同，Gob 是二进制编码的数据流，并且 Gob 流是可以自解释的，
它在保证高效率的同时，也具备完整的表达能力。
作为针对 Go 的数据结构进行编码和解码的专用序列化方法，这意味着 Gob 无法跨语言使
用。在 Go 的net/rpc包中，传输数据所需要用到的编码解码器，默认就是 Gob。由于 Gob 仅局
限于使用 Go 语言开发的程序，这意味着我们只能用 Go 的 RPC 实现进程间通信。然而，大多数
时候，我们用 Go 编写的 RPC 服务端（或客户端），可能更希望它是通用的，与语言无关的，无
论是Python 、 Java 或其他编程语言实现的 RPC 客户端，均可与之通信。

自定义rpc传输时使用的编码解码
Go 的net/rpc很灵活，
它在数据传输前后实现了编码解码器的接口定义可以自定义数据的传输方式以及RPC服务端和客户端之间的交互行为。
RPC 提供的编码解码器接口如下：
type ClientCodec interface { 
   WriteRequest(*Request, interface{}) error 
   ReadResponseHeader(*Response) error 
   ReadResponseBody(interface{}) error 
   Close() error 
} 
type ServerCodec interface { 
   ReadRequestHeader(*Request) error 
   ReadRequestBody(interface{}) error 
   WriteResponse(*Response, interface{}) error 
   Close() error 
}
接口ClientCodec定义了 RPC 客户端如何在一个 RPC 会话中发送请求和读取响应。客户端程
序通过 WriteRequest() 方法将一个请求写入到 RPC 连接中，并通过 ReadResponseHeader()
和 ReadResponseBody() 读取服务端的响应信息。当整个过程执行完毕后，再通过 Close() 方
法来关闭该连接。
接口ServerCodec定义了 RPC 服务端如何在一个 RPC 会话中接收请求并发送响应。服务端
程序通过 ReadRequestHeader() 和 ReadRequestBody() 方法从一个 RPC 连接中读取请求
信息，然后再通过 WriteResponse() 方法向该连接中的 RPC 客户端发送响应。当完成该过程
后，通过 Close() 方法来关闭连接。
通过实现上述接口，我们可以自定义数据传输前后的编码解码方式，而不仅仅局限于 Gob。
同样，可以自定义RPC 服务端和客户端的交互行为。

实际上，Go 标准库提供的net/rpc/json包，
就是一套实现了rpc.ClientCodec和rpc.ServerCodec接口的 JSON-RPC 模块。
*/

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

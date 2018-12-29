package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := ":18000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)
	log.Println("Socket Server Started at :18000...")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		checkError(err)
		log.Printf("=> got a conection from: %s", conn.RemoteAddr().String())
		// handle a conn in a goroutine
		go handler(conn)
	}

}

func checkError(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}

func handler(conn net.Conn) {
	log.Printf("=> Reading data from conn: %s", conn.RemoteAddr().String())

	// ************************************************************************************
	// use tmp as a buffer, append all bytes to total
	// you can see here:
	//	  just copy data bytes from network buffer to process buffer...
	// ************************************************************************************
	buf := make([]byte, 0, 1024) // total buffer, capcity=1024 bytes
	tmp := make([]byte, 10)      // tmp buffer, capcity=10 bytes
	for {
		/*
		https://blog.csdn.net/hguisu/article/details/7453390
		我们把一个SOCKET接口设置为非阻塞就是告诉内核(调用set_non_blocking)
		当所请求的I/O操作无法完成时，不要将进程睡眠，而是返回一个错误。
		这样我们的I/O操作函数将不断的测试数据是否已经准备好，如果没有准备好，继续测试，直到数据准备好为止

		把SOCKET设置为非阻塞模式，即通知系统内核：在调用Sockets API时，不要让线程睡眠，而应该让函数立即返回
		在返回时，该函数返回一个错误代码。图所示，一个非阻塞模式套接字多次调用recv()函数的过程。
		前三次调用recv()函数时，内核数据还没有准备好。因此，该函数立即返回WSAEWOULDBLOCK错误代码。
		第四次调用recv()函数时，数据已经准备好，被复制到程序的缓冲区中，recv()函数返回成功，程序开始处理数据

		*******************************************************************************************
		从以上可知: 其实selector/epoll这些东西本质上做的事也就是: 轮训socket而已,你的数据好没好啊?
		从selector的api可以看出:把一组socket传给它,然后它自己阻塞(内部在轮训),哪个好了就通知你,调用回调函数
		*******************************************************************************************

		默认情况下，socket的connect/accept/recv/send都是阻塞读写的
		*/
		log.Println("-> waiting for conn.Read()...")
		n, err := conn.Read(tmp) // Read will block here until data ready(client have sended data).
		if err != nil {
			if err != io.EOF {
				log.Printf("unexpected read error: %s", err)
			} else {
				log.Println("read EOF. byebye.")
				break
			}
		}
		log.Printf("got %d bytes:", n)
		log.Println(tmp)
		log.Printf("(bytes to string: %s)", tmp)
		buf = append(buf, tmp[:n]...)
	}
	log.Println("end...total bytes is:")
	log.Println(buf)

	daytime := time.Now().String()
	conn.Write([]byte("i have send you the time: " + daytime))
	conn.Close()

}

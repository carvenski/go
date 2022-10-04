package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

var bind string = ":13306"
var backends []string = []string{"21.50.131.33:8080"}

func main() {
	log.Println("TCP proxy start at=", bind)
	log.Println("proxy backend=", backends)
	server()
}

func server() {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Println(err)
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("accept client conn error:%v\n", err)
			continue
		}
		log.Println("got client conn from:", conn.RemoteAddr())
		go handle(conn)
	}
}
func handle(sconn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("handle client conn err:", err)
		}
	}()
	defer sconn.Close() // close sconn at last
	ip, ok := getBackend()
	if !ok {
		return
	}
	dconn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Println("backend connect error:", ip, err)
		return
	}
	defer dconn.Close() // close dconn at last
	ExitChan := make(chan bool, 1)
	go func(sconn net.Conn, dconn net.Conn, Exit chan bool) {
		// The Copy() function in Go language is used to copy from
		// the stated source to the destination till
		// either the EOF i.e, end of file is attained on src or an error is thrown.
		// 协程会阻塞在io.Copy这里直到读到EOF或者conn Closed
		// io.Copy里面其实就是一个for循环会持续的从sconn里Read数据并Send到dconn
		err := ioCopy(dconn, sconn, Exit)
		log.Println("client conn closed.", err)
		ExitChan <- true //情况1,如果client端连接主动关闭则走到这里
	}(sconn, dconn, ExitChan)
	go func(sconn net.Conn, dconn net.Conn, Exit chan bool) {
		tnow := time.Now()
		err := ioCopy(sconn, dconn, Exit)
		log.Println("backend conn closed.", err)
		log.Println("resp Time=", time.Since(tnow))
		ExitChan <- true //情况2,如果backend端连接主动关闭则走到这里
	}(sconn, dconn, ExitChan)
	<-ExitChan //无论client还是backend谁先关闭,只要有一方关了就结束.
	log.Println("clear client and backend conn and exit.")
	// 连接关系 client -- middle conn -- backend
}

func getBackend() (string, bool) {
	// random select
	ip := backends[rand.Intn(len(backends))]
	return ip, true
}

// /* io.Copy() demo
func ioCopy(des net.Conn, src net.Conn, closed chan bool) error {
// +-------------------+--------------+---------------------------------------------------+
// |      3 Bytes      |    1 Byte    |                   N Bytes                         |
// +-------------------+--------------+---------------------------------------------------+
// |<= length of msg =>|<= sequence =>|<==================== data =======================>|
// |<============= header ===========>|<==================== body =======================>|
// MySQL 报文格式如上，消息头包含了 
// A) 报文长度，标记当前请求的实际数据长度，以字节为单位；
// B) 序号，为了保证交互时报文的顺序，每次客户端发起请求时，序号值都会从 0 开始计算。
// 消息体用于存放报文的具体数据，长度由消息头中的长度值决定。	
	buffer := make([]byte, 4096)
	for {
		n1, err := src.Read(buffer)
		// 这里需要按照mysql协议来解析数据然后区分读写sql
		// 打印sql前部分
		log.Println("=bytes=", buffer[:4], n1, buffer[:n1])
		log.Println("=string=", string(buffer[:4]), n1, string(buffer[:n1]))
		if err != nil {
			log.Printf("Read err= %v", err)
			closed <- true
			return nil
		}
		_, err = des.Write(buffer[:n1])
		if err != nil {
			log.Printf("Write err= %v", err)
			closed <- true
			return nil
		}

	}
}

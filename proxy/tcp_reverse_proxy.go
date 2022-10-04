package main

import (
	"io"
	"log"
	"net"
    "time"
    "math/rand"
)

var bind string = ":58888"
var backends []string = []string{"localhost:48888", "localhost:38888"}

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
		_, err := io.Copy(dconn, sconn)
        log.Println("client conn closed.", err)
		ExitChan <- true //情况1,如果client端连接主动关闭则走到这里
	}(sconn, dconn, ExitChan)
	go func(sconn net.Conn, dconn net.Conn, Exit chan bool) {
        tnow := time.Now()
		_, err := io.Copy(sconn, dconn)        
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

/* io.Copy() demo
func io.Copy(des net.Conn, src net.Conn, closed chan bool){
    buffer := make([]byte, 4096)
    for {
        n1, err := src.Read(buffer)
        if err != nil {
            closed <- true
            return
        }
        n2, err := des.Write(buffer[:n1])
        if err != nil {
            closed <- true
            return
        }
 
    }
}
*/




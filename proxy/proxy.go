package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		// 默认监听 10086
		port = "10086"
	}
	li, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println("tcp listent err:", err)
		return
	}
	defer li.Close()
	log.Println("http proxy start at: ", port)
	for {
		client, err := li.Accept()
		if err != nil {
			log.Println("server accept err:", err)
			continue
		}
		go handleConn(client)	
	}
}

func handleConn(client net.Conn){
	// 首次需要先读取一次4k内容,从中解析出目的http req url
	buffer := make([]byte, 4096)
	n, err := client.Read(buffer)
	// log.Println("req body: \n", string(buffer[:n]))
    // 解析address
	var method, host, address string
	fmt.Sscanf(string(buffer[:bytes.IndexByte(buffer, '\n')]), "%s%s", &method, &host)
	hostPortUrl, err := url.Parse(host)
	if err != nil {		
		hostPortUrl, _ = url.Parse("//" + host)
	}
	if hostPortUrl.Opaque == "443" {
		address = hostPortUrl.Scheme + ":443"
	} else {
		if strings.Index(hostPortUrl.Host, ":") == -1 {
			address = hostPortUrl.Host + ":80"
		} else {
			address = hostPortUrl.Host
		}
	}
	//解析出了目的host和port，开始进行tcp连接
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
        client.Close() //没连上server也及时关闭client
		return
	}
	// 代理https的原理：client会先发送个connect请求
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		// 先转发第一次解析所读取的数据
		server.Write(buffer[:n])
	}

	// 后面则使用io.Copy()来转发数据即可
	go ioCopy(server, client)
	go ioCopy(client, server)
}

func ioCopy(des net.Conn, src net.Conn){
	// 记得关闭2个连接
    defer des.Close()
	defer src.Close() 
    _, err := io.Copy(des, src)  //io.Copy,无论client或者server谁先断开,另一个也会紧接着断开
    log.Println("conn close:", err)
}


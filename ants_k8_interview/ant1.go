// Q1: 这段代码做什么事情？
// A：并行发消息，收集错误信息
//
// Q2：可以工作么？有存在什么问题么？
// A： 打印出的err信息不能和自己的addr相对应上。
//
// Q3：如果存在问题，改进一下
// A：直接编辑   修改chan类型，存放addr和err,可打印出err和其对应的addr值。
package main 

import (
	"net"
	"fmt"
	"time"
)

func sendMsg(msg, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = fmt.Fprint(conn, msg)
	return err
}

func broadcastMsg(msg string, addrs []string) error {
	errc := make(chan error) // len=0 阻塞chan
	for _, addr := range addrs {
		go func(addr string) {
			errc <- sendMsg(msg, addr) // 并发发消息，并写入err
			fmt.Println("done")
		}(addr)
	}

	for _ = range addrs { // 这里是 循环addrs, 需要读取其addr对应的错误信息
		if err := <-errc; err != nil {		
			// return err   //这里不能是return,否则有一个err函数就退出了
			fmt.Println(err)
		}
	}
	return nil
}

func main() {
	addr := []string{
		"localhost:8080", 
		"http://www.taobao.com",
		"http://www.cpic.com",
		"localhost:22",
	}
	err := broadcastMsg("hi", addr)

	time.Sleep(time.Second)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("everything went fine")
}
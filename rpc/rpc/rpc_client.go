package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("rpc client connecting...")
	client, err := rpc.DialHTTP("tcp", "localhost"+":35000")
	checkError(err)
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	checkError(err)
	log.Printf("test rpc result: %d*%d=%d", args.A, args.B, reply)
}

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "../helloworld"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "neon", Age: 29})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s %d", r.Message, r.Code)

	r2, err := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: "jack", Age: 30})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s %d", r2.Message, r2.Code)
}

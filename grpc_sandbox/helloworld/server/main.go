package main

import (
	"fmt"
	"os"
	"time"
	"context"
	pb "github.com/slankdev/grpc_client/sandbox/helloworld/helloworld"
)

const (
	address = "localhost:9999"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 引数の準備
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// SayHelloメソッドの呼び出し
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}
	fmt.Printf("Greeting: %s\n", r.Message)
	fmt.Println("slankdev")
}

package main

import (
	//"context"
	//"os"
	//"time"
	"log"
	"fmt"
	"google.golang.org/grpc"
	pb "github.com/slankdev/frr_grpc_client/frr"
)

const (
	address     = "localhost:9999"
	defaultName = "slankdev"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNorthboundClient(conn)

	fmt.Printf("slankdev: %+v\n", c)
}

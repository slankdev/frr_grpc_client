package main

import (
	"io"
	"fmt"
	"time"
	"context"
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
		panic(err)
	}
	defer conn.Close()
	c := pb.NewNorthboundClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetCapabilities(ctx, &pb.GetCapabilitiesRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("frr_version: %+v\n", r.FrrVersion)
	fmt.Printf("rollback_support: %+v\n", r.RollbackSupport)
	for i, module := range(r.SupportedModules) {
		fmt.Printf("module[%d]: %+v\n", i, module)
	}
	for i, encoding := range(r.SupportedEncodings) {
		fmt.Printf("encoding[%d]: %+v\n", i, encoding)
	}
	fmt.Printf("--------------------\n")

	// path := []string{"/frr-interface","lib/frr-interface","interface/frr-interface","nam"}
	path := []string{"lib/frr-interface"}
	// path := []string{"/frr-ripd:clear-rip-route"}
	stream, err := c.Get(ctx, &pb.GetRequest{Type:pb.GetRequest_ALL, Encoding:pb.Encoding_JSON, Path:path})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("RES: %+v\n", res);
	}
}

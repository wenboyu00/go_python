package main

import (
	"context"
	"fmt"
	"go_part/grpc_start/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "wen"})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
}
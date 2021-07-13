package main

import (
	"context"
	"fmt"
	"go_part/grpc_start/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	s := Server{}
	proto.RegisterGreeterServer(g, &s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":8080"))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	g.Serve(lis)
	if err != nil{
		panic("failed to start grpc:" + err.Error())
	}
}

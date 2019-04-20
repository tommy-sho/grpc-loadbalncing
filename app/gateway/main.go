package main

import (
	"fmt"
	pb "github.com/tommy-sho/grpc-loadbalncing/app/gateway/genproto"
	"github.com/tommy-sho/grpc-loadbalncing/app/gateway/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main (){
	//ctx := context.Background()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50001))
	g := server.NewGatewaySerive()
	s := grpc.NewServer()

	pb.RegisterGreetingServerServer(s,g)
	if err != nil {
		panic(fmt.Errorf("new grpc server err: %v", err))
	}
	reflection.Register(s)

	s.Serve(lis)
}
package main

import (
	"fmt"
	"net"
	"os"

	pb "github.com/tommy-sho/grpc-loadbalncing/app/backend/genproto"
	"github.com/tommy-sho/grpc-loadbalncing/app/backend/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//ctx := context.Background()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("PORT_NUMBER")))
	r := server.NewUserRepository()
	g := server.NewBackendServer(r)

	s := grpc.NewServer()
	pb.RegisterGreetingServerServer(s, g)
	if err != nil {
		panic(fmt.Errorf("new grpc server err: %v", err))
	}
	reflection.Register(s)

	s.Serve(lis)
}

package server

import (
	"context"
	proto "github.com/tommy-sho/grpc-loadbalncing/app/gateway/genproto"
)

type GatewayService struct {

}

func NewGatewaySerive() *GatewayService{
	return &GatewayService{}
}

func (g *GatewayService)Greeting(context.Context, *proto.GreetingRequest) (*proto.GreetingResponse, error){
	return &proto.GreetingResponse{Message:"Hello, I'm backend service",}, nil
}
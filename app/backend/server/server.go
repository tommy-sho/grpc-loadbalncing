package server

import (
	"context"
	"fmt"

	proto "github.com/tommy-sho/grpc-loadbalncing/app/backend/genproto"
)

type BackendServer struct {
	userRepo BackendRepository
}

func NewBackendServer(userRepo BackendRepository) *BackendServer {
	return &BackendServer{
		userRepo: userRepo,
	}
}

func (b *BackendServer) Greeting(ctx context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	m, err := b.userRepo.GetMessageByName(req.Name)
	if err != nil {
		return &proto.GreetingResponse{}, fmt.Errorf("Greeting error : %v ", err)
	}

	res := &proto.GreetingResponse{
		Message: m,
	}
	return res, nil
}

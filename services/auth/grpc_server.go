package main

import (
	"context"
	"fmt"
	"net"

	pb "shop/proto"

	"google.golang.org/grpc"
)

func makeGRPCServerAndRun() error {

	s := GRPCAuthServer{
		svc: &AuthServer{},
	}

	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	pb.RegisterAuthServer(server, s)

	fmt.Printf("server is runnig on port: %s\n", ":4000")
	return server.Serve(ln)
}

type GRPCAuthServer struct {
	svc *AuthServer
	pb.UnimplementedAuthServer
}

func (s GRPCAuthServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	response := s.svc.Login(in.Email, in.Password)

	return &pb.LoginResponse{
		Token:   response.Token,
		Message: response.Message,
	}, nil
}

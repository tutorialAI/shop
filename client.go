package main

import (
	"context"
	"log"
	"time"

	pb "shop/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	client pb.AuthClient
}

func NewAuthServiceClient(addr string) *AuthServiceClient {
	// Устанавливаем соединение с Auth Service
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}

	return &AuthServiceClient{
		client: pb.NewAuthClient(conn),
	}
}

// Вызов метода Login в Auth Service
func (s *AuthServiceClient) Login(email, password string) (*pb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	resp, err := s.client.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Вызов метода Register в Auth Service
func (s *AuthServiceClient) Register(username, password string) (*pb.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.RegisterRequest{
		Username: username,
		Password: password,
	}

	resp, err := s.client.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

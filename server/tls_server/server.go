package main

import (
	"google.golang.org/grpc"
	"net"

	"context"
	"log"
	"rpc_example/proto"

	"google.golang.org/grpc/credentials"
)

type SearchService struct {}

func (s *SearchService) Search(ctx context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{}, nil
}

func (s *SearchService) GetUserById(ctx context.Context, request *proto.UserRequest) (*proto.UserResponse, error) {
	return &proto.UserResponse{}, nil
}

const Addr = ":9005"

func main() {
	c, err := credentials.NewServerTLSFromFile("../../conf/server.pem", "../../conf/server.key")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err %s", err.Error())
	}

	serve := grpc.NewServer(grpc.Creds(c))

	proto.RegisterSearchServiceServer(serve, &SearchService{})

	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("net.Listen err %s", err.Error())
	}

	serve.Serve(lis)
}
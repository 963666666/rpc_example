package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"rpc_example/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func (s *SearchService) GetUserById(ctx context.Context, r *proto.UserRequest) (*proto.UserResponse, error) {
	user := &proto.UserResponse{User: "admin", Admin: "admin", Loginci: 100}
	return user, nil
}

const PORT string = "9001"

func main() {
	server := grpc.NewServer()
	proto.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}

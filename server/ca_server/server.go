package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"rpc_example/proto"

	"context"
	"net"

	"io/ioutil"
	"log"
)

type SearchService struct {
	
}

func (s *SearchService) Search(context.Context, *proto.SearchRequest) (*proto.SearchResponse, error) {

	return &proto.SearchResponse{Response: "hello world"}, nil
}

func (s *SearchService) GetUserById(context.Context, *proto.UserRequest) (*proto.UserResponse, error) {
	return &proto.UserResponse{}, nil
}

const Addr = ":9006"

func main() {
	cert, err := tls.LoadX509KeyPair("../../conf/server/server.pem", "../../conf/server/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err %s", err.Error())
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err %s", err.Error())
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err: %s", err.Error())
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	})

	server := grpc.NewServer(grpc.Creds(c))

	proto.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", Addr)

	if err != nil {
		log.Fatalf("net.Listen err: %s", err.Error())
	}

	server.Serve(lis)
}

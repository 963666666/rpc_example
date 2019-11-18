package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net"
	"rpc_example/proto"

	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"log"
)

type SearchService struct {}

func (s *SearchService) Search(ctx context.Context, SearchRequest *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: SearchRequest.Request + " yingyingying"}, nil
}
func (s *SearchService) GetUserById(ctx context.Context, SearchRequest *proto.UserRequest) (*proto.UserResponse, error) {
	return &proto.UserResponse{}, nil
}


func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)

	resp, err := handler(ctx, req)

	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)

	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if e := recover(); e != nil {
			//err = status.Errorf(codes.Internal, "Panic err: %v", e)
			log.Fatalf("recover err: %s", e)
		}
	}()

	return handler(ctx, req)
}

func main() {
	c, err := GetTLSCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %s", err.Error())
	}
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(RecoveryInterceptor, LoggingInterceptor),
	}

	server := grpc.NewServer(opts...)

	proto.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":9007")
	if err != nil {
		log.Fatalf("net.Listen err: %s", err.Error())
	}
	server.Serve(lis)
}

func GetTLSCredentialsByCA() (credentials.TransportCredentials, error) {
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
	return c, nil
}
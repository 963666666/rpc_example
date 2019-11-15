package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"rpc_example/proto"
)

const Addr = ":9005"

func main() {
	c, err := credentials.NewClientTLSFromFile("../../conf/server.pem", "rpc_example")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}

	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()


	client := proto.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())

}

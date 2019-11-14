package main

import (
	"golang.org/x/net/context"
	"log"
	"rpc_example/proto"

	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":" + PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc Dial err: %v", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)

	resp, err := client.Search(context.Background(), &proto.SearchRequest{
		Request: "gRPC",
	})

	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp: %s", resp.Response)

}
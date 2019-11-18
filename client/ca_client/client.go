package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"rpc_example/proto"

	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

const ADDR = ":9007"

func main() {
	cert, err := tls.LoadX509KeyPair("../../conf/client/client.pem", "../../conf/client/client.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %s", err.Error())
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %s", err.Error())
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err: %s", err.Error())
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName: "server",
		RootCAs: certPool,
	})

	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()
	client := proto.NewSearchServiceClient(conn)

	resp, err := client.Search(context.Background(), &proto.SearchRequest{Request: "hello world! henghenghahi"})
	if err != nil {
		log.Fatalf("client.Search err: %s", err)
	}

	log.Fatalf("resp: %s", resp.Response)
}
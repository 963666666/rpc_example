package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"rpc_example/proto"
)

const Addr = ":9003"

func main() {
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())

	client := proto.NewStreamServiceClient(conn)

	err = printList(client, &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gPRC stream Client: List", Value: 2018}})
	if err != nil {
		log.Fatalf("PrintList.err is %s", err.Error())
	}
	err = printRecord(client, &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gPRC stream Client: Record", Value: 2030}})
	if err != nil {
		log.Fatalf("PrintRecord.err is %s", err.Error())
	}
	err = printRoute(client, &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC stream Client: Route", Value: 2020}})
	if err != nil {
		log.Fatalf("PrintRoute.err is %s", err.Error())
	}
}

func printList(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}

	for true {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}

func printRecord(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n ++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func printRoute(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n <= 6; n ++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	stream.CloseSend()

	return nil
}

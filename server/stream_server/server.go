package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"rpc_example/proto"
)

const ADDRESS = ":9003"

type StreamService struct {}

func (s *StreamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	for n := 0; n <= 6; n ++ {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{Name: r.Pt.Name, Value: r.Pt.Value},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StreamService) Record(stream proto.StreamService_RecordServer) error {
	for true {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.StreamResponse{Pt:
				&proto.StreamPoint{Name: "gRPC Stream Server Record", Value: 1},
			})
		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	
	return nil
}

func (s *StreamService) Route(stream proto.StreamService_RouteServer) error {
	n := 0
	for true {
		err := stream.Send(&proto.StreamResponse{Pt: &proto.StreamPoint{Name: "gRPC Stream Server Route", Value: int64(n)}})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		n ++

		log.Printf("stream.Revc pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	return nil
}

func main() {
	server := grpc.NewServer()

	proto.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Fatalf("net.Listen err %s", err.Error())
	}

	server.Serve(lis)
}
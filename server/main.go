package main

import (
	"context"
	"log"
	"net"

	"github.com/root27/go-grpc/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func main() {
	conn, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

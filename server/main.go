package main

import (
	"context"
	"log"
	"net"

	"github.com/root27/go-grpc/pb"
	"github.com/root27/go-grpc/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
	pb.UnimplementedUserServiceServer
}

func main() {
	conn, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &server{})
	pb.RegisterUserServiceServer(s, &server{})

	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	res, err := db.InsertOne(context.Background(), bson.M{"name": in.Name, "email": in.Email})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &pb.CreateResponse{Id: res.InsertedID.(primitive.ObjectID).Hex(), Name: in.Name, Email: in.Email}, nil
}

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/root27/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	client := pb.NewHelloServiceClient(conn)

	r := gin.Default()

	r.GET("/api/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		request := &pb.HelloRequest{Name: name}

		response, err := client.Hello(c, request)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		c.JSON(200, gin.H{"message": response.Message})

	})

	log.Fatal(r.Run(":8081"))

}

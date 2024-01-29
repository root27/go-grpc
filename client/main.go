package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}

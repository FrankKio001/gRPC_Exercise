package main

import (
	"log"
	"net"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
	"google.golang.org/grpc"
)

// grpc ports 50051(tcp/udp)
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.HWServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

package main

import (
	"context"
	"log"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func (s *Server) Hello(ctx context.Context, in *pb.HWRequest) (*pb.HWResponse, error) {
	log.Printf("Hello function %v\n", in)
	return &pb.HWResponse{
		Result: "Hello" + in.Name,
	}, nil
}

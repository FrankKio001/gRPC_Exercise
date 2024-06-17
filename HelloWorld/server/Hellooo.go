package main

import (
	"fmt"
	"log"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func (s *Server) Hellooo(in *pb.HWRequest, stream pb.HWService_HelloooServer) error {
	log.Printf("server streaming Hellooo function %v\n", in)

	for i := 0; i < 3; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.Name, i)

		stream.Send(&pb.HWResponse{
			Result: res,
		})
	}
	return nil
}

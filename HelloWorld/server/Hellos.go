package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func (s *Server) Hellos(stream pb.HWService_HellosServer) error {
	log.Printf("Hello functions")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.HWResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += fmt.Sprintf("Hellos %s\n", req.Name)
	}
}

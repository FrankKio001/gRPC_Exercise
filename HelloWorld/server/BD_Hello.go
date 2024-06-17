package main

import (
	"io"
	"log"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func (s *Server) BD_Hello(stream pb.HWService_BD_HelloServer) error {
	log.Printf("bi direction streaming")

	for {
		//Recv = Receive
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream")
		}

		res := "Hello " + req.Name + "!"
		err = stream.Send(&pb.HWResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}

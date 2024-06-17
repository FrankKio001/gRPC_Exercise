package main

import (
	"context"
	"log"
	"time"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func doHellos(c pb.HWServiceClient) {
	log.Println("dohellos")

	reqs := []*pb.HWRequest{
		{Name: "World 1"},
		{Name: "World 2"},
	}

	stream, err := c.Hellos(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Hellos %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving reponse from Hellos: %v\n", err)
	}

	log.Printf("Hellos : %s\n", res.Result)
}

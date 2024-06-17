package main

import (
	"context"
	"io"
	"log"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func doHellooo(c pb.HWServiceClient) {
	log.Println("dohellooo")

	req := &pb.HWRequest{
		Name: "World",
	}

	stream, err := c.Hellooo(context.Background(), req)

	if err != nil {
		log.Fatalf("error call hellooo :%v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Hellooo : %s\n", msg.Result)
	}
}

package main

import (
	"context"
	"log"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func doHello(c pb.HWServiceClient) {
	log.Println("dohello")
	res, err := c.Hello(context.Background(), &pb.HWRequest{
		Name: "World",
	})

	if err != nil {
		log.Fatalf("Couldn't Hi %v\n", err)
	}

	log.Printf("Hi :%s\n", res.Result)
}

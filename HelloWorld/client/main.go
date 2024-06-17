package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewHWServiceClient(conn)

	//doHello(c)
	//doHellooo(c)
	//doHellos(c)
	doBD_Hello(c)
}

package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/FrankKio001/gRPC_Exercise/HelloWorld/proto"
)

func doBD_Hello(c pb.HWServiceClient) {
	log.Println("dobd_Hello")

	stream, err := c.BD_Hello(context.Background())

	if err != nil {
		//記錄錯誤並終止程序
		log.Fatalf("Error while creating stream %v\n", err)
	}

	reqs := []*pb.HWRequest{
		{Name: "Wrold 1"},
		{Name: "Wrold 2"},
	}

	//channel
	//goroutine
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received :%v\n", res.Result)
		}

		close(waitc)
	}()

	//<-waitc,阻塞,等待goroutine結束
	<-waitc
}

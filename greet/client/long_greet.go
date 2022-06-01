package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dev-hyunsang/study-grpc-video/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{
			FirstName: "HyunSang",
		},
		{
			FirstName: "Marie",
		},
		{
			FirstName: "James",
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreetL %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet : %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}

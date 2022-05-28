package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dev-hyunsang/study-grpc-video/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTime was invoked")

	req := &pb.GreetRequest{
		FirstName: "HyunSang",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Printf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}

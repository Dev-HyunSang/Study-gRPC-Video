package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/study-grpc-video/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})
	if err != nil {
		log.Printf("Cloud not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

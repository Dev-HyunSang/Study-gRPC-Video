package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/study-grpc-video/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FristNumber:  1,
		SecnodNumber: 1,
	})

	if err != nil {
		log.Fatalf("Cloud not sum: %v", err)
	}

	log.Printf("sum: %d\n", res.Result)
}

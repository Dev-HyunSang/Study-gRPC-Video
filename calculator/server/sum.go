package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/study-grpc-video/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FristNumber + in.SecnodNumber,
	}, nil
}

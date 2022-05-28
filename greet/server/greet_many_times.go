package main

import (
	"fmt"
	"log"

	pb "github.com/dev-hyunsang/study-grpc-video/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes fucncitons was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, nubmer %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}

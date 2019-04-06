package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function call with %v \n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hallo " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	fmt.Println("Hallo my course proto")

	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}

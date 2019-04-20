package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"../greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreatEveryone(stream greetpb.GreetService_GreatEveryoneServer) error {
	fmt.Printf("GreatEveryone function was run streaming request \n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while calling server", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Haiii " + firstName + " !"

		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client : %v", err)
		}
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

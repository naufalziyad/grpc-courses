package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"../greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was run streaming request \n")
	result := "Hallo "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream : %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += firstName + " ! "
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

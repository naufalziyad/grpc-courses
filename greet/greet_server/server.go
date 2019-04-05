package main

import ("fmt"
	"log"
	"net"
	"../greetpb"
	"google.golang.org/grpc")

type server struct{}

func main() {
	fmt.Println("Hallo my course proto")

	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil{
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatal("failed to serve : %v", err)
	}
}
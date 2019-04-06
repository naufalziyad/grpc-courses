package main

import (
	"fmt"
	"log"

	"../greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello this is me , your client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("can't connect : %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)

}

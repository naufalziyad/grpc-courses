package main

import (
	"context"
	"fmt"
	"io"
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

	doServerStreaming(c)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Server  Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Naufal",
			LastName:  "Ziyad Luthfiansyah",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet many time %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//we've reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading streaming %v", err)
		}
		log.Printf("Response from GreetManyTimes : %v", msg.GetResult())
	}
}

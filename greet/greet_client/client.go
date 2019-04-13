package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// doServerStreaming(c)
	doClientStreaming(c)

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Client streaming RPC")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Naufal",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ziyad",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Luthfiansyah",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lutfi",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Zahid",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Fian",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		fmt.Printf("Error while calling Longreet : %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req : %v \n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet : %v", err)
	}
	fmt.Printf("LongGreet Response : %v \n ", res)
}

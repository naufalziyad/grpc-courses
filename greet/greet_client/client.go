package main

import (
	"context"
	"fmt"
	"io"
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
	// doClientStreaming(c)
	doBiDiStreaming(c)

}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a BiDi Streaming RPC")
	stream, err := c.GreatEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream : %v\n", err)
		return
	}

	//requests
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Naufal",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ziyad",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Luthfiansyah",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lutfi",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Zahid",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Fian",
			},
		},
	}

	waitc := make(chan struct{})

	go func() {
		//function to send messages
		for _, req := range requests {
			fmt.Printf("Sending Message : %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving: %v\n", err)
				break
			}
			fmt.Printf("Received: %v", res.GetResult())
		}

	}()

	<-waitc

}

// func doClientStreaming(c greetpb.GreetServiceClient) {
// 	fmt.Println("Starting Client streaming RPC")

// 	requests := []*greetpb.LongGreetRequest{
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Naufal",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Ziyad",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Luthfiansyah",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Lutfi",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Zahid",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Fian",
// 			},
// 		},
// 	}

// 	stream, err := c.LongGreet(context.Background())
// 	if err != nil {
// 		fmt.Printf("Error while calling Longreet : %v", err)
// 	}

// 	for _, req := range requests {
// 		fmt.Printf("Sending req : %v \n", req)
// 		stream.Send(req)
// 		time.Sleep(1000 * time.Millisecond)
// 	}

// 	res, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalf("Error while receiving response from LongGreet : %v", err)
// 	}
// 	fmt.Printf("LongGreet Response : %v \n ", res)
// }

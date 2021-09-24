package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-course/greet/greetpb"
	"io"
	"log"
)

func main() {
	fmt.Println("Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	// doUnary(c)

	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC")
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Eliran",
			LastName:  "Turgeman",
		},
	}
	res, err := c.Greet(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling greet rpc %v", err)
	}

	log.Printf("Response from greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Server Streaming RPC")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Eliran",
			LastName:  "Turgeman",
		},
	}
	res_stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling greetmanytimes rpc %v", err)
	}

	for {
		msg, err := res_stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Printf("result from greetmanytimes: %v", msg.GetResult())
	}
}

package main

import (
	"context"
	"fmt"
	"grpc-course/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)
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

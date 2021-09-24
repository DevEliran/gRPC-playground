package main

import (
	"context"
	"fmt"
	"grpc-course/calculator/calculatorpb"
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
	c := calculatorpb.NewCalculatorServiceClient(conn)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Unary RPC")
	request := &calculatorpb.CalculatorRequest{
		Calc: &calculatorpb.Calculator{
			FirstNum:  3,
			SecondNum: 10,
		},
	}
	res, err := c.Calculator(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling greet rpc %v", err)
	}

	log.Printf("Response from greet: %v", res.Result)
}

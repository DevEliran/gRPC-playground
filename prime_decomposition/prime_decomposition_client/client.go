package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-course/prime_decomposition/prime_decompositionpb"
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
	c := prime_decompositionpb.NewPrimeNumberServiceClient(conn)

	doServerStreaming(c)
}

func doServerStreaming(c prime_decompositionpb.PrimeNumberServiceClient) {
	fmt.Println("Starting Server Streaming RPC")

	req := &prime_decompositionpb.PrimeNumberRequest{
		PrimeNumber: &prime_decompositionpb.PrimeNumber{
			Number: 1242656510,
		},
	}

	res_stream, err := c.PrimeDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling prime decomposition rpc %v", err)
	}

	for {
		msg, err := res_stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Printf("Result from prime decomposition: %v", msg.GetResult())
	}
}

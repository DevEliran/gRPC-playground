package main

import (
	"fmt"
	"grpc-course/prime_decomposition/prime_decompositionpb"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeDecomposition(req *prime_decompositionpb.PrimeNumberRequest, stream prime_decompositionpb.PrimeNumberService_PrimeDecompositionServer) error {
	fmt.Printf("Prime Decomposition function invoked with %v", req)
	prime_number := req.GetPrimeNumber().GetNumber()

	var div int32 = 2
	p := prime_number
	for {
		if p <= 1 {
			break
		}
		if p%div == 0 {
			result := "Decomposition of " + strconv.Itoa(int(prime_number)) + " includes " + strconv.Itoa(int(div))
			res := &prime_decompositionpb.PrimeNumberResponse{
				Result: result,
			}
			stream.Send(res)

			p /= div
		} else {
			div += 1
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	prime_decompositionpb.RegisterPrimeNumberServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

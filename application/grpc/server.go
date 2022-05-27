package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("could not connect", err)
	}

	grpcServer := grpc.NewServer()

	log.Println("gRPC server has been started")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("could not connect", err)
	}
}

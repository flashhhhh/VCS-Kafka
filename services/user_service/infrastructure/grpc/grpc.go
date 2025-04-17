package grpc

import (
	"log"
	"net"
	grpc_handler "user_service/internal/handler/grpc"
	"user_service/pb"

	"google.golang.org/grpc"
)

func StartGRPCServer(userHandler *grpc_handler.UserHandler, port string) {
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		panic(err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("gRPC server is running on port: ", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
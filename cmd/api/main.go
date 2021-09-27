package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/koga456/rock-paper-scissors/pb"
	"github.com/koga456/rock-paper-scissors/service"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRockPaperScissorsServiceServer(server, service.NewRockPaperScissorsService())

	reflection.Register(server)
	server.Serve(listenPort)
}

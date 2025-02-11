package main

import (
	"log"
	"net"

	"github.com/Solwery-Veronika/user/internal/repository"
	"github.com/Solwery-Veronika/user/internal/service"

	"github.com/Solwery-Veronika/user/pkg/user"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.New()

	srv := service.New(repo)

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, srv)

	lis, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

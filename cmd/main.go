package main

import (
	"log"
	"net"

	desc "github.com/Dnlbb/auth/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	desc.UnimplementedAuthServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen: 50051 ")
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthServer(s, &server{})

	log.Printf("server listening at #{lis.Addr()}")

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve: #{err}")
	}

}

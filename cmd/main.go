package main

import (
	"context"
	"log"
	"net"

	desc "github.com/Dnlbb/auth/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedAuthServer
}

func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id:%d", req.GetId())
	return &desc.GetResponse{}, nil
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User #+%v\n", req.GetUser())
	log.Printf("Password: %s", req.Password)
	log.Printf("Password confirm: %s", req.PasswordConfirm)
	return &desc.CreateResponse{}, nil
}

func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())
	log.Printf("Username: %s", req.Name.Value)
	log.Printf("Email: %s", req.Email.Value)
	return nil, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())
	return nil, nil
}
func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal("failed to listen: 50051 ")
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

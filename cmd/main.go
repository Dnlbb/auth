package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/joho/godotenv"

	dao "github.com/Dnlbb/auth/postgres/cmd"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

func main() {
	err := godotenv.Load("../postgres/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal("failed to listen: 50051 ")
	}

	storage, err := dao.InitStorage()
	defer storage.CloseCon()
	if err != nil {
		log.Fatal("failed to init storage")
	}
	s := grpc.NewServer()
	reflection.Register(s)
	authv1.RegisterAuthServer(s, &server{storage: storage})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type server struct {
	authv1.UnimplementedAuthServer
	storage dao.PostgresInterface
}

func toTimestampProto(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}

func (s *server) Get(_ context.Context, req *authv1.GetRequest) (*authv1.GetResponse, error) {
	var params dao.GetUserParams
	// Достаем пришедшие параметры
	switch nameOrID := req.NameOrId.(type) {
	case *authv1.GetRequest_Id:
		params.ID = &nameOrID.Id
	case *authv1.GetRequest_Username:
		params.Username = &nameOrID.Username
	default:
		return nil, fmt.Errorf("необходимо указать либо ID, либо Username")
	}

	userProfile, err := s.storage.GetUser(params)
	if err != nil {
		return nil, fmt.Errorf("error when getting the user profile: %w", err)
	}

	response := authv1.GetResponse{Id: userProfile.ID,
		User: &authv1.User{
			Name:  userProfile.Name,
			Email: userProfile.Email,
			Role:  userProfile.Role,
		},
		CreatedAt: toTimestampProto(userProfile.CreatedAt),
		UpdatedAt: toTimestampProto(userProfile.UpdatedAt),
	}

	return &response, nil
}

func (s *server) Create(_ context.Context, req *authv1.CreateRequest) (*authv1.CreateResponse, error) {
	user := dao.User{Name: req.GetUser().GetName(),
		Email:    req.GetUser().GetEmail(),
		Role:     req.GetUser().GetRole(),
		Password: req.GetPassword()}

	err := s.storage.Save(user)
	if err != nil {
		return nil, fmt.Errorf("error when saving the user: %w", err)
	}
	return &authv1.CreateResponse{}, nil
}

func (s *server) Update(_ context.Context, req *authv1.UpdateRequest) (*emptypb.Empty, error) {
	updateUser := dao.UpdateUser{ID: req.GetId(),
		Name:  req.Name.Value,
		Email: req.Email.Value,
		Role:  req.GetRole()}
	err := s.storage.Update(updateUser)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("error updating user: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *authv1.DeleteRequest) (*emptypb.Empty, error) {
	idDel := dao.DeleteID(req.GetId())
	err := s.storage.Delete(idDel)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("error deleting user: %w", err)
	}

	return &emptypb.Empty{}, nil
}

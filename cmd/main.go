package main

import (
	"context"
	"log"
	"net"

	"github.com/joho/godotenv"

	dao "github.com/Dnlbb/auth/postgres/cmd"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Dnlbb/auth/pkg/auth"
)

type server struct {
	desc.UnimplementedAuthServer
	storage dao.PostgresInterface
}

func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id:%d", req.GetId())
	userID := dao.GetID(req.GetId())

	user, err := s.storage.Get(userID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	log.Printf("User:%+v", user)
	return &desc.GetResponse{}, nil
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User #+%v\n", req.GetUser())
	log.Printf("Password: %s", req.Password)
	log.Printf("Password confirm: %s", req.PasswordConfirm)
	user := dao.User{Name: req.GetUser().GetName(),
		Email:    req.GetUser().GetEmail(),
		Role:     req.GetUser().GetRole(),
		Password: req.GetPassword()}

	err := s.storage.Save(user)
	if err != nil {
		log.Printf("Error saving user: %v", err)
		return nil, err
	}
	return &desc.CreateResponse{}, nil
}

func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())
	log.Printf("Username: %s", req.Name.Value)
	log.Printf("Email: %s", req.Email.Value)
	updateUser := dao.UpdateUser{ID: req.GetId(),
		Name:  req.Name.Value,
		Email: req.Email.Value,
		Role:  req.GetRole()}
	err := s.storage.Update(updateUser)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())
	idDel := dao.DeleteID(req.GetId())
	err := s.storage.Delete(idDel)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetProfile(_ context.Context, req *desc.GetProfileRequest) (*desc.GetProfileResponse, error) {
	log.Printf("Имя пользователя: %s", req.GetUsername())
	var UserProfile dao.UserProfile
	UserProfile, err := s.storage.GetProfile(req.GetUsername())
	if err != nil {
		log.Printf("Ошибка при получении профиля пользователя: %v", err)
		return nil, err
	}
	response := &desc.GetProfileResponse{
		Id:    UserProfile.ID,
		Name:  UserProfile.Name,
		Email: UserProfile.Email,
		Role:  UserProfile.Role,
	}
	log.Printf("Профиль пользователя для добавления в чат: %+v", response)
	return response, nil
}

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
	desc.RegisterAuthServer(s, &server{storage: storage})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

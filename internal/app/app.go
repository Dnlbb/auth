package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Dnlbb/auth/internal/closer"
	"github.com/Dnlbb/auth/internal/config"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, fmt.Errorf("init deps: %w", err)
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.LoadEnv("auth.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	authv1.RegisterAuthServer(a.grpcServer, a.serviceProvider.GetAuthImpl(ctx))

	return nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	return a.runGRPCServer()
}

func (a *App) runGRPCServer() error {
	log.Printf("starting gRPC server on %s", a.serviceProvider.GetGRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GetGRPCConfig().Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

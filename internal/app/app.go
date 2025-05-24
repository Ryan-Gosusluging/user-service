package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ryan-Gosusluging/go-common-forum/logger"
	"github.com/Ryan-Gosusluging/go-common-forum/postgres"
	"github.com/Ryan-Gosusluging/user-service/config"
	authgrpc "github.com/Ryan-Gosusluging/user-service/internal/grpc"
	"github.com/Ryan-Gosusluging/user-service/internal/repo"
	"github.com/Ryan-Gosusluging/user-service/internal/usecase"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	//Logger
	logger := logger.New("user-service", cfg.LogLevel)

	//Repository
	pg, err := postgres.New(cfg.PG_URL)
	if err != nil {
		log.Fatalf("app - Run - postgres.New")
	}
	defer pg.Close()

	userRepo := repo.New(pg, logger)

	//Usecase
	userUsecase := usecase.New(userRepo, logger)

	//GRPC-Server
	grpcServer := grpc.NewServer()
	authgrpc.Register(grpcServer, userUsecase, logger)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("app - Run - net.Listen: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(l); err != nil {
			log.Fatalf("app - Run - grpcServer.Serve: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
package main

import (
	"avito/config"
	"avito/internal/handlers"
	"avito/internal/repository"
	"avito/internal/service"
	"avito/proto/proto/pb"
	"context"
	"fmt"
	//"github.com/gofiber/fiber/v2/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
	"net/http"
	"os"
)

type App struct {
	repository *repository.Repository
	service    *service.Service
	handlers   *handlers.Handlers
}

const (
	grpcPort = ":8090"
	httpPort = ":8080"
)

func main() {
	fmt.Println("Start")
	//ctx := context.Background()
	lis, err := net.Listen("tcp", grpcPort)
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)) //засунуть в контекст
	app := &App{}
	cfg := config.Config_load()
	app.repository = repository.New(cfg)
	go func() {
		err := app.repository.CacheRecovery()
		if err != nil {
			log.Error("Fail in cache recovery: " + err.Error())
		}
	}()
	app.service = service.New(app.repository)
	app.handlers = handlers.New(app.service)
	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, app.handlers)
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Error("failed to serve: " + err.Error())
			//log.Error("failed to serve: " + err.Error())
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		grpcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error("failed to dial serve: " + err.Error())
		//log.Error("Failed to dial server: " + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Error("failed to close connection " + err.Error())
		}
	}(conn)

	gwmux := runtime.NewServeMux()
	err = pb.RegisterYourServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Error("failed : " + err.Error())
		log.Error("Failed to register gateway:" + err.Error())
	}

	gwServer := &http.Server{
		Addr:    httpPort,
		Handler: gwmux,
	}

	log.Info("Serving gRPC-Gateway on port " + httpPort)
	if err = gwServer.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Warn("Server closed: " + err.Error())
			os.Exit(0)
		}
		log.Error("Failed to listen and serve: " + err.Error())
	}
	if err != nil {
		return
	}
}

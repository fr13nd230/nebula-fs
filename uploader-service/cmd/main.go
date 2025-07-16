package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/fr13nd230/nebula-fs/uploader-service/cmd/uploader/service"
	"github.com/fr13nd230/nebula-fs/uploader-service/config"
	"github.com/fr13nd230/nebula-fs/uploader-service/grpc/uploader"
	"go.uber.org/zap"
)

func main() {
    lg, _ := zap.NewDevelopment()
    sug := lg.Sugar()
    
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	defer func() {
		if r := recover(); r != nil {
			sug.Infow("[UploaderServer]: Recovering application from panic.", "original panic", r)
		}
	}()

	err := config.InitConfig()
	if err != nil {
		sug.Errorw("[UploaderServer]: Failed to mount config files.", "error", err)
		return
	}

	gRPCPort := config.GetVariable("GRPC_PORT")
	if gRPCPort == "" {
		sug.Error("[UploaderServer]: GRPC_PORT is not set in the config")
		return
	}

	grpcErrs := make(chan error, 1)
	cfg := NewConfig(gRPCPort)
	grpcServer := cfg.NewGrpcServer()
	
	uploader.RegisterFileUploaderServer(grpcServer, &service.UploaderService{})

	go func() {
		grpcErrs <- cfg.Start(grpcServer)
	}()

	sug.Infow("[UploaderServer]: gRPC Server has been succesfully launched.", "port", cfg.port)

	select {
	case <-ctx.Done():
		sug.Error("[UploaderServer]: Gracefull shutdown has been scheduled, shutdown signal received ....")
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		grpcServer.GracefulStop()
		sug.Info("[UploaderServer]: Uploader server has been shutdowned succesfully.")

	case err := <-grpcErrs:
		sug.Errorw("[UploaderServer]: Failed to start new gRPC server.", "error", err)
		sug.Info("[UploaderServer]: Starting to gracefully shutdown the gRPC server ....")
		grpcServer.GracefulStop()
		sug.Info("[UploaderServer]: gRPC server has been gracefully shutdowned.")
	}
}

package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"
	"time"

	"github.com/fr13nd230/nebula-fs/uploader-service/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	defer func() {
		if r := recover(); r != nil {
			slog.Info("[UploaderServer]: Recovering application from panic.", "original panic", r)
		}
	}()

	err := config.InitConfig()
	if err != nil {
		slog.Error("[UploaderServer]: Failed to mount config files.", "error", err)
		return
	}

	gRPCPort := config.GetVariable("GRPC_PORT")
	if gRPCPort == "" {
		slog.Error("[UploaderServer]: GRPC_PORT is not set in the config")
		return
	}

	grpcErrs := make(chan error, 1)
	cfg := NewConfig(gRPCPort)
	grpcServer := cfg.NewGrpcServer()

	go func() {
		grpcErrs <- cfg.Start(grpcServer)
	}()

	slog.Info("[UploaderServer]: gRPC Server has been succesfully launched.", "port", cfg.port)

	select {
	case <-ctx.Done():
		slog.Error("[UploaderServer]: Gracefull shutdown has been scheduled, shutdown signal received ....")
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		grpcServer.GracefulStop()
		slog.Info("[UploaderServer]: Uploader server has been shutdowned succesfully.")

	case err := <-grpcErrs:
		slog.Error("[UploaderServer]: Failed to start new gRPC server.", "error", err)
		slog.Info("[UploaderServer]: Starting to gracefully shutdown the gRPC server ....")
		grpcServer.GracefulStop()
		slog.Info("[UploaderServer]: gRPC server has been gracefully shutdowned.")
	}
}

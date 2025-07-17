package main

import (
	"context"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fr13nd230/nebula-fs/storage-service/config"
	"go.uber.org/zap"
)

func main() {
	lg, _ := zap.NewDevelopment()
	sg := lg.Sugar()
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer func() {
		if r := recover(); r != nil {
			sg.Infow("[StorageService]: Application has recovered from a panic.", "panic", r)
		}
	}()

	err := config.LoadEnv()
	if err != nil {
		sg.Errorw("[StorageService]: Failed to load environement files.", "error", err)
		return
	}

	port := config.GetVar("GRPC_PORT")
	if strings.Trim(port, " ") == " " {
		sg.Error("[ServiceStorage]: Missing port variable for GRPC server.")
		return
	}

	gCfg := NewgRPConfig(port)
	grpcServ := gCfg.NewGRPCServer()
	grpcErr := make(chan error, 1)

	go func() {
		grpcErr <- gCfg.StartGRPCServer(grpcServ)
	}()

	sg.Infow("[StorageService]: New gRPC Server has started.", "port", port)

	select {
	case <-ctx.Done():
		sg.Info("[StorageService]: A gracefull shutdown has been scheduled, 5s timeout.")
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		grpcServ.GracefulStop()
		sg.Info("[StorageService]: Service servers has been shutdown succesfully.")
	case err := <-grpcErr:
		sg.Infow("[StorageService]: gRPC Server has faced some troules preparing for gracefull shutdown.", "error", err)
		grpcServ.GracefulStop()
		sg.Info("[StorageService]: Service servers has been shutdown succesfully.")
	}
}

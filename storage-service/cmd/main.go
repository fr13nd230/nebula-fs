package main

import (
	"context"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fr13nd230/nebula-fs/storage-service/cmd/handler"
	"github.com/fr13nd230/nebula-fs/storage-service/cmd/service"
	"github.com/fr13nd230/nebula-fs/storage-service/config"
	"github.com/fr13nd230/nebula-fs/storage-service/grpc/storage"
	"github.com/fr13nd230/nebula-fs/storage-service/repository/store"
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
	drvPath := config.GetVar("PSQL_DRIVER_PATH")
	if strings.Trim(drvPath, " ") == " " {
		sg.Error("[ServiceStorage]: Missing database driver path variable for GRPC server.")
		return
	}

	q, err := store.NewDB(drvPath)
	if err != nil {
		sg.Errorw("[StorageService]: Service has failed, no database connection.", "error", err)
		return
	}

	gCfg := NewgRPConfig(port)
	grpcServ := gCfg.NewGRPCServer()
	grpcErr := make(chan error, 1)
	defer close(grpcErr)

	s := service.NewStorageService(q, sg)
	h := handler.NewStorageHandler(sg, s)

	storage.RegisterStorageServiceServer(grpcServ, h)

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
		sg.Debug("ERR CHAN CAUSED THIS INFINITE WAIT BLOCK")
		sg.Infow("[StorageService]: gRPC Server has faced some troules preparing for gracefull shutdown.", "error", err)
		grpcServ.GracefulStop()
		sg.Info("[StorageService]: Service servers has been shutdown succesfully.")
	}
}

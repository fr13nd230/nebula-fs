package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"
	"time"

	"github.com/fr13nd230/nebula-fs/uploader-service/config"
)

// main Entry point for any configuration mouting, healthcheck and the
// point where we handle recovery and gracefull shutdowns.
func main() {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
    defer stop()
    defer func() {
		if r := recover(); r != nil {
			slog.Info("[UploaderServer]: Recovering application from panic.", "original panic", r)
		}
	}()
    
    // serverErrs := make(chan error, 1)

	err := config.InitConfig()
	if err != nil {
		slog.Error("[UploaderServer]: Failed to mount config files.", "error", err)
	}
	
	select {
	case <-ctx.Done():
	    slog.Error("[UploaderServer]: Gracefull shutdown has been scheduled, shutdown signal received ....")
		_, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()
        
        // Use shutCtx to shut down any operations in here ... 
        
        slog.Info("[UploaderServer]: Uploader server has been shutdowned succesfully.")
	}
}

package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thesoulless/playground/go/pgo/app"
)

func main() {
	ctx := context.Background()
	run(ctx)
}

func run(ctx context.Context) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	srv := app.New(log)

	go func() {
		log.Info("running app...")
		// Start the service
		if err := srv.Run(ctx); err != nil && err != app.ErrServerClosed {
			log.Error("app Run failed", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown the server
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := srv.Close(cctx)
	if err != nil {
		log.Error("app closed with error", "error", err)
		return
	}

	log.Info("app closed successfully")
}

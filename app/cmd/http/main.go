package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucasti79/meli-interview/cmd/http/router"
	"github.com/lucasti79/meli-interview/config"
	"github.com/lucasti79/meli-interview/internal/factory"
)

// @title Example API
// @version 1.0
// @description This is an example API
// @BasePath /
func main() {
	cfg := config.LoadConfig()

	if err := factory.InitFactory(); err != nil {
		log.Fatalf("failed to initialize AppFactory: %v", err)
	}

	r := router.NewRouter()

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Println("starting server on", addr)

	server := &http.Server{
		Addr:         addr,
		Handler:      r.MapRoutes(factory.GetFactory()),
		ReadTimeout:  cfg.Server.TimeoutRead,
		WriteTimeout: cfg.Server.TimeoutWrite,
		IdleTimeout:  cfg.Server.TimeoutIdle,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()
	log.Println("server is running...")

	<-stop
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exited properly")
}

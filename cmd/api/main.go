package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/ioc"
	"github.com/velosypedno/zlagoda/internal/server"
)

func main() {
	cfg := config.Load()
	handlerContainer, err := ioc.BuildHandlerContainer(cfg)
	if err != nil {
		log.Fatal("Failed to build handler container:", err)
	}
	defer func() {
		if err := handlerContainer.Close(); err != nil {
			log.Printf("Error closing handler container: %v", err)
		}
	}()

	router := server.SetupRoutes(handlerContainer)

	// Setup graceful shutdown
	go func() {
		if err := router.Run(":" + cfg.PORT); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}

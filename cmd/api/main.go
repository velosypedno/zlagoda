package main

import (
	"log"

	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/ioc"
	"github.com/velosypedno/zlagoda/internal/server"
)

func main() {
	cfg := config.Load()
	handlerContainer := ioc.BuildHandlerContainer(cfg)
	router := server.SetupRoutes(handlerContainer)
	err := router.Run(":" + cfg.PORT)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

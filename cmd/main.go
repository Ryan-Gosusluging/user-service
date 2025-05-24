package main

import (
	"log"

	"github.com/Ryan-Gosusluging/user-service/config"
	"github.com/Ryan-Gosusluging/user-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
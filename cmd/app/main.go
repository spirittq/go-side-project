package main

import (
	"log"
	"sideq/config"
	"sideq/internal/app"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error")
	}
	err = app.Run(cfg)
	if err != nil {
		log.Fatalf("App error")
	}
}

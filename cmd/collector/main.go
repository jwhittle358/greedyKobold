package main

import (
	"fmt"
	"log"

	"github.com/jwhittle358/greedyKobold/internal/config"
)

func main() {

	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Application:", cfg.App.Name)
	fmt.Println("Environment:", cfg.App.Environment)
	fmt.Println("Log Level:", cfg.App.LogLevel)
	fmt.Println("HTTP Port:", cfg.Server.Port)

	fmt.Println("Log Collector starting...")
}

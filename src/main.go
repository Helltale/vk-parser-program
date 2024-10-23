package main

import (
	"fmt"
	"log"

	"github.com/Helltale/vk-parser-program/config"
	"github.com/Helltale/vk-parser-program/internal/logger"
)

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v\n", err)
	}

	fmt.Println(conf)

	logger, err := logger.Init(conf)
	if err != nil {
		log.Fatalf("logger error: %v\n", err)
	}
	logger.Info("info: program started")

}

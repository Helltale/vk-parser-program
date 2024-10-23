package main

import (
	"fmt"
	"log"

	"github.com/Helltale/vk-parser-program/config"
	"github.com/Helltale/vk-parser-program/internal/flags"
	"github.com/Helltale/vk-parser-program/internal/logger"
)

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v\n", err)
	}

	fmt.Println(conf)

	slogger := logger.NewSLogger()
	fileLogger, err := logger.NewFLogger(conf.AppLogfile)
	if err != nil {
		slogger.Error("Ошибка создания FileLogger", "error", err)
	}
	defer fileLogger.Close()

	logger := logger.NewCombinedLogger(slogger, fileLogger)
	logger.Info("program started")

	flagEntry := flags.FlagHandler(logger)

	fmt.Println(flagEntry)
}

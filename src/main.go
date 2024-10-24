package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Helltale/vk-parser-program/config"
	"github.com/Helltale/vk-parser-program/internal/fetcher"
	"github.com/Helltale/vk-parser-program/internal/flags"
	"github.com/Helltale/vk-parser-program/internal/logger"
)

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v\n", err)
	}

	slogger := logger.NewSLogger()
	fileLogger, err := logger.NewFLogger(conf.AppLogfile)
	if err != nil {
		slogger.Error("Ошибка создания FileLogger", "error", err)
	}
	defer fileLogger.Close()

	logger := logger.NewCombinedLogger(slogger, fileLogger)
	logger.Info("program started")

	flagManager := flags.NewFlagManager()
	flagEntry := flagManager.FlagHandler(logger)

	response, err := fetcher.Init(
		flagEntry.Flag,
		flagEntry.Value[0], //заглушка
		conf,
		logger,
	)
	if err != nil {
		os.Exit(1)
	}

	if err := fetcher.SaveResponseToJSON(response, filepath.Join(conf.AppResDir, fmt.Sprintf("%s_%s_response.json", flagEntry.Flag, flagEntry.Value))); err != nil {
		logger.Error("error saving json", "flag", flagEntry.Flag, "value", flagEntry.Value, "error", err)
	}
}

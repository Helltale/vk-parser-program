package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Helltale/vk-parser-program/config"
	"github.com/Helltale/vk-parser-program/internal/fetcher"
	"github.com/Helltale/vk-parser-program/internal/flags"
	"github.com/Helltale/vk-parser-program/internal/logger"
)

const (
	TIMEOUT = 10 * time.Second
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

	var wg sync.WaitGroup
	sem := make(chan struct{}, conf.AppMaxGoroutine)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	for _, value := range flagEntry.Value {
		wg.Add(1)
		sem <- struct{}{}

		go func(val string) {
			defer wg.Done()
			defer func() { <-sem }()

			select {
			case <-ctx.Done():
				logger.Error("context canceled", "flag", flagEntry.Flag, "value", val, "timeout", TIMEOUT, "error", ctx.Err())
				return
			default:
				time.Sleep(time.Second / time.Duration(conf.AppMaxResponceToVkToSec))

				response, err := fetcher.Init(flagEntry.Flag, val, conf, logger)
				if err != nil {
					logger.Error("error fetching data", "flag", flagEntry.Flag, "value", val, "error", err)
					return
				}

				if err := fetcher.SaveResponseToJSON(response, conf.AppResDir, val, fmt.Sprintf("%s_responce.json", flagEntry.Flag), logger); err != nil {
					logger.Error("error saving json", "flag", flagEntry.Flag, "value", val, "error", err)
				}
			}
		}(value)
	}

	wg.Wait()
}

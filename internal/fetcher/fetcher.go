package fetcher

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Helltale/vk-parser-program/config"
	"github.com/Helltale/vk-parser-program/internal/flags"
	"github.com/Helltale/vk-parser-program/internal/logger"
)

type Fetcher interface {
	CreateLink(accessToken string, version string) string
	Fetch()
}

func Init(flagEntry *flags.Entry, conf *config.Config, logger *logger.CombinedLogger) (map[string]interface{}, error) {
	switch flagEntry.Flag {
	case "user":
		user := NewUser(flagEntry.Value, "")
		url := user.CreateLink(conf.ApiToken, conf.ApiVersion)
		// fmt.Println(url)
		user.Fetch(url, logger)

		logger.Info("fetcher init get responce", "switch-case", "user", "flag", flagEntry.Flag, "value", flagEntry.Value)

		return user.user_responce, nil
	case "wall":

	default:

		logger.Error("fetcher init error", "switch-case", "default", "flag", flagEntry.Flag, "value", flagEntry.Value)
		return nil, fmt.Errorf("fetcher init error: switch-case - default")
	}

	logger.Error("fetcher init error", "switch-case", "none", "flag", flagEntry.Flag, "value", flagEntry.Value)
	return nil, fmt.Errorf("fetcher init error: switch-case - none")

}

func SaveResponseToJSON(response map[string]interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(response); err != nil {
		return err
	}

	return nil
}

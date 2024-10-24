package fetcher

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Helltale/vk-parser-program/config"
	"github.com/Helltale/vk-parser-program/internal/logger"
)

type Fetcher interface {
	CreateLink(accessToken string, version string) string
	Fetch()
}

func Init(flag string, value string, conf *config.Config, logger *logger.CombinedLogger) (map[string]interface{}, error) {
	switch flag {
	case "user":
		user := NewUser(value, "")
		url := user.CreateLink(conf.ApiToken, conf.ApiVersion)
		// fmt.Println(url)
		user.Fetch(url, logger)

		logger.Info("fetcher init get responce", "switch-case", "user", "flag", flag, "value", value)

		return user.user_responce, nil
	case "wall":

	default:

		logger.Error("fetcher init error", "switch-case", "default", "flag", flag, "value", value)
		return nil, fmt.Errorf("fetcher init error: switch-case - default")
	}

	logger.Error("fetcher init error", "switch-case", "none", "flag", flag, "value", value)
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

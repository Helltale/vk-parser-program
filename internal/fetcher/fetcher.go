package fetcher

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

func SaveResponseToJSON(response map[string]interface{}, path, dir, filename string, logger *logger.CombinedLogger) error {

	if err := directory(path, dir, logger); err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(path, dir, filename))
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

func directory(path, directory string, logger *logger.CombinedLogger) error {
	dirPath := filepath.Join(path, directory)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			logger.Error("error with create dir", "dir", directory)
			return err
		}
		return nil
	} else {
		return nil
	}
}

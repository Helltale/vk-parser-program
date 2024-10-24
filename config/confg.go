package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiToken   string `yaml:"api-token"`
	ApiVersion string `yaml:"api-version"`
	AppLogfile string `yaml:"app_logfile"`
	AppResDir  string `yaml:"app_resdir"`
	AppHost    string `yaml:"app_host"`
	AppPort    string `yaml:"app_port"`
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbName     string `yaml:"db_name"`
}

func NewConfig() (*Config, error) {
	configPath, err := filepath.Abs(filepath.Join("..", "config", "config.yaml"))
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

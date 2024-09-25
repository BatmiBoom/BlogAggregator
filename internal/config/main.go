package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(user_name string) error {
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	cfg.CurrentUserName = user_name

	handler, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ERROR: opening file %v", err)
	}

	jsonString, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("ERROR: converting struct %v", err)
	}

	_, err = handler.Write(jsonString)
	if err != nil {
		return fmt.Errorf("ERROR: writing file %v", err)
	}

	return nil
}

func Read() (*Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	stringContent, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR: reading config file %v", err)
	}

	config := Config{}
	err = json.Unmarshal(stringContent, &config)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Marshalling config %v", err)
	}

	return &config, nil
}

func getConfigFilePath() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("ERROR: wrong path %v", err)
	}
	return fmt.Sprintf("%s/workspace/github.com/batmiboom/blog_aggregator/%s", homeDirectory, configFileName), nil
}

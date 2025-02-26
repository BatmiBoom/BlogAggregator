package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const ConfigFileName string = ".gatorconfig.json"

type Config struct {
	Db_url           string `json:"db_url"`
	Current_username string `json:"current_username"`
}

func (cfg *Config) SetUser(username string) error {
	config_path, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("ERROR: Obtaining the config file %s", err)
	}

	file, err := os.OpenFile(config_path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("ERROR: Opening config file %s", err)
	}
	defer file.Close()

	cfg.Current_username = username

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("ERROR: Writing to file %s", err)
	}

	return nil
}

func ReadConfig() (Config, error) {
	config_path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	fmt.Println(config_path)
	file, err := os.Open(config_path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("ERROR: Decoding json %s", err)
	}

	return cfg, nil
}

func getConfigPath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	config_path := home_dir + "/" + ConfigFileName
	if !fileExists(config_path) {
		return "", fmt.Errorf("Config file doesn't exists, check in path : %s", config_path)
	}

	return config_path, nil
}

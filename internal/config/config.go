package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(c)
	if err != nil {
		return fmt.Errorf("couldn't write user to config: %w", err)
	}
	return nil
}

func Read() (Config, error) {

	fileLocation, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file:%s", err)
	}

	fileByte, err := os.ReadFile(fileLocation)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading file:%w", err)
	}
	var cfg Config

	if err := json.Unmarshal(fileByte, &cfg); err != nil {
		return Config{}, fmt.Errorf("Error transforming to json: %w", err)
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting home director:%s", err)
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg *Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path:%w", err)
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, 0644)
}

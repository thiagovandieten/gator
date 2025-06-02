package config

import (
	"encoding/json"
	"os"
)

const (
	configFileName string = ".gatorconfig.json"
	defaultDB      string = "postgres://example"
)

type Config struct {
	DBinURL  string `json:"db_url"`
	Username string `json:"name"`
}

func GetConfig() (Config, error) {
	homedir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(homedir + configFileName)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homedir = homedir + "/"
	return homedir, err
}

func SetUser(username string, cfg Config) error {
	cfg.Username = username
	err := write(&cfg)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg *Config) error {
	if cfg.DBinURL == "" {
		cfg.DBinURL = defaultDB
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	homepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(homepath+configFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

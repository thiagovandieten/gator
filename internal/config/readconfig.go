package config

import (
	"encoding/json"
	"os"
)

func GetConfig() (Config, error) {
	homedir, err := GetConfigFilePath()
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

	return Config{}, nil
}

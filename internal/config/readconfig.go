package config

import (
	"encoding/json"
	"os"
)

func GetConfig(username string) (Config, error) {
	homedir, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(homedir + configName)
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

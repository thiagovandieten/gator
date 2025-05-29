package config

import (
	"encoding/json"
	"os"
)

func SetUser(username string, cfg Config) error {
	cfg.Username = username
	err := write(&cfg)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg *Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	homepath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(homepath+configFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

package config

import "os"

func GetConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homedir = homedir + "/"
	return homedir, err
}

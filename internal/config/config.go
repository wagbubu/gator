package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
} 

func Read() (Config, error) {
	fullpath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(fullpath)
	if err != nil {
		return Config{}, err	
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "" , err
	}
	fullpath := filepath.Join(home, configFileName)
	return fullpath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	} 

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
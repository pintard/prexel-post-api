package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBHost     string `json:"DB_HOST"`
	DBPort     string `json:"DB_PORT"`
	DBUser     string `json:"DB_USER"`
	DBPassword string `json:"DB_PASSWORD"`
	DBName     string `json:"DB_NAME"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("Could not decode config file: %v", err)
	}

	config.DBHost = GetEnv("DB_HOST", config.DBHost)
	config.DBPort = GetEnv("DB_PORT", config.DBPort)
	config.DBUser = GetEnv("DB_USER", config.DBUser)
	config.DBPassword = GetEnv("DB_PASSWORD", config.DBPassword)
	config.DBName = GetEnv("DB_NAME", config.DBName)

	return config, nil
}

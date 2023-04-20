package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config is a struct to hold the configuration values
type Config struct {
	DBHost        string `json:"db_host"`
	DBPort        string `json:"db_port"`
	DBDatabase    string `json:"db_database"`
	RedisHost     string `json:"redis_host"`
	RedisPort     string `json:"redis_port"`
	RedisPassword string `json:"redis_password"`
	RedisDb       int    `json:"redis_db"`
	APIBaseUrl    string `json:"api_base_url"`
}

// LoadConfig loads the configuration values from a JSON file
func LoadConfig(filePath string) (*Config, error) {
	// Read the file contents
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// Unmarshal the JSON contents into a Config struct
	var config Config
	err = json.Unmarshal(fileContents, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config JSON: %v", err)
	}

	return &config, nil
}

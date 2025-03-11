package config

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration holds all config values
type Configuration struct {
	Database struct {
		URL string `json:"url"`
	} `json:"database"`
}

// Config is the global configuration instance
var Config Configuration

// LoadConfig loads configuration from config.json
func LoadConfig() {
	// Try different paths for the config file
	configPaths := []string{"../config.json", "../../config.json", "./config.json"}

	var configFile *os.File
	var err error

	for _, path := range configPaths {
		configFile, err = os.Open(path)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatal("Error opening config file:", err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
	}
}

package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("../")    // Look for config.json in the parent directory
	viper.AddConfigPath("../../") // Additional path to handle different run locations
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file:", err)
	}
}

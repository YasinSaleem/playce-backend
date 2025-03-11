package config

import (
    "log"
    "github.com/spf13/viper"
)

func LoadConfig() {
    viper.SetConfigFile("config.json")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatal("Error loading config file", err)
    }
}

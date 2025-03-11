package config

import (
	"log"
	"user_service/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open("postgres", viper.GetString("database.url"))
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&models.User{})
}

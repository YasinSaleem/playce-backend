package config

import (
	"log"
	"post_service/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open("postgres", Config.Database.URL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Post{})
}

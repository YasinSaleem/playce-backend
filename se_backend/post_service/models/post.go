package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}

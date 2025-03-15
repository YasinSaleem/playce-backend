package models

import (
	"time"

    "github.com/jinzhu/gorm"
)

type User struct {
	UserID   uint   `json:"user_id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserProfile struct {
    gorm.Model
	ProfilePic string    `json:"profile_pic"`
	Username   string    `json:"username"`
	DOB        time.Time `json:"dob"`
	Bio        string    `json:"bio"`
	City       string    `json:"city"`
	School     string    `json:"school"`
	User       User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

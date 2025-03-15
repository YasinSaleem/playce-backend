package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	PostID     int    `json:"post_id"`    // Unique identifier for the post
	Title      string `json:"title"`      // Title of the post
	Body       string `json:"body"`       // Content of the post
	Pic        string `json:"pic"`        // URL or path to the picture associated with the post
	No_Likes   int    `json:"no_likes"`   // Number of likes for the post
	User_Liked bool   `json:"user_liked"` // Whether the current user has liked the post
	UserID     int    `json:"user_id"`    // ID of the user who created the post
}

type Like struct {
	gorm.Model
	PostID int `json:"post_id"` // ID of the post being liked
	UserID int `json:"user_id"` // ID of the user who liked the post
	Post   Post `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"` // Foreign key constraint
}

type Comment struct {
	gorm.Model
	PostID int    `json:"post_id"` // ID of the post being commented on
	UserID int    `json:"user_id"` // ID of the user who commented
	Body   string `json:"body"`    // Comment content
	Post   Post   `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"` // Foreign key constraint
}


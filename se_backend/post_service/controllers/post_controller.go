package controllers

import (
	"net/http"
	"post_service/config"
	"post_service/models"

	"github.com/gin-gonic/gin"
)

// GetAllPosts - Fetches all posts from the database
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetUserPosts - Fetches all posts of a specific user
func GetUserPosts(c *gin.Context) {
	userId := c.Param("userId")
	var posts []models.Post
	if err := config.DB.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// CreatePost - Creates a new post
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the user ID matches the token claims (assuming user ID is in the token)
	userId := c.GetInt("user_id")
	if post.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID mismatch"})
		return
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create post"})
		return
	}
	c.JSON(http.StatusCreated, post)
}

package controllers

import (
	"net/http"
	"post_service/config"
	"post_service/models"

	"github.com/labstack/echo/v4"
)

// GetAllPosts - Fetches all posts from the database
func GetAllPosts(c echo.Context) error {
	var posts []models.Post
	if err := config.DB.Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to retrieve posts"})
	}
	return c.JSON(http.StatusOK, posts)
}

// GetUserPosts - Fetches all posts of a specific user
func GetUserPosts(c echo.Context) error {
	userId := c.Param("userId")
	var posts []models.Post
	if err := config.DB.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to retrieve user posts"})
	}
	return c.JSON(http.StatusOK, posts)
}

// CreatePost - Creates a new post
func CreatePost(c echo.Context) error {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Get the email from the context (set by the auth middleware)
	// We're not using the email in this simplified example, but in a real app
	// you would use it to associate the post with the correct user
	_ = c.Get("email").(string)

	if err := config.DB.Create(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create post"})
	}
	return c.JSON(http.StatusCreated, post)
}

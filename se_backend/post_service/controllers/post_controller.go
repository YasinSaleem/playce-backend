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

func LikePost(c echo.Context) error {
    var like models.Like
    var post models.Post

    // Bind the post ID and user ID directly from the request body (assuming JSON input)
    if err := c.Bind(&like); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    // Check if the post exists
    if err := config.DB.First(&post, like.PostID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
    }

    // Check if the like already exists (to toggle like/unlike)
    if err := config.DB.Where("post_id = ? AND user_id = ?", like.PostID, like.UserID).First(&like).Error; err == nil {
        // Unlike the post if the like already exists
        if err := config.DB.Delete(&like).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to unlike the post"})
        }
        post.No_Likes--
        if err := config.DB.Save(&post).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update like count"})
        }
        return c.JSON(http.StatusOK, map[string]string{"message": "Post unliked"})
    }

    // Add a new like
    if err := config.DB.Create(&like).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to like the post"})
    }

    // Increment like count
    post.No_Likes++
    if err := config.DB.Save(&post).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update like count"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Post liked"})
}

func CommentOnPost(c echo.Context) error {
	var comment models.Comment

	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create comment"})
	}
	return c.JSON(http.StatusCreated, comment)
}


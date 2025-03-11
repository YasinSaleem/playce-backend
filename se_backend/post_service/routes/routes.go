package routes

import (
	"post_service/controllers"
	"post_service/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Public Routes
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:userId", controllers.GetUserPosts)

	// Protected Routes
	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.POST("/posts", controllers.CreatePost)
	}

	return r
}

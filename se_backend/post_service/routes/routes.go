package routes

import (
	"post_service/controllers"
	"post_service/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes() *echo.Echo {
	// Create a new Echo instance
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.LoggerMiddleware())

	// Public Routes
	e.GET("/posts", controllers.GetAllPosts)
	e.GET("/posts/:user_id", controllers.GetUserPosts)

	// Protected Routes
	posts := e.Group("")
	posts.Use(middlewares.AuthMiddleware())
	posts.POST("/posts", controllers.CreatePost)
	posts.POST("/posts/like", controllers.LikePost)
	posts.POST("/posts/:post_id/comment", controllers.CommentOnPost)

	return e
}

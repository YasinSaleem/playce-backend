package routes

import (
	"user_service/controllers"
	"user_service/middlewares"

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

	// User routes
	e.POST("/user/signup", controllers.SignUp)
	e.POST("/user/signin", controllers.SignIn)
	e.POST("/user/forgot-password", controllers.ForgotPassword)
	e.POST("/user/reset-password", controllers.ResetPassword)
	e.GET("/user/profile/:user_id", controllers.GetUserProfile)

	// Protected routes
	posts := e.Group("")
	posts.Use(middlewares.AuthMiddleware())

	// User Profile routes
	posts.POST("user/post-user-profile", controllers.PostUserProfile)
	
	return e
}

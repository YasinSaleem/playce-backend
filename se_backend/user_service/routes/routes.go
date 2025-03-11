package routes

import (
    "github.com/gin-gonic/gin"
    "user_service/controllers"
)

func SetupRoutes() *gin.Engine {
    router := gin.Default()
    userRoutes := router.Group("/user")
    {
        userRoutes.POST("/signup", controllers.SignUp)
        userRoutes.POST("/signin", controllers.SignIn)
        userRoutes.POST("/forgot-password", controllers.ForgotPassword)
        userRoutes.POST("/reset-password", controllers.ResetPassword)
    }
    return router
}

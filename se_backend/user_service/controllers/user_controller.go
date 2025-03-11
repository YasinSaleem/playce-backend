package controllers

import (
	"net/http"
	"user_service/config"
	"user_service/models"
	"user_service/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPassword(hashed, plain string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
    return err == nil
}

func SignUp(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
        return
    }
    user.Password = hashedPassword

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func SignIn(c *gin.Context) {
    var credentials models.User
    if err := c.BindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    if !CheckPassword(user.Password, credentials.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    token, err := utils.GenerateToken(user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func ForgotPassword(c *gin.Context) {
    var input struct {
        Email string `json:"email"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    resetToken, err := utils.GenerateResetToken(user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Reset token generated", "token": resetToken})
}

func ResetPassword(c *gin.Context) {
    var input struct {
        Email       string `json:"email"`
        Token       string `json:"token"`
        NewPassword string `json:"new_password"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    claims, valid := utils.ValidateResetToken(input.Email, input.Token)
    if !valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }

    hashedPassword, err := HashPassword(input.NewPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
        return
    }

    if err := config.DB.Model(&models.User{}).Where("email = ?", claims.Email).Update("password", hashedPassword).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Password update failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

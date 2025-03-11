package controllers

import (
	"net/http"
	"user_service/config"
	"user_service/models"
	"user_service/utils"

	"github.com/labstack/echo/v4"
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

func SignUp(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Password hashing failed"})
	}
	user.Password = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already exists"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}

func SignIn(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	if !CheckPassword(user.Password, credentials.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func ForgotPassword(c echo.Context) error {
	var request struct {
		Email string `json:"email"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	var user models.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		// Don't reveal that the email doesn't exist
		return c.JSON(http.StatusOK, map[string]string{"message": "If your email is registered, you will receive a reset link"})
	}

	// Generate reset token
	resetToken, err := utils.GenerateResetToken(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate reset token"})
	}

	// In a real application, send an email with the reset token
	// For this example, we'll just return it
	return c.JSON(http.StatusOK, map[string]string{
		"message": "If your email is registered, you will receive a reset link",
		"token":   resetToken, // This would normally be sent via email
	})
}

func ResetPassword(c echo.Context) error {
	var request struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Validate token and get user email
	claims, valid := utils.ValidateResetToken("", request.Token)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid or expired token"})
	}

	// Find user by email
	var user models.User
	if err := config.DB.Where("email = ?", claims.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
	}

	// Hash new password
	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Password hashing failed"})
	}

	// Update user password
	user.Password = hashedPassword
	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update password"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Password reset successfully"})
}

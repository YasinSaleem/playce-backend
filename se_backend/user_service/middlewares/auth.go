package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// JWTSecret is the secret key for JWT token validation
const JWTSecret = "my_secret_key"

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authorization format"})
			}

			tokenString := parts[1]

			// Parse and validate the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Validate the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(JWTSecret), nil
			})

			// Check for parsing errors and token validity
			if err != nil {
				fmt.Println("Token parsing error:", err) // Debug: Print the error
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
			}

			// Extract claims and validate expiration
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				exp, ok := claims["exp"].(float64)
				if !ok {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
				}

				// Check if the token has expired
				if int64(exp) < time.Now().Unix() {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token has expired"})
				}

				// Add user ID to context for further use
				userID, ok := claims["user_id"].(float64)
				if !ok {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID in token"})
				}
				c.Set("user_id", int(userID))

				fmt.Println("Authenticated user ID:", int(userID)) 
				return next(c)
			} else {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
			}
		}
	}
}

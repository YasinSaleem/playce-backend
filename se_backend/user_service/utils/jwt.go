package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key")
var resetSecretKey = []byte("your_reset_secret_key")

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtKey)
}

// GenerateResetToken creates a reset token for password recovery
func GenerateResetToken(email string) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour) // Token expires in 1 hour
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(resetSecretKey)
}

// ValidateResetToken checks the validity of the reset token
func ValidateResetToken(email, tokenString string) bool {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return resetSecretKey, nil
    })

    if err != nil || !token.Valid {
        return false
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || claims.Email != email {
        return false
    }

    return true
}

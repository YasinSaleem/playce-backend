package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "fmt"
)

// Secret keys for JWT tokens
var JwtKey = []byte("my_secret_key")
var resetSecretKey = []byte("your_reset_secret_key")

// Claims structure to store user data and standard claims
type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

// GenerateToken creates a new JWT token for authentication
func GenerateToken(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(JwtKey)
    if err != nil {
        fmt.Println("Error generating token:", err)  // Debug: Print error
        return "", err
    }
    fmt.Println("Generated token for:", email)  // Debug: Print success
    return tokenString, nil
}

// GenerateResetToken creates a reset token for password recovery
func GenerateResetToken(email string) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour) // Token valid for 1 hour
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(resetSecretKey)
    if err != nil {
        fmt.Println("Error generating reset token:", err)  // Debug: Print error
        return "", err
    }
    fmt.Println("Generated reset token for:", email)  // Debug: Print success
    return tokenString, nil
}

// ValidateToken checks the validity of a regular JWT token
func ValidateToken(tokenString string) (*Claims, bool) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })

    if err != nil {
        fmt.Println("Error validating token:", err)  // Debug: Print error
        return nil, false
    }

    if !token.Valid {
        fmt.Println("Invalid token for email:", claims.Email)  // Debug: Print error
        return nil, false
    }

    return claims, true
}

// ValidateResetToken checks the validity of the reset token
func ValidateResetToken(email, tokenString string) (*Claims, bool) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return resetSecretKey, nil
    })

    if err != nil {
        fmt.Println("Error validating reset token:", err)  // Debug: Print error
        return nil, false
    }

    if !token.Valid {
        fmt.Println("Invalid reset token for email:", claims.Email)  // Debug: Print error
        return nil, false
    }

    // Check if the email from claims matches the given email
    if claims.Email != email {
        fmt.Println("Token email mismatch for email:", claims.Email)  // Debug: Print error
        return nil, false
    }

    // Check if the token has expired
    if claims.ExpiresAt < time.Now().Unix() {
        fmt.Println("Reset token has expired for email:", claims.Email)  // Debug: Print expiration message
        return nil, false
    }

    return claims, true
}

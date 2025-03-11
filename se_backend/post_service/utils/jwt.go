package utils

import (
    "errors"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/spf13/viper"
)

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func VerifyToken(tokenString string) (string, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(viper.GetString("jwt.secret")), nil
    })

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims.Email, nil
    }

    return "", err
}

func ValidateUser(userID uint, token string) error {
    resp, err := http.Get("http://localhost:8080/user/validate?user_id=" + string(rune(userID)) + "&token=" + token)
    if err != nil || resp.StatusCode != http.StatusOK {
        return errors.New("user validation failed")
    }
    return nil
}

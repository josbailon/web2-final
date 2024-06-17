package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "parking-management/config"
)

func GenerateJWT(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JwtSecret))
}

func ParseJWT(tokenStr string) (uint, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.JwtSecret), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID := uint(claims["user_id"].(float64))
        return userID, nil
    } else {
        return 0, err
    }
}

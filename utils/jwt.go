package utils

import (
    "time"
    "os"
    "github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userId uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userId,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

package middleware

import (
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
)

func Protect(c *fiber.Ctx) error {
    header := c.Get("Authorization")
    if header == "" || !strings.HasPrefix(header, "Bearer ") {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    tokenStr := strings.TrimPrefix(header, "Bearer ")
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
    }

    claims := token.Claims.(jwt.MapClaims)
    userId := uint(claims["user_id"].(float64))
    c.Locals("user_id", userId)

    return c.Next()
}

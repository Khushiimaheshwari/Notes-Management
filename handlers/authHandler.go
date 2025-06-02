package handlers

import (
    "CodingNinja_Assignment/models"
    "CodingNinja_Assignment/utils"
    "github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
    var input models.User
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    if input.Email == "" || input.Password == "" || input.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing fields"})
    }

    input.Password = utils.HashPassword(input.Password)

    result := models.DB.Create(&input)
    if result.Error != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
    }

    return c.JSON(fiber.Map{"message": "Registered successfully"})
}

func Login(c *fiber.Ctx) error {
    var input models.User
    var user models.User

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    models.DB.Where("email = ?", input.Email).First(&user)
    if user.ID == 0 || !utils.CheckPassword(input.Password, user.Password) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
    }

    return c.JSON(fiber.Map{"token": token})
}

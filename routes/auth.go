package routes

import (
    "CodingNinja_Assignment/handlers"
    "github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
    app.Post("/register", handlers.Register)
    app.Post("/login", handlers.Login)
}

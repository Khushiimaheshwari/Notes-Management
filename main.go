package main

import (
    "log"
    "os"
    "CodingNinja_Assignment/routes"
    "CodingNinja_Assignment/models"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    models.ConnectDB()

    app := fiber.New()

    routes.AuthRoutes(app)
    routes.NoteRoutes(app)

    log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

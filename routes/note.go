package routes

import (
    "CodingNinja_Assignment/handlers"
    "CodingNinja_Assignment/middleware"
    "github.com/gofiber/fiber/v2"
)

func NoteRoutes(app *fiber.App) {
    group := app.Group("/notes", middleware.Protect)
    group.Post("/", handlers.CreateNote)
    group.Get("/", handlers.GetNotes)
    group.Get("/:id", handlers.GetNote)
    group.Put("/:id", handlers.UpdateNote)
    group.Delete("/:id", handlers.DeleteNote)
}

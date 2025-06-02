package handlers

import (
    "CodingNinja_Assignment/models"
    "github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {
    userId := c.Locals("user_id").(uint)
    var note models.Note

    if err := c.BodyParser(&note); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    note.UserID = userId
    models.DB.Create(&note)
    return c.JSON(note)
}

func GetNotes(c *fiber.Ctx) error {
    userId := c.Locals("user_id").(uint)
    var notes []models.Note
    models.DB.Where("user_id = ?", userId).Find(&notes)
    return c.JSON(notes)
}

func GetNote(c *fiber.Ctx) error {
    userId := c.Locals("user_id").(uint)
    var note models.Note

    models.DB.First(&note, c.Params("id"))
    if note.ID == 0 || note.UserID != userId {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
    }

    return c.JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
    userId := c.Locals("user_id").(uint)
    var note models.Note

    models.DB.First(&note, c.Params("id"))
    if note.ID == 0 || note.UserID != userId {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
    }

    var input models.Note
    c.BodyParser(&input)

    note.Title = input.Title
    note.Content = input.Content
    models.DB.Save(&note)

    return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
    userId := c.Locals("user_id").(uint)
    var note models.Note

    models.DB.First(&note, c.Params("id"))
    if note.ID == 0 || note.UserID != userId {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
    }

    models.DB.Delete(&note)
    return c.JSON(fiber.Map{"message": "Deleted"})
}

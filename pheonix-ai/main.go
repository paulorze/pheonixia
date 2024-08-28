package main

import (
	"log"
	"os"

	"pheonix/pdf_processor"
	"pheonix/pheonix"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error al cargar el archivo")
		}

		tempFile, err := os.CreateTemp("", "uploaded-*.pdf")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error al crear archivo temporal")
		}
		defer os.Remove(tempFile.Name())

		if err := c.SaveFile(file, tempFile.Name()); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error al guardar archivo")
		}

		apiKey := "tu-clave-api"
		ai := pheonix.NewPhoenixIA(apiKey)

		response, err := pdf_processor.ProcessPDF(tempFile.Name(), ai)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error al procesar PDF")
		}

		return c.JSON(fiber.Map{"message": response})
	})

	log.Fatal(app.Listen(":8080"))
}

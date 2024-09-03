package engine

import (
	"phoenixia/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) PDFLoader(context *fiber.Ctx) error {
	var file []byte
	err := context.BodyParser(&file)
	if err != nil {
		context.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "request failed",
		})
		return nil
	}

	tableName, err := h.EngineService.PDFLoader(file)
	if err != nil {
		utils.ErrorResponseCreator(context, err)
	}
	context.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message":   "user created successfully",
		"tablename": tableName,
	})

	return nil
}

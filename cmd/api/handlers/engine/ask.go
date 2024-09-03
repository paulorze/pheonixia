package engine

import (
	"phoenixia/internal/domain"
	"phoenixia/utils"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Ask(context *fiber.Ctx) error {
	var req domain.AskRequest
	err := context.BodyParser(&req)
	if err != nil {
		context.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "request failed",
		})
		return nil
	}

	docs, err := h.EngineService.DocumentsRetriever(req.TableName, req.Query)
	if err != nil {
		context.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "request failed",
		})
		return nil
	}

	response, err := h.EngineService.Ask(docs, "", req.Query)
	if err != nil {
		utils.ErrorResponseCreator(context, err)
	}

	context.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"response": response,
	})

	return nil
}

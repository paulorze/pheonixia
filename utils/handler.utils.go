package utils

import (
	customErrors "phoenixia/errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponseCreator(context *fiber.Ctx, err error) error {
	if customErr, ok := err.(*customErrors.CustomError); ok {
		switch customErr.Code {
		case 400:
			context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": customErr.Message,
			})
			return nil
		case 401:
			context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": customErr.Message,
			})
			return nil
		case 404:
			context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": customErr.Message,
			})
			return nil
		case 409:
			context.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": customErr.Message,
			})
			return nil
		case 422:
			context.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": customErr.Message,
			})
			return nil
		case 500:
			context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": customErr.Message,
			})
			return nil
		default:
			context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "something unexpected happened",
			})
			return nil
		}
	} else {
		context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
		return nil
	}

}

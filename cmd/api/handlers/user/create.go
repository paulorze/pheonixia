package user

import (
	"phoenixia/internal/domain"
	"phoenixia/utils"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Create(context *fiber.Ctx) error {
	user := domain.User{}
	err := context.BodyParser(&user)
	if err != nil {
		context.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "request failed",
		})
		return nil
	}
	err = h.UserService.Create(user)
	if err != nil {
		utils.ErrorResponseCreator(context, err)
	}
	context.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "user created successfully",
	})
	return nil
}
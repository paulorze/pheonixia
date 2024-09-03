package user

import (
	"phoenixia/utils"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetAll(context *fiber.Ctx) error {
	usersList, err := h.UserService.GetAll()
	if err != nil {
		utils.ErrorResponseCreator(context, err)
	}
	context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": usersList,
	})
	return nil
}

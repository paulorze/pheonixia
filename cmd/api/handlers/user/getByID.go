package user

import (
	"phoenixia/utils"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetByID(context *fiber.Ctx) error {
	id := context.Params("id")
	user, err := h.UserService.GetByID(id)
	if err != nil {
		utils.ErrorResponseCreator(context, err)
	}
	context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
	return nil
}

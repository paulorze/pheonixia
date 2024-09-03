package user

import (
	"phoenixia/utils"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Delete(context *fiber.Ctx) error {
	id := context.Params("id")
	err := h.UserService.Delete(id)
	if err != nil {
		utils.ErrorResponseCreator(context, err)
	}
	context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "user deleted successfully",
	})
	return nil
}

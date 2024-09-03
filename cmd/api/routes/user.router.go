package routes

import (
	"phoenixia/cmd/api/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App, h *user.Handler) {
	api := app.Group("/api/users")
	api.Post("/register", h.Create)
	api.Get("/", h.GetAll)
	api.Get("/:id", h.GetByID)
	api.Put("/:id", h.Update)
	api.Delete("/:id", h.Delete)
}

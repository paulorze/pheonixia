package routes

import (
	"phoenixia/cmd/api/handlers/engine"

	"github.com/gofiber/fiber/v2"
)

func EngineRouter(app *fiber.App, h *engine.Handler) {
	api := app.Group("/api/engine")
	api.Post("/ask", h.Ask)
	api.Post("/upload", h.PDFLoader)
}

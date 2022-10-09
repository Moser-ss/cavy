package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moser-ss/cavy/handlers"
)

func ReadinessRouter(app fiber.Router) {
	app.Get("/readyz", handlers.GetReady())
}

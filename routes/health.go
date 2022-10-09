package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moser-ss/cavy/handlers"
)

func HealthRouter(app fiber.Router) {
	app.Get("/healthz", handlers.GetHealth())
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moser-ss/cavy/api/handlers"
	"github.com/moser-ss/cavy/pkg/config"
)

func HealthRouter(app fiber.Router, health config.EndpointConfig) {
	app.Get("/healthz", handlers.GetProbe(health))
}

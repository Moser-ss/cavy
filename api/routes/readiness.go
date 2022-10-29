package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moser-ss/cavy/api/handlers"
	"github.com/moser-ss/cavy/pkg/config"
)

func ReadinessRouter(app fiber.Router, ready config.EndpointConfig) {
	app.Get("/readyz", handlers.GetProbe(ready))
}

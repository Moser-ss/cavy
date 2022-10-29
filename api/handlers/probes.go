package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/moser-ss/cavy/pkg/config"
)

func GetProbe(probeConfig config.EndpointConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Printf("Route %s called", c.Path())
		pd := probeConfig.Delay
		if pd != 0 {
			log.Printf("Route %s has a delay of %d seconds", c.Path(), pd)
			time.Sleep(time.Duration(pd) * time.Second)

		}
		c.Status(probeConfig.Status)
		return c.SendString(probeConfig.Body)
	}
}

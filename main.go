package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config, _ := LoadConfig()

	log.Printf("%+v\n", config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World! I am Cavy!")
	})
	log.Println("Starting Cavy")
	sd := config.Server.StartDelay
	if sd != 0 {
		log.Printf("Cavy has a slow start of %d seconds", sd)
		time.Sleep(time.Duration(sd) * time.Second)

	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shuting down Cavy")
		ssd := config.Server.ShutdownDelay
		if ssd != 0 {
			log.Printf("Cavy has a slow shutdown of %d seconds", ssd)
			time.Sleep(time.Duration(ssd) * time.Second)

		}
		app.Shutdown()
	}()

	app.Listen(":3000")
}

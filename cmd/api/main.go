package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.Int("port", 8000, "Port that servert listen to")
	flag.Parse()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Go fiber")
	})
	app.Listen(fmt.Sprintf(":%d", *port))
}

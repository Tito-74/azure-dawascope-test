package main

import (
	// "fmt"

	"log"

	"github.com/gofiber/fiber/v2"
)

func TestDocker(c *fiber.Ctx) error{
	return c.JSON("hello Docker")
}
func main() {
	app:= fiber.New()

	app.Get("/docker", TestDocker)

	log.Fatal(app.Listen(":3000"))
}
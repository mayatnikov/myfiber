package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RouteJson3(c *fiber.Ctx) error {
	jj := MyJson{
		Par1: "P1xxxxxx",
		Par2: "P2yyyyyy",
		Par3: "P3zzzzzz",
	}
	log.Print(jj)
	return c.JSON(jj)
}

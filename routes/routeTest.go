package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteTest2(app *fiber.App) {
	app.Get("/TEST", func(c *fiber.Ctx) error {
		jj := MyJson{Par1: "P1-TEST", Par2: "P2-TEST", Par3: "P3-TEST"}
		return c.JSON(jj)
	})
}

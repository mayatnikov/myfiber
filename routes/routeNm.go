package routes

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RouteNm(c *fiber.Ctx) error {

	hostname, _ := os.Hostname()

	// Render index template
	c.Response().Header.Set("User-Header", "My special value")

	time.Sleep(1000 * time.Millisecond)
	return c.Render("index", fiber.Map{
		"Title":    c.Params("name"),
		"par1":     "Параграф № 1",
		"par2":     "Paragraph N 2",
		"hostname": hostname,
		"BOOLPAR":  "TRUE",
	})
}

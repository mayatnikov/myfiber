package routes

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RouteJson1(c *fiber.Ctx) error {

	hostname, _ := os.Hostname()
	jj := MyJson{
		Par1: "Param-11",
		Par2: "P2",
		Par3: "P3",
		Hostname: hostname,
	}
	log.Print(jj)
	time.Sleep(1000 * time.Millisecond)
	return c.JSON(jj)
}

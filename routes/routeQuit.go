package routes

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RouteQuit(c *fiber.Ctx) error {

	log.Print("request exit from Application")
	time.Sleep(1000 * time.Millisecond)
	os.Exit(1)
	return nil
}

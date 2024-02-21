package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type SomeStruct struct {
	Name string
	Age  uint8
	Myj1 MyJson
}

func RouteJson2(c *fiber.Ctx) error {
	// Create data struct:
	data := SomeStruct{
		Name: "IVAN",
		Age:  68,
		Myj1: MyJson{Par1: "1111", Par2: "2222", Par3: "3333"},
	}
	time.Sleep(500 * time.Millisecond)
	return c.JSON(data)
}

package main

/**
мои опыты:
   Fiber
**/

import (
	"log"
	"os"

	"myfiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/mustache"
)

func main() {
	// Initialize standard Go html template engine
	// engine := html.New("./views", ".html")

	var views string
	if len(os.Args) > 1 {
		views = os.Args[1]
	} else {
		views = "./views"
	}

	app := Setup(views)

	log.Println("Try: curl -X GET  http://minic:8008/json[1|2|3")
	log.Println("Try: curl -X GET  http://minic:8008/nm/any_Name")
	log.Println("Try: curl -X GET  http://minic:8008/TEST")

	// ---------- start server --------
	log.Fatal(app.Listen(":3001"))
}

func Setup(views string) *fiber.App {

	engine := mustache.New(views, ".mustache")

	app := fiber.New(fiber.Config{
		Views:         engine,
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "проверка работы Fiber: App v1.0.1",
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:     "[${time}] ${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
		TimeFormat: "2006 Jan 02 15:04:05 MST", // магическая дата ---> Mon Jan 2 15:04:05 MST 2006
		TimeZone:   "GMT+3",
	}))

	app.Get("/json1", routes.RouteJson1)
	app.Get("/json2", routes.RouteJson2)
	app.Get("/json3", routes.RouteJson3)
	app.Get("/nm/:name", routes.RouteNm) // curl localhost:3001/nm/MyName
	routes.RouteTest2(app)
	app.Get("/quit", routes.RouteQuit)

	return app
}

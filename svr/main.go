package main

import "github.com/gofiber/fiber"
import "github.com/gofiber/cors"
import "github.com/adayswait/mojo/utils"
import "github.com/adayswait/mojo/router"
import "github.com/adayswait/mojo/ws"

func main() {
	app := fiber.New()
	app.Use(cors.New())
	router.Route(app)

	app.Use("/ws", ws.New())
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})
	app.Listen(int(utils.GetListeningPort()))
}

package main

import "github.com/adayswait/mojo/mlog"
import "github.com/gofiber/fiber"
import "github.com/adayswait/mojo/utils"
import "github.com/adayswait/mojo/router"
import "github.com/adayswait/mojo/ws"

func main() {
	app := fiber.New()
	router.Route(app)

	app.Use("/ws", ws.New())
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})
	mlog.Log("hsajdhsajdlhas")
	app.Listen(int(utils.GetListeningPort()))
}

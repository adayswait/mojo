package main

import (
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/router"
	"github.com/adayswait/mojo/utils"
	"github.com/adayswait/mojo/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/glog"
)

func main() {
	glog.Info("welcome to mojo")
	app := fiber.New()
	router.Route(app)

	app.Use("/ws", ws.New())
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})
	err := app.Listen(int(utils.GetWebPort()))
	if err != nil {
		mlog.Log("Listening port:", utils.GetWebPort(), "error ", err)
	}
}

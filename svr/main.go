package main

import (
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/msignal"
	"github.com/adayswait/mojo/router"
	"github.com/adayswait/mojo/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	mlog.Init()
	defer mlog.Flush()

	go msignal.CatchSignal()
	app := fiber.New()
	router.Route(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	mlog.Fatal(app.Listen(utils.GetWebHost() + ":" + utils.GetWebPort()))
}

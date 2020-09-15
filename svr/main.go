package main

import (
	"flag"
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/router"
	"github.com/adayswait/mojo/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("welcome to mojo")
	app := fiber.New()
	router.Route(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	err := app.Listen(utils.GetWebHost() + ":" + utils.GetWebPort())
	if err != nil {
		mlog.Log("Listening port:", utils.GetWebPort(), "error ", err)
	}
}

package router

import "github.com/gofiber/fiber"
import "github.com/gofiber/websocket"
import "github.com/adayswait/mojo/handler"

func Route(app *fiber.App) {
	app.Static("/", "../cli/")
	app.Get("/ws", websocket.New(handler.Websocket))
	app.Get("/test", func(c *fiber.Ctx) {
		c.Send("test")
	})
	api := app.Group("/api")
	web := app.Group("/web")

	authApi := api.Group("/auth")
	authApi.Get("/login", handler.Login)
	authApi.Get("/logout", handler.Logout)
	authApi.Get("/register", handler.Register)
	authWeb := web.Group("/auth")
	authWeb.Get("/login", handler.Login)
}

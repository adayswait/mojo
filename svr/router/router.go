package router

import "github.com/adayswait/mojo/handler"
import "github.com/gofiber/websocket"
import "github.com/gofiber/fiber"

func Route(app *fiber.App) {
	app.Get("/ws", websocket.New(handler.Websocket))
	api := app.Group("/api")
	api.Get("/", handler.Hello)

	auth := api.Group("/auth")
	auth.Get("/login", handler.Login)
}

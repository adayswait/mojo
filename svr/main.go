package main

import "fmt"
import "github.com/gofiber/fiber"
import "github.com/gofiber/cors"
import "github.com/adayswait/mojo/utils"
import "github.com/adayswait/mojo/db"
import "github.com/adayswait/mojo/global"
import "github.com/adayswait/mojo/router"
import "github.com/adayswait/mojo/ws"

func main() {
	app := fiber.New()
	app.Use(cors.New())
	router.Route(app)
	app.Static("/", "../view/")

	// Optional middleware
	app.Use("/ws", ws.New())

	// Upgraded websocket request
	
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})

	// db.Set(global.BUCKET_USER,"jesse","value")
	value, _ := db.Get(global.BUCKET_USER, "jesse")
	fmt.Println("ends")
	fmt.Println(string(value))
	app.Listen(int(utils.GetListeningPort()))
}

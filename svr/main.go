package main

import "log"
import "github.com/gofiber/fiber"
import "github.com/gofiber/cors"
import "github.com/gofiber/session"
import "github.com/gofiber/websocket"
import "github.com/boltdb/bolt"

func main() {
	app := fiber.New()
	sessions := session.New()
	app.Use("/ws", func(c *fiber.Ctx) {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			c.Next()
		}
	})

	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) {
		store := sessions.Get(c) // get/create new session
		defer store.Save()
		c.Send("mojo")
	})

	db, err := bolt.Open("mojo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app.Listen(3000)
}

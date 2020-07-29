package ws

import "github.com/gofiber/fiber"

func New() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		if c.Get("host") == "localhost:3000" {
			c.Locals("Host", "Localhost:3000")
			c.Next()
			return
		}
		c.Status(403).Send("Request origin not allowed")
	}
}

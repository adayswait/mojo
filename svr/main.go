package main

import "github.com/gofiber/fiber"
import "log"
import "github.com/boltdb/bolt"

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
  })
  db, err := bolt.Open("my.db", 0600, nil)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  app.Listen(3000)
}

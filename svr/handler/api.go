package handler

import "time"
import "fmt"
import "github.com/gofiber/fiber"
import "github.com/google/uuid"
import "github.com/gofiber/session"

func Hello(c *fiber.Ctx) {
	fmt.Println(c.Cookies("jesse"))
	c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
	cookie := new(fiber.Cookie)
	cookie.Name = "jesse"
	cookie.Value = uuid.New().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)
}

func Login(c *fiber.Ctx) {
	c.JSON(fiber.Map{"status": "success", "message": "Login ok!", "data": nil})
}
func SessionTest(c *fiber.Ctx){
	sessions := session.New()
	fmt.Println(sessions)
}

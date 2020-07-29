package handler

import "fmt"
import "github.com/gofiber/websocket"

func Websocket(c *websocket.Conn) {
	fmt.Println(c.Locals("Host"))
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s", msg)
		err = c.WriteMessage(mt, msg)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}

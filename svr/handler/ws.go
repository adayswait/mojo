package handler

import (
	"fmt"
	"github.com/adayswait/mojo/mlog"
	"github.com/gofiber/websocket/v2"
)

func Websocket(c *websocket.Conn) {
	mlog.Log(c.Locals("Host"))
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			mlog.Log("read:", err)
			break
		}
		fmt.Printf("recv: %s", msg)
		err = c.WriteMessage(mt, msg)
		if err != nil {
			mlog.Log("write:", err)
			break
		}
	}
}

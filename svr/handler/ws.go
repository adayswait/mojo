package handler

import (
	"github.com/adayswait/mojo/mlog"
	"github.com/gofiber/websocket/v2"
)

func Websocket(c *websocket.Conn) {
	for {
		mt, msg, errr := c.ReadMessage()
		if errr != nil {
			mlog.Errorf("websocker ReadMessage err:%s", errr)
			break
		}
		mlog.Infof("websocker recv:%s", msg)
		errw := c.WriteMessage(mt, msg)
		if errw != nil {
			mlog.Errorf("websocker WriteMessage err:%s", errw)
			break
		}
	}
}

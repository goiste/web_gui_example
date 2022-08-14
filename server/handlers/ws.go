package handlers

import (
	"encoding/json"
	"time"

	"github.com/goiste/web_gui_example/src/sysinfo"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func WSHandler(out chan<- struct{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		for {
			sendUpdates(ws)

			_, _, err = ws.ReadMessage()
			if err == nil {
				out <- struct{}{}
				continue
			}
			return nil
		}
	}
}

func sendUpdates(ws *websocket.Conn) {
	upd := sysinfo.GetUpdate(time.Second)
	b, _ := json.Marshal(upd)
	_ = ws.WriteMessage(websocket.TextMessage, b)
}

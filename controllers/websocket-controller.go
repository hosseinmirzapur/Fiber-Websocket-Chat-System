package controllers

import (
	"github.com/gofiber/websocket/v2"
	"github.com/hosseinmirzapur/second_test_project/handlers"
	"log"
)

func WebsocketHandler(c *websocket.Conn) {
	key := c.Params("key")
	username := handlers.FindUserName(key, handlers.Users)
	if username == "" {
		return
	}

	defer func() {
		handlers.Unregister <- c
		socketCloseError := c.Close()
		if socketCloseError != nil {
			return
		}
	}()
	handlers.Register <- c

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}
			return
		}
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
		if mt == websocket.TextMessage {
			// Broadcast the received message
			handlers.Broadcast <- username + ": " + string(msg)
		}
	}
}

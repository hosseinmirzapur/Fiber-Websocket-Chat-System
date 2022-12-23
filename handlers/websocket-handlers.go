package handlers

import (
	"github.com/gofiber/websocket/v2"
	"github.com/hosseinmirzapur/second_test_project/models"
	"log"
)

var Clients = make(map[*websocket.Conn]*models.Client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var Register = make(chan *websocket.Conn)
var Broadcast = make(chan string)
var Unregister = make(chan *websocket.Conn)

func RunHub() {
	for {
		select {
		case connection := <-Register:
			Clients[connection] = &models.Client{}

		case message := <-Broadcast:
			// Send the message to all Clients
			for connection, c := range Clients {
				go func(connection *websocket.Conn, c *models.Client) { // send to each Client in parallel, so we don't block on a slow Client
					c.Mu.Lock()
					defer c.Mu.Unlock()
					if c.IsClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.IsClosing = true
						log.Println("write error:", err)

						firstErr := connection.WriteMessage(websocket.CloseMessage, []byte{})
						if firstErr != nil {
							log.Fatal("Message could not be sent over channel")
						}
						secondErr := connection.Close()
						if secondErr != nil {
							log.Fatal("Websocket connection could not be closed")
						}
						Unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-Unregister:
			// Remove the Client from the hub
			delete(Clients, connection)

			log.Println("connection unregistered")
		}
	}
}

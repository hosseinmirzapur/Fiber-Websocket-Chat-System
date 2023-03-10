package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/hosseinmirzapur/second_test_project/controllers"
	"github.com/hosseinmirzapur/second_test_project/middlewares"
)

func SetupWsRoutes(app *fiber.App) {
	ws := app.Group("/auth")
	// Middleware on Websocket Routes
	ws = ws.Use(middlewares.NeedsUpgrade)

	// Websocket Routes
	ws.Get("/key/:key/ws", websocket.New(controllers.WebsocketHandler))
}

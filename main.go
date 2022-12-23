package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/second_test_project/config"
	"github.com/hosseinmirzapur/second_test_project/handlers"
	"github.com/hosseinmirzapur/second_test_project/router"
	"log"
	"os"
)

func init() {
	// Initialize environment variables
	config.LoadEnv()
}

func main() {
	// initialize fiber
	app := fiber.New(config.FiberConfig())

	// Goroutine for managing clients
	go handlers.RunHub()

	// Api Routes
	router.SetupApiRoutes(app)

	// Websocket Routes
	router.SetupWsRoutes(app)

	err := app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		log.Println(err.Error())
	}
}

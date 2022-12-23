package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/second_test_project/controllers"
)

func SetupApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Registry
	api.Post("/users", controllers.RegisterUser)

	// Users List
	api.Get("/users", controllers.UsersList)
}

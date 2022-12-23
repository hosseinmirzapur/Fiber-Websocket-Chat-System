package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/hosseinmirzapur/second_test_project/handlers"
	"github.com/hosseinmirzapur/second_test_project/models"
)

func RegisterUser(ctx *fiber.Ctx) error {
	user := models.User{
		ID: utils.UUID(),
	}
	validationError := ctx.BodyParser(&user)
	if validationError != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": fiber.ErrUnprocessableEntity.Error(),
		})
	}
	handlers.Users = append(handlers.Users, user)
	return ctx.JSON(fiber.Map{
		"message": "user created successfully",
		"key":     user.ID,
	})
}

func UsersList(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"users": handlers.Users,
	})
}

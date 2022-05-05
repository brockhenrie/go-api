package controllers

import (
	"api/database"
	"api/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission
	database.DB.Find(&permissions)
	return c.JSON(permissions)
}

func CreatePermission(c *fiber.Ctx)error{
	var permission models.Permission

	if err := c.BodyParser(&permission); err != nil{
		return err
	}

	database.DB.Create(&permission)

	return c.JSON(permission)
}

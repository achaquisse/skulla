package handler

import (
	"github.com/achaquisse/skulla-api/helper"
	"github.com/gofiber/fiber/v2"
)

func ReturnUsers(c *fiber.Ctx) error {
	users := helper.GenerateUsers(5)
	return c.JSON(users)
}

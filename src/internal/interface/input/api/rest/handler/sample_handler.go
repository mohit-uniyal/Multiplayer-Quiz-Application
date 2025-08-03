package handler

import "github.com/gofiber/fiber/v2"

func CheckAPIHandler(c *fiber.Ctx) error {
	return c.SendString("API is running")
}

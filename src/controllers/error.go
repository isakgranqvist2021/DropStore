package controllers

import "github.com/gofiber/fiber/v2"

func Error(c *fiber.Ctx) error {
	return c.Render("pages/error", nil)
}

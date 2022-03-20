package controllers

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx) error {
	return c.Render("pages/success", nil)
}

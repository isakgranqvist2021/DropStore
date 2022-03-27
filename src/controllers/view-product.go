package controllers

import "github.com/gofiber/fiber/v2"

func ViewProduct(c *fiber.Ctx) error {

	return c.Render("pages/view-product", nil)
}

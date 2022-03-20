package controllers

import "github.com/gofiber/fiber/v2"

func Cancel(c *fiber.Ctx) error {
	return c.SendString("Err!")
}

package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func AddToCart(c *fiber.Ctx) error {
	var output models.Order

	if err := c.BodyParser(&output); err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}

	fmt.Println(output)

	return c.Redirect("/")
}

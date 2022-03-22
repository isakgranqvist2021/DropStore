package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func Index(c *fiber.Ctx) error {
	var products []models.Product

	for i := 0; i < 25; i++ {
		products = append(products, models.Product{
			ID:          int64(i + 1),
			Name:        "Great product",
			Description: "Great description",
			Amount:      10,
			Features:    []string{},
		})
	}

	return c.Render("pages/index", fiber.Map{
		"Products": products,
	})
}

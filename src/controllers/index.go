package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func Index(c *fiber.Ctx) error {
	var products []models.Product

	bytes, err := ioutil.ReadFile("./data/products.json")

	if err != nil {
		return c.Redirect("/error")
	}

	if err := json.Unmarshal(bytes, &products); err != nil {
		return c.Redirect("/error")
	}

	return c.Render("pages/index", fiber.Map{
		"Products": products,
	})
}

package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/utils"
)

func ViewProducts(c *fiber.Ctx) error {
	products, err := utils.GetProducts()

	if err != nil {
		return c.Redirect("/error")
	}

	return c.Render("pages/index", fiber.Map{
		"Products": products,
	})
}

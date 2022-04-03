package product

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/services/logger"
)

func ViewProduct(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("ID"), 10, 0)

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	product, err := GetProduct(int(ID))

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Render("pages/product/product", fiber.Map{
		"Product": product,
	})
}

func ViewProducts(c *fiber.Ctx) error {
	products, err := GetProducts()

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Render("pages/product/index", fiber.Map{
		"Products": products,
	})
}

package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/services/database"
	"github.com/isakgranqvist2021/dropstore/src/services/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ViewProduct(c *fiber.Ctx) error {
	var product Product

	objectId, err := primitive.ObjectIDFromHex(c.Params("ID"))

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	readOptions := database.ReadOptions{
		Collection: "products",
		Filter:     bson.M{"_id": objectId},
	}

	if err := readOptions.ReadOne(&product); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Render("pages/product/product", fiber.Map{
		"Product": product,
	})
}

func ViewProducts(c *fiber.Ctx) error {
	getProductOptions := database.ReadOptions{Collection: "products"}

	var products []Product

	if err := getProductOptions.ReadMany(&products); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Render("pages/product/index", fiber.Map{
		"Products": products,
	})
}

func AddProduct(c *fiber.Ctx) error {
	var body []Product

	if err := c.BodyParser(&body); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	for _, v := range body {
		if err := v.InsertProduct(); err != nil {
			go logger.Log(err)

			return c.Redirect("/error")
		}

	}

	return c.JSON(body)
}

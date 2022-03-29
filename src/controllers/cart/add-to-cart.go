package cart

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/helpers/cart"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func AddToCart(c *fiber.Ctx) error {
	var newItem models.CartItem

	if err := c.BodyParser(&newItem); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"message": "Invalid request body",
			"success": false,
			"data":    nil,
		})
	}

	productId, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Invalid query parameter",
			"success": false,
			"data":    nil,
		})
	}

	newItem.ID = int(productId)

	if err := cart.AddToCartAndUpdateSession(c, newItem); err != nil {
		return c.JSON(fiber.Map{
			"message": "Item has not been added to cart",
			"success": false,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item has been added to cart",
		"success": true,
		"data":    nil,
	})
}

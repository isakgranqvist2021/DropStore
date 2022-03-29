package cart

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/helpers/cart"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func AddToCart(c *fiber.Ctx) error {
	var newItem models.CartItem

	if err := c.BodyParser(&newItem); err != nil {
		return c.Redirect("/error")
	}

	productId, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.Redirect("/error")
	}

	newItem.ID = int(productId)

	if err := cart.AddToCartAndUpdateSession(c, newItem); err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect(string(c.Context().Referer()))
}

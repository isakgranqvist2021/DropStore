package cart

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/helpers/cart"
)

func ChangeQuantity(c *fiber.Ctx) error {
	action := c.Params("ACTION")

	productId, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.Redirect("/error")
	}

	err = cart.ChangeQtyAndUpdateSession(c, int(productId), action)

	if err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect("/checkout")
}

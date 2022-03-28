package cart

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func ChangeQuantity(c *fiber.Ctx) error {
	action := c.Params("ACTION")

	productId, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	cartInventory := sess.Get("CART_INVENTORY").([]models.CartItem)

	for i, v := range cartInventory {
		if productId == int64(v.ID) {
			product := &cartInventory[i]

			switch action {
			case "increment":
				product.Quantity += 1
			case "decrement":
				product.Quantity -= 1
			default:
				return c.Redirect("/error")
			}

			if product.Quantity == 0 {
				cartInventory = append(cartInventory[:i], cartInventory[i+1:]...)
			}
		}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect("/checkout")
}

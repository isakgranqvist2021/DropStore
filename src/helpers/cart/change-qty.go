package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func ChangeQtyAndUpdateSession(c *fiber.Ctx, productId int, action string) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		return err
	}

	cartInventory := sess.Get("CART_INVENTORY").([]models.CartItem)

	for i, v := range cartInventory {
		if productId == v.ID {
			product := &cartInventory[i]

			switch action {
			case "increment":
				product.Quantity += 1
			case "decrement":
				product.Quantity -= 1
			default:
				return err
			}

			if product.Quantity == 0 {
				cartInventory = append(cartInventory[:i], cartInventory[i+1:]...)
			}
		}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return err
	}

	return nil
}

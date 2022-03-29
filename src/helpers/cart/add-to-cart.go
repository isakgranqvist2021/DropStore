package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func AddToCartAndUpdateSession(c *fiber.Ctx, newItem models.CartItem) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		return err
	}

	cartInventory := sess.Get("CART_INVENTORY")

	if cartInventory != nil {
		cartInventory = append(cartInventory.([]models.CartItem), newItem)
	} else {
		cartInventory = []models.CartItem{newItem}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return err
	}

	return nil
}

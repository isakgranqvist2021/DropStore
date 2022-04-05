package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/packages/alert"
	"github.com/isakgranqvist2021/dropstore/src/services/logger"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
)

func AddToCart(c *fiber.Ctx) error {
	var newItem CartItem

	if err := c.BodyParser(&newItem); err != nil {
		go logger.Log(err)

		return c.JSON(fiber.Map{
			"message": "Invalid request body",
			"success": false,
			"data":    nil,
		})
	}

	newItem.ID = c.Params("ID")

	if err := AddToCartAndUpdateSession(c, newItem); err != nil {
		go logger.Log(err)

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

func ChangeQuantity(c *fiber.Ctx) error {
	action := c.Params("ACTION")

	if err := ChangeQtyAndUpdateSession(c, c.Params("ID"), action); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Redirect("/cart")
}

func Cart(c *fiber.Ctx) error {
	sess, err := store.GetStore().Get(c)

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	cartInventory := []CartItem{}

	cartInventoryRaw := sess.Get("CART_INVENTORY")

	if cartInventoryRaw != nil {
		cartInventory = cartInventoryRaw.([]CartItem)
	}

	products := JoinCart(cartInventory)

	return c.Render("pages/cart/index", fiber.Map{
		"Products": products,
		"Alert":    alert.GetAlert(c),
	})
}

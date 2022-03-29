package checkout

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
	"github.com/isakgranqvist2021/dropstore/src/utils"
)

func Checkout(c *fiber.Ctx) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	cartInventory := []models.CartItem{}

	cartInventoryRaw := sess.Get("CART_INVENTORY")

	if cartInventoryRaw != nil {
		cartInventory = cartInventoryRaw.([]models.CartItem)
	}

	products := utils.JoinCart(cartInventory)

	return c.Render("pages/checkout", fiber.Map{
		"Products": products,
	})
}

package cart

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
)

func AddToCart(c *fiber.Ctx) error {
	var newItem CartItem

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

	if err := AddToCartAndUpdateSession(c, newItem); err != nil {
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

	productId, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.Redirect("/error")
	}

	err = ChangeQtyAndUpdateSession(c, int(productId), action)

	if err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect("/checkout")
}

func Checkout(c *fiber.Ctx) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	cartInventory := []CartItem{}

	cartInventoryRaw := sess.Get("CART_INVENTORY")

	if cartInventoryRaw != nil {
		cartInventory = cartInventoryRaw.([]CartItem)
	}

	products := JoinCart(cartInventory)

	return c.Render("pages/checkout", fiber.Map{
		"Products": products,
	})
}

package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

func AddToCart(c *fiber.Ctx) error {
	var body models.Order

	if err := c.BodyParser(&body); err != nil {
		return c.Redirect("/error")
	}

	productId, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.Redirect("/error")
	}

	newItem := map[string]int{
		"ID":       int(productId),
		"Quantity": body.Quantity,
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	cartInventory := sess.Get("CART_INVENTORY")

	if err != nil {
		return c.Redirect("/error")
	}

	if cartInventory != nil {
		cartInventory = append(cartInventory.([]map[string]int), newItem)
	} else {
		cartInventory = []map[string]int{newItem}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect("/")
}

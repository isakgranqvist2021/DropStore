package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
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

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	cartInventory := sess.Get("CART_INVENTORY")

	if err != nil {
		return c.Redirect("/error")
	}

	if cartInventory != nil {
		cartInventory = append(cartInventory.([]models.CartItem), newItem)
	} else {
		cartInventory = []models.CartItem{newItem}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect(string(c.Context().Referer()))
}

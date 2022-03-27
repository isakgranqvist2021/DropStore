package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
)

func CompletedPurchase(c *fiber.Ctx) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		c.Redirect("/error")
	}

	fmt.Println(sess.Get("STRIPE_SESSION"))

	return c.Render("pages/completed-purchase", nil)
}

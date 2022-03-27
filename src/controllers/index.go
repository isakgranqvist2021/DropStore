package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/utils"
)

func Index(c *fiber.Ctx) error {
	products, err := utils.GetProducts()

	if err != nil {
		return c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	sess.Set("GO_BACK_HREF", "/")

	if err := sess.Save(); err != nil {
		return c.Redirect("/error")
	}

	return c.Render("pages/index", fiber.Map{
		"Products": products,
	})
}

package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/utils"
)

func ViewProduct(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.Redirect("/error")
	}

	product, err := utils.GetProduct(int64(ID))

	if err != nil {
		return c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	return c.Render("pages/view-product", fiber.Map{
		"Product":    product,
		"GoBackHref": sess.Get("GO_BACK_HREF"),
	})
}

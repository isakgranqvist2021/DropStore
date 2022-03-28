package products

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/utils"
)

func ViewProduct(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("ID"), 10, 0)

	if err != nil {
		return c.Redirect("/error")
	}

	product, err := utils.GetProduct(int(ID))

	if err != nil {
		return c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	return c.Render("pages/view-product", fiber.Map{
		"Product":    product,
		"GoBackHref": sess.Get("GO_BACK_HREF"),
	})
}

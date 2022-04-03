package packages

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/packages/cart"
	"github.com/isakgranqvist2021/dropstore/src/packages/err"
	"github.com/isakgranqvist2021/dropstore/src/packages/order"
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
)

func MainRouter(r fiber.Router) {
	r.Use(func(c *fiber.Ctx) error {
		fmt.Println(fmt.Sprintf("%s %s - %v",
			c.Method(),
			c.OriginalURL(),
			time.Now().Format(time.RFC1123),
		))

		return c.Next()
	})

	err.ErrRouter(r.Group("/error"))
	cart.CartRouter(r.Group("/cart"))
	order.OrderRouter(r.Group("/order"))
	product.ProductRouter(r.Group("/"))
}

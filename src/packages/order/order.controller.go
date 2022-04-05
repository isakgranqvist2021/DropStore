package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/packages/alert"
	"github.com/isakgranqvist2021/dropstore/src/packages/cart"
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
	"github.com/isakgranqvist2021/dropstore/src/services/logger"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func Cancel(c *fiber.Ctx) error {
	orderId, err := GetOrderIDAndCast(c)

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	order := Order{ID: *orderId, Status: "completed"}

	if err := order.UpdateStatus(); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Render("pages/cart/cancel", nil)
}

func Success(c *fiber.Ctx) error {
	orderId, err := GetOrderIDAndCast(c)

	order := Order{ID: *orderId, Status: "completed"}

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	if err := order.UpdateStatus(); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Render("pages/cart/success", nil)
}

func Pay(c *fiber.Ctx) error {
	var body Order

	if err := c.BodyParser(&body); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	sess, err := store.GetStore().Get(c)

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	if err := body.Validate(); err != nil {
		alert := alert.Alert{Message: err.Error(), Severity: "error"}
		alert.SetAlert(c)

		return c.Redirect("/cart")
	}

	body.Products = cart.JoinCart(sess.Get("CART_INVENTORY").([]cart.CartItem))

	params := &stripe.CheckoutSessionParams{
		LineItems:  product.ConvertProductsToStripeLineItems(&body.Products),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(config.GetConfig().GetDomain() + "/order/success"),
		CancelURL:  stripe.String(config.GetConfig().GetDomain() + "/order/cancel"),
	}

	insertedID, err := body.CreateOrder()

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	s, err := session.New(params)

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	sess.Set("STRIPE_SESSION", s.ID)
	sess.Set("ORDER_ID", insertedID.Hex())

	if err := sess.Save(); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Redirect(s.URL)
}

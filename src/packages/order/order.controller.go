package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/packages/cart"
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
	"github.com/isakgranqvist2021/dropstore/src/services/logger"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func GetOrderIDAndCast(c *fiber.Ctx) (*string, error) {
	sess, err := store.GetStore().Get(c)

	if err != nil {
		return nil, err
	}

	orderId, ok := sess.Get("ORDER_ID").(string)

	if !ok {
		return nil, err
	}

	return &orderId, nil
}

func Cancel(c *fiber.Ctx) error {
	orderId, err := GetOrderIDAndCast(c)

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	if err := UpdateOrderStatus(orderId, "canceled"); err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	return c.Render("pages/cart/cancel", nil)
}

func Success(c *fiber.Ctx) error {
	orderId, err := GetOrderIDAndCast(c)

	if err != nil {
		go logger.Log(err)

		return c.Redirect("/error")
	}

	if err := UpdateOrderStatus(orderId, "completed"); err != nil {
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

	body.Products = cart.JoinCart(sess.Get("CART_INVENTORY").([]cart.CartItem))

	params := &stripe.CheckoutSessionParams{
		LineItems:  product.ConvertProductsToStripeLineItems(&body.Products),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(config.GetConfig().GetDomain() + "/order/success"),
		CancelURL:  stripe.String(config.GetConfig().GetDomain() + "/order/cancel"),
	}

	insertedID, err := CreateOrder(body)

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

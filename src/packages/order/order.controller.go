package order

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/packages/cart"
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func Cancel(c *fiber.Ctx) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	fmt.Println(sess.Get("STRIPE_SESSION"))

	return c.Redirect("/cart")
}

func Success(c *fiber.Ctx) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	fmt.Println(sess.Get("STRIPE_SESSION"))

	return c.Render("pages/completed-purchase", nil)
}

func Pay(c *fiber.Ctx) error {
	var body Order

	if err := c.BodyParser(&body); err != nil {
		log.Fatal(err)
		return c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	body.Products = cart.JoinCart(sess.Get("CART_INVENTORY").([]cart.CartItem))

	params := &stripe.CheckoutSessionParams{
		LineItems:  product.ConvertProductsToStripeLineItems(&body.Products),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(config.GetConfig().GetDomain() + "/order/success"),
		CancelURL:  stripe.String(config.GetConfig().GetDomain() + "/order/cancel"),
	}

	s, err := session.New(params)

	if err != nil {
		return c.Redirect("/error")
	}

	sess.Set("STRIPE_SESSION", s.ID)

	if err := sess.Save(); err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect(s.URL)
}

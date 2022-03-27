package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func Pay(c *fiber.Ctx) error {
	domain := config.GetConfig().GetDomain()

	paymentMode := string(stripe.CheckoutSessionModePayment)

	cancelUrl := domain + "/cancel"
	successUrl := domain + "/success"

	params := &stripe.CheckoutSessionParams{
		LineItems: models.ConvertProductsToStripeLineItems(&[]models.Product{
			{Amount: 200, Name: "Sporting Pants"},
			{Amount: 250, Name: "Sporting Shirt"},
		}),
		Mode:       &paymentMode,
		SuccessURL: &successUrl,
		CancelURL:  &cancelUrl,
	}

	s, err := session.New(params)

	if err != nil {
		c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		c.Redirect("/error")
	}

	sess.Set("STRIPE_SESSION", s.ID)

	err = sess.Save()

	if err != nil {
		c.Redirect("/error")
	}

	return c.Redirect(s.URL)
}

func Checkout(c *fiber.Ctx) error {

	return c.Render("pages/checkout", nil)
}

package purchase

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func Pay(c *fiber.Ctx) error {
	var body models.Order

	if err := c.BodyParser(&body); err != nil {
		log.Fatal(err)
		return c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	// TODO - inject cart items in LineItems
	params := &stripe.CheckoutSessionParams{
		LineItems: models.ConvertProductsToStripeLineItems(&[]models.Product{
			{Amount: 200, Name: "Sporting Pants"},
			{Amount: 250, Name: "Sporting Shirt"},
		}),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(config.GetConfig().GetDomain() + "/success"),
		CancelURL:  stripe.String(config.GetConfig().GetDomain() + "/cancel"),
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

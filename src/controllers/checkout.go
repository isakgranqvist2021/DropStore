package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/models"
	"github.com/isakgranqvist2021/dropstore/src/utils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func ChangeQuantity(c *fiber.Ctx) error {
	action := c.Params("ACTION")

	productId, err := strconv.ParseInt(c.Params("ID"), 10, 64)

	if err != nil {
		return c.Redirect("/error")
	}

	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	cartInventory := sess.Get("CART_INVENTORY").([]models.CartItem)

	for i, v := range cartInventory {
		if productId == int64(v.ID) {
			product := &cartInventory[i]

			switch action {
			case "increment":
				product.Quantity += 1
			case "decrement":
				product.Quantity -= 1
			default:
				return c.Redirect("/error")
			}

			if product.Quantity == 0 {
				cartInventory = append(cartInventory[:i], cartInventory[i+1:]...)
			}
		}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return c.Redirect("/error")
	}

	return c.Redirect("/checkout")
}

func Checkout(c *fiber.Ctx) error {
	sess, err := config.GetStore().Get(c)

	if err != nil {
		return c.Redirect("/error")
	}

	cartInventory := []models.CartItem{}

	cartInventoryRaw := sess.Get("CART_INVENTORY")

	if cartInventoryRaw != nil {
		cartInventory = cartInventoryRaw.([]models.CartItem)
	}

	products := utils.JoinCart(cartInventory)

	sess.Set("GO_BACK_HREF", "/checkout")

	if err := sess.Save(); err != nil {
		return c.Redirect("/error")
	}

	return c.Render("pages/checkout", fiber.Map{
		"Products": products,
	})
}

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

	if err := sess.Save(); err != nil {
		c.Redirect("/error")
	}

	return c.Redirect(s.URL)
}

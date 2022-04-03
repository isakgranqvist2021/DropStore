package cart

import "github.com/gofiber/fiber/v2"

func CartRouter(r fiber.Router) {
	r.Get("/", Checkout)
	r.Post("/add-to-cart/:ID", AddToCart)
	r.Get("/change-quantity/:ACTION/:ID", ChangeQuantity)
}

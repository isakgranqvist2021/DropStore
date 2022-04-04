package product

import "github.com/gofiber/fiber/v2"

func ProductRouter(r fiber.Router) {
	r.Get("/", ViewProducts)
	r.Post("/add", AddProduct)
	r.Get("/:PRODUCT_NAME/:ID", ViewProduct)
}

package order

import "github.com/gofiber/fiber/v2"

func OrderRouter(r fiber.Router) {
	r.Post("/", Pay)
	r.Get("/cancel", Cancel)
	r.Get("/success", Success)
}

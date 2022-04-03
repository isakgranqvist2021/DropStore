package err

import "github.com/gofiber/fiber/v2"

func ErrRouter(r fiber.Router) {
	r.Get("/", Err)
}

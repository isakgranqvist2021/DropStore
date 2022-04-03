package err

import "github.com/gofiber/fiber/v2"

func Err(c *fiber.Ctx) error {
	return c.Render("pages/error", nil)
}

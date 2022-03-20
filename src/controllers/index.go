package controllers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "text/html")

	return c.SendString(`
			<div>
				<h1>Hello, World</h1>
				<form method='POST' action='/buy'>
					<button>Buy</button>
				</form>
			</div>`,
	)
}

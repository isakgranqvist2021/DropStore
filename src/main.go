package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/controllers"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	godotenv.Load(".env")

	config := config.GetConfig()

	stripe.Key = config.StripeKey

	engine := html.New("./src/views", ".html").Reload(true)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", controllers.Index)
	app.Get("/cancel", controllers.Cancel)
	app.Get("/success", controllers.Success)
	app.Post("/checkout", controllers.Checkout)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}

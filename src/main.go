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

	serverConfig := config.GetConfig()

	stripe.Key = serverConfig.StripeKey

	engine := html.New("./src/views", ".html").Reload(true)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	config.NewStore()

	app.Get("/", controllers.Index)
	app.Get("/cancel", controllers.Cancel)
	app.Get("/success", controllers.Success)
	app.Post("/checkout", controllers.Checkout)
	app.Get("/error", controllers.Error)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", serverConfig.Port)))
}

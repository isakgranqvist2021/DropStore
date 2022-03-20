package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/controllers"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	godotenv.Load(".env")

	config := config.GetConfig()

	stripe.Key = config.StripeKey

	app := fiber.New()

	app.Get("/", controllers.Index)
	app.Get("/cancel", controllers.Cancel)
	app.Get("/success", controllers.Success)
	app.Post("/checkout", controllers.Checkout)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}

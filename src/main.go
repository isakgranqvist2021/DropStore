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

const BASEDIR = "./src"

func main() {
	godotenv.Load(".env")

	serverConfig := config.GetConfig()

	stripe.Key = serverConfig.StripeKey

	engine := html.New(BASEDIR+"/views", ".html").Reload(true)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/public", BASEDIR+"/public")

	config.NewStore()

	app.Get("/", controllers.Index)
	app.Get("/error", controllers.Error)
	app.Get("/cancel", controllers.Cancel)
	app.Get("/success", controllers.Success)
	app.Post("/checkout", controllers.Checkout)
	app.Post("/add-to-cart", controllers.AddToCart)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", serverConfig.Port)))
}

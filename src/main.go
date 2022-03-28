package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/controllers"
	"github.com/isakgranqvist2021/dropstore/src/controllers/cart"
	"github.com/isakgranqvist2021/dropstore/src/controllers/checkout"
	"github.com/isakgranqvist2021/dropstore/src/controllers/products"
	"github.com/isakgranqvist2021/dropstore/src/controllers/purchase"
	"github.com/isakgranqvist2021/dropstore/src/utils"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

const BASEDIR = "./src"

func main() {
	godotenv.Load(".env")

	serverConfig := config.GetConfig()

	stripe.Key = serverConfig.StripeKey

	engine := html.New(BASEDIR+"/views", ".html").Reload(true).AddFunc("CutStr", utils.CutStr)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/public", BASEDIR+"/public")

	config.NewStore()

	app.Get("/error", controllers.Error)

	app.Get("/", products.ViewProducts)
	app.Get("/product/:ID", products.ViewProduct)

	app.Post("/add-to-cart/:ID", cart.AddToCart)
	app.Get("/change-quantity/:ACTION/:ID", cart.ChangeQuantity)

	app.Post("/pay", purchase.Pay)
	app.Get("/cancel", purchase.CancelPurchase)
	app.Get("/success", purchase.CompletedPurchase)

	app.Get("/checkout", checkout.Checkout)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", serverConfig.Port)))
}

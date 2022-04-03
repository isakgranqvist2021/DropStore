package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/packages/alert"
	"github.com/isakgranqvist2021/dropstore/src/packages/cart"
	"github.com/isakgranqvist2021/dropstore/src/packages/error"
	"github.com/isakgranqvist2021/dropstore/src/packages/order"
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
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

	store := config.NewStore()
	store.RegisterType([]cart.CartItem{})
	store.RegisterType(alert.Alert{})

	app.Get("/error", error.Error)

	app.Get("/", product.ViewProducts)
	app.Get("/product/:ID", product.ViewProduct)

	app.Post("/add-to-cart/:ID", cart.AddToCart)
	app.Get("/change-quantity/:ACTION/:ID", cart.ChangeQuantity)

	app.Post("/pay", order.Pay)
	app.Get("/cancel", order.Cancel)
	app.Get("/success", order.Success)

	app.Get("/checkout", cart.Checkout)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", serverConfig.Port)))
}

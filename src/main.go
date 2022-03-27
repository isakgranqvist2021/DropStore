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

func CutStr(value string) string {

	subStr := value[:100]

	return subStr

}

func main() {
	godotenv.Load(".env")

	serverConfig := config.GetConfig()

	stripe.Key = serverConfig.StripeKey

	engine := html.New(BASEDIR+"/views", ".html").Reload(true).AddFunc("CutStr", CutStr)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/public", BASEDIR+"/public")

	config.NewStore()

	app.Get("/", controllers.Index)
	app.Post("/pay", controllers.Pay)
	app.Get("/error", controllers.Error)
	app.Get("/checkout", controllers.Checkout)
	app.Get("/cancel", controllers.CancelPurchase)
	app.Get("/product/:ID", controllers.ViewProduct)
	app.Get("/success", controllers.CompletedPurchase)
	app.Post("/add-to-cart/:ID", controllers.AddToCart)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", serverConfig.Port)))
}

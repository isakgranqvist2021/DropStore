package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/packages"
	"github.com/isakgranqvist2021/dropstore/src/packages/alert"
	"github.com/isakgranqvist2021/dropstore/src/packages/cart"
	"github.com/isakgranqvist2021/dropstore/src/services/database"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
	"github.com/isakgranqvist2021/dropstore/src/utils"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

func Run() error {
	godotenv.Load(".env")

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	defer database.Disconnect()

	serverConfig := config.GetConfig()

	stripe.Key = serverConfig.StripeKey

	engine := html.New(config.BASEDIR+"/views", ".html").Reload(true).AddFunc("CutStr", utils.CutStr)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/public", config.BASEDIR+"/public")

	store := store.NewStore()
	store.RegisterType([]cart.CartItem{})
	store.RegisterType(alert.Alert{})

	packages.MainRouter(app.Group("/"))

	return app.Listen(fmt.Sprintf(":%d", serverConfig.Port))
}

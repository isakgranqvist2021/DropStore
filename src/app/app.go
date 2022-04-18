package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/packages"
	"github.com/isakgranqvist2021/dropstore/src/services/database"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

func Run() error {
	godotenv.Load(".env")

	if err := database.Connect(); err != nil {
		return err
	}

	defer database.Disconnect()

	serverConfig := config.GetConfig()

	stripe.Key = serverConfig.StripeKey

	engine := Setup()

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/public", config.BASEDIR+"/public")

	packages.MainRouter(app.Group("/"))

	return app.Listen(fmt.Sprintf(":%d", serverConfig.Port))
}

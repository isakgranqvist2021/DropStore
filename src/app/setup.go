package app

import (
	"github.com/gofiber/template/html"
	"github.com/isakgranqvist2021/dropstore/src/config"
	"github.com/isakgranqvist2021/dropstore/src/packages/alert"
	"github.com/isakgranqvist2021/dropstore/src/packages/cart"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
	"github.com/isakgranqvist2021/dropstore/src/utils/stringm"
)

func Setup() *html.Engine {
	engine := html.New(config.BASEDIR+"/views", ".html").Reload(true)

	engine = engine.AddFunc("CutStr", stringm.CutStr)
	engine = engine.AddFunc("Replace", stringm.ReplaceWhiteSpaceWithUnderscore)

	store := store.NewStore()
	store.RegisterType([]cart.CartItem{})
	store.RegisterType(alert.Alert{})

	return engine
}

package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/isakgranqvist2021/dropstore/src/models"
)

var Store *session.Store

func NewStore() *session.Store {
	Store = session.New()

	Store.RegisterType([]models.CartItem{})
	Store.RegisterType(models.Alert{})

	return Store
}

func GetStore() *session.Store {
	return Store
}

package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func NewStore() *session.Store {

	Store = session.New()

	Store.RegisterType([]map[string]int{})

	return Store
}

func GetStore() *session.Store {
	return Store
}

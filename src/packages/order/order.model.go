package order

import (
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
)

type Order struct {
	ID         string
	Status     string
	Products   []product.Product
	FirstName  string `form:"firstName"`
	LastName   string `form:"lastName"`
	Email      string `form:"email"`
	Country    string `form:"country"`
	Phone      string `form:"phone"`
	Address    string `form:"address"`
	PostalCode string `form:"postalCode"`
}

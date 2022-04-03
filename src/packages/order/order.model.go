package order

import (
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
)

type Order struct {
	ID         int
	Products   []product.Product
	FirstName  string `form:"firstName"`
	LastName   string `form:"lastName"`
	Email      string `form:"email"`
	Phone      string `form:"phone"`
	Address    string `form:"address"`
	PostalCode string `form:"postalCode"`
}

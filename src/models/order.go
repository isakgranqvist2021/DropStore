package models

type Order struct {
	Currency string
	Product  Product
	Quantity int
}

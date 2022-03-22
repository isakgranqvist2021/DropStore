package models

type Order struct {
	Currency string
	Products []Product
	Quantity int
}

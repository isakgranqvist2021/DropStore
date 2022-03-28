package models

type Order struct {
	ID         int
	Products   []Product
	FirstName  string `form:"firstName"`
	LastName   string `form:"lastName"`
	Email      string `form:"email"`
	Phone      string `form:"phone"`
	Address    string `form:"address"`
	PostalCode string `form:"postalCode"`
}

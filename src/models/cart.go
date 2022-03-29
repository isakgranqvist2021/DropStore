package models

type CartItem struct {
	ID       int
	Quantity int `json:"quantity"`
}

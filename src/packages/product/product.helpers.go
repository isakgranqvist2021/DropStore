package product

import (
	"errors"

	"github.com/isakgranqvist2021/dropstore/src/services/database"
)

func FindProduct(productId string, products []*Product) *Product {
	for _, v := range products {
		if v.ID == productId {
			return v
		}
	}

	return nil
}

// Find product from raw without qty
func GetProduct(ID string) (*Product, error) {
	products, err := GetProducts()

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(products); i++ {
		if products[i].ID == ID {
			return products[i], nil
		}
	}

	return nil, errors.New("No product found")
}

// Find products from raw without qty
func GetProducts() ([]*Product, error) {
	var products []*Product

	readOptions := database.ReadOptions{Collection: "products"}

	if err := readOptions.ReadMany(&products); err != nil {
		return nil, err
	}

	return products, nil
}

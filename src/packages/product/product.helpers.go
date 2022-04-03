package product

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func FindProduct(productId int, products *[]Product) *Product {
	for _, v := range *products {
		if v.ID == productId {
			return &v
		}
	}

	return nil
}

// Find product from raw without qty
func GetProduct(ID int) (*Product, error) {
	products, err := GetProducts()

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(products); i++ {
		if products[i].ID == ID {
			return &products[i], nil
		}
	}

	return nil, errors.New("No product found")
}

// Find products from raw without qty
func GetProducts() ([]Product, error) {
	var products []Product

	bytes, err := ioutil.ReadFile("./data/products.json")

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &products); err != nil {
		return nil, err
	}

	return products, nil
}

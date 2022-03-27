package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/isakgranqvist2021/dropstore/src/models"
)

func GetProducts() ([]models.Product, error) {
	var products []models.Product

	bytes, err := ioutil.ReadFile("./data/products.json")

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(ID int64) (*models.Product, error) {
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

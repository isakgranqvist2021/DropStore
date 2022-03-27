package utils

import (
	"github.com/isakgranqvist2021/dropstore/src/models"
)

/*
	returnValue=[]

	[
		{
			ID: 1,
			Quantity: 2,
		},
		{
			ID: 2,
			Quantity: 1,
		},
		{
			ID: 2,
			Quantity: 4,
		}
	]

	if product not in ref {
		add product to returnValue
	} else {
		returnValue[productIndex].Quantity += product.Quantity
	}
*/

func alreadyExists(productId int, returnValue *[]models.Product) *int {
	if len(*returnValue) == 0 {
		return nil
	}

	for i, v := range *returnValue {
		if v.ID == productId {
			return &i
		}
	}

	return nil
}

func findProduct(productId int, products *[]models.Product) *models.Product {
	for _, v := range *products {
		if v.ID == productId {
			return &v
		}
	}

	return nil
}

func JoinCart(cart []models.CartItem) []models.Product {
	returnValue := []models.Product{}

	products, err := GetProducts()

	if err != nil {
		return returnValue
	}

	for _, v := range cart {
		product := findProduct(v.ID, &products)

		existingProductId := alreadyExists(v.ID, &returnValue)

		if existingProductId != nil {
			returnValue[*existingProductId].Quantity += v.Quantity
		} else {
			returnValue = append(returnValue, models.Product{
				ID:          product.ID,
				Amount:      product.Amount,
				Description: product.Description,
				Image:       product.Image,
				Features:    product.Features,
				Name:        product.Name,
				Stock:       product.Stock,
				Quantity:    v.Quantity,
			})
		}
	}

	return returnValue
}

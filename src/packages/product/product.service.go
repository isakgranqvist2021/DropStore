package product

import "github.com/isakgranqvist2021/dropstore/src/services/database"

func (product *Product) InsertProduct() error {
	payload := map[string]interface{}{
		"name":        product.Name,
		"amount":      product.Amount,
		"description": product.Description,
		"features":    product.Features,
		"stock":       product.Stock,
		"image":       product.Image,
	}

	insertOptions := database.CreateOptions{
		Payload:    payload,
		Collection: "products",
	}

	_, err := insertOptions.InsertOne()

	return err
}

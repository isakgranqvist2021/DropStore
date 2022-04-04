package product

import (
	"github.com/isakgranqvist2021/dropstore/src/packages/image"
	"github.com/isakgranqvist2021/dropstore/src/utils"
	"github.com/stripe/stripe-go/v72"
)

type Product struct {
	ID          string      `bson:"_id"`
	Amount      int         `bson:"amount"`
	Description string      `bson:"description"`
	Image       image.Image `bson:"image"`
	Features    []string    `bson:"features"`
	Name        string      `bson:"name"`
	Stock       int         `bson:"stock"`
	Quantity    int         `bson:"quantity"`
}

func (product *Product) ConvertToStripeProduct() *stripe.CheckoutSessionLineItemParams {
	description := utils.CutStr(product.Description)

	return &stripe.CheckoutSessionLineItemParams{
		Description: &description,
		Name:        &product.Name,
		Currency:    stripe.String("SEK"),
		Quantity:    stripe.Int64(int64(product.Quantity)),
		Amount:      stripe.Int64(int64(product.Amount * 100)),
		Images:      []*string{stripe.String(product.Image.Src)},
	}
}

func ConvertProductsToStripeLineItems(products *[]Product) []*stripe.CheckoutSessionLineItemParams {
	var lineItems []*stripe.CheckoutSessionLineItemParams

	for i := 0; i < len(*products); i++ {
		lineItems = append(
			lineItems,
			(*products)[i].ConvertToStripeProduct(),
		)
	}

	return lineItems
}

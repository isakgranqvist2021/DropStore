package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/packages/product"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
)

func AddToCartAndUpdateSession(c *fiber.Ctx, newItem CartItem) error {
	sess, err := store.GetStore().Get(c)

	if err != nil {
		return err
	}

	cartInventory := sess.Get("CART_INVENTORY")

	if cartInventory != nil {
		cartInventory = append(cartInventory.([]CartItem), newItem)
	} else {
		cartInventory = []CartItem{newItem}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return err
	}

	return nil
}

func ChangeQtyAndUpdateSession(c *fiber.Ctx, productId string, action string) error {
	sess, err := store.GetStore().Get(c)

	if err != nil {
		return err
	}

	cartInventory := sess.Get("CART_INVENTORY").([]CartItem)

	for i, v := range cartInventory {
		if productId == v.ID {
			product := &cartInventory[i]

			switch action {
			case "increment":
				product.Quantity += 1
			case "decrement":
				product.Quantity -= 1
			default:
				return err
			}

			if product.Quantity == 0 {
				cartInventory = append(cartInventory[:i], cartInventory[i+1:]...)
			}
		}
	}

	sess.Set("CART_INVENTORY", cartInventory)

	if err := sess.Save(); err != nil {
		return err
	}

	return nil
}

func alreadyExists(productId string, returnValue *[]product.Product) *int {
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

func JoinCart(cart []CartItem) []product.Product {
	returnValue := []product.Product{}

	products, err := product.GetProducts()

	if err != nil {
		return returnValue
	}

	for _, v := range cart {
		_product := product.FindProduct(v.ID, products)

		existingProductId := alreadyExists(v.ID, &returnValue)

		if existingProductId != nil {
			returnValue[*existingProductId].Quantity += v.Quantity
		} else {
			returnValue = append(returnValue, product.Product{
				ID:          _product.ID,
				Amount:      _product.Amount,
				Description: _product.Description,
				Images:      _product.Images,
				Features:    _product.Features,
				Name:        _product.Name,
				Stock:       _product.Stock,
				Quantity:    v.Quantity,
			})
		}
	}

	return returnValue
}

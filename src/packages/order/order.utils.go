package order

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/services/logger"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
)

func GetOrderIDAndCast(c *fiber.Ctx) (*string, error) {
	sess, err := store.GetStore().Get(c)

	if err != nil {
		return nil, err
	}

	_orderId := sess.Get("ORDER_ID")

	if _orderId != nil {
		orderId, ok := _orderId.(string)

		if !ok {
			go logger.Log(errors.New("Cast failed ORDER_ID -> string"))

			return nil, err
		}

		if len(orderId) <= 0 {
			return nil, errors.New("Empty string")
		}

		return &orderId, nil
	}

	return nil, errors.New("Order id not found")
}

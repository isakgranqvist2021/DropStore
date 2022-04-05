package order

import (
	"errors"

	"github.com/isakgranqvist2021/dropstore/src/services/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (order *Order) CreateOrder() (*primitive.ObjectID, error) {
	order.Status = "pending"

	createOptions := database.CreateOptions{Collection: "orders", Payload: order}

	return createOptions.InsertOne()
}

func (order *Order) UpdateStatus() error {
	payload := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "status",
				Value: order,
			},
		},
	}}

	createOptions := database.UpdateOptions{Collection: "orders", Payload: payload}

	return createOptions.UpdateOneByID(&order.ID)
}

func (order *Order) Validate() error {

	if len(order.Address) <= 0 {
		return errors.New("You must enter an address")
	}

	return nil
}

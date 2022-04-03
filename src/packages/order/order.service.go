package order

import (
	"github.com/isakgranqvist2021/dropstore/src/services/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(order Order) (*primitive.ObjectID, error) {
	order.Status = "pending"

	insertOptions := database.Options{Collection: "orders", Payload: order}

	return insertOptions.InsertOne()
}

func UpdateOrderStatus(ID *string, status string) error {
	payload := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "status",
				Value: status,
			},
		},
	}}

	insertOptions := database.Options{Collection: "orders", Payload: payload}

	return insertOptions.UpdateOneByID(ID)
}

package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (options Options) UpdateOneByID(ID *string) error {
	db, err := GetDb()

	if err != nil {
		return err
	}

	coll := db.Collection(options.Collection)

	objectId, err := primitive.ObjectIDFromHex(*ID)

	if err != nil {
		return err
	}

	result, err := coll.UpdateByID(
		context.TODO(),
		objectId,
		options.Payload,
	)

	fmt.Println(result.ModifiedCount)

	return err
}

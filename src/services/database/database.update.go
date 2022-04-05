package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (options UpdateOptions) UpdateOneByID(ID *string) error {
	db, err := GetDb()

	if err != nil {
		return err
	}

	coll := db.Collection(options.Collection)

	objectId, err := primitive.ObjectIDFromHex(*ID)

	if err != nil {
		return err
	}

	_, err := coll.UpdateByID(
		context.TODO(),
		objectId,
		options.Payload,
	)

	return err
}

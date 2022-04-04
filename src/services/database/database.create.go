package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (options CreateOptions) InsertOne() (*primitive.ObjectID, error) {
	db, err := GetDb()

	if err != nil {
		return nil, err
	}

	collection := db.Collection(options.Collection)

	result, err := collection.InsertOne(context.TODO(), options.Payload)

	if err != nil {
		return nil, err
	}

	insertedId := result.InsertedID.(primitive.ObjectID)

	return &insertedId, err
}

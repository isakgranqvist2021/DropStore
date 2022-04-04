package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (options ReadOptions) ReadMany(out interface{}) error {
	db, err := GetDb()

	if err != nil {
		return err
	}

	collection := db.Collection(options.Collection)

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return err
	}

	if err = cursor.All(context.TODO(), out); err != nil {
		return err
	}

	return nil
}

func (options ReadOptions) ReadOne(out interface{}) error {
	db, err := GetDb()

	if err != nil {
		return err
	}

	collection := db.Collection(options.Collection)

	result := collection.FindOne(context.TODO(), options.Filter)

	return result.Decode(out)
}

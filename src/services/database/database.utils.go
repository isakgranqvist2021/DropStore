package database

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetDb() (*mongo.Database, error) {
	if Client != nil {
		return Client.Database("dropstore"), nil
	}

	return nil, errors.New("You must connect to mongoDB")
}

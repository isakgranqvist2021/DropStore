package database

import (
	"context"
	"fmt"
	"os"

	"github.com/isakgranqvist2021/dropstore/src/services/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = nil

func Disconnect() error {
	fmt.Println("Disconnect")

	if err := Client.Disconnect(context.TODO()); err != nil {
		go logger.Log(err)

		return err
	}

	return nil
}

func Connect() error {
	fmt.Println("Connect")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		return err
	}

	Client = client

	return nil
}

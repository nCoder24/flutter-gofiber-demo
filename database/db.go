package database

import (
	"context"
	"demo/database/constant"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	client *mongo.Client
	UserDB UserDB
}

func GetDB(client *mongo.Client) DB {
	return DB{
		client: client,
		UserDB: getUserDB(client.Database(constant.DatabaseName)),
	}
}

func (db DB) Disconnect() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

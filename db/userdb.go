package db

import (
	"context"
	"demo/db/schema"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDB struct {
	accounts mongo.Collection
}

func getUserDB(db *mongo.Database) UserDB {
	return UserDB{*db.Collection("accounts")}
}

func (db UserDB) GetAccountByUsername(username string) (schema.Account, error) {
	filter := bson.D{{Key: "username", Value: username}}
	var account schema.Account

	err := db.accounts.FindOne(context.TODO(), filter).Decode(&account)
	return account, err
}

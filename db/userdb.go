package db

import (
	"context"
	"demo/core/models"
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
	account := new(schema.Account)

	err := db.accounts.FindOne(context.TODO(), filter).Decode(account)
	return *account, err
}

func (db UserDB) InsertNewAccount(acData models.AccountDetails) error {
	_, err := db.accounts.InsertOne(context.TODO(), schema.Account{
		Username: acData.Username,
		Password: acData.Password,
	})

	return err
}

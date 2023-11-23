package database

import (
	"context"
	"demo/core/models"
	"demo/database/schema"
	"demo/pkg/errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDB struct {
	accounts mongo.Collection
}

func getUserDB(db *mongo.Database) UserDB {
	return UserDB{*db.Collection("accounts")}
}

func (db UserDB) GetUserByUsername(username string) (schema.User, error) {
	filter := bson.D{{Key: "username", Value: username}}
	account := new(schema.User)

	err := db.accounts.FindOne(context.TODO(), filter).Decode(account)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return *account, errors.AccountNotFound
		}

		log.Println(err)
		return *account, errors.FailedToFetchAccount
	}

	return *account, err
}

func (db UserDB) InsertNewUser(acData models.UserDetails) error {
	_, err := db.accounts.InsertOne(context.TODO(), schema.User{
		Username: acData.Username,
		Password: acData.Password,
	})

	if err != nil {
		log.Println(err)
		return errors.FailedToCreateAccount
	}

	return nil
}

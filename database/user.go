package database

import (
	"context"
	"demo/core/models"
	"demo/database/constant"
	"demo/database/errors"
	"demo/database/schema"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDB struct {
	user mongo.Collection
}

func getUserDB(db *mongo.Database) UserDB {
	return UserDB{*db.Collection(constant.UserCollName)}
}

func (db UserDB) FindUserByUsername(username string) (schema.User, error) {
	filter := bson.D{{Key: "username", Value: username}}
	account := new(schema.User)

	err := db.user.FindOne(context.TODO(), filter).Decode(account)

	if err == nil {
		return *account, err
	}

	log.Println(err)
	if err == mongo.ErrNoDocuments {
		return *account, errors.ErrDocumentNotFound
	}

	return *account, errors.ErrFailedToFetch
}

func (db UserDB) InsertUser(acData models.UserData) error {
	_, err := db.user.InsertOne(context.TODO(), schema.User{
		Username: acData.Username,
		Password: acData.Password,
	})

	if err != nil {
		log.Println(err)
		return errors.ErrFailedToInsert
	}

	return nil
}

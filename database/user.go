package database

import (
	"context"
	"demo/core/models"
	"demo/database/constant"
	dbErrs "demo/database/errors"
	"demo/database/schema"
	"errors"
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
		return *account, nil
	}

	log.Println(err)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return *account, dbErrs.ErrDocumentNotFound
	}

	return *account, dbErrs.ErrFailedToFetch
}

// CheckUserExists TODO: Implement using aggregation instead of relying on errors
func (db UserDB) CheckUserExists(username string) (bool, error) {
	_, err := db.FindUserByUsername(username)

	if err == nil || errors.Is(err, dbErrs.ErrDocumentNotFound) {
		return true, nil
	}

	return false, err
}

func (db UserDB) InsertUser(acData models.UserData) error {
	_, err := db.user.InsertOne(context.TODO(), schema.User{
		Username: acData.Username,
		Password: acData.Password,
	})

	if err != nil {
		log.Println(err)
		return dbErrs.ErrFailedToInsert
	}

	return nil
}

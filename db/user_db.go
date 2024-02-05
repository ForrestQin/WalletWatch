package db

import (
	"WalletWatch/pkg/user"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var userDatabase *mongo.Database
var userCollection *mongo.Collection

func InitializeUserDatabaseSetting(client *mongo.Client) {
	userDatabase = client.Database("WalletWatch")
	userCollection = userDatabase.Collection("Users")
}

func CreateUser(newUser user.User) error {
	newUser.UserId = primitive.NewObjectID()
	newUser.CreatedDate = time.Now()

	_, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}
	return nil
}

func GetUserList(limit, offset int64) ([]user.User, error) {
	if userDatabase == nil || userCollection == nil {
		return nil, errors.New("Database not initialize!")
	}
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	cursor, err := userCollection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, errors.New("User collection has error.")
	}
	defer cursor.Close(context.TODO())

	var users []user.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, errors.New("Retrieve data has error.")
	}
	return users, nil
}

func GetUserId(email, password string) (string, error) {
	filter := bson.D{{"email", email}, {"password", password}}
	var userEntity user.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		return "", err
	}
	return userEntity.UserId.String(), nil
}

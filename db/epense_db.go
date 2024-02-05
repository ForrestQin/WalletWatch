package db

import "go.mongodb.org/mongo-driver/mongo"

var expenseCollection *mongo.Collection

func InitializeDatabaseSetting(client *mongo.Client) {
	expenseCollection = client.Database("WalletWatch").Collection("expense")
}

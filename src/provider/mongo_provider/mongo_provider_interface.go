package mongo_provider

import "go.mongodb.org/mongo-driver/mongo"

type IMongoProvider interface {
	GetMongoClient(timeout int) (*mongo.Client, error)
	GetDatabase(dbName string, timeout int) (*mongo.Database, error)
	GetCollection(dbName string, collectionName string, timeout int) (*mongo.Collection, error)
}
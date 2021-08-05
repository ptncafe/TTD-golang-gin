package mongo_provider

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var mongoClient *mongo.Client

type mongoProvider struct {
	uri    string
	logger *logrus.Logger
}

func NewMongoProvider(uri string, logger *logrus.Logger) *mongoProvider {
	return &mongoProvider{uri: uri, logger: logger}
}


// https://github.com/cavdy-play/go_mongo/blob/master/config/db.go
func (provider *mongoProvider) GetMongoClient(timeout int) (*mongo.Client, error) {
	if mongoClient != nil {
		return mongoClient, nil
	}
	provider.logger.Infof("GetMongoClient url %v",provider.uri)

	clientOptions := options.Client().ApplyURI(provider.uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		provider.logger.Errorf("GetMongoClient GetMongoClient %v", err)
		return client, errors.Wrap(err, "NewDataWarehouseProvider GetMongoClient")
	}

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	err = client.Connect(ctx)
	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		provider.logger.Errorf("NewMongoProvider Ping %v", err)
		return client, errors.Wrap(err, "NewMongoProvider Ping")
	} else {
		provider.logger.Infof("NewMongoProvider Ping good")
	}

	mongoClient = client
	return client, err
}


func (provider *mongoProvider) GetDatabase(dbName string, timeout int) (*mongo.Database, error) {
	client, err := provider.GetMongoClient(timeout)
	if err != nil {
		provider.logger.Errorf("NewMongoProvider GetDatabase %v", err)
		return nil, err
	}
	db := client.Database(dbName)
	return db, err
}


func (provider *mongoProvider) GetCollection(dbName string, collectionName string, timeout int) (*mongo.Collection, error) {
	db, err := provider.GetDatabase(dbName, timeout)
	if err != nil {
		provider.logger.Errorf("NewMongoProvider GetCollection %v", err)
		return nil, err
	}
	collection := db.Collection(collectionName)
	return collection, err
}

package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type MongoClient struct {
	URI    string
	Ctx    context.Context
	client *mongo.Client
}

var clientMutex sync.Mutex

func (c *MongoClient) Connect() {
	c.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	c.client, err = mongo.Connect(c.Ctx, options.Client().ApplyURI(c.URI))
	if err != nil {
		panic("mongo connect error")
	}
}

func (c *MongoClient) InitClient() {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	if c.client == nil {
		c.Connect()
	}
}

func (c *MongoClient) DB(dbname string) *mongo.Database {
	c.InitClient()
	return c.client.Database(dbname)
}

func (c *MongoClient) CL(dbname, clname string) *mongo.Collection {
	return c.DB(dbname).Collection(clname)
}

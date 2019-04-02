package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoClient struct {
	URI    string
	Ctx    context.Context
	client *mongo.Client
}

func (c *MongoClient) Connect() {
	c.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	c.client, err = mongo.Connect(c.Ctx, options.Client().ApplyURI(c.URI))
	if err != nil {
		panic("mongo connect error")
	}
}

func (c *MongoClient) DB(dbname string) *mongo.Database {
	if c.client == nil {
		c.Connect()
	}
	return c.client.Database(dbname)
}

func (c *MongoClient) CL(dbname, clname string) *mongo.Collection {
	return c.DB(dbname).Collection(clname)
}

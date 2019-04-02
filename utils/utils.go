package utils

import (
	"context"
	"github.com/regod/gwt"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoCollection(ctx *gwt.Context, dbName string, collectionName string) (collection *mongo.Collection, mongoctx context.Context) {
	mongo := ctx.GetStore("mongo").(*mongo.Client)
	mongoctx = ctx.GetStore("mongoctx").(context.Context)
	db := mongo.Database(dbName)
	collection = db.Collection(collectionName)
	return
}

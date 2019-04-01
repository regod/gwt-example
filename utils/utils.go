package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gwt"
)

func GetMongoCollection(ctx *gwt.Context, db_name string, collection_name string) (collection *mongo.Collection, mongoctx context.Context) {
	mongo := ctx.GetStore("mongo").(*mongo.Client)
	mongoctx = ctx.GetStore("mongoctx").(context.Context)
	db := mongo.Database(db_name)
	collection = db.Collection(collection_name)
	return
}

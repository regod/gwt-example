package middlewares

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/regod/gwt"
	"time"
)

func MongoDBInit(uri string) gwt.MiddlewareFunc {
	return func(handler gwt.HandlerFunc) gwt.HandlerFunc {
		return func(ctx *gwt.Context) error {
			mongoctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			client, err := mongo.Connect(mongoctx, options.Client().ApplyURI(uri))
			if err != nil {
				panic("mongo connect error")
			}
			ctx.SetStore("mongo", client)
			ctx.SetStore("mongoctx", mongoctx)

			return handler(ctx)
		}
	}
}

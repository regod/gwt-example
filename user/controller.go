package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gwt"
	"gwt_example/utils"
)

// RespData return structure
type RespData map[string]interface{}

func Create(ctx *gwt.Context) error {
	collection, mongoctx := utils.GetMongoCollection(ctx, DB_NAME, COLLECTION_NAME)
	user := User{
		Name:  ctx.PostForm().Get("name"),
		Phone: ctx.PostForm().Get("phone"),
	}
	res, err := collection.InsertOne(mongoctx, user)
	var data RespData
	if err == nil {
		data = RespData{
			"status": 0,
			"data":   map[string]string{"id": res.InsertedID.(primitive.ObjectID).Hex()},
		}
	} else {
		data = RespData{
			"status": 1,
			"errmsg": err.Error(),
		}
	}
	ctx.RespJson(200, data)
	return err
}

func UpdatePhone(ctx *gwt.Context) error {
	collection, mongoctx := utils.GetMongoCollection(ctx, DB_NAME, COLLECTION_NAME)
	userid := ctx.GetParam("id")
	phone := ctx.PostForm().Get("phone")

	user_objectid, _ := primitive.ObjectIDFromHex(userid)
	res, err := collection.UpdateOne(mongoctx, bson.M{"_id": user_objectid}, bson.M{"$set": bson.M{"phone": phone}})

	var data RespData
	if err == nil {
		data = RespData{
			"status": 0,
			"data":   map[string]string{"id": userid},
		}
	} else {
		data = RespData{
			"status": 1,
			"errmsg": err.Error(),
		}
	}
	ctx.RespJson(200, data)

	return err
}

func List(ctx *gwt.Context) error {
	collection, mongoctx := utils.GetMongoCollection(ctx, DB_NAME, COLLECTION_NAME)
	var users []User
	cursor, err := collection.Find(mongoctx, bson.D{})
	var data RespData
	if err == nil {
		defer cursor.Close(mongoctx)
		for cursor.Next(mongoctx) {
			var user User
			if err := cursor.Decode(&user); err == nil {
				users = append(users, user)
			}
		}
		data = RespData{
			"status": 0,
			"data":   users,
		}
	} else {
		data = RespData{
			"status": 1,
			"errmsg": err.Error(),
		}
	}
	ctx.RespJson(200, data)
	return err
}

func Delete(ctx *gwt.Context) error {
	collection, mongoctx := utils.GetMongoCollection(ctx, DB_NAME, COLLECTION_NAME)
	userid := ctx.GetParam("id")
	user_objectid, _ := primitive.ObjectIDFromHex(userid)
	var data RespData
	_, err := collection.DeleteOne(mongoctx, bson.M{"_id": user_objectid})
	if err == nil {
		data = RespData{
			"status": 0,
			"data":   map[string]string{"id": userid},
		}
	} else {
		data = RespData{
			"status": 1,
			"errmsg": err.Error(),
		}
	}
	ctx.RespJson(200, data)
	return err
}

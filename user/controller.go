package user

import (
	"github.com/regod/gwt"
	"github.com/regod/gwt-example/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RespData return structure
type RespData map[string]interface{}

const MongoURI = "mongo://127.0.0.1:27017"

func Create(ctx *gwt.Context) error {
	m := utils.MongoClient{URI: MongoURI}
	cl := m.CL(DBName, CollectionName)
	user := User{
		Name:  ctx.PostForm().Get("name"),
		Phone: ctx.PostForm().Get("phone"),
	}
	res, err := cl.InsertOne(m.Ctx, user)
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
	m := utils.MongoClient{URI: MongoURI}
	cl := m.CL(DBName, CollectionName)
	userid := ctx.GetParam("id")
	phone := ctx.PostForm().Get("phone")

	user_objectid, _ := primitive.ObjectIDFromHex(userid)
	_, err := cl.UpdateOne(m.Ctx, bson.M{"_id": user_objectid}, bson.M{"$set": bson.M{"phone": phone}})

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
	m := utils.MongoClient{URI: MongoURI}
	cl := m.CL(DBName, CollectionName)
	var users []User
	cursor, err := cl.Find(m.Ctx, bson.D{})
	var data RespData
	if err == nil {
		defer cursor.Close(m.Ctx)
		for cursor.Next(m.Ctx) {
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
	m := utils.MongoClient{URI: MongoURI}
	cl := m.CL(DBName, CollectionName)
	userid := ctx.GetParam("id")
	user_objectid, _ := primitive.ObjectIDFromHex(userid)
	var data RespData
	_, err := cl.DeleteOne(m.Ctx, bson.M{"_id": user_objectid})
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

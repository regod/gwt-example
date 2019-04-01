package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		ID    primitive.ObjectID "_id,omitempty"
		Name  string
		Phone string
	}

	//UserModel struct {
	//    user User
	//    db_name string
	//    collection_name string
	//    collection *mongo.Collection
	//    ctx context.Context
	//}
)

const DB_NAME = "gwt_example"
const COLLECTION_NAME = "user"

//func (u *UserModel) Create(user User) error {
//    _, err := u.collection.InsertOne(u.ctx, user)
//    return err
//}

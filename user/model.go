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
)

const DBName = "gwt_example"
const CollectionName = "user"

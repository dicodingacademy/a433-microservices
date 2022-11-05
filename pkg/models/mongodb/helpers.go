package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToObjID(s string) primitive.ObjectID {
	result, _ := primitive.ObjectIDFromHex(s)

	return result
}

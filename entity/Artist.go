package entity

import "github.com/globalsign/mgo/bson"

type Artist struct {
	Id      bson.ObjectId `json:"_id" bson:"_id"`
	Title   string        `json:"title" bson:"title"`
	Cover   string        `json:"cover" bson:"cover"`
	Details string        `json:"details" bson:"details"`
}

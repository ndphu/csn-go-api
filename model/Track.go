package model

import "github.com/globalsign/mgo/bson"

type Track struct {
	Id       bson.ObjectId `json:"_id" bson:"_id"`
	Title    string        `json:"title"  bson:"title"`
	Artist   string        `json:"artist" bson:"artists"`
	Quality  string        `json:"quality" bson:"quality"`
	Duration int           `json:"duration" bson:"duration"`
}

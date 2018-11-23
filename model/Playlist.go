package model

import (
	"github.com/globalsign/mgo/bson"
)

type Playlist struct {
	Id bson.ObjectId `json:"_id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Owner bson.ObjectId `json:"owner" bson:"owner,omitempty"`
	Shared string `json:"shared" bson:"shared"`
}


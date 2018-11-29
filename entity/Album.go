package entity

import "github.com/globalsign/mgo/bson"

type Album struct {
	Id        bson.ObjectId   `json:"_id" bson:"_id"`
	Title     string          `json:"title" bson:"title"`
	Artist    string          `json:"artist" bson:"artist"`
	Year      int             `json:"year" bson:"year"`
	PicMIME   string          `json:"picMIME" bson:"picMIME" `
	PicData   string          `json:"picData" bson:"picData"`
	Tracks    []bson.ObjectId `json:"tracks" bson:"tracks"`
	TrackList []*Track        `json:"trackList" bson:"trackList"`
}

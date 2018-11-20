package entity

import "github.com/globalsign/mgo/bson"

type DriveFile struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	TrackId string `json:"trackId" bson:"trackId"`
	Quality string `json:"quality" bson:"quality"`
	Name string `json:"name" bson:"name"`
	Size int64 `json:"size" bson:"size"`
	DriveFileId string `json:"driveId" bson:"driveId"`
	DriveAccount string `json:"driveAccount" bson:"driveAccount"`
}

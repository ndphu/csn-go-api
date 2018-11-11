package dao

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/entity"
)

func GetAllPlaylist(page int) ([]entity.Playlist, error) {
	var result []entity.Playlist
	err := Collection("playlist").Find(bson.M{}).Skip((page-1) * 24).Limit(24).All(&result)
	return result, err
}

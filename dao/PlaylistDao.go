package dao

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/entity"
)

func GetAllPlaylist(page int, size int) ([]entity.Playlist, error) {
	var result []entity.Playlist
	err := Collection("playlist").Find(bson.M{}).Skip((page - 1) * size).Limit(size).All(&result)
	return result, err
}

func SavePlaylist(playlist *entity.Playlist) (error) {
	err := Collection("playlist").Insert(playlist)
	if err != nil {
		fmt.Println("fail to insert playlist", err.Error())
	}
	return err
}

func FindPlaylistById(id bson.ObjectId) (*entity.Playlist, error) {
	var p entity.Playlist
	err := Collection("playlist").FindId(id).One(&p)
	return &p, err
}

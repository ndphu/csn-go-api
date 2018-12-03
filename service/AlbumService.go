package service

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
)

type AlbumService struct {
}

func (s *AlbumService) GetAlbumById(id string) (*entity.Album, error) {
	album := entity.Album{}
	err := dao.Collection("album").Pipe([]bson.M{
		{
			"$match": bson.M{
				"_id": bson.ObjectIdHex(id),
			},
		},
		{
			"$lookup": bson.M{
				"from":         "track",
				"foreignField": "_id",
				"localField":   "tracks",
				"as":           "trackList",
			},
		},
		{
			"$project": bson.M{
				"tracks":         0,
				"picData":        0,
				"picMIME":        0,
				"trackList.id3":  0,
				"trackList.link": 0,
			},
		},
	}).One(&album)

	return &album, err
}
func (s *AlbumService) GetRandomAlbums(number int) ([]entity.Album, error) {
	albums := make([]entity.Album, 0)

	err := dao.Collection("album").Pipe([]bson.M{
		{
			"$project": bson.M{
				"title": 1,
				"artist": 1,
				"year": 1,
				"trackCount": bson.M{ "$size": "$tracks" },
			},
		},
		{
			"$match": bson.M{
				"trackCount": bson.M{
					"$gt": 8,
				},
			},
		},
		{
			"$sample": bson.M{
				"size": number,
			},
		},

	}).All(&albums)

	return albums, err
}

var albumService *AlbumService

func GetAlbumService() (*AlbumService) {
	if albumService == nil {
		albumService = &AlbumService{
		}
	}

	return albumService
}

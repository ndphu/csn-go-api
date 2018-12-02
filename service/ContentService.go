package service

import (
	"encoding/base64"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
)

type ContentService struct{}

var contentService *ContentService

func GetContentService() (*ContentService) {
	if contentService == nil {
		contentService = &ContentService{}
	}

	return contentService
}

func (s*ContentService) GetAlbumCover(id string) ([]byte, string, error) {
	var album entity.Album
	err :=dao.Collection("album").FindId(bson.ObjectIdHex(id)).Select(bson.M{"picData": 1, "picMIME": 1}).One(&album)
	if err != nil {
		return nil, "", err
	}
	raw, err := base64.StdEncoding.DecodeString(album.PicData)
	return raw, album.PicMIME, err
}
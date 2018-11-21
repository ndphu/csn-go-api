package service

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
)

type ArtistService struct {

}

func (service ArtistService) FindArtistById(id string) (*entity.Artist, error) {
	var artist entity.Artist
	err := dao.Collection("artist").FindId(bson.ObjectIdHex(id)).One(&artist)
	return &artist, err
}

var artistService *ArtistService

func GetArtistService() *ArtistService {
	if artistService == nil {
		artistService = &ArtistService{

		}
	}

	return artistService
}
package service

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
)

type TrackService struct {

}

func (s *TrackService) GetTrackById(id string) (*entity.Track, error) {
	var track *entity.Track
	err := dao.Collection("track").FindId(bson.ObjectIdHex(id)).One(&track)
	return track, err
}
func (s *TrackService) GetTrackFiles(id string) ([]*entity.DriveFile, error) {
	var files = make([]*entity.DriveFile, 0)
	err := dao.Collection("file").Find(bson.M{"trackId": bson.ObjectIdHex(id)}).All(&files)
	return files, err
}

var trackService *TrackService

func GetTrackService() *TrackService {
	if trackService == nil {
		trackService = &TrackService{

		}
	}

	return trackService
}

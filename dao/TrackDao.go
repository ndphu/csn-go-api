package dao

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/entity"
	"github.com/ndphu/csn-go-api/model"
)

func SaveTrackModel(m *model.Track) (*entity.Track, error) {
	e := entity.Track{
		Title: m.Title,
		Quality: m.Quality,
		Artist: m.Artist,
		Duration: m.Duration,
		Id: bson.NewObjectId(),
	}
	err := Collection("track").Insert(&e)
	return &e, err
}
package service

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
)

type DBSearchService struct {
}

var dbSearchService *DBSearchService

func GetDBSearchService() *DBSearchService {
	if dbSearchService == nil {
		dbSearchService = &DBSearchService{
		}
	}

	return dbSearchService
}

type QuickSearchEntry struct {
	Id       string `json:"_id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

func (*DBSearchService) QuickSearch(query string) ([]QuickSearchEntry, error) {
	var tracks []entity.Track
	err := dao.Collection("track").
		Find(bson.M{"title": bson.RegEx{Pattern: query, Options: "i"},}).
		Select(bson.M{"link": 0, "quality": 0, "duration": 0, "file": 0}).
		Limit(10).All(&tracks)
	if err != nil {
		return nil, err
	}
	//entries := [len(tracks)]QuickSearchEntry{}
	var entries []QuickSearchEntry
	for _, track := range tracks {
		entries = append(entries, QuickSearchEntry{
			Id:       track.Id.Hex(),
			Title:    track.Title,
			Subtitle: track.Artist,
		})
	}
	return entries, nil
}

func (*DBSearchService) SearchTracks(query string, page int, size int) ([]entity.Track, error) {
	var tracks []entity.Track
	err := dao.Collection("track").
		Find(bson.M{"title": bson.RegEx{Pattern: query, Options: "i"},}).
		Skip((page - 1) * size).Limit(size).All(&tracks)
	fmt.Println("found", len(tracks), "items")
	return tracks, err
}

func (*DBSearchService) SearchArtist(query string, page int, size int) ([]entity.Artist, error) {
	var artists []entity.Artist
	err := dao.Collection("artist").
		Find(bson.M{"title": bson.RegEx{Pattern: query, Options: "i"},}).
		Skip((page - 1) * size).Limit(size).All(&artists)
	fmt.Println("found", len(artists), "artists")
	return artists, err
}

func (*DBSearchService) SearchTracksByArtist(query string, page int, size int) ([]entity.Track, error) {
	var tracks []entity.Track
	err := dao.Collection("track").
		Find(bson.M{"artists": bson.RegEx{Pattern: query, Options: "i"},}).
		Skip((page - 1) * size).Limit(size).All(&tracks)
	fmt.Println("found", len(tracks), "items")
	return tracks, err
}

func condition(query string) []bson.M {
	return []bson.M{
		{"title": bson.RegEx{Pattern: ".*" + query + ".*", Options: "i"}},
	}
}

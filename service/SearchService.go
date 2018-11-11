package service

import (
	"fmt"
	"github.com/ndphu/csn-go-api/model"
	"net/url"
)

type SearchService struct {
	crawService *CrawService
}

var searchService *SearchService

func GetSearchService() *SearchService {
	if searchService == nil {
		searchService = &SearchService{
			crawService: GetCrawService(),
		}
	}

	return searchService
}

func (s *SearchService) Search(name string, p int) ([]model.Track, error) {
	fmt.Println("Search for name", name)
	raw := fmt.Sprintf(SearchByTrackName, url.QueryEscape(name), p)
	return s.crawService.crawTracksFromUrl(raw)
}

func (s *SearchService) SearchByArtist(a string, p int) ([]model.Track, error) {
	raw := fmt.Sprintf(SearchByArtist, url.QueryEscape(a), p)
	return s.crawService.crawTracksFromUrl(raw)
}

package service

import "github.com/ndphu/csn-go-api/model"

type SourceService struct {
	CrawService *CrawService
}

var sourceService *SourceService

func GetSourceService() *SourceService {
	if sourceService == nil {
		sourceService = &SourceService{
			CrawService: GetCrawService(),
		}
	}

	return sourceService
}

func (*SourceService) GetSourcesFromTrackUrl(trackUrl string) ([]model.Source, error) {
	return crawService.crawSources(trackUrl)
}

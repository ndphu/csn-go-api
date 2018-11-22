package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ndphu/csn-go-api/model"
	"strings"
)

var (
	SearchUrl         = "http://chiasenhac.vn/search.php"
	SearchByArtist    = SearchUrl + "?mode=artist&s=%s&order=quality&cat=music&page=%d"
	SearchByAlbum    = SearchUrl + "?mode=album&s=%s&order=quality&cat=music&page=%d"
	SearchByTrackName = SearchUrl + "?mode=&s=%s&order=quality&cat=music&page=%d"
)

type CrawService struct{}

var crawService *CrawService

func GetCrawService() *CrawService {
	if crawService == nil {
		crawService = &CrawService{}
	}

	return crawService
}

func (s *CrawService) crawSources(trackUrl string) ([]model.Source, error)  {
	downloadUrl := strings.Replace(trackUrl, ".html", "_download.html", 1)
	fmt.Println("download url:", downloadUrl)
	doc, err := goquery.NewDocument(downloadUrl)
	if err != nil {
		return nil, err
	}
	var sources []model.Source
	doc.Find("#downloadlink2 a").Each(func(__ int, link *goquery.Selection) {
		if link.Find("span").Length() > 0 {
			source := model.Source{
				Source:  link.AttrOr("href", ""),
				Quality: link.Find("span").First().Text(),
			}
			fmt.Println("found source:", source.Quality)
			sources = append(sources, source)
		}
	})
	return sources, nil
}
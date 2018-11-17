package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ndphu/csn-go-api/model"
	"github.com/ndphu/csn-go-api/utils"
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

func (s *CrawService) crawTracksFromUrl(raw string) ([]model.Track, error) {
	fmt.Println("Querying " + raw)
	doc, err := goquery.NewDocument(raw)
	if err != nil {
		return nil, err
	}
	tracks := make([]model.Track, 0)
	doc.Find(".page-dsms tbody tr").EachWithBreak(func(idx int, s *goquery.Selection) bool {
		if idx == 0 {
			return true
		}
		track := model.Track{}
		s.Find("td").Each(func(col int, r *goquery.Selection) {
			if col == 1 {
				r.Find("p").Each(func(idx int, p *goquery.Selection) {
					if idx == 0 {
						anchor := p.Find("a").First()
						track.Title = anchor.Text()
						track.Link = anchor.AttrOr("href", "")
					} else if idx == 1 {
						track.Artist = p.Text()
					}
				})
			} else if col == 2 {
				r.Find("span").Each(func(iii int, span *goquery.Selection) {
					if iii == 1 {
						track.Quality = span.Text()
					}
				})

				str := r.Find("span").First().Text()
				track.Duration = utils.GetSecondFromString(strings.Replace(str, track.Quality, "", -1))
			}
		})

		tracks = append(tracks, track)
		if idx >= 25 {
			return false
		}
		return true
	})

	return tracks, nil
}

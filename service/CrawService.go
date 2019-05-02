package service

import (
	"fmt"
	"net/http"
	"net/url"
	"github.com/ndphu/csn-go-api/model"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/ndphu/csn-go-api/utils"
	"time"
)

var (
	SearchUrl         = "http://search.chiasenhac.vn/search.php"
	SearchByArtist    = SearchUrl + "?mode=artist&s=%s&order=quality&cat=music&page=%d"
	SearchByTrackName = SearchUrl + "?mode=&s=%s&order=quality&cat=music&page=%d"
)

type CrawService struct{}

func (s *CrawService) CrawByArtist(a string, p int) ([]*model.Track, error) {
	raw := fmt.Sprintf(SearchByArtist, url.QueryEscape(a), p)
	return s.CrawTracksFromUrl(raw)
}

func (s *CrawService) Search(name string, p int) ([]*model.Track, error) {
	raw := fmt.Sprintf(SearchByTrackName, url.QueryEscape(name), p);
	return s.CrawTracksFromUrl(raw)
}

func (s *CrawService) CrawSources(trackUrl string) ([]model.Source, error)  {
	//downloadUrl := strings.Replace(trackUrl, ".html", "_download.html", 1)
	downloadUrl := strings.Replace(trackUrl, "beta.", "", 1)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(downloadUrl)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	var sources []model.Source
	doc.Find("ul.download_status > li > a.download_item").Each(func(__ int, link *goquery.Selection) {
		if link.Find("span").Length() > 0 {
			mp3Link := link.AttrOr("href", "")
			if mp3Link == "" || strings.Index(mp3Link, "javascript") == 0 {
				return
			}
			source := model.Source{
				Source:  mp3Link,
				Quality: link.Find("span").First().Text(),
			}
			sources = append(sources, source)
		}
	})
	return sources, nil
}

func (s *CrawService) CrawTracksFromUrl(raw string) ([]*model.Track, error) {
	fmt.Println("Querying " + raw)
	doc, err := goquery.NewDocument(raw)
	if err != nil {
		return nil, err
	}
	tracks := make([]*model.Track, 0)
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

		tracks = append(tracks, &track)
		if idx >= 25 {
			return false
		}
		return true
	})

	return tracks, nil
}

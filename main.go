package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/model"
	"github.com/ndphu/csn-go-api/service"
	"log"
	"sync"
)

type CrawSourceRequest struct {
	URL string `json:"url"`
}

type SearchInput struct {
	Query string `json:"query" bson:"type"`
	Type string `json:"type" bson:"type"`
}


func main() {
	r := gin.Default()

	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowCredentials = true
	c.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	c.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-Requested-With"}

	r.Use(cors.New(c))

	crawService := service.CrawService{}

	api := r.Group("/api")
	search := api.Group("/search")
	{
		search.POST("", func(c *gin.Context) {
			si := SearchInput{}
			if err := c.BindJSON(&si); err != nil {
				c.JSON(400, gin.H{"err": err})
				return
			}

			log.Printf("searching for [%s] with type=[%s]", si.Query, si.Type)

			resultMap := make(map[int][]*model.Track)

			wg := sync.WaitGroup{}
			wg.Add(5)

			for i := 1; i <=5; i++ {
				go func(page int) {
					defer wg.Done()
					tracks, err := crawService.Search(si.Query, page)
					if err == nil {
						resultMap[page] = tracks
					}
				}(i)
			}

			wg.Wait()

			var tracks []*model.Track

			for i := 1; i <=5; i++ {
				if _track, ok := resultMap[i]; ok {
					tracks = append(tracks, _track...)
				}
			}


			fmt.Printf("Found %d result\n", len(tracks))

			c.JSON(200, gin.H{
				"request": si,
				"result": tracks,
			})
		})
		search.GET("/byArtist/:artistName/tracks", func(c *gin.Context) {
			//name, err := base64.StdEncoding.DecodeString(c.Param("artistName"))
			//if err != nil {
			//	c.JSON(500, gin.H{"err": err})
			//} else {
			//	//page := utils.GetIntQuery(c, "page", 1)
			//	//tracks, err := crawService.CrawByArtist(string(name), page)
			//	//returnTracksOrError(c, tracks, err)
			//}
		})
	}

	source := api.Group("/source")
	{
		source.POST("", func(c *gin.Context) {
			request := CrawSourceRequest{}
			err := c.BindJSON(&request)
			if err != nil {
				c.JSON(500, gin.H{"err": err})
			} else {
				fmt.Println("requesting base64 url: " + request.URL)
				realUrl, err := base64.StdEncoding.DecodeString(request.URL)
				if err != nil {
					c.JSON(500, gin.H{"err": err})
				} else {
					fmt.Println("real url: " + string(realUrl))
					sources, err := crawService.CrawSources(string(realUrl))
					if err != nil {
						c.JSON(500, gin.H{"err": err})
					} else {
						c.JSON(200, sources)
					}
				}
			}
		})
	}

	fmt.Println("Starting server")
	r.Run()
}

func returnTracksOrError(c *gin.Context, tracks []model.Track, err error) {
	if err != nil {
		c.JSON(500, gin.H{"err": err})
	} else {
		c.JSON(200, tracks)
	}
}

package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/service"
	"github.com/ndphu/csn-go-api/model"
	"github.com/ndphu/csn-go-api/utils"
	"encoding/base64"
)

type CrawSourceRequest struct {
	URL string `json:"url"`
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
		search.GET("/q/:query", func(c *gin.Context) {
			query, err := base64.StdEncoding.DecodeString(c.Param("query"))
			page := utils.GetIntQuery(c, "page", 1)
			tracks, err := crawService.Search(string(query), page)
			returnTracksOrError(c, tracks, err)
		})
		search.GET("/byArtist/:artistName/tracks", func(c *gin.Context) {
			name, err := base64.StdEncoding.DecodeString(c.Param("artistName"))
			if err != nil {
				c.JSON(500, gin.H{"err": err})
			} else {
				page := utils.GetIntQuery(c, "page", 1)
				tracks, err := crawService.CrawByArtist(string(name), page)
				returnTracksOrError(c, tracks, err)
			}
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

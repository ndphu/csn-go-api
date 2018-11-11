package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/service"
	"github.com/ndphu/csn-go-api/utils"
)

func SearchController(api *gin.RouterGroup) {

	searchService := service.GetSearchService()

	api.GET("/q/:query", func(c *gin.Context) {
		query, err := base64.StdEncoding.DecodeString(c.Param("query"))
		page := utils.GetIntQuery(c, "page", 1)
		tracks, err := searchService.Search(string(query), page)
		utils.ReturnTracksOrError(c, tracks, err)
	})

	api.GET("/byArtist/:artistName/tracks", func(c *gin.Context) {
		name, err := base64.StdEncoding.DecodeString(c.Param("artistName"))
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			page := utils.GetIntQuery(c, "page", 1)
			tracks, err := searchService.SearchByArtist(string(name), page)
			utils.ReturnTracksOrError(c, tracks, err)
		}
	})

}

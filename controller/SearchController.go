package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/entity"
	"github.com/ndphu/csn-go-api/service"
	"github.com/ndphu/csn-go-api/utils"
)

func SearchController(api *gin.RouterGroup) {

	searchService := service.GetDBSearchService()

	api.GET("/q/:query", func(c *gin.Context) {
		query, err := base64.StdEncoding.DecodeString(c.Param("query"))
		fmt.Println("search", string(query))
		page := utils.GetIntQuery(c, "page", 1)
		size := utils.GetIntQuery(c, "size", 25)
		tracks, err := searchService.SearchTracks(string(query), page, size)
		if tracks == nil {
			tracks = make([]entity.Track, 0)
		}
		utils.ReturnResponseOrError(c, tracks, err)
	})

	api.GET("/q/:query/quick", func(c *gin.Context) {
		query, err := base64.StdEncoding.DecodeString(c.Param("query"))
		fmt.Println("quick search", string(query))
		entries, err := searchService.QuickSearch(string(query))
		if entries == nil {
			entries = make([]service.QuickSearchEntry, 0)
		}
		utils.ReturnResponseOrError(c, entries, err)
	})

	api.GET("/q/:query/artist" , func(c *gin.Context) {
		query, err := base64.StdEncoding.DecodeString(c.Param("query"))
		fmt.Println("search", string(query))
		page := utils.GetIntQuery(c, "page", 1)
		size := utils.GetIntQuery(c, "size", 25)
		artists, err := searchService.SearchArtist(string(query), page, size)
		if artists == nil {
			artists = make([]entity.Artist, 0)
		}
		utils.ReturnResponseOrError(c, artists, err)
	})

	api.GET("/q/:query/artist/tracks" , func(c *gin.Context) {
		query, err := base64.StdEncoding.DecodeString(c.Param("query"))
		fmt.Println("search", string(query))
		page := utils.GetIntQuery(c, "page", 1)
		size := utils.GetIntQuery(c, "size", 25)
		tracks, err := searchService.SearchTracksByArtist(string(query), page, size)
		if tracks == nil {
			tracks = make([]entity.Track, 0)
		}
		utils.ReturnResponseOrError(c, tracks, err)
	})

	//api.GET("/byArtist/:artistName/tracks", func(c *gin.Context) {
	//	name, err := base64.StdEncoding.DecodeString(c.Param("artistName"))
	//	if err != nil {
	//		c.JSON(500, gin.H{"err": err})
	//	} else {
	//		page := utils.GetIntQuery(c, "page", 1)
	//		tracks, err := searchService.SearchByArtist(string(name), page)
	//		utils.ReturnTracksOrError(c, tracks, err)
	//	}
	//})
	//
	//api.GET("/byAlum/:albumName/tracks", func(c *gin.Context) {
	//	name, err := base64.StdEncoding.DecodeString(c.Param("albumName"))
	//	if err != nil {
	//		c.JSON(500, gin.H{"err": err})
	//	} else {
	//		page := utils.GetIntQuery(c, "page", 1)
	//		tracks, err := searchService.SearchByArtist(string(name), page)
	//		utils.ReturnTracksOrError(c, tracks, err)
	//	}
	//})
}

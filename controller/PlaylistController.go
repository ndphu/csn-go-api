package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/entity"
	"github.com/ndphu/csn-go-api/service"
	"strconv"
)

func PlaylistController(r *gin.RouterGroup) {
	playlistService := service.GetPlaylistService()
	r.GET("", func(c *gin.Context) {
		page, e := strconv.Atoi(c.Query("page"))
		if e != nil || page < 1{
			page = 1
		}
		size, e := strconv.Atoi(c.Query("size"))
		if e != nil {
			size = 24
		}

		playlist, e := playlistService.FindAllPlaylist(page, size)
		if e != nil {
			c.Error(e)
			return
		}
		if playlist == nil {
			playlist = []entity.Playlist{}
		}

		c.JSON(200, playlist)
	})

	r.POST("", func(c *gin.Context) {
		fmt.Println("creating new playlist...")
		playlist := entity.Playlist{}
		err := c.BindJSON(&playlist)
		if err != nil {
			c.JSON(400, "Bad request: "+err.Error())
		} else {
			err = playlistService.CreatePlaylist(&playlist)
			if err != nil {
				c.JSON(500, err.Error())
			} else {
				c.JSON(201, playlist)
			}
		}
	})

	r.GET("/:id", func(c *gin.Context) {
		fmt.Println("get playlist", c.Param("id"))
		playlist, err := playlistService.FindPlaylistById(c.Param("id"))
		if err != nil {
			c.JSON(500, err.Error())
		} else {
			c.JSON(200, playlist)
		}
	})
}

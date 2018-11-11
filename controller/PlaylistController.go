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
		if e != nil {
			c.Error(e)
			return
		}

		playlist, e := playlistService.FindAllPlaylist(page)
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
		// new playlist
		fmt.Print("creating new playlist...")
	})
}

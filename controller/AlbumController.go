package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/service"
)

func AlbumController(api *gin.RouterGroup) {
	albumService := service.GetAlbumService()
	api.GET("/id/:id", func(c *gin.Context) {
		album, err := albumService.GetAlbumById(c.Param("id"))
		if err != nil {
			ServerError("fail to get album with id "+c.Param("id"), err, c)
			return
		}
		c.JSON(200, album)
	})
}

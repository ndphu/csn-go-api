package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/service"
	"github.com/ndphu/csn-go-api/utils"
)

func RecommendController(api *gin.RouterGroup) {
	albumService := service.GetAlbumService()
	api.GET("/album/:number", func(c *gin.Context) {
		albums, err := albumService.GetRandomAlbum(utils.GetIntParam(c, "size", 16))
		if err != nil {
			ServerError("fail to get random album", err, c)
			return
		}
		c.JSON(200, albums)
	})
}

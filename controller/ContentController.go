package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/service"
)

func ContentController(g * gin.RouterGroup)  {
	contentService :=  service.GetContentService()
	g.GET("/album/:id/cover", func(c *gin.Context) {
		albumId := c.Param("id")
		if albumId == "" {
			BadRequest("missing albumId", nil, c)
			return
		}

		raw, mime, err := contentService.GetAlbumCover(albumId)
		if err != nil {
			ServerError("fail to get album cover", err, c)
			return
		}

		c.Data(200, mime, raw)
	})
}
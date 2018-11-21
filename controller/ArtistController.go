package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/service"
	"github.com/ndphu/csn-go-api/utils"
)

func ArtistController(api *gin.RouterGroup)  {

	artistService := service.GetArtistService()

	api.GET("/:id", func(c *gin.Context) {
		artist, err := artistService.FindArtistById(c.Param("id"))
		utils.ReturnResponseOrError(c, artist, err)
	})

}

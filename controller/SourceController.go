package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/controller/request"
	"github.com/ndphu/csn-go-api/service"
)

func SourceController(api *gin.RouterGroup) {

	sourceService := service.GetSourceService()

	api.POST("", func(c *gin.Context) {
		r := request.SourceRequest{}
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			fmt.Println("requesting base64 url: " + r.URL)
			realUrl, err := base64.StdEncoding.DecodeString(r.URL)
			if err != nil {
				c.JSON(500, gin.H{"err": err})
			} else {
				fmt.Println("real url: " + string(realUrl))
				sources, err := sourceService.GetSourcesFromTrackUrl(string(realUrl))
				if err != nil {
					c.JSON(500, gin.H{"err": err})
				} else {
					c.JSON(200, sources)
				}
			}
		}
	})
}

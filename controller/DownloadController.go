package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/service"
)

func DownloadController(g*gin.RouterGroup)  {
	accountService := service.GetAccountService()

	g.GET("/file/:id", func(c *gin.Context) {
		link, err := accountService.GetDownloadLink(c.Param("id"))
		if err != nil {
			ServerError("Fail to get download link", err, c)
			return
		}
		c.JSON(200, gin.H{"link": link})
	})
}

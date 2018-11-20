package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/entity"
	"github.com/ndphu/csn-go-api/service"
)

func AccountController(r *gin.RouterGroup) {
	accountService := service.GetAccountService()

	r.POST("/", func(c *gin.Context) {
		da := entity.DriveAccount{}
		err := c.BindJSON(&da)
		if err != nil {
			BadRequest("Fail to parse body", err, c)
		}
		err = accountService.Save(&da)
		if err != nil {
			ServerError("Fail to save drive account", err, c)
		}
	})

	r.GET("/", func(c *gin.Context) {
		accList, err := accountService.FindAll()
		if err != nil {
			ServerError("Fail to get account list", err, c)
		}
		c.JSON(200, accList)
	})

	r.POST("/:id/key", func(c *gin.Context) {
		body, err:= c.GetRawData()
		if err != nil {
			BadRequest("Request required body as base64",err,c)
		}
		var keyDecoded []byte
		count ,err := base64.StdEncoding.Decode(keyDecoded, body)
		if err != nil || count == 0{
			BadRequest("Fail to decode base64 key data", err, c)
		}
		err = accountService.InitializeKey(c.Param("id"), keyDecoded)
		if err != nil {
			ServerError("Fail to initialize key for account", err, c)
		}

		account, err := accountService.FindAccount(c.Param("id"))
		ServerError("Fail to query account", err, c)

		c.JSON(200, account)
	})
}

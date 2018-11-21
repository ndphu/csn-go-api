package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/entity"
	"github.com/ndphu/csn-go-api/service"
	helper "github.com/ndphu/google-api-helper"
	"strconv"
)

func AccountController(r *gin.RouterGroup) {
	accountService := service.GetAccountService()

	r.POST("", func(c *gin.Context) {
		da := entity.DriveAccount{}
		err := c.BindJSON(&da)
		if err != nil {
			BadRequest("Fail to parse body", err, c)
			return
		}

		keyDecoded ,err := base64.StdEncoding.DecodeString(da.Key)
		if err != nil {
			BadRequest("Fail to decode base64 key data", err, c)
			return
		}

		driveService, err := helper.GetDriveService([]byte(keyDecoded))
		if err != nil {
			BadRequest("Invalid JSON Key", err, c)
			return
		}

		if _, err := driveService.GetQuotaUsage(); err != nil {
			BadRequest("Invalid JSON Key", err, c)
			return
		}

		accountService.InitializeKey(&da, keyDecoded)
		if err != nil {
			ServerError("Fail to key data", err, c)
			return
		}

		err = accountService.Save(&da)
		if err != nil {
			ServerError("Fail to save drive account", err, c)
			return
		}
		c.JSON(200, da)
	})

	r.GET("", func(c *gin.Context) {
		accList, err := accountService.FindAll()
		if err != nil {
			ServerError("Fail to get account list", err, c)
			return
		}
		if accList == nil {
			accList = []entity.DriveAccount{}
		}
		c.JSON(200, accList)
	})

	r.GET("/:id", func(c *gin.Context) {
		acc, err := accountService.FindAccount(c.Param("id"))
		if err != nil {
			ServerError("Fail to get account", err, c)
			return
		}
		driveService, err := helper.GetDriveService([]byte(acc.Key))
		if err != nil {
			ServerError("Fail to initialize drive service", err, c)
			return
		}
		quota, err := driveService.GetQuotaUsage()
		if err != nil {
			ServerError("Fail to get quota for account", err, c)
			return
		}
		c.JSON(200, gin.H{
			"_id":   acc.Id,
			"name":  acc.Name,
			"desc":  acc.Desc,
			"quota": quota,
		})
	})

	r.POST("/:id/key", func(c *gin.Context) {
		body, err:= c.GetRawData()
		if err != nil {
			BadRequest("Request required body as base64",err,c)
			return
		}
		keyDecoded ,err := base64.StdEncoding.DecodeString(string(body))
		if err != nil {
			BadRequest("Fail to decode base64 key data", err, c)
			return
		}
		err = accountService.UpdateKey(c.Param("id"), keyDecoded)
		if err != nil {
			ServerError("Fail to initialize key for account", err, c)
			return
		}

		account, err := accountService.FindAccount(c.Param("id"))
		if err != nil {
			ServerError("Fail to query account", err, c)
			return
		}
		c.JSON(200, account)
	})

	r.GET("/:id/files", func(c *gin.Context) {
		page,err := strconv.Atoi(c.Query("page"))
		if err != nil {
			BadRequest("Invalid page parameter", err, c)
			return
		}
		size,err := strconv.Atoi(c.Query("size"))
		if err != nil {
			BadRequest("Invalid size parameter", err, c)
			return
		}
		acc, err := accountService.FindAccount(c.Param("id"))
		if err != nil {
			ServerError("Fail to get account", err, c)
			return
		}
		driveService, err := helper.GetDriveService([]byte(acc.Key))
		if err != nil {
			ServerError("Fail to initialize drive service", err, c)
			return
		}
		files, err := driveService.ListFiles(page, int64(size))
		if err != nil {
			ServerError("Fail to list file", err, c)
			return
		}
		c.JSON(200, files)
	})

	r.GET("/:id/file/:fileId/download", func(c *gin.Context) {
		acc, err := accountService.FindAccount(c.Param("id"))
		if err != nil {
			ServerError("Fail to get account", err, c)
			return
		}
		driveService, err := helper.GetDriveService([]byte(acc.Key))
		if err != nil {
			ServerError("Fail to initialize drive service", err, c)
			return
		}
		link, err := driveService.GetDownloadLink(c.Param("fileId"))
		if err != nil {
			ServerError("Fail to get download link", err, c)
			return
		}
		c.JSON(200, gin.H{"link": link})
	})
}
package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/controller"
	"github.com/ndphu/csn-go-api/service"
)

func main() {
	r := gin.Default()

	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowCredentials = true
	c.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	c.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-Requested-With"}

	r.Use(cors.New(c))

	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	controller.SearchController(api.Group("/search"))
	controller.SourceController(api.Group("/source"))
	controller.PlaylistController(api.Group("/playlist"))
	controller.ArtistController(api.Group("/artist"))
	controller.AccountController(api.Group("/manage/driveAccount"))
	controller.TrackController(api.Group("/track"))
	controller.DownloadController(api.Group("/download"))
	controller.AlbumController(api.Group("/album"))

	scheduleService, err := service.GetScheduleService()
	if err != nil {
		panic(err)
	}
	scheduleService.Start()

	fmt.Println("Starting server")
	r.Run()
}
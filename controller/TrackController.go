package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/entity"
	"github.com/ndphu/csn-go-api/service"
	"sync"
)

func TrackController(g *gin.RouterGroup)  {
	trackService := service.GetTrackService()
	accountService := service.GetAccountService()
	g.GET("/:id", func(c *gin.Context) {
		track, err := trackService.GetTrackById(c.Param("id"))
		if err != nil {
			ServerError("fail to get track", err, c)
			return
		}
		c.JSON(200, track)
	})

	g.GET("/:id/files", func(c *gin.Context) {
		files, err := trackService.GetTrackFiles(c.Param("id"))
		if err != nil {
			ServerError("fail to get files of track", err, c)
			return
		}
		c.JSON(200, files)
	})

	g.GET("/:id/sources", func(c *gin.Context) {
		files, err := trackService.GetTrackFiles(c.Param("id"))
		if err != nil {
			ServerError("fail to get files of track", err, c)
			return
		}

		type Source struct {
			Id string `json:"_id"`
			FileName string `json:"fileName"`
			Quality string `json:"quality"`
			Source string `json:"source"`

		}

		res := make([]Source, len(files))

		wg := sync.WaitGroup{}

		for idx, file := range files {
			wg.Add(1)
			go func(index int, f *entity.DriveFile) {
				link, _ := accountService.GetDownloadLink(file.Id.Hex())
				res[index] = Source{
					Id: f.Id.Hex(),
					FileName: f.Name,
					Quality: f.Quality,
					Source: link,
				}
				wg.Done()
			}(idx, file)
		}

		wg.Wait()

		c.JSON(200, res)
	})
}

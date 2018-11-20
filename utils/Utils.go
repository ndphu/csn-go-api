package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/csn-go-api/model"
	"log"
	"strconv"
	"strings"
)

func ReturnTracksOrError(c *gin.Context, tracks []model.Track, err error) {
	if err != nil {
		c.JSON(500, gin.H{"err": err})
	} else {
		c.JSON(200, tracks)
	}
}

func GetSecondFromString(input string) int {
	chunks := strings.Split(input, ":")
	min, _ := strconv.Atoi(chunks[0])
	sec, _ := strconv.Atoi(chunks[1])
	return min*60 + sec
}

func GetIntQuery(c *gin.Context, key string, defaultValue int) int {
	page, parseError := strconv.Atoi(c.DefaultQuery(key, strconv.Itoa(defaultValue)))
	if parseError == nil {
		return page
	}
	return defaultValue
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
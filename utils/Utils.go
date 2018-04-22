package utils

import (
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
)

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

package plain

import "github.com/gin-gonic/gin"

func Route(router *gin.RouterGroup) {
	router.GET("/", plainFunc)
}

func plainFunc(c *gin.Context) {
	c.JSON(200, gin.H{"Hello":"World"})
}
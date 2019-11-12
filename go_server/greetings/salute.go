package greetings

import "github.com/gin-gonic/gin"

func Salute(router *gin.RouterGroup) {
	router.GET("/:name", getSalute)
}


func getSalute(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{"Hello, friend": name})
}

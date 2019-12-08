package router

import (
	"fmt"
	"log"
	"net/http"

	"Go_Gingonic_Redis/models"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	router.POST("/", loginIdCard)
}


func loginIdCard(c *gin.Context) {
	var token models.UserDTO
	err := c.BindJSON(&token)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	fmt.Println(token)
}
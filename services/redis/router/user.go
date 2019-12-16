package router

import (
	"fmt"
	"log"
	"net/http"

	"Go_Gingonic_Redis/models"
	"Go_Gingonic_Redis/redis"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	router.POST("/", loginIdCard)
	router.GET("/:user", getUserToken)
}


func loginIdCard(c *gin.Context) {
	var user models.UserDTO
	err := c.BindJSON(&token)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
    client := redis.Connect()
    fmt.Print(client)
    name := user.Name
    token := user.Token
    fmt.Printf("%T & %T", name, token)
    e := client.Set(name, token, 0)
    if e != nil {
        fmt.Println(admin)
    }

	fmt.Println(token)
}

func getUserToken(c *gin.Context) {
	fmt.Print("getting redis token")
	user := c.Param("user")
	fmt.Print(admin)
	client := redis.Connect()
	result, e := client.Get(user).Result()
	if e != nil {
		panic(e)
	}
	fmt.Print(result)
}

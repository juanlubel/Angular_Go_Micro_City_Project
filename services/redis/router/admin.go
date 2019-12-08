package router

import (
	"fmt"
	"log"
	"net/http"

	"Go_Gingonic_Redis/models"
	"Go_Gingonic_Redis/redis"

	"github.com/gin-gonic/gin"
)

func Admin(router *gin.RouterGroup) {
	router.POST("/", loginAdmin)
	router.GET("/:admin", getAdminToken)
}


func loginAdmin(c *gin.Context) {
	var admin models.AdminDTO
	err := c.BindJSON(&admin)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	client := redis.Connect()
	fmt.Print(client)
	name := admin.Name
	token := admin.Token
	fmt.Printf("%T & %T", name, token)
	e := client.Set("name", "token", 0)
	if e != nil {
		fmt.Println(admin)
	}

	fmt.Println(admin)
}

func getAdminToken(c *gin.Context) {
	fmt.Print("getting redis token")
	admin := c.Param("admin")
	fmt.Print(admin)
	client := redis.Connect()
	result, e := client.Get(admin).Result()
	if e != nil {
		panic(e)
	}
	fmt.Print(result)
}
package main

import (
	"Go_Gingonic_Redis/router"

	"github.com/gin-gonic/gin"
/*	"github.com/go-redis/redis"*/
)

type LoggedDTO struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}


func main() {

	r := gin.Default() //creamos el router

	//client := redis.NewClient(&redis.Options{
	//	Addr:     "redis:6379",
	//	Password: "",
	//	DB:       0,
	//})
	//
	//_, err1 := client.Ping().Result()
	//if err1 != nil {
	//	panic(err1)
	//}

	v1 := r.Group("/")

	router.Router(v1)

	err := r.Run(":3015") //arrancamos el servidor por el puerto indicado
	if err != nil {
		panic(err)
	}
}

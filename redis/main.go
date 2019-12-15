package main

import (
	"Go_Gingonic_Redis/redis_server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //creamos el router
	redisClient := redis_server.NewClient()
	redisAPI := redis_server.MainAPi{}
	cache := initTopicsAPI(redisClient, &redisAPI)
	r.GET("/test", cache.GetTopics)
	r.GET("/keys", cache.GetKeys)
	r.GET("/remove", cache.DeleteKeys)
	err := r.Run(":3015") //arrancamos el servidor por el puerto indicado
	if err != nil {
		panic(err)
	}
}

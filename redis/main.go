package main

import (
	"Go_Gingonic_Redis/redis_server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default() //creamos el router
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://bank.localhost:3010", "http://forum.localhost:3010"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4500"
		},
	}))
	redisClient := redis_server.NewClient()
	redisAPI := redis_server.MainAPi{}
	cache := initTopicsAPI(redisClient, &redisAPI)
	r.GET("/topics", cache.GetTopics)
	r.GET("/keys", cache.GetKeys)
	r.GET("/remove", cache.DeleteKeys)
	err := r.Run(":3015") //arrancamos el servidor por el puerto indicado
	if err != nil {
		panic(err)
	}
}

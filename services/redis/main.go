package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	r := gin.Default() //creamos el router

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err1 := client.Ping().Result()
	if err1 != nil {
		panic(err1)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Test: "+pong+"  pong "+fmt.Sprint(time.Now().Unix()))
	})
	err := r.Run(":3015") //arrancamos el servidor por el puerto indicado
	if err != nil {
		panic(err)
	}
}

package main

import (
	"Go_Gingonic_Server/greetings"
	"Go_Gingonic_Server/plain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)


func main() {
	//Ignore return
	_ = godotenv.Load()
	//Import port value
	port := os.Getenv("PORT")

	//Load app
	app := gin.New()
/*	app.Use(favicon.New("./favicon.ico"))*/

	//Create a group route
	v1 := app.Group("/api")

	plain.Route(v1)

	//Load route greetings con /api => localhost:${port}/api/salute
	greetings.Salute(v1.Group("/salute"))

	//Start server
	fmt.Println("listen and serve on localhost"+port)
	_ = app.Run(port) // listen and serve on 0.0.0.0:3030
}
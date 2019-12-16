package main

import (
	 "idcards/db"
	_ "idcards/docs"
	"idcards/idCards"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//Ignore return

	port := os.Getenv("GO_SERVER")
	fmt.Println(port)

	//Load app
	app := gin.New()
	//app.Use(favicon.New("./favicon.ico"))

	// create table
	 err := db.CreateTable()
	 	if err != nil {
		panic(err)
	}

	app.Use(cors.Default())
	//CORS
	/*	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))*/
	//Swagger
	url := ginSwagger.URL("http://localhost" + port + "/swagger/doc.json") // The url pointing to API definition
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := app.Group("/api")

	//plain.Route(v1)
	//Load route greetings con /api => localhost:${port}/api/salute
	//greetings.Salute(v1.Group("/salute"))

	idCards.UserRegister(v1.Group("/idCard"))
	idCards.UsersList(v1.Group("/idCards"))
	idCards.UsersSeed(v1.Group("/seeder"))
	idCards.UserLogin(v1.Group("/idCards/login"))

	//Start server
	fmt.Println("listen and serve on localhost " + port)
	_ = app.Run(port)
}

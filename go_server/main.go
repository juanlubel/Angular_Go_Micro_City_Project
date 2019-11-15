package main

import (
	//"Go_Gingonic_Server/driver"
	//"Go_Gingonic_Server/greetings"
	"Go_Gingonic_Server/idCards"
	//"Go_Gingonic_Server/plain"
	//"./user"
	//user "Go_Gingonic_Server/user/routers.go"

	"fmt"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	"os"
)


func main() {
	//Ignore return

	port := os.Getenv("GO_SERVER")
	fmt.Println(port)
	//Load app
	app := gin.New()
	//app.Use(favicon.New("./favicon.ico"))


	v1 := app.Group("/api")


	//plain.Route(v1)
	//Load route greetings con /api => localhost:${port}/api/salute
	//greetings.Salute(v1.Group("/salute"))

	idCards.UserRegister(v1.Group("/user"))
	idCards.UsersList(v1.Group("/users"))

	//Start server
	fmt.Println("listen and serve on localhost "+port)
	_ = app.Run(port)
}



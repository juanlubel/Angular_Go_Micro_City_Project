package main

import (
	router "Go_Gingonic_Server/routers"
	"Go_Gingonic_Server/topics"
	"Go_Gingonic_Server/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB { //Gorm Abre la conexion a base de datos
	var DBHOST = "mysql"
	var DBPORT = "3306"
	var DBUSER = "root"
	var DBPASS = "test"
	var DBNAME = "microcity"
	var DRIVER = "mysql"
	db, err := gorm.Open(DRIVER, DBUSER+":"+DBPASS+"@("+DBHOST+":"+DBPORT+")/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local")
	/* gorm.Open(DRIVER, fmt.Sprintf("%s:%s@tcp("+DBHOST+":"+DBPORT+")/%s", DBUSER, DBPASS, DBNAME)) */
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&topics.Topic{}) //genera las tablas dependiendo de los modelos que le pasemos

	return db
}

func main() {
	db := initDB()   //abrimos la conexion a base de datos
	defer db.Close() //a cerramos

	r := gin.Default() //creamos el router
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://bank.localhost:3010"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4500"
		},
	}))
	jwt := utils.JWT{}
	router.MakePublicTopicRoutes(r, db, &jwt)

	err := r.Run(":3010") //arrancamos el servidor por el puerto indicado
	if err != nil {
		panic(err)
	}
}

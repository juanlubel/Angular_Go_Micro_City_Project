package main

import (
	"os"


	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"Go_Gingonic_Server/admin"
)

func initDB() *gorm.DB { //Gorm Abre la conexion a base de datos

	var DRIVER = "mysql"

	DBNAME := os.Getenv("MYSQL_DATABASE")
	DBUSER := os.Getenv("MYSQL_USER")
	DBPASS := os.Getenv("MYSQL_PASSWORD")
	DBHOST := os.Getenv("MYSQL_HOST")
	DBPORT := os.Getenv("MYSQL_PORT")


	db, err := gorm.Open(DRIVER, DBUSER+":"+DBPASS+"@("+DBHOST+":"+DBPORT+")/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&admin.Admin{}) //genera las tablas dependiendo de los modelos que le pasemos

	return db
}

func main() {
	db := initDB()   //abrimos la conexion a base de datos
	defer db.Close() //a cerramos

	adminApi := InitAdminAPI(db) //illectamos todas las dependencias que necesitamos con wire.

	r := gin.Default() //creamos el router

	v1 := r.Group("/")

	//CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://admin.docker.localhost:3010"},
		AllowMethods:     []string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
	}))
/*	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200/"}*/
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	//r.Use(cors.New(config))


	//creamos las rutas y anexamos el capturador el cual se encontrara en <nombre>API
	//r.Use()
	//v1.Use(admin.AuthMiddleware(true))

	v1.GET("/admin/:id", adminApi.FindByID)
	v1.POST("/admin", adminApi.Create)
	v1.PUT("/admin/:id", adminApi.Update)
	v1.DELETE("/admin/:id", adminApi.Delete)

	//v1.Use(admin.AuthMiddleware(false))
	v1.POST("/admin/login", adminApi.LogIn)

	err := r.Run(":3010") //arrancamos el servidor por el puerto indicado
	if err != nil {
		panic(err)
	}
}

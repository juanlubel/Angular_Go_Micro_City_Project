package main

import (
	"os"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"

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

	//CORS
	makeRoutes(r)
	v1 := r.Group("/")




	v1.GET("/admin/:id", adminApi.FindByID)
	//v1.GET("/ws", adminApi.WSHandler)
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

func makeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()

		/*
			fmt.Printf("c.Request.Method \n")
			fmt.Printf(c.Request.Method)
			fmt.Printf("c.Request.RequestURI \n")
			fmt.Printf(c.Request.RequestURI)
		*/
	}
	r.Use(cors)
}
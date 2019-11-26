package main

import (
	"Go_Gingonic_Server/bank"

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

	db.AutoMigrate(&bank.Bank{}) //genera las tablas dependiendo de los modelos que le pasemos

	return db
}

func main() {
	db := initDB()   //abrimos la conexion a base de datos
	defer db.Close() //a cerramos

	bankAPI := initBankAPI(db) //illectamos todas las dependencias que necesitamos con wire.

	r := gin.Default() //creamos el router

	//creamos las rutas y anexamos el capturador el cual se encontrara en <nombre>API
	r.GET("/banks", bankAPI.FindAll)
	r.GET("/banks/:id", bankAPI.FindByID)
	r.POST("/banks", bankAPI.Create)
	r.PUT("/banks/:id", bankAPI.Update)
	r.DELETE("/banks/:id", bankAPI.Delete)

	err := r.Run(":3010") //arrancamos el servidor por el puerto indicado
	if err != nil {
		panic(err)
	}
}

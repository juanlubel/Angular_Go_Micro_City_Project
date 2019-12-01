package router

import (
	"Go_Gingonic_Server/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MakePublicBankRoutes(r *gin.Engine, db *gorm.DB) {
	bankAPI := initBankAPI(db) //inllectamos todas las dependencias que necesitamos con wire.
	//creamos las rutas y anexamos el capturador el cual se encontrara en <nombre>API
	/* v1.Use() */
	r.GET("/banks", bankAPI.FindAll)
	r.GET("/banks/:id", bankAPI.FindByID)
	r.POST("/banks", bankAPI.Create)
	r.PUT("/banks/:id", bankAPI.Update)
	r.DELETE("/banks/:id", bankAPI.Delete)
}
func MakeAuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.Use(JWTAuthMiddleware(false, utils.Secret))
	r.GET("/auth/:provider", redirectHandler)
	r.GET("/auth/:provider/callback", callbackHandler)
}
func MakePublicAccountsRoutes(r *gin.Engine, db *gorm.DB, jwt *utils.JWT) {
	accountAPI := initAccountsAPI(db, jwt) //inllectamos todas las dependencias que necesitamos con wire.
	//creamos las rutas y anexamos el capturador el cual se encontrara en <nombre>API
	/* v1.Use() */
	r.Use(JWTAuthMiddleware(false, utils.Secret))
	r.GET("/accounts", accountAPI.FindAll)
	r.GET("/account/:name", accountAPI.FindByOwner)
	r.Use(JWTAuthMiddleware(true, utils.Secret))
	r.POST("/account", accountAPI.Create)
	r.POST("/account/login", accountAPI.LogIn)
	/* 	r.PUT("/account/:name", accountAPI.Update)
	   	r.DELETE("/account/:name", accountAPI.Delete) */
}

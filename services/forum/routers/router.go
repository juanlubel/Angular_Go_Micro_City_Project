package router

import (
	"Go_Gingonic_Server/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MakePublicTopicRoutes(r *gin.Engine, db *gorm.DB, jwt *utils.JWT) {
	topicAPI := initTopicsAPI(db, jwt) //inllectamos todas las dependencias que necesitamos con wire.
	//creamos las rutas y anexamos el capturador el cual se encontrara en <nombre>API
	/* r.Use(JWTAuthMiddleware(false, utils.Secret)) */
	r.GET("/topics", topicAPI.FindAll)
	r.GET("/topics/:name", topicAPI.FindByOwner)
	/* r.Use(JWTAuthMiddleware(true, utils.Secret)) */
	r.POST("/topics", topicAPI.Create)
	r.POST("/topic/comment", topicAPI.CreateComment)

	/* r.PUT("/account/:name", topicAPI.Update)
	r.DELETE("/account/:name", topicAPI.Delete) */
}

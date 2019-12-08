package router

import (
	//"Go_Gingonic_Redis/redis"

	"github.com/gin-gonic/gin"
)
func Router(router *gin.RouterGroup) {
	Admin(router.Group("/admin"))
	User(router.Group("/user"))
}

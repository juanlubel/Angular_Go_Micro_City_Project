package router

import (
	"github.com/gin-gonic/gin"
)
func router(router *gin.RouterGroup) {
	Admin(router.Group("/admin"))
}
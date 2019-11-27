package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"Go_Gingonic_Server/common"
)

//func stripBearerPrefixFromTokenString(tok string) (string, error) {
//	// Should be a bearer token
//	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
//		return tok[6:], nil
//	}
//	return tok, nil
//}
//
//
//
//func __AuthMiddleware(c *gin.Context) string {
//
//	id, _ := common.ParseJWT(c.Request.Header.Get(key))
//	return id
//}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	key := "Token"
	return func(c *gin.Context) {

		id, err := common.ParseJWT(c.Request.Header.Get(key))

		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		c.Set("requestBy", id)
	}
}
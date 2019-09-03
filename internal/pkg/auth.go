package app

import (
	"github.com/gin-gonic/gin"
	gorsp "github.com/standard-rsp/gorsp"
	"net/http"
)

var notAuth = []string{"/api/v1/user/login"}

// 验证 token 中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path

		// 无需验证权限
		if authorizationNotRequired(requestPath) {
			c.Next()
			return
		}

		tokenHeader := c.GetHeader("Authorization") //Grab the token from the header

		if !isAuthorized(tokenHeader) {
			response := gorsp.RspBaseData(
				nil,
				gorsp.StatusInvalidToken,
				gorsp.NewMessage("Invalid Token", "error"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Next()
	}
}

// 无需验证权限
func authorizationNotRequired(requestPath string) bool {
	for _, value := range notAuth {
		if value == requestPath {
			return true
		}
	}
	return false
}

// token 是否有效
func isAuthorized(tokenHeader string) bool {
	if tokenHeader == "" {
		return false
	}
	// TODO
	return true
}

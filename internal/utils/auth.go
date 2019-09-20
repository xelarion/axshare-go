package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
	"os"
)

var notAuth = []string{"/api/v1/user/login"}

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

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
			response := ogs.RspBase(ogs.StatusInvalidToken, ogs.ErrorMessage("Invalid Token"))
			c.AbortWithStatusJSON(http.StatusOK, response)
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
	tk := &Token{}

	token, err := jwt.ParseWithClaims(tokenHeader, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_KEY")), nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}

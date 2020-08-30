package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/ogs-go"
	"net/http"
	"os"
	"regexp"
)

var notAuthPattern = []string{
	`/api/v1/user/login`,
	`^/api/v1/axures/.*\d`,
}

/*
JWT claims struct
*/
type TokenClaims struct {
	UserId uint
	jwt.StandardClaims
}

// 验证 token 中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path

		// 无需验证权限
		if isAuthorizationNotRequired(requestPath) {
			c.Next()
			return
		}

		if !isAuthorized(c) {
			response := ogs.RspBase(ogs.StatusInvalidToken, ogs.ErrorMessage("Invalid Token"))
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}

		c.Next()
	}
}

func GetUserIdByToken(tokenString string) (userId uint, err error) {
	tokenClaims := &TokenClaims{}
	_, err = jwt.ParseWithClaims(tokenString, tokenClaims, tokenSecretKeyFunc())

	userId = tokenClaims.UserId
	return userId, err
}

// 无需验证权限
func isAuthorizationNotRequired(requestPath string) bool {
	for _, pattern := range notAuthPattern {
		// 这里循环匹配 requestPath，先添加的先匹配
		reg, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if reg.MatchString(requestPath) {
			return true
		}
	}
	return false
}

// token 是否有效
func isAuthorized(c *gin.Context) bool {
	headerToken := GetHeaderToken(c)
	token, err := getParseToken(headerToken)

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}

func tokenSecretKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_KEY")), nil
	}
}

func getParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, tokenSecretKeyFunc())
	return token, err
}

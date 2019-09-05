package utils

import "github.com/gin-gonic/gin"

func GetHeaderToken(c *gin.Context) string {
	return c.Request.Header.Get("Authorization")
}

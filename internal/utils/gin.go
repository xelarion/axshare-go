package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetHeaderToken(c *gin.Context) string {
	return c.GetHeader("Authorization")
}

func GetBodyParams(c *gin.Context) interface{} {
	var params interface{}
	_ = json.NewDecoder(c.Request.Body).Decode(&params)
	return params
}

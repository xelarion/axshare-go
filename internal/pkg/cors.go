package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

// 解决跨域问题
func AllowRouterCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(os.Getenv("allow_origins"), ",")
	router.Use(cors.New(config))
}

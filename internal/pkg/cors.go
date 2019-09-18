package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

// 解决跨域问题
func AllowRouterCors(router *gin.Engine) {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	config.AllowOrigins = strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	router.Use(cors.New(config))
}

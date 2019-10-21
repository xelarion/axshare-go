package api

import (
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func RunHttpServer() {
	// 设置日志文件
	f, _ := os.OpenFile("log/http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	gin.DefaultWriter = io.MultiWriter(f)

	// 定义路由
	router := gin.Default()

	utils.AllowRouterCors(router)

	router.Use(utils.TokenAuthMiddleware())

	SetV1Router(router)

	port := os.Getenv("HTTP_PORT")
	_ = router.Run(":" + port)

}

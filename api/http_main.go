package api

import (
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
)

func HttpServerRun() {
	// 设置日志文件
	f, _ := os.OpenFile("log/axshare_go_http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	gin.DefaultWriter = io.MultiWriter(f)

	// 定义路由
	router := gin.Default()

	utils.AllowRouterCors(router)

	router.Use(utils.TokenAuthMiddleware())

	SetV1Router(router)

	port := viper.GetString("http.port")
	_ = router.Run(":" + port)

}

package qiniu

import (
	app "axshare_go/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

func CreateUploadToken(c *gin.Context) {
	token := app.GenUpToken()
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), token))
}

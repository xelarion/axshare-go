package qiniu

import (
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

func CreateUploadToken(c *gin.Context) {
	token := utils.GenUpToken()
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), token))
}

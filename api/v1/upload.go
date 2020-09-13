package v1

import (
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/ogs-go"
	"net/http"
	"os"
)

func CreateUploadToken(c *gin.Context) {
	token := utils.GenUpToken()
	data := gin.H{"token": token, "upload_url": os.Getenv("QINIU_UPLOAD_URL")}
	c.JSON(http.StatusOK, ogs.RspOKWithData("", data))
}

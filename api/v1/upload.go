package v1

import (
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/ogs-go"
	"net/http"
)

func CreateUploadToken(c *gin.Context) {
	data := gin.H{
		"token":      models.GenerateUploadToken(),
		"upload_url": models.CacheConfig.QiniuUploadUrl,
	}
	c.JSON(http.StatusOK, ogs.RspOKWithData("", data))
}

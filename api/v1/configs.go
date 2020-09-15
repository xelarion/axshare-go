package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
	"github.com/xandercheung/ogs-go"
	"net/http"
)

func GetPublicConfig(c *gin.Context) {
	c.JSON(http.StatusOK, ogs.RspOKWithData("", models.CacheConfig.PublicConfig))
}

func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, ogs.RspOKWithData("", models.CacheConfig))
}

func UpdateConfig(c *gin.Context) {
	models.CacheConfigLock.Lock()
	defer models.CacheConfigLock.Unlock()

	_ = json.NewDecoder(c.Request.Body).Decode(&models.CacheConfig)
	models.CacheConfig.IsValid = true

	if err := db.AxshareDb.Save(&models.CacheConfig).Error; err != nil {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusUpdateFailed, err.Error()))
	} else {
		acct.Utils.JSON(c, ogs.RspOK("Update Successfully"))
	}
}

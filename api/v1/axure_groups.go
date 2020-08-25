package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/ogs-go"
	"net/http"
)

func GetAxureGroups(c *gin.Context) {
	var axureGroups []models.AxureGroup
	db.AxshareDb.Model(&models.AxureGroup{}).Find(&axureGroups)
	c.JSON(http.StatusOK, ogs.RspOKWithData(
		ogs.BlankMessage(), axureGroups))
}

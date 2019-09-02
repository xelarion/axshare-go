package axure

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	gorsp "github.com/standard-rsp/gorsp"
	"net/http"
	"strconv"
)

func GetAxures(c *gin.Context) {
	if c.Param("axure_group_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	axureGroupId, _ := strconv.ParseUint(c.Param("axure_group_id"), 10, 64)
	var axures []models.Axure
	db.AxshareDb.Where(&models.Axure{AxureGroupID: uint(axureGroupId)}).Find(&axures)
	c.JSON(http.StatusOK, gorsp.RspPagData(
		FormatList(axures), gorsp.OK,
		gorsp.NewMessage("", ""),
		gorsp.NewPaginate(1, 102, 10)))
}

package axure

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
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
	db.AxshareDb.Where(&models.Axure{AxureGroupId: uint(axureGroupId)}).Order("id desc").Find(&axures)
	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		FormatList(axures),
		ogs.NewPaginate(1, 101, 10)))
}

func GetAxure(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	axure := models.Axure{}
	db.AxshareDb.First(&axure, id)
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure))
}

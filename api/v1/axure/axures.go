package axure

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	go_rsp "github.com/standard-rsp/go-rsp"
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
	c.JSON(http.StatusOK, go_rsp.RspPagData(
		FormatList(axures), go_rsp.OK,
		go_rsp.NewBaseMessage("", ""),
		go_rsp.BasePaginate{CurrentPage: 1, TotalPages: 20, TotalCount: 201, PerPage: 10}))
}

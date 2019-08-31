package attachment

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	go_rsp "github.com/standard-rsp/go-rsp"
	"net/http"
	"strconv"
)

func GetAttachments(c *gin.Context) {
	_ = errors.New("fsdf")

	if c.Param("axure_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	axureId, _ := strconv.ParseUint(c.Param("axure_id"), 10, 64)
	axure := models.Axure{}
	db.AxshareDb.First(&axure, axureId)
	var attachments []models.Attachment
	//db.AxshareDb.Model(&axure).Related(&attachments, "Attachments")
	db.AxshareDb.Model(&axure).Preload("User").Association("Attachments").Find(&attachments)
	c.JSON(http.StatusOK, go_rsp.RspPagData(
		FormatList(attachments), go_rsp.OK,
		go_rsp.NewBaseMessage("", ""),
		go_rsp.BasePaginate{CurrentPage: 1, TotalPages: 20, TotalCount: 201, PerPage: 10}))
}

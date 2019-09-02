package attachment

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	gorsp "github.com/standard-rsp/gorsp"
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
	c.JSON(http.StatusOK, gorsp.RspPagData(
		FormatList(attachments), gorsp.OK,
		gorsp.NewMessage("", ""),
		gorsp.NewPaginate(1, 102, 10)))
}

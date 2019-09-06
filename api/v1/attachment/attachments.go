package attachment

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
	"strconv"
)

func GetAttachments(c *gin.Context) {
	if c.Param("axure_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	axureId, _ := strconv.ParseUint(c.Param("axure_id"), 10, 64)
	axure := models.Axure{}
	db.AxshareDb.First(&axure, axureId)
	var attachments []models.Attachment
	//db.AxshareDb.Model(&axure).Related(&attachments, "Attachments")
	db.AxshareDb.Model(&axure).Preload("User").Order("id desc").Association("Attachments").Find(&attachments)
	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		FormatList(attachments),
		ogs.NewPaginate(1, 101, 10)))
}

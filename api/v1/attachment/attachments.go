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
	var attachments []models.Attachment
	db.AxshareDb.Debug().Model(&models.Attachment{}).Where(
		"axure_id = ?", axureId).Order("id desc").Preload("User").Find(&attachments)
	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		FormatList(attachments),
		ogs.NewPaginate(1, 101, 10)))
}

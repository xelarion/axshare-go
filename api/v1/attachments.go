package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

func GetAttachments(c *gin.Context) {
	if c.Param("axure_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	axureId, _ := utils.ParseUint(c.Param("axure_id"))
	var attachments []models.Attachment
	db.AxshareDb.Model(&models.Attachment{}).Where(
		"axure_id = ?", axureId).Order("id desc").Preload("User").Find(&attachments)

	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		FormatAttachmentList(attachments),
		ogs.NewPaginate(1, 101, 10)))
}

func FormatAttachmentList(attachments []models.Attachment) []map[string]interface{} {
	var json = make([]map[string]interface{}, len(attachments))
	for i, attachment := range attachments {
		var data = make(map[string]interface{})
		data["id"] = attachment.ID
		data["desc"] = attachment.Desc
		data["download_url"] = attachment.DownloadUrl()
		data["is_released"] = attachment.IsReleased()
		data["web_link"] = attachment.WebLink()
		data["created_at"] = utils.FormatDateTime(attachment.CreatedAt)
		data["updated_at"] = utils.FormatDateTime(attachment.UpdatedAt)
		data["user"] = map[string]interface{}{"username": attachment.User.Username}
		json[i] = data
	}
	return json
}

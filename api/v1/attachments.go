package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
	"github.com/xandercheung/ogs-go"
	"net/http"
)

func GetAttachments(c *gin.Context) {
	if c.Param("axure_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	axureId := utils.ParseUint(c.Param("axure_id"))
	var attachments []models.Attachment
	db.AxshareDb.Model(&models.Attachment{}).Where(
		"axure_id = ?", axureId).Order("id desc").Preload("User").Find(&attachments)

	c.JSON(http.StatusOK,
		ogs.RspOKWithData(
			ogs.BlankMessage(),
			FormatAttachmentList(attachments)))
}

func GetAllAttachments(c *gin.Context) {
	var attachments []models.Attachment
	relation := db.AxshareDb.Model(&models.Attachment{}).Order("id desc")
	relation, paginate := acct.Utils.PaginateGin(relation, c)
	relation.Preload("User").Preload("Axure").Preload("Axure.AxureGroup").Find(&attachments)

	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		FormatAttachmentActivityList(attachments),
		paginate))
}

func FormatAttachmentActivityList(attachments []models.Attachment) []map[string]interface{} {
	var json = make([]map[string]interface{}, len(attachments))
	for i, attachment := range attachments {
		var data = make(map[string]interface{})
		data["id"] = attachment.ID
		data["desc"] = attachment.Desc

		data["axure"] = map[string]interface{}{"name": attachment.Axure.Name}

		axureGroup := attachment.Axure.AxureGroup
		data["axure_group"] = map[string]interface{}{
			"name": axureGroup.Name,
			"id":   axureGroup.ID,
		}

		data["download_url"] = attachment.DownloadUrl()
		data["is_released"] = attachment.IsReleased()
		data["web_link"] = attachment.WebLink()
		data["created_at"] = utils.FormatDateTime(attachment.CreatedAt)
		data["updated_at"] = utils.FormatDateTime(attachment.UpdatedAt)

		data["user"] = map[string]interface{}{"nickname": attachment.User.Nickname}
		json[i] = data
	}
	return json
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
		data["user"] = map[string]interface{}{"nickname": attachment.User.Nickname}
		json[i] = data
	}
	return json
}

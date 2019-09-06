package attachment

import (
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
)

func FormatList(attachments []models.Attachment) []map[string]interface{} {
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

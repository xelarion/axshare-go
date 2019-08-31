package attachment

import "axshare_go/internal/models"

func FormatList(attachments []models.Attachment) []map[string]interface{} {
	var json = make([]map[string]interface{}, len(attachments))
	for i, attachment := range attachments {
		var data = make(map[string]interface{})
		data["id"] = attachment.ID
		data["desc"] = attachment.Desc
		data["download_url"] = attachment.DownloadUrl()
		data["is_released"] = attachment.IsReleased()
		data["web_link"] = attachment.WebLink()
		data["updated_at"] = attachment.UpdatedAt
		data["user"] = map[string]interface{}{"nickname": attachment.User.Nickname}
		json[i] = data
	}
	return json
}

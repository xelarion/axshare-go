package axure

import "axshare_go/internal/models"

func FormatList(axures []models.Axure) []map[string]interface{} {
	var json = make([]map[string]interface{}, len(axures))
	for i, axure := range axures {
		var data = make(map[string]interface{})
		data["id"] = axure.ID
		data["desc"] = axure.Desc
		data["updated_at"] = axure.UpdatedAt
		data["is_released"] = axure.IsReleased()
		data["web_link"] = axure.WebLink()
		data["permanent_link"] = axure.PermanentLink()
		json[i] = data
	}
	return json
}

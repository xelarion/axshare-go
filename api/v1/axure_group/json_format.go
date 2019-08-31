package axure_group

import "axshare_go/internal/models"

func FormatList(axureGroups []models.AxureGroup) []map[string]interface{} {
	var datas = make([]map[string]interface{}, len(axureGroups))
	for i, axureGroup := range axureGroups {
		var data = make(map[string]interface{})
		data["id"] = axureGroup.ID
		data["name"] = axureGroup.Name
		data["desc"] = axureGroup.Desc
		datas[i] = data
	}
	return datas
}

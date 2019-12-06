package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/pg"
	"axshare_go/internal/task"
	"axshare_go/internal/utils"
	"encoding/json"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

func GetAxures(c *gin.Context) {
	if c.Param("axure_group_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	var axures []models.Axure
	axureGroupId := utils.ParseUint(c.Param("axure_group_id"))
	relation := db.AxshareDb.Model(&models.Axure{}).Where(&models.Axure{AxureGroupId: uint(axureGroupId)}).Order("id desc")

	searchConditions := utils.StringToMap(c.Query("search_conditions"))
	if len(searchConditions) > 0 && len(searchConditions["name"].(string)) > 0 {
		relation = relation.Where("name LIKE ?", "%"+searchConditions["name"].(string)+"%")
	}

	relation, paginate := pg.PaginateGin(relation, c)
	relation.Find(&axures)

	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(), FormatAxureList(axures), paginate))
}

func GetAxure(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	id := utils.ParseUint(c.Param("id"))
	axure := models.Axure{}
	db.AxshareDb.First(&axure, id)
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure))
}

func GetAxureWebInfo(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	id := utils.ParseUint(c.Param("id"))
	axure := models.Axure{}
	db.AxshareDb.First(&axure, id)

	if c.Query("key") != axure.SecretKey {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure.WebLink()))
}

func UpdateAxure(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	tx := db.AxshareDb.Begin()

	id := utils.ParseUint(c.Param("id"))
	params := utils.GetBodyParams(c)

	axure := models.Axure{}
	db.AxshareDb.First(&axure, id)
	jsonBody, _ := json.Marshal(params)
	_ = json.Unmarshal(jsonBody, &axure)
	db.AxshareDb.Model(&axure).Updates(params)

	attachment := models.Attachment{AxureId: axure.ID, UserId: CurrentUserId(c)}
	jsonBody1, _ := json.Marshal(params["attachment"])
	_ = json.Unmarshal(jsonBody1, &attachment)
	db.AxshareDb.Create(&attachment)

	tx.Commit()
	releaseAttachment(attachment.ID)

	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure))
}

func CreateAxure(c *gin.Context) {
	tx := db.AxshareDb.Begin()

	axureGroupId := utils.ParseUint(c.Param("axure_group_id"))
	params := utils.GetBodyParams(c)

	//axure := models.Axure{Name: params["name"].(string)}
	axure := models.Axure{AxureGroupId: uint(axureGroupId)}
	jsonBody1, _ := json.Marshal(params)
	_ = json.Unmarshal(jsonBody1, &axure)
	db.AxshareDb.Create(&axure)

	attachment := models.Attachment{AxureId: axure.ID, UserId: CurrentUserId(c)}
	jsonBody, _ := json.Marshal(params["attachment"])
	_ = json.Unmarshal(jsonBody, &attachment)
	db.AxshareDb.Create(&attachment)

	tx.Commit()
	releaseAttachment(attachment.ID)

	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.SuccessMessage("创建成功！"), axure))
}

func FormatAxureList(axures []models.Axure) []map[string]interface{} {
	var json = make([]map[string]interface{}, len(axures))
	for i, axure := range axures {
		var data = make(map[string]interface{})
		data["id"] = axure.ID
		data["name"] = axure.Name
		data["updated_at"] = utils.FormatDateTime(axure.UpdatedAt)
		data["is_released"] = axure.IsReleased()
		data["web_link"] = axure.WebLink()
		data["permanent_link"] = axure.PermanentLink()
		data["axure_group_id"] = axure.AxureGroupId
		json[i] = data
	}
	return json
}

func releaseAttachment(attachmentId uint) {
	newTask := tasks.Signature{
		Name: "release_attachment",
		Args: []tasks.Arg{
			{
				Type:  "uint",
				Value: attachmentId,
			},
		},
	}
	_ = task.Send(&newTask)
}

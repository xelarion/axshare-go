package axure

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
	"strconv"
)

func GetAxures(c *gin.Context) {
	if c.Param("axure_group_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	axureGroupId, _ := strconv.ParseUint(c.Param("axure_group_id"), 10, 64)
	var axures []models.Axure
	db.AxshareDb.Where(&models.Axure{AxureGroupId: uint(axureGroupId)}).Order("id desc").Find(&axures)
	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		FormatList(axures),
		ogs.NewPaginate(1, 101, 10)))
}

func GetAxure(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	axure := models.Axure{}
	db.AxshareDb.First(&axure, id)
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure))
}

func UpdateAxure(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	tx := db.AxshareDb.Begin()

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	params := utils.GetBodyParams(c).(map[string]interface{})

	axure := models.Axure{}
	db.AxshareDb.First(&axure, id)
	jsonBody, _ := json.Marshal(params)
	_ = json.Unmarshal(jsonBody, &axure)
	db.AxshareDb.Debug().Model(&axure).Updates(params)

	attachment := models.Attachment{AxureId: axure.ID}
	jsonBody1, _ := json.Marshal(params["attachment"])
	_ = json.Unmarshal(jsonBody1, &attachment)
	db.AxshareDb.Debug().Create(&attachment)

	tx.Commit()

	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure))
}

func CreateAxure(c *gin.Context) {
	tx := db.AxshareDb.Begin()

	axureGroupId, _ := strconv.ParseUint(c.Param("axure_group_id"), 10, 64)
	params := utils.GetBodyParams(c).(map[string]interface{})

	//axure := models.Axure{Name: params["name"].(string)}
	axure := models.Axure{AxureGroupId: uint(axureGroupId)}
	jsonBody1, _ := json.Marshal(params)
	_ = json.Unmarshal(jsonBody1, &axure)
	db.AxshareDb.Debug().Create(&axure)

	attachment := models.Attachment{AxureId: axure.ID}
	jsonBody, _ := json.Marshal(params["attachment"])
	_ = json.Unmarshal(jsonBody, &attachment)
	db.AxshareDb.Debug().Create(&attachment)

	tx.Commit()

	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.SuccessMessage("创建成功！"), axure))
}

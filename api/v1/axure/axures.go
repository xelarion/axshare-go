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

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	axure := models.Axure{}
	params := utils.GetBodyParams(c)
	db.AxshareDb.Debug().Model(&axure).First(&axure, id).Update(params)
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure))
}

func CreateAxure(c *gin.Context) {
	tx := db.AxshareDb.Begin()
	axureGroupId, _ := strconv.ParseUint(c.Param("axure_group_id"), 10, 64)
	params := utils.GetBodyParams(c).(map[string]interface{})
	axure := models.Axure{Name: params["name"].(string)}
	axure.AxureGroupId = uint(axureGroupId)
	attachment := models.Attachment{}
	jsonBody, _ := json.Marshal(params["attachment"])
	_ = json.Unmarshal(jsonBody, &attachment)
	db.AxshareDb.Debug().Create(&axure)
	attachment.AxureId = axure.ID
	db.AxshareDb.Debug().Create(&attachment)
	tx.Commit()
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), axure))
}

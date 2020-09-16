package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
	"github.com/xandercheung/ogs-go"
	"gorm.io/gorm"
	"net/http"
)

func GetAxureGroups(c *gin.Context) {
	var axureGroups []models.AxureGroup
	db.AxshareDb.Model(&models.AxureGroup{}).Find(&axureGroups)
	acct.Utils.JSON(c, ogs.RspOKWithData("", axureGroups))
}

func GetAxureGroup(c *gin.Context) {
	axureGroup, err := loadAxureGroup(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, ogs.RspOKWithData("", axureGroup))
}

func CreateAxureGroup(c *gin.Context) {
	axureGroup := models.AxureGroup{}
	if err := json.NewDecoder(c.Request.Body).Decode(&axureGroup); err != nil {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusBadParams, "Bad Params"))
		return
	}

	if err := db.AxshareDb.Create(&axureGroup).Error; err != nil {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusCreateFailed, err.Error()))
		return
	}

	acct.Utils.JSON(c, ogs.RspOKWithData("Create Successfully", axureGroup))
}

func UpdateAxureGroup(c *gin.Context) {
	axureGroup, err := loadAxureGroup(c)
	if err != nil {
		return
	}
	params := utils.GetBodyParams(c)
	if err = db.AxshareDb.Model(&axureGroup).Updates(params).Error; err != nil {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusUpdateFailed, err.Error()))
	} else {
		acct.Utils.JSON(c, ogs.RspOK("Update Successfully"))
	}
}

func loadAxureGroup(c *gin.Context) (axureGroup models.AxureGroup, err error) {
	id := utils.ParseUint(c.Param("id"))
	db.AxshareDb.First(&axureGroup, id)

	if axureGroup.ID == 0 {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusUserNotFound, "AxureGroup Not Found"))
		return axureGroup, gorm.ErrRecordNotFound
	}

	return axureGroup, nil
}

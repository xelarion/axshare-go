package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/task"
	"axshare_go/internal/utils"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
	"github.com/xandercheung/ogs-go"
	"gorm.io/gorm"
	"net/http"
)

type AxureFormParams struct {
	Name         string `json:"name"`
	AxureGroupId uint   `json:"axure_group_id"`
	Attachment   struct {
		Desc     string `json:"desc"`
		FileHash string `json:"file_hash"`
	}
}

func GetAxures(c *gin.Context) {
	if c.Param("axure_group_id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	var axures []models.Axure
	axureGroupId := utils.ParseUint(c.Param("axure_group_id"))
	relation := db.AxshareDb.Model(&models.Axure{}).Where(&models.Axure{AxureGroupId: axureGroupId}).Order("id desc")

	searchConditions := utils.StringToMap(c.Query("search_conditions"))
	if len(searchConditions) > 0 && len(searchConditions["name"].(string)) > 0 {
		relation = relation.Where("name LIKE ?", "%"+searchConditions["name"].(string)+"%")
	}

	relation, paginate := acct.Utils.PaginateGin(relation, c)
	relation.Find(&axures)

	c.JSON(http.StatusOK, ogs.RspOKWithPaginate("", FormatAxureList(axures), paginate))
}

func GetAxure(c *gin.Context) {
	axure, err := loadAxure(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, ogs.RspOKWithData("", axure))
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

	c.JSON(http.StatusOK, ogs.RspOKWithData("", axure.WebLink()))
}

func CreateAxure(c *gin.Context) {
	axureGroupId := utils.ParseUint(c.Param("axure_group_id"))
	params := utils.GetBodyParams(c)

	axure := models.Axure{AxureGroupId: axureGroupId}
	axure.Name, _ = params["name"].(string)

	attachment := models.Attachment{}
	attachmentParams, _ := params["attachment"].(map[string]interface{})
	attachment.Desc, _ = attachmentParams["desc"].(string)
	attachment.FileHash, _ = attachmentParams["file_hash"].(string)
	attachment.AccountId = CurrentAccountId(c)

	if err := createAxureAndAttachment(&axure, &attachment); err != nil {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusCreateFailed, err.Error()))
	} else {
		if attachment.IsFileUploaded() {
			releaseAttachment(attachment.ID)
		}
		acct.Utils.JSON(c, ogs.RspOK("Create Successfully"))
	}
}

func UpdateAxure(c *gin.Context) {
	axure, err := loadAxure(c)
	if err != nil {
		return
	}

	params := utils.GetBodyParams(c)
	axure.Name, _ = params["name"].(string)

	attachment := models.Attachment{
		AccountId: CurrentAccountId(c),
		AxureId:   axure.ID,
	}
	attachmentParams, _ := params["attachment"].(map[string]interface{})
	attachment.Desc, _ = attachmentParams["desc"].(string)
	attachment.FileHash, _ = attachmentParams["file_hash"].(string)

	if err := updateAxureAndAttachment(&axure, &attachment); err != nil {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusUpdateFailed, "Update Failed"))
	} else {
		if attachment.IsFileUploaded() {
			releaseAttachment(attachment.ID)
		}

		acct.Utils.JSON(c, ogs.RspOK("Update Successfully"))
	}
}

func DestroyAxure(c *gin.Context) {
	axure, err := loadAxure(c)
	if err != nil {
		return
	}

	if err := db.AxshareDb.Delete(&axure).Error; err != nil {
		acct.Utils.JSON(c, ogs.RspOK(err.Error()))
	} else {
		cleanAxure(axure.ID)
		acct.Utils.JSON(c, ogs.RspOK("Destroy Successfully"))
	}
}

func FormatAxureList(axures []models.Axure) []map[string]interface{} {
	var json = make([]map[string]interface{}, len(axures))
	for i, axure := range axures {
		var data = make(map[string]interface{})
		attachment := axure.LatestAttachment()

		data["id"] = axure.ID
		data["name"] = axure.Name
		data["updated_at"] = utils.FormatDateTime(axure.UpdatedAt)
		data["release_status"] = attachment.ReleaseStatus
		data["release_error"] = attachment.ReleaseError
		data["web_link"] = axure.WebLink()
		data["permanent_link"] = axure.PermanentLink()
		data["axure_group_id"] = axure.AxureGroupId
		json[i] = data
	}
	return json
}

func createAxureAndAttachment(axure *models.Axure, attachment *models.Attachment) error {
	tx := db.AxshareDb.Begin()

	if err := tx.Create(&axure).Error; err != nil {
		return err
	}

	attachment.AxureId = axure.ID
	if err := tx.Create(&attachment).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func updateAxureAndAttachment(axure *models.Axure, attachment *models.Attachment) error {
	tx := db.AxshareDb.Begin()

	if attachment.IsFileUploaded() {
		if err := tx.Model(&axure).Update("name", axure.Name).Error; err != nil {
			return err
		}

		if err := tx.Create(&attachment).Error; err != nil {
			return err
		}
	} else {
		if err := tx.Model(&axure).UpdateColumn("name", axure.Name).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
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

func cleanAxure(axureId uint) {
	newTask := tasks.Signature{
		Name: "clean_deleted_axure_attachments",
		Args: []tasks.Arg{
			{
				Type:  "uint",
				Value: axureId,
			},
		},
	}
	_ = task.Send(&newTask)
}

func loadAxure(c *gin.Context) (axure models.Axure, err error) {
	id := utils.ParseUint(c.Param("id"))
	db.AxshareDb.First(&axure, id)

	if axure.ID == 0 {
		acct.Utils.JSON(c, ogs.RspError(ogs.StatusUserNotFound, "Axure Not Found"))
		return axure, gorm.ErrRecordNotFound
	}

	return axure, nil
}

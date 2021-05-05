package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
	"github.com/xandercheung/ogs-go"
	"gorm.io/gorm"
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
		"axure_id = ?", axureId).Order("id desc").Preload("Account").Find(&attachments)

	c.JSON(http.StatusOK, ogs.RspOKWithData("", FormatAttachmentList(attachments)))
}

func GetAllAttachments(c *gin.Context) {
	var attachments []models.Attachment
	relation := db.AxshareDb.Model(&models.Attachment{}).Order("id desc")
	relation, paginate := acct.Utils.PaginateGin(relation, c)
	relation.
		Preload("Account", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Preload("Axure", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Preload("Axure.AxureGroup").
		Find(&attachments)

	c.JSON(http.StatusOK, ogs.RspOKWithPaginate("",
		FormatAttachmentActivityList(attachments),
		paginate))
}

func ReleaseAttachment(c *gin.Context) {
	id := utils.ParseUint(c.Param("id"))
	attachment := models.FindAttachment(id)
	if attachment.ID == 0 {
		acct.Utils.JSON(c, ogs.RspError(1, "操作失败，该记录不存在"))
	}

	if attachment.IsFileUploaded() {
		releaseAttachment(attachment.ID)
		acct.Utils.JSON(c, ogs.RspOK("操作成功，请稍后刷新页面查看"))
	} else {
		acct.Utils.JSON(c, ogs.RspError(1, "操作失败，该附件文件不存在"))
	}
}

func CleanAttachment(c *gin.Context) {
	id := utils.ParseUint(c.Param("id"))
	attachment := models.FindAttachment(id)
	if attachment.ID == 0 {
		acct.Utils.JSON(c, ogs.RspError(1, "操作失败，该记录不存在"))
	}

	if err := attachment.CleanAxureFileDir(); err != nil {
		acct.Utils.JSON(c, ogs.RspError(1, err.Error()))
	} else {
		acct.Utils.JSON(c, ogs.RspOK("操作成功"))
	}
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
		data["release_status"] = attachment.ReleaseStatus
		data["release_error"] = attachment.ReleaseError
		data["web_link"] = attachment.WebLink()
		data["created_at"] = utils.FormatDateTime(attachment.CreatedAt)
		data["updated_at"] = utils.FormatDateTime(attachment.UpdatedAt)

		data["user"] = map[string]interface{}{
			"nickname": attachment.Account.Nickname,
			"username": attachment.Account.Username,
		}
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
		data["release_status"] = attachment.ReleaseStatus
		data["release_error"] = attachment.ReleaseError
		data["web_link"] = attachment.WebLink()
		data["created_at"] = utils.FormatDateTime(attachment.CreatedAt)
		data["updated_at"] = utils.FormatDateTime(attachment.UpdatedAt)
		data["user"] = map[string]interface{}{
			"nickname": attachment.Account.Nickname,
			"username": attachment.Account.Username,
		}
		json[i] = data
	}
	return json
}

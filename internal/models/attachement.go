package models

import (
	"axshare_go/internal/db"
	tasks "axshare_go/internal/task"
	"axshare_go/internal/utils"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Attachment struct {
	gorm.Model
	Desc     string `json:"desc" xml:"desc" binding:"required"`
	Link     string `json:"link"`
	FileHash string `json:"file_hash"`
	AxureId  uint   `json:"axure_id" gorm:"index" xml:"axure_id" binding:"required"`
	UserId   uint   `gorm:"index" json:"user_id"`
	User     User   `json:"user"`
}

func ReleaseFile(attachmentId uint) {
	tx := db.AxshareDb.Begin()

	axure := Axure{}
	attachment := Attachment{}
	db.AxshareDb.First(&attachment, attachmentId)
	db.AxshareDb.Model(&attachment).Related(&axure)

	webLink := tasks.DeployAxure(attachment.DownloadUrl(), attachment.genFileName())

	db.AxshareDb.Model(&axure).Update("link", webLink)
	db.AxshareDb.Model(&attachment).Update("link", webLink)
	tx.Commit()
}

func (c *Attachment) genFileName() string {
	axure := Axure{}
	db.AxshareDb.Model(&c).Related(&axure)

	fileName := strings.Join([]string{
		utils.FormatUint(axure.AxureGroupId),
		utils.StrftimeDateTime(c.CreatedAt, "20060102150405"),
		c.FileHash,
	}, "_")

	return fileName
}

// 文件是否解压
func (c *Attachment) IsReleased() bool {
	isReleased := len(c.Link) > 0
	return isReleased
}

// 原型静态web链接
func (c *Attachment) WebLink() string {
	if !c.IsReleased() {
		return ""
	}
	webHost := viper.GetString("web_host")
	return webHost + c.Link
}

// 原型压缩包下载地址
func (c *Attachment) DownloadUrl() string {
	domain := os.Getenv("QINIU_BUCKET_DOMAIN")
	return domain + "/" + c.FileHash
}

package models

import (
	"axshare_go/internal/utils"
	"github.com/xandercheung/acct"
	"gorm.io/gorm"
	"os"
	"strings"
)

type Attachment struct {
	gorm.Model
	Desc     string       `json:"desc" xml:"desc" binding:"required"`
	Link     string       `json:"link"`
	FileHash string       `json:"file_hash"`
	AxureId  uint         `json:"axure_id" gorm:"index" xml:"axure_id" binding:"required"`
	Axure    Axure        `json:"axure" gorm:"foreignKey:AxureId"`
	UserId   uint         `gorm:"index" json:"user_id"`
	User     acct.Account `json:"user" gorm:"foreignKey:UserId"`
}

func (c *Attachment) GenFileName() string {
	axure := FindAxure(c.AxureId)
	fileName := strings.Join([]string{
		utils.FormatUint(axure.AxureGroupId),
		utils.Strftime(c.CreatedAt, "20060102150405"),
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
	webHost := os.Getenv("AXURE_HOST")
	return webHost + c.Link
}

// 原型压缩包下载地址
func (c *Attachment) DownloadUrl() string {
	domain := os.Getenv("QINIU_BUCKET_DOMAIN")
	return domain + "/" + c.FileHash
}
